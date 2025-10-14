package vpc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"

	cvmSDK "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"

	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

func ResourceTencentCloudVpcRouteEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudVpcRouteEntryCreate,
		Read:   resourceTencentCloudVpcRouteEntryRead,
		Update: resourceTencentCloudVpcRouteEntryUpdate,
		Delete: resourceTencentCloudVpcRouteEntryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of routing table to which this entry belongs.",
			},
			"destination_cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: tccommon.ValidateCIDRNetworkAddress,
				Description:  "Destination address block.",
			},
			"next_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: tccommon.ValidateAllowedStringValue(ALL_GATE_WAY_TYPES),
				Description:  "Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP`, `LOCAL_GATEWAY`, `INTRANAT` and `USER_CCN`.",
			},
			"next_hub": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of next-hop gateway. Note: when `next_type` is EIP, `next_hub` should be `0`. when `next_type` is NORMAL_CVM, `next_hub` should be instance-id",
			},
			// Name enabled will lead to exist route table diff fail (null -> false cannot diff).
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled, default is `false`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Description of the routing table entry.",
			},
			"route_item_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of route table entry.",
			},
		},
	}
}

func resourceTencentCloudVpcRouteEntryCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_route_table_entry.create")()

	var (
		logId   = tccommon.GetLogId(tccommon.ContextNil)
		ctx     = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		service = VpcService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	)

	var (
		description          = ""
		routeTableId         = ""
		destinationCidrBlock = ""
		nextType             = ""
		nextHub              = ""
		disabled             = false
	)

	if temp, ok := d.GetOk("description"); ok {
		description = temp.(string)
	}

	if temp, ok := d.GetOk("route_table_id"); ok {
		routeTableId = temp.(string)
	}

	if temp, ok := d.GetOk("destination_cidr_block"); ok {
		destinationCidrBlock = temp.(string)
	}

	if temp, ok := d.GetOk("next_type"); ok {
		nextType = temp.(string)
	}

	if temp, ok := d.GetOk("next_hub"); ok {
		nextHub = temp.(string)
	}

	if temp, ok := d.GetOkExists("disabled"); ok {
		disabled = temp.(bool)
	}

	if routeTableId == "" || destinationCidrBlock == "" || nextType == "" || nextHub == "" {
		return fmt.Errorf("some needed fields is empty string")
	}

	if nextType == GATE_WAY_TYPE_EIP && nextHub != "0" {
		return fmt.Errorf("if next_type is %s, next_hub can only be \"0\" ", GATE_WAY_TYPE_EIP)
	}

	// we accept instance or network interface attachment id as next_hub for nextType is NORMAL_CVM
	// network interface attachment id is necessary as any change to the attachment will invalidate route entry
	if nextType == GATE_WAY_TYPE_NORMAL_CVM {
		vpcService := VpcService{
			client: meta.(tccommon.ProviderMeta).GetAPIV3Conn(),
		}
		eniId, _, eniUsed := strings.Cut(nextHub, "+")

		var instances []*cvmSDK.Instance
		var interfaces []*vpc.NetworkInterface
		var errRet error
		err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
			if !eniUsed {
				instances, errRet = vpcService.DescribeCvmInstanceByFilter(ctx, []*string{&nextHub}, nil)
			} else {
				interfaces, errRet = vpcService.DescribeEniById(ctx, []string{eniId})
			}
			if errRet != nil {
				return tccommon.RetryError(errRet, tccommon.InternalError)
			}
			return nil
		})
		if err != nil {
			return err
		}

		if len(instances) == 1 {
			nextHub = *(instances[0].PrivateIpAddresses[0]) // 此处为实例主网卡的第一个ip，当前cdk主网卡只有一个内网ip
		} else if len(interfaces) == 1 {
			for _, ip := range interfaces[0].PrivateIpAddressSet {
				if ip.Primary != nil && *ip.Primary && ip.PrivateIpAddress != nil {
					nextHub = *ip.PrivateIpAddress
					interfaces = nil
					break
				}
			}
			if interfaces != nil {
				return fmt.Errorf("cannot find primary ip address of interface %s", nextHub)
			}
		} else {
			return fmt.Errorf("cannot find exact instance by id %s", nextHub)
		}
	}

	// route cannot disable on create
	entryId, err := service.CreateRoutes(ctx, routeTableId, destinationCidrBlock, nextType, nextHub, description, true)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d.%s", entryId, routeTableId))

	if disabled {
		request := vpc.NewDisableRoutesRequest()
		request.RouteTableId = &routeTableId
		request.RouteIds = []*uint64{helper.Int64Uint64(entryId)}
		err = service.DisableRoutes(ctx, request)
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudVpcRouteEntryRead(d, meta)
}

func resourceTencentCloudVpcRouteEntryRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_route_table_entry.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	var (
		logId   = tccommon.GetLogId(tccommon.ContextNil)
		ctx     = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		service = VpcService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	)

	items := strings.Split(d.Id(), ".")
	if len(items) != 2 {
		return fmt.Errorf("entry id be destroyed, we can not get route table id")
	}

	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		info, has, e := service.DescribeRouteTable(ctx, items[1])
		if e != nil {
			return tccommon.RetryError(e)
		}

		if has == 0 {
			d.SetId("")
			return nil
		}

		if has != 1 {
			e = fmt.Errorf("one routeTable id get %d routeTable infos", has)
			return resource.NonRetryableError(e)
		}

		for _, v := range info.entryInfos {
			if fmt.Sprintf("%d", v.routeEntryId) == items[0] {
				_ = d.Set("description", v.description)
				_ = d.Set("route_table_id", items[1])
				_ = d.Set("destination_cidr_block", v.destinationCidr)
				_ = d.Set("next_type", v.nextType)
				// convert next_hub to instance or network interface attachment id if nextType is NORMAL_CVM
				if v.nextType == GATE_WAY_TYPE_NORMAL_CVM {
					vpcService := VpcService{
						client: meta.(tccommon.ProviderMeta).GetAPIV3Conn(),
					}
					filter := make(map[string]string)
					// filter by vpc-id and private ip, private ip is only unique in vpc
					filter["private-ip-address"] = v.nextBub
					filter["vpc-id"] = info.vpcId

					var instances []*cvmSDK.Instance
					var interfaces []*vpc.NetworkInterface
					var errRet error
					err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
						instances, errRet = vpcService.DescribeCvmInstanceByFilter(ctx, nil, filter)
						if errRet != nil {
							return tccommon.RetryError(errRet, tccommon.InternalError)
						}
						return nil
					})
					if err != nil {
						return resource.NonRetryableError(err)
					}
					err = resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
						interfaces, errRet = vpcService.DescribeEniByFilters(ctx, &info.vpcId, nil, nil, nil, nil, nil, &v.nextBub, nil)
						if errRet != nil {
							return tccommon.RetryError(errRet, tccommon.InternalError)
						}
						return nil
					})
					if err != nil {
						return resource.NonRetryableError(err)
					}

					if len(instances) == 1 {
						_ = d.Set("next_hub", *(instances[0].InstanceId))
					} else if len(interfaces) == 1 {
						cvmId := ""
						if interfaces[0].Attachment != nil && interfaces[0].Attachment.InstanceId != nil {
							cvmId = *(interfaces[0].Attachment.InstanceId)
						}
						_ = d.Set("next_hub", *(interfaces[0].NetworkInterfaceId)+"+"+cvmId)
					} else {
						return resource.NonRetryableError(fmt.Errorf("cannot find exact instance by ip %s in vpc %s", v.nextBub, info.vpcId))
					}
				} else {
					_ = d.Set("next_hub", v.nextBub)
				}
				_ = d.Set("route_item_id", v.routeItemId)
				_ = d.Set("disabled", !v.enabled)
				return nil
			}
		}

		d.SetId("")
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudVpcRouteEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_route_table_entry.update")()

	var (
		logId   = tccommon.GetLogId(tccommon.ContextNil)
		ctx     = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		client  = meta.(tccommon.ProviderMeta).GetAPIV3Conn()
		service = VpcService{client}
	)

	items := strings.Split(d.Id(), ".")
	if len(items) != 2 {
		return fmt.Errorf("entry id be destroyed, we can not get route table id")
	}

	id := items[0]
	routeTableId := items[1]
	routeEntryId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("parse route entry id %s fail: %s", id, routeTableId)
	}

	if d.HasChange("disabled") {
		disabled := d.Get("disabled").(bool)
		if err := service.SwitchRouteEnabled(ctx, routeTableId, routeEntryId, !disabled); err != nil {
			return err
		}
	}

	return resourceTencentCloudVpcRouteEntryRead(d, meta)
}

func resourceTencentCloudVpcRouteEntryDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_route_table_entry.delete")()

	var (
		logId   = tccommon.GetLogId(tccommon.ContextNil)
		ctx     = context.WithValue(context.TODO(), tccommon.LogIdKey, logId)
		service = VpcService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	)

	items := strings.Split(d.Id(), ".")
	if len(items) != 2 {
		return fmt.Errorf("entry id be destroyed, we can not get route table id")
	}

	routeTableId := items[1]
	entryId, err := strconv.ParseUint(items[0], 10, 64)
	if err != nil {
		return fmt.Errorf("entry id be destroyed, we can not get route entry id")
	}

	err = resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteRoutes(ctx, routeTableId, entryId); err != nil {
			if sdkErr, ok := err.(*errors.TencentCloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}

			return resource.RetryableError(err)
		}

		return nil
	})

	return err
}
