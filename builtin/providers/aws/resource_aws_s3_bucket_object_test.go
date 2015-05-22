package aws

import (
	"testing"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
)
var tf,err = ioutil.TempFile("", "tf")
var _ = log.Printf
func TestAccAWSS3BucketObject_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { 
			if err != nil {
				panic(err)
			} 
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: testAccCheckAWSS3BucketObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSS3BucketObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSS3BucketObjectExists("aws_s3_bucket_object.object"),
				),
			},
		},
	})
}

func testAccCheckAWSS3BucketObjectDestroy(s *terraform.State) error {
	//s3conn := testAccProvider.Meta().(*AWSClient).s3conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_s3_bucket_object" {
			continue
		}
		
	}
	return nil
}

func testAccCheckAWSS3BucketObjectExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		defer os.Remove(tf.Name());

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No S3 Bucket Object ID is set")
		}

		s3conn := testAccProvider.Meta().(*AWSClient).s3conn
		_, err := s3conn.GetObject(
			&s3.GetObjectInput{
				Bucket:  aws.String(rs.Primary.Attributes["bucket"]),
				Key:     aws.String(rs.Primary.Attributes["key"]),
				IfMatch: aws.String(rs.Primary.ID),
		})
		if err != nil {
			return fmt.Errorf("S3Bucket Object error: %s", err)
		}
		return nil
	}
}

var randomBucket = randInt;
var testAccAWSS3BucketObjectConfig = fmt.Sprintf(`
resource "aws_s3_bucket" "bucket" {
	bucket = "tf-object-test-bucket-%d"
}
resource "aws_s3_bucket_object" "object" {
	depends_on = "aws_s3_bucket.bucket"
	bucket = "tf-object-test-bucket-%d"
	key = "test-key"
	source = "%s"
}
`, randomBucket, randomBucket, tf.Name())

