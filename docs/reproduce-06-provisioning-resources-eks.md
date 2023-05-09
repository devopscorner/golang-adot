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

## Provisioning EKS

### Clean Cache

- Goto `terraform/environment/providers/aws/infra/resources/eks` folder

      cd terraform/environment/providers/aws/infra/resources/eks

- Cleanup Cache Terraform

      rm -rf .terraform .terraform.lock.hcl terraform.eks.d


### Initialize Terraform

    terraform init


### Using Terraform Workspace

- Production Environment

      WORKSPACE_ENV=prod
      terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV

- Staging Environment

      WORKSPACE_ENV=staging
      terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV

- Lab Environment

      WORKSPACE_ENV=lab
      terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV


### Terraform Plan

- Save Binary Plan

      terraform plan --out tfplan.binary

- Export to JSON

      terraform show -json tfplan.binary > tfplan.json


### Terraform Security

- Using: [`terrascan`](https://github.com/tenable/terrascan)

      terrascan init
      terrascan scan -o human

- Using: [`tfsec`](https://aquasecurity.github.io/tfsec/v1.28.1/)

      tfsec .

- Using: [`checkov`](https://github.com/bridgecrewio/checkov)

      checkov -f tfplan.json


### Terraform Cost Review

- Using: [`infracost`](https://github.com/infracost/infracost)

      infracost breakdown --path tfplan.json


### Terraform Apply

      terraform apply -auto-approve


## CleanUp All

- Cleanup Terraform Cache

      rm -rf .terraform .terraform.lock.hcl terraform.eks.d

- Destroy All EKS

      terraform destroy -auto-approve
