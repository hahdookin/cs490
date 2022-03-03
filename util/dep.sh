#!/bin/bash

# Run as sudo 

# Install dependencies
apt install nginx -y

# Download go1.17.7
wget go.dev/dl/go1.17.7.linux-amd64.tar.gz

# Removes old go and uncompress new Go
rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz

# Add Go Path
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile

# Update PATH
source $HOME/.profile
