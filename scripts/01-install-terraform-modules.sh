# Install Terraform Modules

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

line1="----------------------------------------------------------------------------------------------------"
line2="===================================================================================================="

get_time() {
  DATE=$(date '+%Y-%m-%d %H:%M:%S')
}

module_cleanup() {
    echo $line2
    echo " Cleanup All Previous Modules..."
    echo " \$ rm -rf $TF_MODULE_PATH"
    echo $line2
    sleep 2
    rm -rf $TF_MODULE_PATH
    echo ' - DONE - '
    echo ''
    sleep 1
}

install_module_officials() {
    echo $line2
    echo " Download Officials Modules..."
    echo " \$ make sub-officials"
    echo $line2
    sleep 2
    cd $REPO_PATH
    make sub-officials
    echo ' - DONE - '
    echo ''
    sleep 1
}

install_module_community() {
    echo $line2
    echo " Download Community Modules..."
    echo " \$ make sub-community"
    echo $line2
    sleep 2
    cd $REPO_PATH
    make sub-community
    echo ' - DONE - '
    echo ''
    sleep 1
}

show_modules() {
    echo $line2
    echo " Check Module List..."
    echo " \$ ls -alR $ALL_MODULES_PATH"
    echo $line2
    sleep 2
    ls -alR $ALL_MODULES_PATH
    echo ' - DONE - '
}

main() {
    module_cleanup
    install_module_officials
    install_module_community
    show_modules
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main

