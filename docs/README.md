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

## INDEX

- Build Container `devopscorner/cicd`
  - Build Container `devopscorner/cicd` for DockerHub, detail [here](https://github.com/devopscorner/devopscorner-container/blob/master/docs/container-cicd-dockerhub.md)
  - Build Container `devopscorner/cicd` for ECR, detail [here](https://github.com/devopscorner/devopscorner-container/blob/master/docs/container-cicd-ecr.md)

- Build Container `devopscorner/bookstore-adot`
  - Build Container `devopscorner/bookstore-adot` for DockerHub, detail [here](container-bookstore-adot-dockerhub.md)
  - Build Container `devopscorner/bookstore-adot` for ECR, detail [here](container-bookstore-adot-ecr.md)

- Workflow CI/CD Pipeline, detail [here](workflow-cicd-bookstore-adot-pipeline.md)

- Deployments:
  - **AWS Developer Tools** (AWS CodeCommit, AWS CodeBuild & AWS CodePipeline), detail [here](deployment-aws-developer-tools.md) link
  - **Jenkins CI/CD**, detail [here](deployment-jenkins.md) link
  - **Terraform AWS CodeBuild, AWS CodePipeline & Amazon SNS**, detail [here](deployment-terraform.md) link

- Reproduce Provisioning
  1. Install Terraform Modules, detail [here](reproduce-01-teraform-modules.md)
  2. Provisioning HelmChart Template, detail [here](reproduce-02-helm-template.md)
  3. Provisioning TFState DB & Bucket, detail [here](reproduce-03-provisioning-tfstate-db-bucket.md)
  4. Provisioning Infra Core, detail [here](reproduce-04-provisioning-infra-core.md)
  5. Provisioning Infra CI/CD, detail [here](reproduce-05-provisioning-infra-cicd.md)
  6. Provisioning Resources Amazon EKS, detail [here](reproduce-06-provisioning-resources-eks.md)
  7. Provisioning Resources Managed Services Prometheus (AMP), detail [here](reproduce-07-provisioning-resources-amp.md)
  8. Provisioning Resources Managed Services Grafana (AMG), detail [here](reproduce-08-provisioning-resources-amg.md)
  9. Provisioning Resources Managed Services OpenSearch, detail [here](reproduce-09-provisioning-resources-opensearch.md)

## Reproduce Testing

- RESTful API Testing, detail [here](test-restful-api.md) link
