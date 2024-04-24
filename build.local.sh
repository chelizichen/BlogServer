#!/bin/bash  

# if permission denied
# run script with ` chmod +x build.sh ` 
readonly ServerName="BlogServer"

# rm
rm ./$ServerName.tar.gz ./sgrid_app

# compile
go build -o sgrid_app

# build
tar -cvf $ServerName.tar.gz ./sgrid.yml ./sgrid_app ./dist
