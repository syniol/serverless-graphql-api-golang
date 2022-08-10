FROM node:16

RUN apt-get update -y && apt-get install -y \
    git \
    zip \
    && curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && aws --version \
    && npm install -g aws-cdk

RUN mkdir -p /var/local/cdk

VOLUME /var/local/cdk

WORKDIR /var/local/cdk
