# Golang Bookstore ADOT - Amazon ECR (Elastic Container Registry)

ADOT (AWS Distro for OpenTelemetry) Implementation for Simple Golang RESTful API Application (Bookstore)

![goreport](https://goreportcard.com/badge/github.com/devopscorner/golang-adot/src)
![all contributors](https://img.shields.io/github/contributors/devopscorner/golang-adot)
![tags](https://img.shields.io/github/v/tag/devopscorner/golang-adot?sort=semver)
[![docker pulls](https://img.shields.io/docker/pulls/devopscorner/bookstore-adot.svg)](https://hub.docker.com/r/devopscorner/bookstore-adot/)
![download all](https://img.shields.io/github/downloads/devopscorner/golang-adot/total.svg)
![view](https://views.whatilearened.today/views/github/devopscorner/golang-adot.svg)
![clone](https://img.shields.io/badge/dynamic/json?color=success&label=clone&query=count&url=https://github.com/devopscorner/golang-adot/blob/master/clone.json?raw=True&logo=github)
![issues](https://img.shields.io/github/issues/devopscorner/golang-adot)
![pull requests](https://img.shields.io/github/issues-pr/devopscorner/golang-adot)
![forks](https://img.shields.io/github/forks/devopscorner/golang-adot)
![stars](https://img.shields.io/github/stars/devopscorner/golang-adot)
[![license](https://img.shields.io/github/license/devopscorner/golang-adot)](https://img.shields.io/github/license/devopscorner/golang-adot)

---

## Build Container Image

- Clone this repository

  ```
  git clone https://github.com/devopscorner/golang-adot.git
  ```

- Replace "YOUR_AWS_ACCOUNT" with your AWS ACCOUNT ID

  ```
  find ./ -type f -exec sed -i 's/YOUR_AWS_ACCOUNT/123456789012/g' {} \;
  ```

- Set Environment Variable

  ```
  export ALPINE_VERSION=3.17   # 3.15 | 3.16 | 3.17
  export BASE_IMAGE="alpine"
  export IMAGE="YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot"
  export TAG="latest"
  ```

- Execute Build Image

  ```
  # Golang 1.19.3 - Alpine 3.15
  docker build -f Dockerfile -t ${IMAGE}:alpine .
  docker build -f Dockerfile.alpine-3.15 -t ${IMAGE}:alpine-3.15 .
  docker build -f Dockerfile.alpine-3.15 -t ${IMAGE}:golang1.19.3-alpine3.15 .

  # Golang 1.19.5 - Alpine 3.16
  docker build -f Dockerfile -t ${IMAGE}:alpine .
  docker build -f Dockerfile.alpine-3.16 -t ${IMAGE}:alpine-3.16 .
  docker build -f Dockerfile.alpine-3.16 -t ${IMAGE}:golang1.19.5-alpine3.16 .

  # Golang 1.19.5 - Alpine 3.17
  docker build -f Dockerfile -t ${IMAGE}:alpine .
  docker build -f Dockerfile.alpine-3.17 -t ${IMAGE}:alpine-3.17 .
  docker build -f Dockerfile.alpine-3.17 -t ${IMAGE}:golang1.19.5-alpine3.17 .

  -- or --

  ecr-build.sh ${YOUR_AWS_ACCOUNT} alpine Dockerfile ${ALPINE_VERSION}
  ecr-build.sh ${YOUR_AWS_ACCOUNT} alpine Dockerfile.alpine-3.15 ${ALPINE_VERSION}
  ecr-build.sh ${YOUR_AWS_ACCOUNT} alpine Dockerfile.alpine-3.16 ${ALPINE_VERSION}
  ecr-build.sh ${YOUR_AWS_ACCOUNT} alpine Dockerfile.alpine-3.17 ${ALPINE_VERSION}

  -- or --

  # default: 3.17
  make ecr-build-alpine ARGS=YOUR_AWS_ACCOUNT
  ```

## Push Image to Amazon ECR (Elastic Container Registry)

- Create Tags Image
  - Example:

    ```
    # Alpine
    docker tag YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:latest

    docker tag YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine-latest

    docker tag YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine-3.16
    ```

  - With Script:

    ```
    # default: 3.17
    docker tag ${IMAGE}:${ALPINE_VERSION}

    -- or --

    # default: 3.17
    ./ecr-tag.sh ARGS=YOUR_AWS_ACCOUNT alpine ${ALPINE_VERSION} CI_PATH=devopscorner/bookstore-adot

    -- or --

    make ecr-tag-alpine ARGS=YOUR_AWS_ACCOUNT CI_PATH=devopscorner/bookstore-adot
    ```

 Push Image to **Amazon ECR** with Tags

- Example:

    ```
    # Alpine
    docker push YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine

    docker push YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine-latest

    docker push YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/bookstore-adot:alpine-3.16
    ```

- With Script:

    ```
    ./ecr-push.sh ARGS=YOUR_AWS_ACCOUNT alpine CI_PATH="devopscorner/bookstore-adot"

    -- or --

    make ecr-push-alpine ARGS=YOUR_AWS_ACCOUNT CI_PATH="devopscorner/bookstore-adot"
    ```
