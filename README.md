# aws-nquire
Tools to interact with AWS APIs to discover resources

How to use this tool

Query for the name of latest ami with prefix 'ca'
aws-nquire -ami-prefix ca Name

Query for the ami id of the latest ami with prefix 'ca'
aws-nquire -ami-prefix ca ImageId

Query for the vpc id in a stack
aws-nquire -stack $STACK_NAME VpcId

To download and build:
go get github.com/cultureamp/aws-nquire


Copyright © 2016 Culture Amp Pty Ltd, all rights reserved.
