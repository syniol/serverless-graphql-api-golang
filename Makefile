.PHONY: deploy
deploy:
	docker build -t cdk:node . -f deploy/cdk/Dockerfile \
	--build-arg AWS_KEY=${AWS_KEY} \
	--build-arg AWS_SECRET=${AWS_SECRET} \
	--build-arg CDK_DEFAULT_ACCOUNT=${CDK_DEFAULT_ACCOUNT} \
	--build-arg CDK_DEFAULT_REGION=${CDK_DEFAULT_REGION}


.PHONY: bash
bash:
	#docker run -v $(pwd):/var/local/pipeline --rm -it cdk:node bash
	docker run --rm -it cdk:node bash
