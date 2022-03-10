#!/bin/bash

# alias update-aws="<LOCATION>/update-aws.sh"
curr=`pwd`

cd ~/cs490
git pull
cd middleware
# make compile
# sudo systemctl daemon-reload
# sudo systemctl restart nginx
# sudo systemctl restart goweb

cd $(curr)