#!/bin/bash

cd /usr/app
task wait-for -- rabbitmq:5672 --timeout=30
task wait-for -- postgres:5432 --timeout=30
task run:$APP_NAME
