# Twitter snowflake ID 算法之 golang 实现
snowflake ID 算法是 twitter 使用的唯一 ID 生成算法，主要解决了高并发时 ID 生成不重复的问题，为了满足 Twitter 每秒上万条消息的请求，使每条消息有唯一、有一定顺序的 ID ，且支持分布式生成。

# 结构
snowflake ID 的结构是一个 64 bit 的 int 型数据。

1 bit：不使用，可以是 1 或 0

41 bit：记录时间戳 (当前时间戳减去用户设置的初始时间，毫秒表示)，可记录最多 69 年的时间戳数据

10 bit：用来记录分布式节点 ID，一般每台机器一个唯一 ID，也可以多进程每个进程一个唯一 ID，最大可部署 1024 个节点

12 bit：序列号，用来记录不同 ID 同一毫秒时的序列号，最多可生成 4096 个序列号

# 用法

```
  go run ./bin/main.go
  
  curl http://localhost:1323/get/10
  
  {
    status: "ok",
    data: [
      30959110987976704,
    ]
  }
```
