# Golang Bookstore ADOT

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

## Workflow CI/CD Pipeline

- Create Container Image CI/CD CodeBuild, refer to this repository: [DevOpsCorner CI/CD CodeBuild](https://github.com/devopscorner/devopscorner-container/tree/main/compose/docker/cicd-codebuild)

- Create HelmChart Template Global, go to [this](https://github.com/devopscorner/devopscorner-helm) section
  - `api`
  - `backend`
  - `frontend`
  - `stateful`
  - `secretref`
  - `svcrole`
- Deploy HelmChart Template Global to S3
  - `AWS_REGION=us-west-2 helm repo add devopscorner-lab s3://devopscorner-adot-chart/lab`
  - `AWS_REGION=us-west-2 helm repo add devopscorner-staging s3://devopscorner-adot-chart/staging`
  - `AWS_REGION=us-west-2 helm repo add devopscorner-prod s3://devopscorner-adot-chart/prod`
- Create HelmChart Template for GO App
  - Helm template GO App (`_infra/{env}/helm-template.yml`)
  - Helm value GO App (`_infra/{env}/helm-value.yml`)

- Create `Dockerfile` for Container GO App CI/CD
- Create Script CI/CD
  - `ecr-build.sh`
  - `ecr-push.sh`
  - `ecr-tag.sh`
  - `git-clone.sh`
  - `Makefile`

- Register Container Image GO App to Amazon ECR (Container Registry)
- Create script for Building Container Image GO App (`.aws/buildspec-build.yml`)
- Create script for Deployment Container Image GO App to EKS (`.aws/buildspec-deploy.yml`)
- Setup Variable Environment / Using Config Secret with **AWS Systems Manager (Parameter Store)**

- Create Pipeline with AWS CodePipeline
  - **Source**: Reference from CodeCommit and/or 3rd Party Repository (GitHub, GitLab, BitBucket, Azure DevOps)
  - **Build**: Building Container GO App
    - Golang Unit Test
    - Code Quality and Code Security
    - Build Container GO App
    - Tagging Container GO App
    - Push Container GO App to Container Registry (ECR / Dockerhub)
  - **Deploy**: Deploy Container GO App
    - **Static Application Security Testing (SAST)**, or static analysis
    - Manual Approval
    - **Deploy-DEV**
    - Manual Approval
    - **Deploy-UAT**
    - **Dynamic Application Security Testing (DAST)**
    - Manual Approval
    - **Deploy-PROD**

- Running Deployment: Commit -> Push -> Webhook (CodeCommit)

- CodeBuild Process
  - Build Container
    - Environment Image: `aws/codebuild/amazonlinux2-x86_64-standard:4.0`
    - Environment Type: `Linux`
    - Buildspec: `.aws/buildspec-build.yml`
  - Deploy Container
    - Environment Image: `devopscorner/cicd:codebuild-4.0` or `YOUR_AWS_ACCOUNT.dkr.ecr.us-west-2.amazonaws.com/devopscorner/cicd:codebuild-4.0`
    - Environment Type: `Linux`
    - Buildspec: `.aws/buildspec-deploy.yml`
