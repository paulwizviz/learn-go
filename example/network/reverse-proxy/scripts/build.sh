#!/bin/bash

go get github.com/GeertJohan/go.rice/rice

function node_build(){
    docker build -f ./build/Dockerfile -t paulwizviz/reactjs:current .
    id=$(docker create paulwizviz/reactjs:current)
    container_id="${id:0:12}"
    if [ -d ./web/public ]; then
        rm -rf ./web/public 
    fi
    docker cp $container_id:/opt/public ./web/public
    docker rm -f $container_id
}

function rice_single_html_build(){
    pushd ./internal/singlehtml
        rice embed-go
    popd
}

function rice_react_build(){
    pushd ./internal/react
        rice embed-go
    popd
}

node_build
rice_single_html_build
rice_react_build