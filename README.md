# Artifact Uploader from gradle caches to S3
This sis CLI use to upload the artifact binary like .jar and .pom from gradle caches to S3

## How to use
Using configuration file located in

```shell
$ $HOME/.s3-binary-upload/config.yaml
```
You can use any type of extension file like `.yaml`,`.toml`,`.json` choose whaterver config extension you like.
after installation you can run the command in the following order :

```shell
$ cd ~/.gradle/caches/modules-2/
$ s3-binary-upload #press enter
enter your package name = "TYPE_YOUR_PACKAGE_PREFIX"

```

NOTE use `all` to upload all artifact 
