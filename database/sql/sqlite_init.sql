-- 用户表
CREATE TABLE IF NOT EXISTS users
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    user_account  CHAR(64)   NOT NULL UNIQUE,                    -- 账号
    user_password CHAR(128)  NOT NULL,                           -- 密码
    union_id      CHAR(128)  NULL,                               -- 微信开发平台ID
    open_id       CHAR(128)  NULL,                               -- 公众号ID
    user_name     CHAR(64)   NULL,                               -- 用户昵称
    user_avatar   CHAR(1024) NULL,                               -- 用户头像
    user_profile  CHAR(512)  NULL,                               -- 用户简介
    user_role     CHAR(64)   NULL,                               -- 用户角色
    is_deleted    BOOLEAN    NOT NULL DEFAULT 0,                 -- 是否删除，1：删除；0：正常
    create_time   DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time   DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP  -- 更新时间
);

-- 创建账号的索引
CREATE INDEX idx_user_account ON users (user_account);


-- 由于sqlite数据表不能增加字段注释，所以需要创建字段注释表，来维护字段描述
CREATE TABLE IF NOT EXISTS field_comments
(
    table_name CHAR(128) NOT NULL,       -- 表名
    field_name CHAR(128) NOT NULL,       -- 字段名
    comment    CHAR(512),                -- 字段描述
    PRIMARY KEY (table_name, field_name) -- 主键
);

PRAGMA encoding = 'UTF-8';

-- 插入字段注释
INSERT INTO field_comments (table_name, field_name, comment)
VALUES ('users', 'id', '用户ID，自增主键'),
       ('users', 'user_account', '账号，唯一且不能为空，不能包含中文字符'),
       ('users', 'user_password', '密码，不能为空'),
       ('users', 'union_id', '微信开发平台ID'),
       ('users', 'open_id', '公众号ID'),
       ('users', 'user_name', '用户昵称'),
       ('users', 'user_avatar', '用户头像'),
       ('users', 'user_profile', '用户简介'),
       ('users', 'user_role', '用户角色'),
       ('users', 'is_deleted', '是否删除，布尔值，默认值为0（未删除）'),
       ('users', 'create_time', '创建时间，默认值为当前时间'),
       ('users', 'update_time', '更新时间，默认值为当前时间');

