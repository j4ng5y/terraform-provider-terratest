package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/j4ng5y/terraform-provider-terratest/resources/aws/s3"
)

// Provider is a function that returns a terraform provider instance pointer
//
// Arguments:
//     None
//
// Returns:
//     (*schema.Provider): The pointer to the instance of schema.Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"terratest_aws_s3_bucket": s3.ResourceTerratestAWSS3Bucket(),
		},
	}
}
