-- 用户表
CREATE TABLE IF NOT EXISTS users
(
    id            INT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID，自增主键',
    user_account  VARCHAR(64)  NOT NULL UNIQUE COMMENT '账号，唯一且不能为空',
    user_password VARCHAR(128) NOT NULL COMMENT '密码，不能为空',
    union_id      VARCHAR(128)          DEFAULT NULL COMMENT '微信开发平台ID',
    open_id       VARCHAR(128)          DEFAULT NULL COMMENT '公众号ID',
    user_name     VARCHAR(64)           DEFAULT NULL COMMENT '用户昵称',
    user_avatar   VARCHAR(1024)         DEFAULT NULL COMMENT '用户头像',
    user_profile  VARCHAR(512)          DEFAULT NULL COMMENT '用户简介',
    user_role     VARCHAR(64)           DEFAULT NULL COMMENT '用户角色',
    is_deleted    BOOLEAN      NOT NULL DEFAULT 0 COMMENT '是否删除，1：删除；0：正常',
    create_time   TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time   TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY idx_user_account (user_account) -- 创建唯一索引
) COMMENT '用户表' DEFAULT CHARSET = utf8;