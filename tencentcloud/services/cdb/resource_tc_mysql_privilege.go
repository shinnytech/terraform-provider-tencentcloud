package cdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	sdkError "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"

	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type ResourceTencentCloudMysqlPrivilegeId struct {
	MysqlId     string
	AccountName string
	AccountHost string `json:"AccountHost,omitempty"`
}

func ResourceTencentCloudMysqlPrivilege() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMysqlPrivilegeCreate,
		Read:   resourceTencentCloudMysqlPrivilegeRead,
		Update: resourceTencentCloudMysqlPrivilegeUpdate,
		Delete: resourceTencentCloudMysqlPrivilegeDelete,
		Schema: map[string]*schema.Schema{
			"mysql_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Instance ID.",
			},
			"account_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errs []error) {
					if map[string]bool{
						"root":        true,
						"mysql.sys":   true,
						"tencentroot": true,
					}[v.(string)] {
						errs = append(errs, errors.New("account_name is forbidden to be "+v.(string)))
					}
					return
				},
				Description: "Account name.the forbidden value is:root,mysql.sys,tencentroot.",
			},
			"account_host": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     MYSQL_DEFAULT_ACCOUNT_HOST,
				Description: "Account host, default is `%`.",
			},
			"global": {
				Type: schema.TypeSet,
				// Workaround： 将Required变为Optional+Computed.
				// Issue: plugin sdk 会错误地在create时候把赋值给required的`[]`变为nullVal, 然后在read变回`[]`,
				//        见 https://github.com/hashicorp/terraform-plugin-sdk/issues/766
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return helper.HashString(v.(string))
				},
				Description: `Global privileges. available values for Privileges:` + strings.Join(MYSQL_GlOBAL_PRIVILEGE, ",") + ".",
			},
			"database": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Database privileges list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Database name.",
						},
						"privileges": {
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: `Database privilege.available values for Privileges:` + strings.Join(MYSQL_DATABASE_PRIVILEGE, ",") + ".",
						},
					},
				},
			},
			"table": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Table privileges list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Database name.",
						},
						"table_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Table name.",
						},
						"privileges": {
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: `Table privilege.available values for Privileges:` + strings.Join(MYSQL_TABLE_PRIVILEGE, ",") + ".",
						},
					},
				},
			},
			"column": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Column privileges list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Database name.",
						},
						"table_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Table name.",
						},
						"column_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Column name.",
						},
						"privileges": {
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: `Column privilege.available values for Privileges:` + strings.Join(MYSQL_COLUMN_PRIVILEGE, ",") + ".",
						},
					},
				},
			},
		},
	}
}

func (me *ResourceTencentCloudMysqlPrivilegeId) update(ctx context.Context, d *schema.ResourceData, meta interface{}) error {

	if me.AccountHost == "" {
		me.AccountHost = MYSQL_DEFAULT_ACCOUNT_HOST
	}

	sp := func(str interface{}) *string { v := str.(string); return &v }

	request := cdb.NewModifyAccountPrivilegesRequest()
	request.InstanceId = &me.MysqlId

	var accountInfo = cdb.Account{User: &me.AccountName, Host: helper.String(me.AccountHost)}

	request.Accounts = []*cdb.Account{&accountInfo}

	if d != nil {
		sliceInterface := d.Get("global").(*schema.Set).List()
		if len(sliceInterface) > 0 {
			request.GlobalPrivileges = make([]*string, 0, len(sliceInterface))
			for _, v := range sliceInterface {
				ptr := sp(v)
				if !tccommon.IsContains(MYSQL_GlOBAL_PRIVILEGE, *ptr) {
					return errors.New("global privileges not support " + *ptr)
				}
				request.GlobalPrivileges = append(request.GlobalPrivileges, ptr)
			}
		}

		if v, ok := d.GetOk("database"); ok {
			for _, item := range v.(*schema.Set).List() {
				dMap := item.(map[string]interface{})
				privilege := cdb.DatabasePrivilege{}
				if v, ok := dMap["database_name"]; ok {
					privilege.Database = helper.String(v.(string))
				}
				if v, ok := dMap["privileges"]; ok {
					privilegeList := []*string{}
					for _, v := range v.(*schema.Set).List() {
						privilegeList = append(privilegeList, helper.String(v.(string)))
					}
					privilege.Privileges = privilegeList
				}
				request.DatabasePrivileges = append(request.DatabasePrivileges, &privilege)
			}
		}

		if v, ok := d.GetOk("table"); ok {
			for _, item := range v.(*schema.Set).List() {
				dMap := item.(map[string]interface{})
				privilege := cdb.TablePrivilege{}
				if v, ok := dMap["database_name"]; ok {
					privilege.Database = helper.String(v.(string))
				}
				if v, ok := dMap["table_name"]; ok {
					privilege.Table = helper.String(v.(string))
				}
				if v, ok := dMap["privileges"]; ok {
					privilegeList := []*string{}
					for _, v := range v.(*schema.Set).List() {
						privilegeList = append(privilegeList, helper.String(v.(string)))
					}
					privilege.Privileges = privilegeList
				}
				request.TablePrivileges = append(request.TablePrivileges, &privilege)
			}
		}

		if v, ok := d.GetOk("column"); ok {
			for _, item := range v.(*schema.Set).List() {
				dMap := item.(map[string]interface{})
				privilege := cdb.ColumnPrivilege{}
				if v, ok := dMap["database_name"]; ok {
					privilege.Database = helper.String(v.(string))
				}
				if v, ok := dMap["table_name"]; ok {
					privilege.Table = helper.String(v.(string))
				}
				if v, ok := dMap["column_name"]; ok {
					privilege.Column = helper.String(v.(string))
				}
				if v, ok := dMap["privileges"]; ok {
					privilegeList := []*string{}
					for _, v := range v.(*schema.Set).List() {
						privilegeList = append(privilegeList, helper.String(v.(string)))
					}
					privilege.Privileges = privilegeList
				}
				request.ColumnPrivileges = append(request.ColumnPrivileges, &privilege)
			}
		}
	}

	ratelimit.Check(request.GetAction())
	response, err := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseMysqlClient().ModifyAccountPrivileges(request)
	if err != nil {
		return err
	}
	if response.Response == nil || response.Response.AsyncRequestId == nil {
		return errors.New("sdk action ModifyAccountPrivileges return error,miss AsyncRequestId")
	}
	asyncRequestId := *response.Response.AsyncRequestId
	mysqlService := MysqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	err = resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		taskStatus, message, err := mysqlService.DescribeAsyncRequestInfo(ctx, asyncRequestId)
		if err != nil {
			return tccommon.RetryError(err)
		}
		if taskStatus == MYSQL_TASK_STATUS_SUCCESS {
			return nil
		}
		if taskStatus == MYSQL_TASK_STATUS_INITIAL || taskStatus == MYSQL_TASK_STATUS_RUNNING {
			return resource.RetryableError(fmt.Errorf("modify account privilege  task  status is %s", taskStatus))
		}

		if taskStatus == MYSQL_TASK_STATUS_FAILED || taskStatus == MYSQL_TASK_STATUS_REMOVED {
			return resource.NonRetryableError(errors.New("sdk ModifyAccountPrivileges task running fail," + message))
		}
		err = fmt.Errorf("modify account privilege task status is %s,we won't wait for it finish ,it show message:%s\n", taskStatus, message)
		return resource.NonRetryableError(err)
	})
	return err
}

func resourceTencentCloudMysqlPrivilegeCreate(d *schema.ResourceData, meta interface{}) error {

	defer tccommon.LogElapsed("resource.tencentcloud_mysql_privilege.update")()

	var (
		logId       = tccommon.GetLogId(tccommon.ContextNil)
		ctx         = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		mysqlId     = d.Get("mysql_id").(string)
		accountName = d.Get("account_name").(string)
		accountHost = d.Get("account_host").(string)

		privilegeId = ResourceTencentCloudMysqlPrivilegeId{MysqlId: mysqlId,
			AccountName: accountName,
			AccountHost: accountHost}
	)
	privilegeIdStr, err := json.Marshal(privilegeId)
	if err != nil {
		return errors.New("json encode to id fail," + err.Error())
	}
	err = privilegeId.update(ctx, d, meta)
	if err != nil {
		return err
	}
	d.SetId(string(privilegeIdStr))
	return nil
}

func resourceTencentCloudMysqlPrivilegeRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_privilege.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	var privilegeId ResourceTencentCloudMysqlPrivilegeId

	if err := json.Unmarshal([]byte(d.Id()), &privilegeId); err != nil {
		err = fmt.Errorf("Local data[terraform.tfstate] corruption,can not got old account privilege id")
		log.Printf("[CRITAL]%s %s\n ", logId, err.Error())
		return err
	}

	if privilegeId.AccountHost == "" {
		privilegeId.AccountHost = MYSQL_DEFAULT_ACCOUNT_HOST
	}

	mysqlService := MysqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	var accountInfo *cdb.AccountInfo = nil
	var onlineHas = true
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		accountInfos, err := mysqlService.DescribeAccounts(ctx, privilegeId.MysqlId)
		if err != nil {
			if sdkerr, ok := err.(*sdkError.TencentCloudSDKError); ok && sdkerr.GetCode() == "InvalidParameter" &&
				strings.Contains(sdkerr.GetMessage(), "instance not found") {
				d.SetId("")
				onlineHas = false
				return nil
			}
			return tccommon.RetryError(err)
		}

		for _, account := range accountInfos {
			if *account.User == privilegeId.AccountName && *account.Host == privilegeId.AccountHost {
				accountInfo = account
				break
			}
		}

		if accountInfo == nil {
			d.SetId("")
			onlineHas = false
			return nil
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Describe mysql acounts fails, reaseon %s", err.Error())
	}
	if !onlineHas {
		return nil
	}

	request := cdb.NewDescribeAccountPrivilegesRequest()
	request.InstanceId = &privilegeId.MysqlId
	request.User = &privilegeId.AccountName
	request.Host = helper.String(privilegeId.AccountHost)

	var response *cdb.DescribeAccountPrivilegesResponse
	err = resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err = meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseMysqlClient().DescribeAccountPrivileges(request)
		if err != nil {
			if sdkErr, ok := err.(*sdkError.TencentCloudSDKError); ok {
				if sdkErr.Code == MysqlInstanceIdNotFound {
					onlineHas = false
				}
				if sdkErr.Code == "InvalidParameter" && strings.Contains(sdkErr.GetMessage(), "instance not found") {
					onlineHas = false
				}
				if sdkErr.Code == "InternalError.TaskError" && strings.Contains(sdkErr.Message, "User does not exist") {
					onlineHas = false
				}
				if !onlineHas {
					return nil
				}
			}
			return tccommon.RetryError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if !onlineHas {
		return nil
	}

	if response == nil || response.Response == nil {
		return errors.New("sdk DescribeAccountPrivileges return error,miss Response")
	}
	globals := make([]string, 0, len(response.Response.GlobalPrivileges))
	for _, v := range response.Response.GlobalPrivileges {
		globals = append(globals, *v)
	}

	databases := make([]map[string]interface{}, 0, len(response.Response.DatabasePrivileges))
	for _, v := range response.Response.DatabasePrivileges {
		privileges := make([]string, 0, len(v.Privileges))
		for _, p := range v.Privileges {
			privileges = append(privileges, *p)
		}
		m := map[string]interface{}{}
		m["database_name"] = *v.Database
		m["privileges"] = privileges
		databases = append(databases, m)
	}

	tables := make([]map[string]interface{}, 0, len(response.Response.TablePrivileges))
	for _, v := range response.Response.TablePrivileges {
		privileges := make([]string, 0, len(v.Privileges))
		for _, p := range v.Privileges {
			privileges = append(privileges, *p)
		}
		m := map[string]interface{}{}
		m["database_name"] = *v.Database
		m["table_name"] = *v.Table
		m["privileges"] = privileges
		tables = append(tables, m)
	}

	columns := make([]map[string]interface{}, 0, len(response.Response.ColumnPrivileges))

	for _, v := range response.Response.ColumnPrivileges {
		privileges := make([]string, 0, len(v.Privileges))
		for _, p := range v.Privileges {
			privileges = append(privileges, *p)
		}
		m := map[string]interface{}{}
		m["database_name"] = *v.Database
		m["table_name"] = *v.Table
		m["column_name"] = *v.Column
		m["privileges"] = privileges
		columns = append(columns, m)
	}
	_ = d.Set("global", globals)
	_ = d.Set("database", databases)
	_ = d.Set("table", tables)
	_ = d.Set("column", columns)
	_ = d.Set("mysql_id", privilegeId.MysqlId)
	_ = d.Set("account_name", privilegeId.AccountName)
	_ = d.Set("account_host", privilegeId.AccountHost)

	return nil
}

func resourceTencentCloudMysqlPrivilegeUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_privilege.update")()

	var (
		logId       = tccommon.GetLogId(tccommon.ContextNil)
		ctx         = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		privilegeId = ResourceTencentCloudMysqlPrivilegeId{}
	)
	if err := json.Unmarshal([]byte(d.Id()), &privilegeId); err != nil {
		err = fmt.Errorf("Local data[terraform.tfstate] corruption,can not got old account privilege id")
		log.Printf("[CRITAL]%s %s\n ", logId, err.Error())
		return err
	}
	if privilegeId.AccountHost == "" {
		privilegeId.AccountHost = MYSQL_DEFAULT_ACCOUNT_HOST
	}

	if d.HasChange("global") || d.HasChange("database") || d.HasChange("table") || d.HasChange("column") {
		err := privilegeId.update(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceTencentCloudMysqlPrivilegeDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_privilege.delete")()

	var (
		logId       = tccommon.GetLogId(tccommon.ContextNil)
		ctx         = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		privilegeId = ResourceTencentCloudMysqlPrivilegeId{}
	)
	if err := json.Unmarshal([]byte(d.Id()), &privilegeId); err != nil {
		err = fmt.Errorf("Local data[terraform.tfstate] corruption,can not got old account privilege id")
		log.Printf("[CRITAL]%s %s\n ", logId, err.Error())
		return err
	}
	if privilegeId.AccountHost == "" {
		privilegeId.AccountHost = MYSQL_DEFAULT_ACCOUNT_HOST
	}
	err := privilegeId.update(ctx, nil, meta)
	if err != nil {
		return err
	}
	return nil
}
