#!/bin/bash

echo "build app..."
docker build -t simplebet:local -f $(pwd)/docker/build.dockerfile .
if [ $? -eq 0 ]; then
    echo "build success"
else
    echo "build fail"
    exit 1
fi

echo "run services"
docker-compose -f $(pwd)/docker-compose.yml up --build -d
if [ $? -eq 0 ]; then
    echo "run success"
else
    echo "run fail"
    exit 1
fi
echo "all services are running..."