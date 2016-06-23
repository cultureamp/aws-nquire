package main

import (
	"flag"
	"fmt"
	"os"

	amiCmd "cultureamp/aws-nquire/command/ami"
)

func main() {
	prefix := flag.String("ami-prefix", "", "specify a ami prefix to search")
	stack := flag.String("stack", "", "specify stack name to search")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	if len(*prefix) > 0 {
		fmt.Printf("prefix = %s\n", *prefix)
		amiCmd.Run(*prefix)
	}

	if len(*stack) > 0 {
		fmt.Println("stack = %s\n", *stack)
	}

	fmt.Printf("other args: %+v\n", flag.Args())
}
