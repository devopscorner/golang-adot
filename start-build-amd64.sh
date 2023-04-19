#!/usr/bin/env sh
# -----------------------------------------------------------------------------
#  Docker Build Container (Linux/AMD64)
# -----------------------------------------------------------------------------
#  Author     : Dwi Fahni Denni
#  License    : Apache v2
# -----------------------------------------------------------------------------
set -e

export CI_PROJECT_PATH="devopscorner"
export CI_PROJECT_NAME="bookstore-adot"

export IMAGE="$CI_PROJECT_PATH/$CI_PROJECT_NAME"

export TARGETPLATFORM="linux/amd64"
export ARCH="amd64"

build_golang_adot() {
    TAGS="3.17 \
        alpine-3.17 \
        alpine-latest \
        alpine \
        latest "

    for TAG in $TAGS; do
        echo " Build Image => $IMAGE:$TAG"
        docker build \
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
        docker build \
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
        docker build \
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
    docker_build
    docker_clean
    echo ''
    echo '-- ALL DONE --'
}

### START HERE ###
main