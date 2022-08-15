# REST~~ful~~ in Peace and Graph Up your API __(Golang Edition)__
In this presentation I will demonstrate how you could create 
a simple scalable GraphQL API using AWS Lambda and GraphQL 
library with Neptune Graph database. I will also take you 
through deployment process using AWS CDK to deploy our lambda 
function. In order to trigger the lambda we will also need to 
create an AWS API Gateway to trigger the deployed lambda.


## Architecture
todo: add Diagram


## AWS Lambda & API Gateway
todo: add information about trigger point


## AWS CDK
Please follow steps below to deploy this project on AWS. Two main 
prerequisites to deploy this application are:
 * to have an AWS account with Administrator Privileges
 * to have a Docker Desktop installed and running on your machine


### Step I - Build Args _(Docker Image Pipeline to Deploy with CDK)_
Please ensure you have the following environment variables set before 
running a command described in Step II.

* `AWS_KEY` Your AWS Public Key, could be obtained from [_Security Credentials_](https://us-east-1.console.aws.amazon.com/iam/home?region=eu-west-2#/security_credentials) page 
* `AWS_SECRET` This will be your AWS Secret key, could be obtained from [_Security Credentials_](https://us-east-1.console.aws.amazon.com/iam/home?region=eu-west-2#/security_credentials) page
* `CDK_DEFAULT_ACCOUNT` Your AWS Account Number
* `CDK_DEFAULT_REGION` Preferred region to deploy resources


### Step II - Build Docker Image _(Local Pipeline)_
There is `Makefile` located at the root of this repository; please use a command 
below to build a docker image which acts as a pipeline to deploy AWS resources 
using Environment variables set in Step I.

    make deploy


#### Credits
Copyright &copy; 2022 Syniol Limited. All Rights Reserved.
