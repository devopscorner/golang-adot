# Provisioning Infra Core

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

export WORKSPACE_ENV="prod"

line1="----------------------------------------------------------------------------------------------------"
line2="===================================================================================================="

get_time() {
  DATE=$(date '+%Y-%m-%d %H:%M:%S')
}

terraform_workspace() {
    echo $line2
    echo " Create / Select Terraform Workspace..."
    echo " \$ terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV"
    echo $line2
    cd $TF_CORE_PATH
    terraform workspace select $WORKSPACE_ENV || terraform workspace new $WORKSPACE_ENV
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
    cd $TF_CORE_PATH
    # ======================== #
    #  Terraform Provisioning  #
    # ======================== #
    terraform apply -auto-approve
    echo ''
    echo ' - DONE - '
    echo ''
    sleep 1
}

main() {
    terraform_workspace
    terraform_apply
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main