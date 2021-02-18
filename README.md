# gmock
[![Build Status](https://travis-ci.com/andy-zhangtao/gmock.svg?branch=main)](https://travis-ci.com/andy-zhangtao/gmock)
[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/dashboard?id=andy-zhangtao_gmock)


A sidecar use for mock test

## Run

`gmock`启动后加载配置文件，然后当请求uri与配置文件中设定的uri相同时，返回配置文件中设定好的数据。

## Configure

`gmock`启动时通过`CONF_PATH`读取配置文件内容，配置文件格式如下：

```toml
[[conf]]
uri = "接口url"
# 返回响应码
status = 200
[conf.header]
# 返回的响应header
content-type = "application/json"
cookie = "a unit test"

[conf.body]
# 设定返回body数据格式
# 如果是string，则data是一个字符串

type = "string"
data = "sample string body"

# 如果是json，则data应该是一个map格式
type = "json"
[conf.body.data]
result = "a json body"
```