package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudMysqlInstanceCharsetDataSource_basic -v
func TestAccTencentCloudMysqlInstanceCharsetDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMysqlInstanceCharsetDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_mysql_instance_charset.instance_charset"),
					resource.TestCheckResourceAttr("data.tencentcloud_mysql_instance_charset.instance_charset", "charset", "UTF8"),
				),
			},
		},
	})
}

const testAccMysqlInstanceCharsetDataSourceVar = `
variable "instance_id" {
	default = "` + defaultDbBrainInstanceId + `"
}
`

const testAccMysqlInstanceCharsetDataSource = testAccMysqlInstanceCharsetDataSourceVar + `

data "tencentcloud_mysql_instance_charset" "instance_charset" {
	instance_id = var.instance_id
}

`
