#!/bin/bash

# usage: prints usage from afs-send script
usage() {
    echo ""
    echo "afs-send : Meant to zip and send files to afs with minimal effort"
    echo "================================================================="
    echo "$ ./afs-send.sh <ucid>"
    echo "$ ./util/afs-send.sh <ucid>"
}

# checkset: checks what directory the script is being run from and switches it
checkset() {
    if [[ `pwd` == *"cs490/util"* ]];
    then 
        # echo "In util"
        cd ../..
    elif [[ `pwd` == *"cs490"* ]];
    then
        # echo "in base dir"
        cd ..
    else
        echo "Go to either cs490/ or cs490/util/ directory"
        exit 1
    fi
    echo `pwd`
}

# compress: compresses all files in the cs490 directory into single file
compress() {
    tar -cvf $1 cs490/*
}


# send: sends cs490.tar to afs via scp
send() {
    # Set ucid to first system argument, and the zip file (contains the github code) to second system argument
    ucid=$1 
    zipfile=$2

    # Sends the zipfile to $ucid njit's afs (the ${ucid:0:1} -> specifies where in afs the ucid lies)
    scp $zipfile $ucid@afsconnect1.njit.edu:/afs/cad/u/${ucid:0:1}/${ucid:1:1}/$ucid
}


# checks if the script passes a system argument
if [ $# -lt 1 ]; then
    usage
else
    ucid=$1
    checkset
    filename="cs490.tar"
    
    # compress cs490 directory -> cs490.tar
    compress $filename

    # send cs490.tar to afs 
    send $ucid $filename

    echo -e "\nTo decompress:\n\t$ tar xvf cd490.tar"

    # removes the cs490.tar file locally
    rm $filename
    
    # move back into cs490 directory
    cd cs490
fi