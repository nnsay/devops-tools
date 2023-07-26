# 1. Devops-TOOLS

本项目编写一些日常运维的小工具, 本项目使用`spf13/cobra`为命令行开发开发, 目前支持的工具如下

- AWS 类
- Monorepo 类

# 2. AWS 类

该类型的工具需要具有 aws 访问权限, 如果需要发送通知, 还需要配置 Slack:

- AWS

  本工具引用了 aws sdk 所以权限上依赖 sdk 自己的设置, 根据文档支持: AWS\_\*环境变量和配置文件, 关于这块的配置请查看 AWS 文档: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html

- Slack

  Slack 消息是通过安装自定义 Slack 应用, 然后利用该应用发送消息的途径, 而发消息需要`calls:write`权限, 详情参考 [Using the Calls API](https://api.slack.com/apis/calls); 之所以使用该种方式而不是[Incoming Webhooks](https://api.slack.com/apps/A052AEV2S68/incoming-webhooks)的原因是 webhook 不支持指定 Channel

## 2.1 删除过期的 IAM 证书

作用: 删除已经过期的 IAM 证书

使用方法:

```bash
# use alias command
devops-tools iam dec
# use fullname command
devops-tools iam delete-expired-certification
```

参数说明:

| 参数名称    | 别名 | 描述                       | 默认值       |
| ----------- | ---- | -------------------------- | ------------ |
| expiration  | e    | 指定过期时间, 格式为时间戳 | time.Now()   |
| path-prefix | p    | 证书路径                   | /cloudfront/ |

## 2.2 检查临期 IAM 证书

作用: 检查即将过期的证书并提醒, 提醒消息发送到 Slack(需配置 **SLACK_HOOK**)

使用方法:

```bash
# use alias command
devops-tools iam ccd
# use fullname command
devops-tools iam check-certification-date
```

参数说明:

| 参数名称        | 别名 | 描述                                               | 默认值       |
| --------------- | ---- | -------------------------------------------------- | ------------ |
| expire-hours    | e    | 剩余过期小时数                                     | 72           |
| path-prefix     | p    | 证书路径                                           | /cloudfront/ |
| channel         | c    | Slack Channel                                      | #devops      |
| **SLACK_TOKEN** | 无   | 必选, Slack 自定义应用 Auto Token 地址**环境变量** | 无           |
| **ENV_NAME**    | 无   | 可选, 如果有多个环境可以指定环境名称**环境变量**   | 无           |

提醒消息:
![提醒消息](https://raw.githubusercontent.com/nnsay/gist/main/img20230629183823.png)

## 2.3 检查未变更的 Cloudformation

作用: 检查超过指定天数的未变更的 Cloudformation Stack 并发送提醒消息

使用方法:

```bash
# use alias command
devops-tools cloudformation cec
# use fullname command
devops-tools cloudformation checkExpirationCloudformation
```

参数说明:

| 参数名称              | 别名 | 描述                                               | 默认值  |
| --------------------- | ---- | -------------------------------------------------- | ------- |
| days                  | d    | 多少天未更新                                       | 10      |
| channel               | c    | Slack Channel                                      | #devops |
| **SLACK_TOKEN**       | 无   | 必选, Slack 自定义应用 Auto Token 地址**环境变量** | 无      |
| **ENV_NAME**          | 无   | 可选, 如果有多个环境可以指定环境名称**环境变量**   | 无      |
| **WHITE_STACK_NAMES** | 无   | 可选, stack 白名单,多个以逗号分隔**环境变量**      | 无      |

提醒消息:
![提醒消息](https://raw.githubusercontent.com/nnsay/gist/main/img20230630104222.png)

# 3. Monorepo 类

Monorepo 基于[Nx](https://nx.dev/)的实践, 不过该类工具设计时与具体哪种 Monorepo 无关, 主要是解决 Monorepo 中的痛点问题.

## 3.1 代码覆盖率报告

依赖前提:

测试覆盖率报告需要是`json-summary`格式, Istanbul 是事实上的代码测试覆盖率标准, 其支持产生的代码覆盖率报告格式有很多种, 常见的如 json, json-summary, text, lcov 等, 详情可以查看[这里](https://istanbul.js.org/docs/advanced/alternative-reporters/)

作用:

- 基于 Monorepo 的多项目的代码覆盖率报告, 覆盖率报告以项目分组
- 支持检测代码覆盖率阈值检查, 目前检查的是 statement 指标

使用方法:

```bash
# use alias command
devops-tools monorepo ccr
# use fullname command
devops-tools monorepo codeCoverageReport
```

参数说明:

| 参数名称    | 别名 | 描述         | 默认值 |
| ----------- | ---- | ------------ | ------ |
| coverageDir | d    | 多少天未更新 | 10     |
| limitTarget | l    | 多少天未更新 | 10     |
| reportPath  | r    | 多少天未更新 | 10     |

扩展使用:

该工具可以结合 gh 一起使用, 可以在流水线中显示覆盖率报告或者让低覆盖流水任务失败, 这部分技巧可以参考: [3. 配合 Github Workflow 使用](https://nnsay.cn/2023/07/17/code-coverage/#3-%E9%85%8D%E5%90%88-Github-Workflow-%E4%BD%BF%E7%94%A8)
