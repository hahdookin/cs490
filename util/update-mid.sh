#!/bin/bash

# Commands to update middleware server

git pull
make compile
sudo systemctl daemon-reload
sudo systemctl restart nginx
sudo systemctl restart goweb