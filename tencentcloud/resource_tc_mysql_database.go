/*
新建mysql数据库
// TODO: Add test cases.
*/
package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudMysqlDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTencentCloudMysqlDatabaseCreate,
		ReadContext:   resourceTencentCloudMysqlDatabaseRead,
		DeleteContext: resourceTencentCloudMysqlDatabaseDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"mysql_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Instance ID to which the account belongs.",
			},
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Database name.",
			},
			"character_set": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "utf8mb4",
				Description: "Database character set, default is `utf8mb4`.",
			},
		},
	}
}

func resourceTencentCloudMysqlDatabaseCreate(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.tencentcloud_mysql_database.create")()

	logId := getLogId(tfCtx)
	ctx := context.WithValue(tfCtx, logIdKey, logId)

	mysqlService := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		mysqlId      = d.Get("mysql_id").(string)
		name         = d.Get("name").(string)
		characterSet = d.Get("character_set").(string)
	)

	requestId, err := mysqlService.CreateDatabase(ctx, mysqlId, name, characterSet)
	if err != nil {
		return diag.Errorf("创建 mysql 数据库失败, 错误为: %s, 请求ID：%v", err, requestId)
	}

	resourceId := fmt.Sprintf("%s%s%s", mysqlId, FILED_SP, name)

	d.SetId(resourceId)

	return nil
}

func resourceTencentCloudMysqlDatabaseRead(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.tencentcloud_mysql_database.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(tfCtx)
	ctx := context.WithValue(tfCtx, logIdKey, logId)

	mysqlService := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		mysqlId = d.Get("mysql_id").(string)
		name    = d.Get("name").(string)
	)

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		resp, e := mysqlService.DescribeDatabase(ctx, mysqlId, fmt.Sprintf("^%s$", name))
		if e != nil {
			return retryError(nil)
		}
		if *(resp.TotalCount) == 0 {
			return retryError(nil)
		}
		database := *(resp.DatabaseList)[0]
		setErr := d.Set("character_set", strings.ToLower(*(database.CharacterSet)))
		if setErr != nil {
			return retryError(fmt.Errorf("设置数据库字符集失败，错误为: %w", setErr))
		} else {
			return nil
		}
	})

	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

// 本资源并不会直接删除database，数据库会随实例一起删除。如果保留实例则需要手动删除数据库。
func resourceTencentCloudMysqlDatabaseDelete(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.tencentcloud_mysql_database.delete")()

	d.SetId("")
	return diag.Diagnostics{{Severity: diag.Warning, Summary: "本资源并不会直接删除database，数据库会随实例一起删除。如果保留实例则需要手动删除数据库。"}}
}
