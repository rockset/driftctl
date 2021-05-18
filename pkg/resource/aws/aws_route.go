// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/zclconf/go-cty/cty"
)

const AwsRouteResourceType = "aws_route"

type AwsRoute struct {
	DestinationCidrBlock     *string `cty:"destination_cidr_block"`
	DestinationIpv6CidrBlock *string `cty:"destination_ipv6_cidr_block"`
	DestinationPrefixListId  *string `cty:"destination_prefix_list_id" computed:"true"`
	EgressOnlyGatewayId      *string `cty:"egress_only_gateway_id" computed:"true"`
	GatewayId                *string `cty:"gateway_id" computed:"true"`
	Id                       string  `cty:"id" computed:"true"`
	InstanceId               *string `cty:"instance_id" computed:"true"`
	InstanceOwnerId          *string `cty:"instance_owner_id" computed:"true"`
	LocalGatewayId           *string `cty:"local_gateway_id" computed:"true"`
	NatGatewayId             *string `cty:"nat_gateway_id" computed:"true"`
	NetworkInterfaceId       *string `cty:"network_interface_id" computed:"true"`
	Origin                   *string `cty:"origin" computed:"true"`
	RouteTableId             *string `cty:"route_table_id"`
	State                    *string `cty:"state" computed:"true"`
	TransitGatewayId         *string `cty:"transit_gateway_id"`
	VpcEndpointId            *string `cty:"vpc_endpoint_id"`
	VpcPeeringConnectionId   *string `cty:"vpc_peering_connection_id"`
	Timeouts                 *struct {
		Create *string `cty:"create"`
		Delete *string `cty:"delete"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsRoute) TerraformId() string {
	return r.Id
}

func (r *AwsRoute) TerraformType() string {
	return AwsRouteResourceType
}

func (r *AwsRoute) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsRouteMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsRouteResourceType, func(val *resource.Attributes) {
		val.SafeDelete([]string{"timeouts"})

		val.DeleteIfDefault("vpc_endpoint_id")
		val.DeleteIfDefault("local_gateway_id")
		val.DeleteIfDefault("destination_cidr_block")
		val.DeleteIfDefault("destination_ipv6_cidr_block")
		val.DeleteIfDefault("egress_only_gateway_id")
		val.DeleteIfDefault("nat_gateway_id")
		val.DeleteIfDefault("instance_id")
		val.DeleteIfDefault("network_interface_id")
		val.DeleteIfDefault("transit_gateway_id")
		val.DeleteIfDefault("vpc_peering_connection_id")
		val.DeleteIfDefault("destination_prefix_list_id")
		val.DeleteIfDefault("instance_owner_id")
	})
}
