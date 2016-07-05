package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
	"strings"

	log "github.com/cultureamp/aws-nquire/logging"
)

func Run(prefix string, field string, region string, key string, value string) string {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
	rsp := queryByAccount(svc, prefix)
	images := filter(rsp.Images, key, value)
	//	for _, v := range rsp.Images {
	//		filterByBranchTag(v, "branch", "setup-ami-baking-on-ci")
	//	}

	//	amis := filter(rsp.Images, func(i *ec2.Image) bool {
	//		return strings.Contains(*i.Name, prefix)
	//	})

	if len(images) == 0 {
		log.Error(fmt.Sprintf("Expecting 1 image, but found: %d", len(images)))
		os.Exit(1)
	}

	id, name := latest(images)
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

func filterByBranchTag(i *ec2.Image, key string, value string) bool {
	for _, tag := range i.Tags {
		if strings.EqualFold(*tag.Key, key) && strings.EqualFold(*tag.Value, value) {
			fmt.Println("found one " + *i.Name)
			return true
		}
	}
	return false
}

func filter(imgs []*ec2.Image, key string, value string) []*ec2.Image {
	var amis []*ec2.Image
	for _, img := range imgs {
		if filterByBranchTag(img, key, value) {
			amis = append(amis, img)
		}
	}
	return amis
}

func queryByAccount(svc *ec2.EC2, prefix string) *ec2.DescribeImagesOutput {
	inputs := params2(prefix)
	resp, err := svc.DescribeImages(inputs)
	if err != nil {
		log.Error("Error in describing images")
		os.Exit(1)
	}
	return resp
}

//func params() *ec2.DescribeImagesInput {
//	return &ec2.DescribeImagesInput{
//		Owners: []*string{
//			aws.String("self"),
//		},
//	}
//}

func params2(prefix string) *ec2.DescribeImagesInput {
	nameRegex := prefix + "*"
	return &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String("self"),
		},
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("tag-key"),
				Values: []*string{aws.String("Name")},
			},
			&ec2.Filter{
				Name:   aws.String("tag-value"),
				Values: []*string{aws.String(nameRegex)},
			},
		},
	}
}

func latest(imgs []*ec2.Image) (string, string) {
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
