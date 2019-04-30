#!/usr/bin/env bash
ruby_image="ruby:2.6.3"
docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app "$ruby_image" bundle install
