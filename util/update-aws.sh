#!/bin/bash

# alias update-aws="<LOCATION>/update-aws.sh"

git pull
sudo systemctl daemon-reload
sudo systemctl restart nginx
sudo systemctl restart goweb