# Artifact Uploader from gradle caches to S3 [![Build Status](https://secure.travis-ci.org/Gujarats/s3-binary-upload.png)](http://travis-ci.org/Gujarats/s3-binary-upload)

This sis CLI use to upload the artifact binary like .jar and .pom from gradle caches to S3

## How to use
Here is some config example using `.yaml`
Using configuration file located in

```shell
$ touch $HOME/.s3-binary-upload/config.yaml
---
s3Buckets:
 - PUT_YOUT_MULTIPLE_BUCKET_HERE
region: YOUR_REGION 
s3Bucket: YOUR_SINGLE_BUCKET 
```

You can use any type of extension file like `.yaml`,`.toml`,`.json` choose whaterver config extension you like.
after installation you can run the command in the following order :

```shell
$ cd ~/.gradle/caches/modules-2/
$ s3-binary-upload #press enter
enter your package name = "TYPE_YOUR_PACKAGE_PREFIX"

```

NOTE use `all` to upload all artifact 

## Configuration

```yaml
---
region: your-region 
s3Bucket: your-bucket 
profile: your-aws-profile 

artfactsDirectories:
  - /artifactory/ext-release-local/
  - /artifactory/libs-release-local/

uploadArtifacs: true
downloadArtifacs: false
linkArtifacts:
  - https://your-artifact-host/artifactory/ext-release-local/
  - https://your-artifact-host/artifactory/libs-release-local/
username: username-for-your-host
password: password 

```
