#!/bin/sh
#
# run.sh
# Copyright (C) 2016 weirdgiraffe <giraffe@cyberzoo.xyz>
#
# Distributed under terms of the MIT license.
#
docker run -d -p 5000:5000 --name insecure_registry registry:2

