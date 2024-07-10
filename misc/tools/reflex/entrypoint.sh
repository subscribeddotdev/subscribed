#!/bin/bash

cd /usr/app
go run ./misc/tools/wait-for/ &&
  task run:$APP_NAME
