#!/bin/sh
#
# make-some-registry-entries.sh
# Copyright (C) 2016 weirdgiraffe <giraffe@cyberzoo.xyz>
#
# Distributed under terms of the MIT license.
#

docker tag alpine localhost:5000/example:1
docker tag alpine localhost:5000/example:2
docker tag alpine localhost:5000/example:3
docker tag alpine localhost:5000/example:latest
docker push localhost:5000/example
