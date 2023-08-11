#!/bin/bash

docker build -t dbocker .

docker run -v $(pwd)/build:/app/build dbocker

cp build/main .