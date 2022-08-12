# IaC with AWS CDK in Golang
lorem


## Dev Setup Commands

 * `go mod download`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile`
 * `docker run -v (pwd):/var/local/pipeline --rm -it cdk:node bash`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile --build-arg AWS_KEY=opwhb2232 --build-arg AWS_SECRET=7862ghhhgy3`





## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests


### Useful Docs
 * [AWS CDK Reference Documentation](https://docs.aws.amazon.com/cdk/api/v1/)