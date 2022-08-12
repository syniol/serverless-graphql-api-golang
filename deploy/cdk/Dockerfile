FROM node:16

ARG AWS_KEY
ARG AWS_SECRET
ARG CDK_DEFAULT_ACCOUNT
ARG CDK_DEFAULT_REGION=eu-west-2

RUN apt-get update -y  \
    && apt-get install -y \
    git \
    zip

# Installing AWS & CDK CLI
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && rm awscliv2.zip \
    && ./aws/install \
    && aws --version \
    && npm install -g aws-cdk

# Intsalling Golang in Node Docker Image and adding it to PATH
RUN curl "https://go.dev/dl/go1.19.linux-amd64.tar.gz" -L -o "go1.19.linux-amd64.tar.gz" \
    && tar -zxvf go1.19.linux-amd64.tar.gz \
    && chmod +x go/bin/go \
    && mv go/bin/go /bin/go
#    && rm -rf go

ENV GOROOT=/go
ENV GO111MODULE=on

RUN mkdir -p /var/local/pipeline

VOLUME /var/local/pipeline

RUN mkdir -p ~/.aws

RUN echo "[default]\n" \
    "aws_access_key_id = $AWS_KEY\n" \
    "aws_secret_access_key = $AWS_SECRET\n" \
    > ~/.aws/credentials

RUN echo "[default]\n" \
    "region = $CDK_DEFAULT_REGION\n" \
    "output = json\n" \
    > ~/.aws/config

WORKDIR /var/local/pipeline

COPY . /var/local/pipeline

WORKDIR /var/local/pipeline/deploy/cdk

# Deploying all Stacks using AWS CDK
RUN go mod vendor \
    && cdk doctor \
    && cdk bootstrap \
    && cdk deploy --all --require-approval never
