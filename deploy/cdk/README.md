# IaC with AWS CDK in Golang


## Dev Setup Commands

 * `go mod download`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile`
 * `docker run -v (pwd):/var/local/pipeline --rm -it cdk:node bash`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile --build-arg AWS_KEY=opwhb2232 --build-arg AWS_SECRET=7862ghhhgy3`


## Build Args (Docker Image Pipeline to Deploy with CDK)
 * `AWS_KEY`
 * `AWS_SECRET`
 * `CDK_DEFAULT_ACCOUNT`
 * `CDK_DEFAULT_REGION`


## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
