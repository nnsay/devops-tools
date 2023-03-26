# AWS-TOOLS

本项目编写一些日常 aws 相关运维的小工具, 本项目使用`spf13/cobra`为命令行开发开发, 目前支持的工具如下:

- 删除过期的 IAM 证书

```
# use alias command
aws-tools iam dec
# use the aws profile to provide the credentials
AWS_PROFILE=dev aws-tools iam dec
```
