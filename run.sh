#!/bin/sh

bundle exec rake deis:create_database
bundle exec puma -C config/puma.rb
