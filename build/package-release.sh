#!/bin/bash
VERSION=0.4.5
BUILD_FOLDER=bin/*
PROCESSING_FOLDER=release

# remove old folder
rm -fr $PROCESSING_FOLDER

# make new temp folder
mkdir $PROCESSING_FOLDER

# move new files
cp -R $BUILD_FOLDER $PROCESSING_FOLDER

cd $PROCESSING_FOLDER

for f in *; do
  NAME=$(basename $f)
  echo "Processing $NAME file: $f"
  tar -czvf $NAME-$VERSION.tar.gz $NAME
done

find . -type f ! -name "*.tar.gz" -delete