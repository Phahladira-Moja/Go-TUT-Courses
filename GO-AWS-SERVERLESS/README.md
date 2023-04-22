Project Brief: - This will be a serverless application - The serverless application will make use of AWS Infrastructure - We will be using AWS-Lambda, API-Gateway and DynamoDB - The Lambda function will be written in Go, which we will use to query DynamoDB - We will use API-Gateway to make our Lambda function publicly accessible.

Project Commands: 1. go mod init <project-name> (go mod init githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api) 2. go get github.com/aws/aws-sdk-go/aws 3. go get github.com/aws/aws-sdk-go/aws/session 4. go get github.com/aws/aws-sdk-go/lambda 5. go get github.com/aws/aws-sdk-go-v2/service/dynamodb 6. go mod tidy

Windows Env Commands: 1. $env:CGO_ENABLED = "0" 2. $env:GOOS = "linux" 3. $env:GOARCH = "amd64"
