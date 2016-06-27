package main

import (
	"flag"
	"fmt"
	"os"

	ami "github.com/cultureamp/aws-nquire/command/ami"
	cfn "github.com/cultureamp/aws-nquire/command/cfn"
	log "github.com/cultureamp/aws-nquire/logging"
)

func main() {
	prefix := flag.String("ami-prefix", "", "specify a ami prefix to search")
	stack := flag.String("stack", "", "specify stack name to search")
	flag.Parse()

	//hard code region for now
	region := "us-west-2"

	//	if len(os.Args) < 2 {
	//		flag.Usage()
	//		os.Exit(1)
	//	}

	var id string

	if len(*prefix) > 0 {
		log.Debug("ami-prefix=" + *prefix)
		id = ami.Run(*prefix, region)
	} else if len(*stack) > 0 {
		args := flag.Args()
		log.Debug("stack=" + *stack)
		id = cfn.Run(*stack, args[0], region)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(id)
}
