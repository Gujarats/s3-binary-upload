# Artifact Uploader from gradle caches to S3 [![Build Status](https://secure.travis-ci.org/Gujarats/s3-binary-upload.png)](http://travis-ci.org/Gujarats/s3-binary-upload)

This sis CLI use to upload the artifact binary like .jar and .pom from gradle caches to S3

## How to use

```shell
$ cd ~/.gradle/caches/modules-2/
$ s3-binary-upload #press enter
enter your package name = "TYPE_YOUR_PACKAGE_PREFIX"

```

NOTE use `all` to upload all artifact 

## Configuration

```yaml
---
region: ap-southeast-1
s3Bucket: artifact-repository-test1

```
