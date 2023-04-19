#!/usr/bin/env sh
# -----------------------------------------------------------------------------
#  Docker Build Multi Architecture Container
# -----------------------------------------------------------------------------
#  Author     : Dwi Fahni Denni
#  License    : Apache v2
# -----------------------------------------------------------------------------
set -e

export CI_PROJECT_PATH="devopscorner"
export CI_PROJECT_NAME="bookstore-adot"

export IMAGE="$CI_PROJECT_PATH/$CI_PROJECT_NAME"

export TARGETPLATFORM="linux/amd64"
export STACKS_NAME="devopscorner-amd64"
export ARCH="amd64"
# List PLATFORM:
# docker buildx inspect $STACKS_NAME

line1="----------------------------------------------------------------------------------------------------"
line2="===================================================================================================="

create_stack() {
    echo $line2
    echo " Build Stacks Multiplatform"
    echo " Stacks: $STACKS_NAME"
    echo $line2
    echo " -> docker buildx create --name $STACKS_NAME --driver docker-container --bootstrap"
    echo $line1
    docker buildx create --name $STACKS_NAME \
        --driver docker-container \
        --bootstrap
    echo " - DONE -"
    echo ''
}

use_stack() {
    echo $line2
    echo " Use Stacks Multiplatform"
    echo " Stacks: $STACKS_NAME"
    echo $line2
    echo " -> docker buildx use $STACKS_NAME"
    echo $line1
    docker buildx use $STACKS_NAME
    echo " - DONE -"
    echo ''
}

build_golang_adot() {
    TAGS="3.17 \
        alpine-3.17 \
        alpine-latest \
        alpine \
        latest "

    for TAG in $TAGS; do
        echo " Build Image => $IMAGE:$TAG"
        docker buildx build --push \
            --platform $TARGETPLATFORM \
            -f Dockerfile-$ARCH \
            -t $IMAGE:$TAG .
        echo ''
    done
}

build_cicd_alpine_317() {
    TAGS="cicd-alpine \
        cicd-alpine-latest \
        cicd-alpine-3.17 "

    for TAG in $TAGS; do
        echo " Build Image => $IMAGE:$TAG"
        docker buildx build --push \
            --platform $TARGETPLATFORM \
            -f Dockerfile-CICD.alpine-3.17-$ARCH \
            -t $IMAGE:$TAG .
        echo ''
    done
}

build_cicd_codebuild_40() {
    TAGS="cicd-codebuild \
        cicd-codebuild-latest \
        cicd-codebuild-4.0 \
        cicd-latest "

    for TAG in $TAGS; do
        echo " Build Image => $IMAGE:$TAG"
        docker buildx build --push \
            --platform $TARGETPLATFORM \
            -f Dockerfile-CICD.codebuild-4.0-$ARCH \
            -t $IMAGE:$TAG .
        echo ''
    done
}

docker_build() {
    build_golang_adot
    build_cicd_alpine_317
    build_cicd_codebuild_40
}

docker_clean() {
    echo "Cleanup Unknown Tags"
    echo "docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi"
    docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi
    echo ''
}

main() {
    create_stack
    use_stack
    docker_build
    docker_clean
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main