.PHONY: deploy
.deploy:
	docker build -t cdk:node . -f deploy/cdk/Dockerfile --build-arg AWS_KEY=opwhb2232 --build-arg AWS_SECRET=7862ghhhgy3
	docker run -v (pwd):/var/local/pipeline --rm -it cdk:node bash
