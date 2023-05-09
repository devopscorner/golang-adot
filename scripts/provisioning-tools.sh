#!/bin/sh

# bitnami             	https://charts.bitnami.com/bitnami
# opensearch          	https://opensearch-project.github.io/helm-charts/
# prometheus-community	https://prometheus-community.github.io/helm-charts
# opensearch-operator 	https://opster.github.io/opensearch-k8s-operator/
# fluent              	https://fluent.github.io/helm-charts
# jetstack            	https://charts.jetstack.io

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add opensearch https://opensearch-project.github.io/helm-charts/
helm repo add opensearch-operator https://opster.github.io/opensearch-k8s-operator/
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add fluent https://fluent.github.io/helm-charts
helm repo add jetstack https://charts.jetstack.io
helm repo add aws-otel-collector https://aws-observability.github.io/aws-otel-helm-charts
helm repo update

helm search repo prometheus-community
helm search repo opensearch
helm search repo opensearch-operator
helm search repo aws-otel-collector

kubectl create namespace observability

helm upgrade --install prometheus prometheus-community/prometheus -n observability --set server.nodeSelector."node"="devopscorner-monitoring"
helm upgrade --install opensearch opensearch/opensearch -n observability --set server.nodeSelector."node"="devopscorner-monitoring"
helm upgrade --install opensearch-operator opensearch-operator/opensearch-operator -n observability --set server.nodeSelector."node"="devopscorner-monitoring"

# kubectl get pods --namespace=observability -l app.kubernetes.io/component=opensearch-cluster-master -w

kubectl -f manifest-grafana.yml -n observability apply

helm upgrade --install opensearch-dashboard opensearch/opensearch-dashboards -n observability --set server.nodeSelector."node"="devopscorner-monitoring"
# export POD_NAME=$(kubectl get pods --namespace observability -l "app.kubernetes.io/name=opensearch-dashboards,app.kubernetes.io/instance=opensearch-dashboard" -o jsonpath="{.items[0].metadata.name}")
# export CONTAINER_PORT=$(kubectl get pod --namespace observability $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
# echo "Visit http://127.0.0.1:8080 to use your application"
# kubectl --namespace observability port-forward $POD_NAME 8080:$CONTAINER_PORT

helm upgrade --install fluent-bit fluent/fluent-bit -n observability --set server.nodeSelector."node"="devopscorner-monitoring"
# export POD_NAME=$(kubectl get pods --namespace observability -l "app.kubernetes.io/name=fluent-bit,app.kubernetes.io/instance=fluent-bit" -o jsonpath="{.items[0].metadata.name}")
# kubectl --namespace observability port-forward $POD_NAME 2020:2020
# curl http://127.0.0.1:2020

# kubectl create secret tls observability-tls \
#     --namespace observability \
#     --key server.key \
#     --cert server.crt

kubectl expose deployment jumppods --port=80 --target-port=8080 \
    --name=jumppods-lb --type=LoadBalancer --namespace=devopscorner

kubectl expose deployment grafana --port=3000 --target-port=3000 \
    --name=grafana-lb --type=LoadBalancer --namespace=observability

kubectl expose deployment prometheus-server --port=9090 --target-port=9090 \
    --name=prometheus-server-lb --type=LoadBalancer --namespace=observability

kubectl expose deployment opensearch-dashboard-opensearch-dashboards --port=5601 --target-port=5601 \
    --name=opensearch-dashboard-lb --type=LoadBalancer --namespace=observability


### Update HelmChart ###
echo '' > opensearch-template.yaml

helm template \
  opensearch opensearch/opensearch \
  --namespace observability \
  --create-namespace >> opensearch-template.yaml

kubectl -f opensearch-template.yaml -n observability apply


echo '' > fluentbit-template.yaml

helm template \
  fluent-bit fluent/fluent-bit \
  --namespace observability \
  --create-namespace >> fluentbit-template.yaml

kubectl -f fluentbit-template.yaml -n observability apply

### Network Ingress ###
kubectl -f ingress-nginx-goapp-prod.yaml -n observability apply
kubectl -f ingress-nginx-goapp-staging.yaml -n observability apply
kubectl -f ingress-nginx-opensearch-prod.yaml -n observability apply
kubectl -f ingress-nginx-opensearch-staging.yaml -n observability apply