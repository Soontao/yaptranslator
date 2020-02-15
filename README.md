# Yet Another Properties (file) Translator

[![CircleCI](https://circleci.com/gh/Soontao/yaptranslator.svg?style=shield)](https://circleci.com/gh/Soontao/yaptranslator)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/Soontao/yaptranslator.svg)](https://github.com/Soontao/yaptranslator/releases)
[![](https://godoc.org/github.com/Soontao/yaptranslator?status.svg)](http://godoc.org/github.com/Soontao/yaptranslator)

Automated translate properties files by machine translation providers.

```properties
# origin en example.2.properties
hello=hello
world=World
a.b.c.d=Hello World
# translated zh example.2_zh.properties
world = 世界
a.b.c.d = 你好世界
hello = 你好
```

## Install

Install by `go get` or download binary from [release page](https://github.com/Soontao/yaptranslator/releases)

## Usage

```bash
./cli --help
NAME:
   Yet Another Properties (file) Translator - A Command Line Tool

USAGE:
   cli [global options] command [command options] [arguments...]

VERSION:
   SNAPSHOT

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value      file or directory to be processed (default: ".")
   --lang value, -l value      Source Language (default: "en") [$YAPT_SOURCE_LANG]
   --target value, -t value    Target Language (default: "zh") [$YAPT_TARGET_LANG]
   --service value, -s value   Service Provider for YAPT Tool (default: "ALICLOUD") [$YAPT_SERVICE_PROVIDER]
   --username value, -u value  Auth User [$YAPT_AUTH_USER]
   --password value, -p value  Auth User Password [$YAPT_AUTH_PASSWORD]
   --region value, -r value    AUTH Region [$YAPT_REGION]
   --key value, -k value       Auth Access Key [$YAPT_ACCESS_KEY]
   --secret value              Auth Access Secret [$YAPT_ACCESS_SECRET]
   --token value               Auth Token [$YAPT_AUTH_TOKEN]
   --help, -h                  show help
   --version, -v               print the version
```

## Providers

### Ali Cloud Translator

Translator powered by Alibaba Cloud

- required options:
    - region (Service region, like `cn-hangzhou`)
    - key (Access Key ID)
    - secret (Access Key Secret)


## [CHANGELOG](./CHANGELOG.md)

## [LICENSE](./LICENSE)