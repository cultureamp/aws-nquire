package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"strings"

	"cultureamp/aws-nquire/logging"
)

func Run(prefix string) {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
	rsp := queryByAccount(svc, prefix)
	amis := Filter(rsp.Images, func(i *ec2.Image) bool {
		return strings.Contains(*i.Name, prefix)
	})
	fmt.Println(len(amis))
	fmt.Println(latest(amis))
}

func queryByAccount(svc *ec2.EC2, prefix string) *ec2.DescribeImagesOutput {
	inputs := params()
	resp, err := svc.DescribeImages(inputs)
	if err != nil {
		logging.Info("Error in describing images")
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

func Filter(imgs []*ec2.Image, f func(*ec2.Image) bool) []*ec2.Image {
	var amis []*ec2.Image
	for _, img := range imgs {
		if f(img) {
			amis = append(amis, img)
		}
	}
	return amis
}

func latest(imgs []*ec2.Image) string {
	name := *imgs[0].Name
	id := *imgs[0].ImageId
	for _, img := range imgs {
		if strings.Compare(name, *img.Name) < 1 {
			name = *img.Name
			id = *img.ImageId
		}
	}
	return id
}
