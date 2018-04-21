# API 协议
$ip = www.corner.today:10086
## Test Group
### 连接测试
    请求:
    url: $ip/api/test/ping
    响应:
    json:
        {"message": "pong"}
