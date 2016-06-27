package command

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
	"strings"

	log "github.com/cultureamp/aws-nquire/logging"
)

func Run(prefix string, field string, region string) string {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
	rsp := queryByAccount(svc, prefix)
	amis := filter(rsp.Images, func(i *ec2.Image) bool {
		return strings.Contains(*i.Name, prefix)
	})
	if len(amis) == 0 {
		log.Error("Unable to find ami by prefix: " + prefix)
		os.Exit(1)
	}
	id, name := latest(amis)
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

func queryByAccount(svc *ec2.EC2, prefix string) *ec2.DescribeImagesOutput {
	inputs := params()
	resp, err := svc.DescribeImages(inputs)
	if err != nil {
		log.Info("Error in describing images")
		panic(err)
	}
	return resp
}

func params() *ec2.DescribeImagesInput {
	return &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String("self"),
		},
	}
}

func filter(imgs []*ec2.Image, f func(*ec2.Image) bool) []*ec2.Image {
	var amis []*ec2.Image
	for _, img := range imgs {
		if f(img) {
			amis = append(amis, img)
		}
	}
	return amis
}

func latest(imgs []*ec2.Image) (string, string) {
	name := *imgs[0].Name
	id := *imgs[0].ImageId
	for _, img := range imgs {
		if strings.Compare(name, *img.Name) < 1 {
			name = *img.Name
			id = *img.ImageId
		}
	}
	return id, name
}
