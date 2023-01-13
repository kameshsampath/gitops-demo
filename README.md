# GitOps Greeter

A hello world demo to used to demonstrate GitOps applications.

## Pre-requisites

- [ko](https://ko.build)

## Quick build and test

```shell
export KO_DOCKER_REPO=<your container repo> # e.g. docker.io/example/gitops-greeter
docker run -p 9090:8080 $(ko build --bare --tags=latest --platform=linux/amd64 --platform=linux/arm64)
```
