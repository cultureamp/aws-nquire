package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
	"regexp"
	"strings"

	log "github.com/cultureamp/aws-nquire/logging"
)

func Run(prefix string, field string, region string, account string, key string, value string) string {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))

	if len(account) == 0 {
		log.Debug("set account to default 'self'")
		account = "self"
	}

	rsp := query(svc, account)
	log.Debug(fmt.Sprintf("total # of imgs retrieved: %d", len(rsp.Images)))

	images := filter(rsp.Images, prefix, key, value)

	validateResult(images)
	id, name := getLatest(images)
	fieldInLower := strings.ToLower(field)

	switch fieldInLower {
	case "name":
		return name
	case "imageid":
		return id
	default:
		log.Error("Invalid argument (expected: ImageId or Name), found: " + field)
		os.Exit(1)
	}
	return ""
}

func dump(obj interface{}) {
	fmt.Printf("%+v\n", obj)
}

func validateResult(images []*ec2.Image) {
	if len(images) == 0 {
		log.Error("Unable to find any image")
		os.Exit(1)
	}
}

// return true if both tag key and value match
func matchByTag(i *ec2.Image, key string, value string) bool {
	for _, tag := range i.Tags {
		if strings.EqualFold(*tag.Key, key) && matchByRegex(value, *tag.Value) {
			dump(i)
			return true
		}
	}
	return false
}

func matchByRegex(pattern string, value string) bool {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		log.Error("Error matching value by regex")
		os.Exit(1)
	}
	log.Debug(fmt.Sprintf("regex pattern: %s, value %s, result %v", pattern, value, matched))
	return matched
}

func filter(imgs []*ec2.Image, prefix string, branchTagName string, branchTagValue string) []*ec2.Image {
	nameTagKey := "Name"
	nameTagValueRegex := prefix + "*"
	var amis []*ec2.Image
	for _, img := range imgs {
		if matchByTag(img, branchTagName, branchTagValue) {
			if matchByTag(img, nameTagKey, nameTagValueRegex) {
				amis = append(amis, img)
			}
		}
	}
	return amis
}

func query(svc *ec2.EC2, account string) *ec2.DescribeImagesOutput {
	inputs := params(account)
	resp, err := svc.DescribeImages(inputs)
	if err != nil {
		log.Error("Error in describing images")
		os.Exit(1)
	}
	return resp
}

func params(account string) *ec2.DescribeImagesInput {
	return &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}
}

func getLatest(imgs []*ec2.Image) (string, string) {
	name := *imgs[0].Name
	id := *imgs[0].ImageId
	for _, img := range imgs {
		if strings.Compare(name, *img.Name) < 0 {
			name = *img.Name
			id = *img.ImageId
		}
	}
	return id, name
}
