package main

import (
	"flag"
	ami "github.com/cultureamp/aws-nquire/command/ami"
	cfn "github.com/cultureamp/aws-nquire/command/cfn"
	log "github.com/cultureamp/aws-nquire/logging"
	"os"
)

func main() {
	prefix := flag.String("ami-prefix", "", "specify a ami prefix to search")
	stack := flag.String("stack", "", "specify stack name to search")
	flag.Parse()

	//hard code region for now
	region := "us-west-2"

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	if len(*prefix) > 0 {
		log.Debug("ami-prefix=" + *prefix)
		ami.Run(*prefix, region)
	}

	args := flag.Args()

	if len(*stack) > 0 {
		log.Debug("stack=" + *stack)
		cfn.Run(*stack, args[0], region)
	}
}
