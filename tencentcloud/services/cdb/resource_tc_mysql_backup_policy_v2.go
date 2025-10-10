/*
Provides a mysql policy resource to create a backup policy.

~> **NOTE:** This attribute `backup_model` only support 'physical' in Terraform TencentCloud provider version 1.16.2

# Example Usage

```hcl

	resource "tencentcloud_mysql_backup_policy" "default" {
	  mysql_id         = "cdb-dnqksd9f"
	  retention_period = 7
	  backup_model     = "physical"
	  backup_time      = "02:00-06:00"
	}

```
*/
package cdb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTencentCloudMysqlBackupPolicyV2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTencentCloudMysqlBackupPolicyV2Create,
		ReadContext:   resourceTencentCloudMysqlBackupPolicyV2Read,
		UpdateContext: resourceTencentCloudMysqlBackupPolicyV2Update,
		DeleteContext: resourceTencentCloudMysqlBackupPolicyV2Delete,

		Schema: map[string]*schema.Schema{
			"mysql_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Instance ID to which policies will be applied.",
			},
			"retention_period": {
				Type:         schema.TypeInt,
				ValidateFunc: tccommon.ValidateIntegerInRange(7, 1830),
				Optional:     true,
				Default:      7,
				Description:  "The retention time of backup files, in days. The minimum value is 7 days and the maximum value is 1830 days. And default value is `7`.",
			},
			"backup_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "physical",
				ValidateFunc: tccommon.ValidateAllowedStringValue([]string{"physical"}),
				Description:  "Backup method. Supported values include: `physical` - physical backup.",
			},
			"binlog_period": {
				Type:         schema.TypeInt,
				ValidateFunc: tccommon.ValidateIntegerInRange(7, 1830),
				Optional:     true,
				Default:      7,
				Description:  "Binlog retention days. Valid value ranges: [7~1830]. And default value is `7`.",
			},
			"backup_window": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Backup window, in the format of {'Monday':'HH:mm-HH:mm'}. See https://cloud.tencent.com/document/api/236/15878#CommonTimeWindow for more details.",
			},
			"enable_backup_period_save": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				ValidateFunc: tccommon.ValidateAllowedStringValue([]string{
					"off",
					"on",
				}),
				Description: "定期保留开关，off - 不开启定期保留策略，on - 开启定期保留策略，默认为off。首次开启定期保留策略时，BackupPeriodSaveDays，BackupPeriodSaveInterval，BackupPeriodSaveCount，StartBackupPeriodSaveDate参数为必填项，否则定期保留策略不会生效",
			},
			"backup_period_save_days": {
				Type:         schema.TypeInt,
				ValidateFunc: tccommon.ValidateIntegerInRange(90, 3650),
				Optional:     true,
				Default:      400,
				Description:  "Backup period save days. Valid value ranges: [90~3650]. And default value is `400`.",
			},
			"backup_period_save_interval": {
				Type:         schema.TypeString,
				ValidateFunc: tccommon.ValidateAllowedStringValue([]string{"weekly", "monthly", "quarterly", "yearly"}),
				Optional:     true,
				Default:      "monthly",
				Description:  "Backup period save interval. Supported values include: `weekly`, `monthly`, `quarterly`, `yearly`. And default value is `monthly`.",
			},
			"backup_period_save_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "定期保留的备份数量，最小值为1，最大值不超过定期保留策略周期内常规备份个数。默认为1.",
			},
			"start_backup_period_save_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "定期保留策略周期起始日期，格式：YYYY-MM-dd HH:mm:ss",
				Deprecated:  "腾讯云已废弃此参数",
			},
			"enable_backup_standby": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				ValidateFunc: tccommon.ValidateAllowedStringValue([]string{
					"off",
					"on",
				}),
				Description: "是否开启数据备份标准存储策略。默认为off，表示关闭。",
			},
			"backup_standby_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "数据备份标准存储起始天数，数据备份达到标准存储起始天数时进行转换，最小为30天，不得大于数据备份保留天数。如果开启备份归档，不得大于等于备份归档天数",
			},
			"enable_binlog_standby": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				ValidateFunc: tccommon.ValidateAllowedStringValue([]string{
					"off",
					"on",
				}),
				Description: "是否开启日志备份标准存储策略，off-关闭，on-打开，默认为off",
			},
			"binlog_standby_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "日志备份标准存储起始天数，日志备份达到标准存储起始天数时进行转换，最小为30天，不得大于日志备份保留天数。如果开启备份归档，不得大于等于备份归档天数",
			},
		},
	}
}

func resourceTencentCloudMysqlBackupPolicyV2Create(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_backup_policy_v2.create")()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(tfCtx, tccommon.LogIdKey, logId)
	d.SetId(d.Get("mysql_id").(string))

	return resourceTencentCloudMysqlBackupPolicyV2Update(ctx, d, meta)
}

func resourceTencentCloudMysqlBackupPolicyV2Read(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_backup_policy_v2.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	mysqlService := MysqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	err := resource.RetryContext(ctx, tccommon.ReadRetryTimeout, func() *resource.RetryError {
		desResponse, e := mysqlService.DescribeBackupConfigByMysqlId(ctx, d.Id())
		if e != nil {
			if mysqlService.NotFoundMysqlInstance(e) {
				d.SetId("")
				return nil
			}
			return tccommon.RetryError(e)
		}
		_ = d.Set("mysql_id", d.Id())
		_ = d.Set("retention_period", int(*desResponse.Response.BackupExpireDays))
		_ = d.Set("backup_method", *desResponse.Response.BackupMethod)
		_ = d.Set("binlog_period", int(*desResponse.Response.BinlogExpireDays))
		_ = d.Set("backup_window", map[string]string{
			"Monday":    parseTimeWindow(desResponse.Response.BackupTimeWindow.Monday),
			"Tuesday":   parseTimeWindow(desResponse.Response.BackupTimeWindow.Tuesday),
			"Wednesday": parseTimeWindow(desResponse.Response.BackupTimeWindow.Wednesday),
			"Thursday":  parseTimeWindow(desResponse.Response.BackupTimeWindow.Thursday),
			"Friday":    parseTimeWindow(desResponse.Response.BackupTimeWindow.Friday),
			"Saturday":  parseTimeWindow(desResponse.Response.BackupTimeWindow.Saturday),
			"Sunday":    parseTimeWindow(desResponse.Response.BackupTimeWindow.Sunday),
		})
		_ = d.Set("enable_backup_period_save", *desResponse.Response.EnableBackupPeriodSave)
		_ = d.Set("backup_period_save_days", int(*desResponse.Response.BackupPeriodSaveDays))
		_ = d.Set("backup_period_save_interval", *desResponse.Response.BackupPeriodSaveInterval)
		_ = d.Set("backup_period_save_count", int(*desResponse.Response.BackupPeriodSaveCount))
		_ = d.Set("enable_backup_standby", *desResponse.Response.EnableBackupStandby)
		_ = d.Set("backup_standby_days", int(*desResponse.Response.BackupStandbyDays))
		_ = d.Set("enable_binlog_standby", *desResponse.Response.EnableBinlogStandby)
		_ = d.Set("binlog_standby_days", int(*desResponse.Response.BinlogStandbyDays))
		return nil
	})
	if err != nil {
		return diag.Errorf("[API]Describe mysql backup policy fail,reason:%s", err.Error())
	}
	return nil
}

func resourceTencentCloudMysqlBackupPolicyV2Update(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_backup_policy_v2.update")()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(tfCtx, tccommon.LogIdKey, logId)

	mysqlService := MysqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	var (
		mysqlId                  = d.Get("mysql_id").(string)
		retentionPeriod          = int64(d.Get("retention_period").(int))
		backupMethod             = d.Get("backup_method").(string)
		binlogExpireDays         = int64(d.Get("binlog_period").(int))
		backupWindow             = d.Get("backup_window").(map[string]interface{})
		enableBackupPeriodSave   = d.Get("enable_backup_period_save").(string)
		backupPeriodSaveDays     = int64(d.Get("backup_period_save_days").(int))
		backupPeriodSaveInterval = d.Get("backup_period_save_interval").(string)
		backupPeriodSaveCount    = int64(d.Get("backup_period_save_count").(int))
		enableBackupStandby      = d.Get("enable_backup_standby").(string)
		backupStandbyDays        = int64(d.Get("backup_standby_days").(int))
		enableBinlogStandby      = d.Get("enable_binlog_standby").(string)
		binlogStandbyDays        = int64(d.Get("binlog_standby_days").(int))
	)

	if backupMethod != "physical" {
		return diag.Errorf("`backup_model` only support 'physical'")
	}
	if d.HasChangeExcept("mysql_id") {
		monday, _ := backupWindow["Monday"].(string)
		tuesday, _ := backupWindow["Tuesday"].(string)
		wednesday, _ := backupWindow["Wednesday"].(string)
		thursday, _ := backupWindow["Thursday"].(string)
		friday, _ := backupWindow["Friday"].(string)
		saturday, _ := backupWindow["Saturday"].(string)
		sunday, _ := backupWindow["Sunday"].(string)
		timeWindow := cdb.CommonTimeWindow{
			Monday:    &monday,
			Tuesday:   &tuesday,
			Wednesday: &wednesday,
			Thursday:  &thursday,
			Friday:    &friday,
			Saturday:  &saturday,
			Sunday:    &sunday,
		}
		err := mysqlService.ModifyBackupConfigByMysqlIdV2(ctx, mysqlId,
			retentionPeriod, backupMethod, binlogExpireDays, timeWindow,
			enableBackupPeriodSave, backupPeriodSaveDays, backupPeriodSaveInterval,
			backupPeriodSaveCount, enableBackupStandby, backupStandbyDays,
			enableBinlogStandby, binlogStandbyDays)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceTencentCloudMysqlBackupPolicyV2Read(ctx, d, meta)
}

// set all config to default
func resourceTencentCloudMysqlBackupPolicyV2Delete(tfCtx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer tccommon.LogElapsed("resource.tencentcloud_mysql_backup_policy.delete")()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(tfCtx, tccommon.LogIdKey, logId)

	mysqlService := MysqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	var (
		retentionPeriod     int64 = 7
		backupModel               = MYSQL_ALLOW_BACKUP_MODEL[1]
		backupTime                = MYSQL_ALLOW_BACKUP_TIME[0]
		binlogExpireDays    int64 = 7
		enableBinlogStandby       = "off"
		binlogStandbyDays   int64 = 180
	)
	err := mysqlService.ModifyBackupConfigByMysqlId(ctx, d.Id(), retentionPeriod, backupModel, backupTime, binlogExpireDays, enableBinlogStandby, binlogStandbyDays)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

// parseTimeWindow 经腾讯云api返回的时间格式转换成terraform默认的字符串空值
// **注意** 空字符串也可以被腾讯云api接受，但是空字符串经过腾讯云处理后从api返回的是"00:00-00:00"
func parseTimeWindow(window *string) string {
	if *window == "00:00-00:00" {
		return ""
	}
	return *window
}
