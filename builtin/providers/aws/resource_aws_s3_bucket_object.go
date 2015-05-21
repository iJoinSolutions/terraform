package aws

import (
	"os"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
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
	s3conn := meta.(*AWSClient).s3conn

	bucket := d.Get("bucket").(string)
	key    := d.Get("key").(string)
	source := d.Get("source").(string)

	file, err := os.Open(source)

	if err != nil {
		return fmt.Errorf("Error opening S3 bucket object source(%s): %s", source, err)
	}

	resp, err := s3conn.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   file,
		})

	if err != nil {
		return fmt.Errorf("Error putting object in S3 bucket (%s): %s", bucket, err)
	}

	log.Printf(awsutil.StringValue(resp))
	d.SetId(*resp.ETag)
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
