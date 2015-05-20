package aws

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
)

func resourceAwsS3BucketObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsS3BucketObjectCreate,
		Read: resourceAwsS3BucketObjectRead,
		Update: resourceAwsS3BucketObjectUpdate,
		Delete: resourceAwsS3BucketObjectDelete,

		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

		},
	}
}

func resourceAwsS3BucketObjectCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsS3BucketObjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsS3BucketObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsS3BucketObjectDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
