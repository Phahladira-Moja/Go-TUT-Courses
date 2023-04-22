Project Breakdown:
    - In this project we will be building a bulletproof CRUD API using GO + DynamoDB + CHI Router.
    - The project will make use of:
        1.  Defined error messages
        2.  Using Interfaces
        3.  Use logger
        4.  Handling CORS  
        5.  Health Check
        6.  Recoverer middleware in CHI
        7.  Ozzo Validation

GO Commands:
    - go mod init githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api
    - go get github.com/go-chi/chi
    - go get github.com/go-chi/cors
    - go get github.com/aws/aws-sdk-go/aws/session
    - go get github.com/aws/aws-sdk-go/service/dynamodb