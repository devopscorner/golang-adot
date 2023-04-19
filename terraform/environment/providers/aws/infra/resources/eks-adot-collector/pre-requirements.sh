#!/bin/sh

helm repo add aws-otel-collector https://aws-observability.github.io/aws-otel-helm-charts
helm repo update

helm search repo aws-otel-collector
