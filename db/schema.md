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
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    birthdate TEXT,
    createtime TEXT NOT NULL,
    gender TEXT);
```

```
--帖子id
--用户id
--创建时间
--内容
--路径

CREATE TABLE posts
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    uid        INTEGER,
    createtime TEXT NOT NULL,
    content    TEXT NOT NULL,
    path       TEXT NOT NULL
);
```