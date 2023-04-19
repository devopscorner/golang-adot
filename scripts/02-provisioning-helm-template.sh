# Provisioning Helm Template

#!/usr/bin/env sh

# set -e

# Path xterm-256-color
export TERM="xterm-256color"

# Partition Data
# export REPO_PATH="/data/repo/golang-adot"

# Cloud9
export REPO_PATH="$HOME/environment/golang-adot"
export HELM_PATH="helm"
export HELM_SRC="$REPO_PATH/$HELM_PATH/template"
export BUCKET_NAME="devopscorner-adot-chart"

line1="----------------------------------------------------------------------------------------------------"
line2="===================================================================================================="

get_time() {
  DATE=$(date '+%Y-%m-%d %H:%M:%S')
}

s3_empty_bucket() {
    echo $line2
    echo " Empty All Bucket Objects..."
    echo " \$ aws s3 rm s3://$BUCKET_NAME --recursive"
    echo $line2
    sleep 2
    aws s3 rm s3://$BUCKET_NAME --recursive
    echo ' - DONE - '
    echo ''
    sleep 1
}

s3_delete_bucket() {
    echo $line2
    echo " Delete Bucket..."
    echo " \$ aws s3 rb s3://$BUCKET_NAME --region us-west-2"
    echo $line2
    sleep 2
    aws s3 rb s3://$BUCKET_NAME --region us-west-2
    echo ' - DONE - '
    echo ''
    sleep 1
}

s3_create_bucket() {
    echo $line2
    echo " Create Bucket..."
    echo " \$ aws s3 mb s3://$BUCKET_NAME --region us-west-2"
    echo $line2
    sleep 2
    aws s3 mb s3://$BUCKET_NAME --region us-west-2
    echo ' - DONE - '
    echo ''
    sleep 1
}

s3_check_bucket() {
    echo $line2
    echo " Check Bucket Created..."
    echo " \$ aws s3 ls s3://$BUCKET_NAME --recursive --human-readable --summarize"
    echo $line2
    sleep 2
    aws s3 ls s3://$BUCKET_NAME --recursive --human-readable --summarize
    echo ' - DONE - '
    echo ''
    sleep 1
}

cleanup_helm_repo() {
    echo $line2
    echo " Cleanup Helm Repository..."
    echo $line2
    helm repo rm stable
    helm repo rm devopscorner-adot
    helm repo rm devopscorner-adot-staging
    helm repo rm devopscorner-adot-lab
    echo ' - DONE - '
    echo ''
    sleep 1
}

check_helm_repo() {
    echo $line2
    echo " Check Existing Helm Repository..."
    echo " \$ helm repo list"
    echo $line2
    helm repo list
    echo ' - DONE - '
    echo ''
    sleep 1
}

init_helm_repo_prod() {
    echo $line1
    echo " Initialize Helm Repo Production..."
    echo " \$ helm s3 init s3://$BUCKET_NAME/prod"
    echo $line1
    helm s3 init s3://$BUCKET_NAME/prod
    echo ' - DONE - '
    echo ''
    sleep 1
}

init_helm_repo_staging() {
    echo $line1
    echo " Initialize Helm Repo Staging..."
    echo " \$ helm s3 init s3://$BUCKET_NAME/staging"
    echo $line1
    helm s3 init s3://$BUCKET_NAME/staging
    echo ' - DONE -'
    echo ''
    sleep 1
}

init_helm_repo_lab() {
    echo $line1
    echo " Initialize Helm Repo Lab..."
    echo " \$ helm s3 init s3://$BUCKET_NAME/lab"
    echo $line1
    helm s3 init s3://$BUCKET_NAME/lab
    echo ' - DONE - '
    echo ''
    sleep 1
}

init_helm() {
    echo $line2
    echo " Initialize Helm Repository..."
    echo $line2
    sleep 1
    echo ''
    init_helm_repo_prod
    init_helm_repo_staging
    init_helm_repo_lab
    echo ''
}


added_helm_repo_stable() {
    echo $line2
    echo " Initialize Helm Repo Stable..."
    echo " \$ helm repo add stable https://charts.helm.sh/stable"
    echo $line2
    helm repo add stable https://charts.helm.sh/stable
    helm repo update
    echo ' - DONE -'
    echo ''
    sleep 1
}


register_helm_repo_prod() {
    echo $line1
    echo " Register Helm Repo Production..."
    echo " \$ AWS_REGION=us-west-2 helm repo add devopscorner-adot s3://$BUCKET_NAME/prod"
    echo $line1
    sleep 1
    AWS_REGION=us-west-2 helm repo add devopscorner-adot s3://$BUCKET_NAME/prod
    echo ' - DONE -'
    echo ''
}

register_helm_repo_staging() {
    echo $line1
    echo " Register Helm Repo Staging..."
    echo " \$ AWS_REGION=us-west-2 helm repo add devopscorner-adot-staging s3://$BUCKET_NAME/staging"
    echo $line1
    sleep 1
    AWS_REGION=us-west-2 helm repo add devopscorner-adot-staging s3://$BUCKET_NAME/staging
    echo ' - DONE -'
    echo ''
}

register_helm_repo_lab() {
    echo $line1
    echo " Register Helm Repo Lab..."
    echo " \$ AWS_REGION=us-west-2 helm repo add devopscorner-adot-lab s3://$BUCKET_NAME/lab"
    echo $line1
    sleep 1
    AWS_REGION=us-west-2 helm repo add devopscorner-adot-lab s3://$BUCKET_NAME/lab
    echo ' - DONE -'
    echo ''
}

register_helm() {
    echo $line2
    echo " Register Helm Repository..."
    echo $line2
    sleep 1
    echo ''
    register_helm_repo_prod
    register_helm_repo_staging
    register_helm_repo_lab
    check_helm_repo
}

packing_push_helm_prod() {
    echo $line1
    echo " Packaging Helm Template Production..."
    echo " \$ cd $HELM_SRC/prod && ./helm-pack-prod.sh"
    echo $line1
    cd $HELM_SRC/prod
    sh ./helm-pack-prod.sh
    echo ' - DONE -'
    echo ''
    sleep 1

    echo $line1
    echo " Push Helm Template Production..."
    echo " \$ cd $HELM_SRC/prod && ./helm-push-prod.sh"
    echo $line1
    cd $HELM_SRC/prod
    sh ./helm-push-prod.sh
    echo ' - DONE -'
    echo ''
    sleep 1
}

packing_push_helm_staging() {
    echo $line1
    echo " Packaging Helm Template Staging..."
    echo " \$ cd $HELM_SRC/staging && ./helm-pack-staging.sh"
    echo $line1
    cd $HELM_SRC/staging
    sh ./helm-pack-staging.sh
    echo ' - DONE -'
    echo ''
    sleep 1

    echo $line1
    echo " Push Helm Template Staging..."
    echo " \$ cd $HELM_SRC/staging && ./helm-push-staging.sh"
    echo $line1
    cd $HELM_SRC/staging
    sh ./helm-push-staging.sh
    echo ' - DONE -'
    echo ''
    sleep 1
}

packing_push_helm_lab() {
    echo $line1
    echo " Packaging Helm Template Lab..."
    echo " \$ cd $HELM_SRC/lab && ./helm-pack-lab.sh"
    echo $line1
    cd $HELM_SRC/lab
    sh ./helm-pack-lab.sh
    echo ' - DONE -'
    echo ''
    sleep 1

    echo $line1
    echo " Push Helm Template Lab..."
    echo " \$ cd $HELM_SRC/lab && ./helm-push-lab.sh"
    echo $line1
    cd $HELM_SRC/lab
    sh ./helm-push-lab.sh
    echo ' - DONE -'
    echo ''
    sleep 1
}

packing_push_helm() {
    echo $line2
    echo " Packing & Push Helm Template..."
    echo $line2
    echo ''
    sleep 2
    packing_push_helm_prod
    packing_push_helm_staging
    packing_push_helm_lab
}

cleanup_existing()  {
    s3_empty_bucket
    s3_delete_bucket
    cleanup_helm_repo
}


main() {
    cleanup_existing
    s3_create_bucket
    s3_check_bucket
    added_helm_repo_stable
    check_helm_repo
    init_helm
    register_helm
    packing_push_helm
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main
