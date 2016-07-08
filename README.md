# aws-nquire
Tools to interact with AWS APIs to discover resources

How to use this tool

Query for the name of latest ami with prefix 'ca-docker'
```
aws-nquire ami imageid --prefix ca-docker --branch master
```

Query for the vpc id in a stack
```
aws-nquire stack VpcId --name $stack_name
```

To download and build:
go get github.com/cultureamp/aws-nquire


Copyright © 2016 Culture Amp Pty Ltd, all rights reserved.
