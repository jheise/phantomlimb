#!/bin/sh

sudo docker run --rm -it --mount type=bind,source="$(pwd)",target=/go/src/github.com/jheise/phantomlimb jheise/gophantomjs /bin/bash
