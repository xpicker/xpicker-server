# API 协议
$ip = www.corner.today:10086
## Test Group
### 连接测试
    请求 GET:
    url: $ip/api/test/ping
    响应:
    json:
        {"message": "pong"}
### 登录测试账号
    请求 GET：
    url: $ip/api/test/login
    响应：
    json:
        {
            "id": user.Id,
            "password": user.Password,
            "username": user.Username,
            "register_time": user.RegisterTime,
            "last_time": user.LastTime,
            "email": user.Email,
            "mobile": user.Mobile,
            "type": user.Type,
        }
    cookie:
        {"LOGIN_" + username: HASH}
### 登录测试
    请求 POST：
    url: $ip/api/test/login
    postform:
    {
        "username": username | "email": email
        "password": password
    }
    响应：
    json:
        {
            "id": user.Id,
            "username": user.Username,
            "register_time": user.RegisterTime,
            "last_time": user.LastTime,
            "email": user.Email,
            "mobile": user.Mobile,
            "type": user.Type,
        }
    cookie:
        {"LOGIN_" + username: HASH}
