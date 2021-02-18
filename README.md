# gmock

A sidecar use for mock test

## Configure

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