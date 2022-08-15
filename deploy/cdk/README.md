# IaC with AWS CDK in Golang
Deployment of described architecture in the main documentation is created using AWS CDK with `Golang` as a language.
They are grouped under two main packages:

 * __Stack__
 * __Construct__ 

In this example I hardcoded props for each of resources; However, this could be dynamic with a use of templating 
processing and JSON/YAML key matching definition inside Props struct.

**Example of Struct Definition for Step Functions Props**:

```go
type StepFunctionsRestApiProps struct {
    DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
    DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
    DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
    CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
    ...
}
```

**Example of JSON for  Step Functions Props Struct Definition**:

```json
{
  "defaultCorsPreflightOptions": {
    "allowCredentials": true
  },
  "defaultIntegration": {
    "type": "AWS",
    "deploymentToken": "Zhu3hHs72267Vz.."
  }
}
```
**Example of YAML for  Step Functions Props Struct Definition**:

```yaml
defaultCorsPreflightOptions:
  allowCredentials": true
defaultIntegration":
  type: "AWS"
  deploymentToken: "Zhu3hHs72267Vz.."
```


### Useful Dev Setup Commands

 * `go mod download`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile`
 * `docker run -v (pwd):/var/local/pipeline --rm -it cdk:node bash`
 * `docker build -t cdk:node . -f deploy/cdk/Dockerfile --build-arg AWS_KEY=opwhb2232 --build-arg AWS_SECRET=7862ghhhgy3`


### Useful CDK Commands

 * `cdk bootstrap`   bootstrap your default AWS account/region for cdk use
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests


#### Useful Docs
 * [AWS CDK Reference Documentation](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-construct-library.html)
 * [AWS CLI Reference Documentation](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/index.html)


#### Credits
Copyright &copy; 2022 Syniol Limited. All Rights Reserved.