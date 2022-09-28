#!/bin/bash

registry_host_port=$1
image_name=$2
image_tag=$3

if [ -z $registry_host_port ]; then
    exit "registry_host_port is not found"
fi

if [ -z $image_name ]; then
    exit "image_name is not found"
fi

if [ -z $image_tag ]; then
    exit "image_tag is not found"
fi
oldcurl=`curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
    -X GET http://$registry_host_port/v2/$image_name/manifests/$image_tag 2>&1 | \
    grep Docker-Content-Digest | awk -F ' ' '{print$3}'`

if [ -z $oldcurl ]; then
    exit "registry can't find this image"
fi

newcurl=$(echo $oldcurl | sed -e 's/\r//g')

echo "Deleting $image_name:$image_tag"

curl -X DELETE http://$registry_host_port/v2/$image_name/manifests/$newcurl

#set -x 相当于显示每一条执行
###执行回收机制
###docker exec -it registry sh
###registry garbage-collect /etc/docker/registry/config.yml 