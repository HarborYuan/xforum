userdata.db

```
-- 用户id
-- 用户名
-- 密码
-- 电子邮件
-- 电话号码 13位数字
-- 出生日期 YYYY-MM-DD
-- 性别 M/F
-- 创建时间 YYYY-MM-DD HH:MM:SS
-- 权限* 待定
CREATE TABLE userinfo (
    uid INTEGER PRIMARY KEY NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    birthdate TEXT,
    createtime TEXT NOT NULL,
    gender TEXT);
```
