# Golang Bookstore ADOT

ADOT (AWS Distro for OpenTelemetry) Implementation for Simple Golang RESTful API Application (Bookstore)

[![goreport](https://goreportcard.com/badge/github.com/devopscorner/golang-adot/src)](https://goreportcard.com/badge/github.com/devopscorner/golang-adot/src)
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

## Provisioning HelmChart Template

### Create S3 Bucket for Helm

- Create Bucket S3

      BUCKET_NAME="devopscorner-adot-chart"

      aws s3 mb "s3://${BUCKET_NAME}" --region us-west-2

- Check If Bucket S3 Created

      aws s3 ls "s3://${BUCKET_NAME}" --recursive --human-readable --summarize


### Check Helm Repo

    helm repo list


### Initialize Helm Repo

- Production Environment

      helm s3 init "s3://${BUCKET_NAME}/prod"

- Staging Environment

      helm s3 init "s3://${BUCKET_NAME}/staging"

- Lab Environment

      helm s3 init "s3://${BUCKET_NAME}/lab"


### Register Helm Repo

- Production Environment

      AWS_REGION=us-west-2 helm repo add devopscorner-adot "s3://${BUCKET_NAME}/prod"

- Staging Environment

      AWS_REGION=us-west-2 helm repo add devopscorner-adot-staging "s3://${BUCKET_NAME}/staging"

- Lab Environment

      AWS_REGION=us-west-2 helm repo add devopscorner-adot-lab "s3://${BUCKET_NAME}/lab"


### Update Helm Repo

    helm repo update


### Check If Helm Repo Already Registered

    helm repo list


### Goto `helm/template/[environment]` folder

- Production Environment

      cd helm/template/prod

- Staging Environment

      cd helm/template/staging

- Lab Environment

      cd helm/template/lab


### Execute Helm Pack

- Production Environment

      ./helm-pack-prod.sh

- Staging Environment

      ./helm-pack-staging.sh

- Lab Environment

      ./helm-pack-lab.sh


### Execute Helm Push

- Production Environment

      ./helm-push-prod.sh

- Staging Environment

      ./helm-push-staging.sh

- Lab Environment

      ./helm-push-lab.sh


## CleanUp All

- Empty Bucket

      aws s3 rm "s3://${BUCKET_NAME}" --recursive

- Delete Bucket

      aws s3 rb "s3://${BUCKET_NAME}" --region us-west-2

- Cleanup Helm Repo

      helm repo rm devopscorner-adot
      helm repo rm devopscorner-adot-staging
      helm repo rm devopscorner-adot-lab
      helm repo update
