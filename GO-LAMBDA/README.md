NB - Project Brief:
    - This project is simply used to show how to create a basic lambda function using go
    - Then we deploy the lambda function to AWS

NB - Packages:
    - We will be using the AWS Go-Lambda Library so we need to install it using the following command:
    - go get "github.com/aws/aws-lambda-go/lambda"

AWS SET-UP:
    1. Make sure you have the AWS CLI configured on your local device.
    2. Once configured in your working directory, add the trust-policy.json file (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-awscli.html)
    3. NB the link provided will guide you in creating an execution role for the lambda function as well as help you attach a policy-to-role. The policy will give your lambda permission for basic execution role.
    4. Command for Execution role (aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole)

NB - This is the command to deploy your lambda to AWS
    - aws lambda create-function --function-name go-lambda-function --zip-file fileb://function.zip --handler main --runtime go1.x --role arn:aws:iam::803846098123:role/lambda-ex

To Invoke The Lambda Function:
    - aws lambda invoke --function-name go-lambda-function --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "Jim", "How old are you?": 33}' output.txt

NB - When building a Go executable be sure to set the following environment variables: (This applies when you're not building it on a linux system)
    - $env:GOOS = "linux"
    - $env:GOARCH = "amd64"
    - $env:CGO_ENABLED = "0"
