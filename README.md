# AWS-TOOLS

本项目编写一些日常 aws 相关运维的小工具, 本项目使用`spf13/cobra`为命令行开发开发, 目前支持的工具如下:

- 删除过期的 IAM 证书

  作用: 删除已经过期的 IAM 证书

  使用方法:

  ```bash
  # use alias command
  aws-tools iam dec
  # use fullname command
  aws-tools iam delete-expired-certification
  ```

  参数说明:
  |参数名称|别名|描述|默认值|
  |---|---|---|---|
  |expiration|e|指定过期时间, 格式为时间戳| time.Now()|
  |path-prefix|p|证书路径| /cloudfront/|

- 检查临期 IAM 证书

  作用: 检查即将过期的证书并提醒, 提醒消息发送到 Slack(需配置 **SLACK_HOOK**)

  使用方法:

  ```bash
  # use alias command
  aws-tools iam ccd
  # use fullname command
  aws-tools iam check-certification-date
  ```

  参数说明:
  |参数名称|别名|描述|默认值|
  |---|---|---|---|
  |expire-hours|e|剩余过期小时数| 72|
  |path-prefix|p|证书路径| /cloudfront/|
  |channel|c|Slack Channel| #devops|
  |**SLACK_HOOK**|无|必选, Slack webhook 地址**环境变量**| 无|
  |**ENV_NAME**|无|可选, 如果有多个环境可以指定环境名称**环境变量**| 无|

  提醒消息:
  ![提醒消息](https://raw.githubusercontent.com/nnsay/gist/main/imgimgimage-20230425160005352.png)

# AWS 权限

本工具引用了 aws sdk 所以权限上依赖 sdk 自己的设置, 根据文档支持: AWS\_\*环境变量和配置文件, 关于这块的配置请查看 AWS 文档: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
