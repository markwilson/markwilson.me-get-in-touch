build:
	go get -t ./...
	GOOS=linux go build -o main

local: build
	sam local start-api

package:
	aws cloudformation package \
		--template-file template.yaml \
		--s3-bucket markwilson.me-get-in-touch \
		--output-template-file packaged-template.yaml

deploy:
	aws cloudformation deploy \
		--template-file packaged-template.yaml \
		--stack-name markwilsonme-get-in-touch \
		--capabilities CAPABILITY_IAM
