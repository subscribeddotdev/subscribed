#!/bin/bash

echo "Waiting for the server"

while ! nc -z localhost 8080; do   
  sleep 0.1 # wait for 1/10 of the second before check again
done

echo "Server launched"