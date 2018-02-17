build:
	GOOS=linux go build -o main

local: build
	sam local start-api
