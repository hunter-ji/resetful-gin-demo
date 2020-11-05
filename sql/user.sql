-- 用户表
create table user
(
    user_id      int auto_increment comment '用户id'
        primary key,
    username     varchar(10)                          not null comment '用户名',
    password     varchar(30)                          not null comment '密码',
    phone_number varchar(20)                          null comment '手机号码',
    created_at   timestamp  default CURRENT_TIMESTAMP null comment '创建时间戳',
    updated_at   timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间戳',
    is_deleted   tinyint(1) default 0                 not null comment '是否被删除,0:未删,1:已删'
);

