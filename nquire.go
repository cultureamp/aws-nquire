package main

import (
	//	"flag"
	//	"fmt"
	//	"os"

	"github.com/cultureamp/aws-nquire/cmd"
	//ami "github.com/cultureamp/aws-nquire/command/ami"
	//cfn "github.com/cultureamp/aws-nquire/command/cfn"
	//log "github.com/cultureamp/aws-nquire/logging"
)

func main() {
	cmd.Execute()
	//	prefix := flag.String("ami-prefix", "", "specify a ami prefix to search")
	//	stack := flag.String("stack", "", "specify stack name to search")
	//	flag.Parse()
	//
	//	//hard code region for now
	//	region := "us-west-2"
	//
	//	//check remainging args
	//	args := flag.Args()
	//	check_args(args)
	//	jsonKey := args[0]
	//
	//	var id string
	//
	//	if len(*prefix) > 0 {
	//		log.Debug("ami-prefix=" + *prefix)
	//		log.Debug("jsonKey=" + jsonKey)
	//		id = ami.Run(*prefix, jsonKey, region)
	//	} else if len(*stack) > 0 {
	//		log.Debug("stack=" + *stack)
	//		id = cfn.Run(*stack, jsonKey, region)
	//	} else {
	//		flag.Usage()
	//		os.Exit(1)
	//	}
	//
	//	fmt.Println(id)
}

//func check_args(args []string) {
//	if len(args) == 0 {
//		log.Info("Insufficient arguments")
//		flag.Usage()
//		os.Exit(1)
//	}
//}
