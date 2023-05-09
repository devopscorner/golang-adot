# Provisioning EKS ADOT Collector

#!/usr/bin/env sh

set -e

# Path xterm-256-color
export TERM="xterm-256color"

# Partition Data
# export REPO_PATH="/data/repo/golang-adot"

# Cloud9
export REPO_PATH="$HOME/environment/golang-adot"
export TF_PATH="terraform"
export TF_MODULE_PATH="$REPO_PATH/$TF_PATH/modules"
export ALL_MODULES_PATH="$TF_MODULE_PATH/providers/aws"

export TF_INFRA_PATH="$REPO_PATH/$TF_PATH/environment/providers/aws/infra"
export TF_CORE_PATH="$TF_INFRA_PATH/core"
export TF_RESOURCES_PATH="$TF_INFRA_PATH/resources"
export TF_STATE_PATH="$TF_INFRA_PATH/tfstate"
export TF_AMG_PATH="$TF_RESOURCES_PATH/eks-adot-collector"

export WORKSPACE_ENV="prod"

line1="----------------------------------------------------------------------------------------------------"
line2="===================================================================================================="

get_time() {
  DATE=$(date '+%Y-%m-%d %H:%M:%S')
}

cleanup_terraform() {
    echo $line2
    echo " Cleanup Terraform Core..."
    echo " \$ rm -rf .terraform .terraform.lock.hcl terraform.eks-adot-collector.d"
    echo $line2
    cd $TF_AMG_PATH
    rm -rf .terraform .terraform.lock.hcl terraform.eks-adot-collector.d
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_init() {
    echo $line2
    echo " Initialize Terraform..."
    echo " \$ terraform init"
    echo $line2
    cd $TF_AMG_PATH
    terraform init
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_workspace() {
    echo $line2
    echo " Create / Select Terraform Workspace..."
    echo " \$ terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV"
    echo $line2
    cd $TF_AMG_PATH
    terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_plan() {
    echo $line2
    echo " Terraform Plan & Save Binary Plan..."
    echo " \$ terraform plan --out tfplan.binary"
    echo $line2
    cd $TF_AMG_PATH
    terraform plan --out tfplan.binary
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_save_plan() {
    echo $line2
    echo " Export JSON Terraform Plan..."
    echo " \$ terraform show -json tfplan.binary > tfplan.json"
    echo $line2
    cd $TF_AMG_PATH
    terraform show -json tfplan.binary > tfplan.json
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_security() {
    echo $line2
    echo " Running Security Inspect Terraform..."
    echo " \$ terraform show -json tfplan.binary > tfplan.json"
    echo $line2
    cd $TF_AMG_PATH
    # ================== #
    #  Terraform Addons  #
    # ================== #
    # ~ Terrascan ~
    echo $line1
    echo " Using: terrascan ..."
    echo " \$ terrascan init"
    echo " \$ terrascan scan -o human"
    echo $line1
    terrascan init
    terrascan scan -o human
    echo ''
    sleep 1

    # ~ Tfsec ~
    echo $line1
    echo " Using: tfsec ..."
    echo " \$ tfsec ."
    echo $line1
    tfsec .
    echo ''
    sleep 1

    # ~ Checkov
    echo $line1
    echo " Using: checkov ..."
    echo " \$ checkov -f tfplan.json"
    echo $line1
    tfsec .
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_cost_review() {
    echo $line2
    echo " Cost Review Terraform..."
    echo " \$ infracost breakdown --path tfplan.json"
    echo $line2
    cd $TF_AMG_PATH
    # ~ Infracost
    infracost breakdown --path tfplan.json
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_apply() {
    echo $line2
    echo " Apply Terraform..."
    echo " \$ terraform apply -auto-approve"
    echo $line2
    cd $TF_AMG_PATH
    # ======================== #
    #  Terraform Provisioning  #
    # ======================== #
    terraform apply -auto-approve
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

terraform_destroy() {
    echo $line2
    echo " Cleanup TFState..."
    echo " \$ terraform destroy -auto-approve"
    echo $line2
    cd $TF_AMG_PATH
    # =================== #
    #  Terraform Destroy  #
    # =================== #
    terraform destroy -auto-approve
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

main() {
    cleanup_terraform
    terraform_init
    terraform_workspace
    terraform_plan
    terraform_save_plan
    # terraform_security
    terraform_cost_review
    # terraform_apply
    # terraform_destroy
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main