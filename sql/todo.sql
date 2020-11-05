-- todo表
create table todo
(
    todo_id    int auto_increment comment 'todo id'
        primary key,
    title      varchar(30)                          not null comment 'todo标题',
    user_id    int                                  not null comment 'todo所属用户id',
    created_at timestamp  default CURRENT_TIMESTAMP null comment '创建时间戳',
    updated_at timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间戳',
    is_deleted tinyint(1) default 0                 not null comment '是否被删除,0:未删,1:已删'
);