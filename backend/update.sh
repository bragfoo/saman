#! /bin/bash

mane='saman'
date=`date +"%Y%m%d%H%M"`
final_name=$name$date

echo 'building backend'
cd api
echo $final_name
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../$final_name
cp $final_name /data/server/saman
cd /data/server/saman
rm -f saman
ln -s $final_name saman
