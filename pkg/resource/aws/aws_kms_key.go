// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/helpers"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/zclconf/go-cty/cty"
)

const AwsKmsKeyResourceType = "aws_kms_key"

type AwsKmsKey struct {
	Arn                   *string           `cty:"arn" computed:"true"`
	CustomerMasterKeySpec *string           `cty:"customer_master_key_spec"`
	DeletionWindowInDays  *int              `cty:"deletion_window_in_days" diff:"-"`
	Description           *string           `cty:"description" computed:"true"`
	EnableKeyRotation     *bool             `cty:"enable_key_rotation"`
	Id                    string            `cty:"id" computed:"true"`
	IsEnabled             *bool             `cty:"is_enabled"`
	KeyId                 *string           `cty:"key_id" computed:"true"`
	KeyUsage              *string           `cty:"key_usage"`
	Policy                *string           `cty:"policy" jsonstring:"true" computed:"true"`
	Tags                  map[string]string `cty:"tags"`
	CtyVal                *cty.Value        `diff:"-"`
}

func (r *AwsKmsKey) TerraformId() string {
	return r.Id
}

func (r *AwsKmsKey) TerraformType() string {
	return AwsKmsKeyResourceType
}

func (r *AwsKmsKey) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsKmsKeyMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.UpdateSchema(AwsKmsKeyResourceType, map[string]func(attributeSchema *resource.AttributeSchema){
		"policy": func(attributeSchema *resource.AttributeSchema) {
			attributeSchema.JsonString = true
		},
	})
	resourceSchemaRepository.SetNormalizeFunc(AwsKmsKeyResourceType, func(val *resource.Attributes) {
		val.SafeDelete([]string{"deletion_window_in_days"})
		jsonString, err := helpers.NormalizeJsonString((*val)["policy"])
		if err != nil {
			return
		}
		val.SafeSet([]string{"policy"}, jsonString)
	})
}
