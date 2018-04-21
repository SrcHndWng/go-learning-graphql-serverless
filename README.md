# About This

[Building a Golang GraphQL API on AWS Lambda](https://medium.com/a-man-with-no-server/building-a-golang-graphql-api-on-aws-lambda-b5278b7afc8c)を写経したものです。

struct毎へのファイルの分割、Serverless FrameworkでのAPI Gateway + Lambdaの実装を行うなどの改変をしています。

# Quick Start

1. Compile function

```
cd go-learning-graphql-serverless
GOOS=linux go build -o bin/main
```
2. Deploy!

```
serverless deploy
```
