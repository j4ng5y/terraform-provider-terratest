package s3

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func ResourceTerratestAWSS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceTerratestAWSS3BucketCreate,
		Read:   resourceTerratestAWSS3BucketRead,
		Update: resourceTerratestAWSS3BucketUpdate,
		Delete: resourceTerratestAWSS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTerratestAWSS3BucketCreate(d *schema.ResourceData, m interface{}) error {
	return resourceTerratestAWSS3BucketRead(d, m)
}

func resourceTerratestAWSS3BucketRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTerratestAWSS3BucketUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTerratestAWSS3BucketRead(d, m)
}

func resourceTerratestAWSS3BucketDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
