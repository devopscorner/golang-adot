# Golang Bookstore ADOT

ADOT (AWS Distro for OpenTelemetry) Implementation for Simple Golang RESTful API Application (Bookstore)

[goreport](https://goreportcard.com/badge/github.com/devopscorner/golang-adot/src)
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

## Install Dependencies Terraform Modules

### Check Existing Modules

    ls -al terraform

### Using Makefile Installation

    make sub-officials   # install terraform officials modules dependencies only

    make sub-community   # install terraform community modules dependencies only

    make sub-all         # install all / both modules (officials & community)

### Check Modules Installed

    ls -alR terraform/modules/providers/aws/

