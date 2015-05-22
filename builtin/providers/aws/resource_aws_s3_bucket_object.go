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
		Create: resourceAwsS3BucketObjectPut,
		Read: resourceAwsS3BucketObjectRead,
		Update: resourceAwsS3BucketObjectPut,
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

func resourceAwsS3BucketObjectPut(d *schema.ResourceData, meta interface{}) error {
	s3conn := meta.(*AWSClient).s3conn

	bucket := d.Get("bucket").(string)
	key    := d.Get("key").(string)
	source := d.Get("source").(string)

	file, err := os.Open(source)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error opening S3 bucket object source(%s): %s", source, err)
	}

	resp, err := s3conn.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   file,
		})

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error putting object in S3 bucket (%s): %s", bucket, err)
	}

	log.Printf(awsutil.StringValue(resp))
	d.SetId(*resp.ETag)
	return nil
}

func resourceAwsS3BucketObjectRead(d *schema.ResourceData, meta interface{}) error {
	s3conn := meta.(*AWSClient).s3conn

	bucket := d.Get("bucket").(string)
	key    := d.Get("key").(string)

	resp, err := s3conn.GetObject(
		&s3.GetObjectInput{
			Bucket:  aws.String(bucket),
			Key:     aws.String(key),
			// we don't really want to download entire object just see if it is there.
			Range:   aws.String("bytes=0-0"),
			IfMatch: aws.String(d.Id()),
		})

	if err != nil {
		// if there is an error reading the object we assume it's not there.
		d.SetId("")
		log.Printf("Error Reading Object (%s): %s", key, err)
	}

	log.Printf(awsutil.StringValue(resp))
	return nil
}

func resourceAwsS3BucketObjectDelete(d *schema.ResourceData, meta interface{}) error {
	s3conn := meta.(*AWSClient).s3conn

	bucket := d.Get("bucket").(string)
	key    := d.Get("key").(string)

	_, err := s3conn.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		return fmt.Errorf("Error deleting S3 bucket object: %s", err)
	}
	return nil
}
