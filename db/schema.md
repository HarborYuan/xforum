userdata.db

```Sqlite
-- 用户id
-- 用户名
-- 密码
-- 电子邮件
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

```Sqlite
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

```Sqlite
---板块id
---板块名称
---板块路径

CREATE TABLE boards
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT NOT NULL,
    path       TEXT NOT NULL UNIQUE
);
```

```Sqlite
---回复id
---用户id
---帖子id
---创建时间
---内容
CREATE TABLE response
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    uid        INTEGER,
    pid        INTEGER,
    createtime TEXT NOT NULL,
    content    TEXT NOT NULL
);
```


```Sqlite
---私信id
---发送者id
---接收者id
---时间
---内容
CREATE TABLE  message
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    sender      INTEGER,
    sendee      INTEGER,
    createtime  TEXT NOT NULL,
    content     TEXT NOT NULL 
);
```