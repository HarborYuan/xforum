userdata.db

```
-- 用户id 固定8字节大小
-- 用户名
-- 密码
-- 电子邮件
-- 电话号码 13位数字
-- 出生日期 XXXX-XX-XX
-- 性别 M/F
-- 权限* 待定
CREATE TABLE userinfo (
    uid INTEGER PRIMARY KEY,
    username TEXT,
    password TEXT,
    email TEXT,
    phone TEXT,
    birthdate TEXT,
    gender TEXT);
```
