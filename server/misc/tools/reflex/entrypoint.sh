#!/bin/bash

cd /usr/app
go run ./misc/tools/wait/ &&
task run:$APP_NAME
