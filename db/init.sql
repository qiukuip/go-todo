drop table if exists my_demo.todo;

create table my_demo.todo
(
    todo_id     int auto_increment comment '主键id',
    content     varchar(255) not null comment '待办事项内容',
    category    varchar(50)  not null comment '分类',
    is_complete char(1)      not null comment '完成标识 Y-已完成 N-未完成',
    deadline    datetime comment '过期时间',
    create_at   timestamp default current_timestamp,
    update_at   timestamp default current_timestamp on update current_timestamp,
    primary key (todo_id)
) engine = InnoDB
  charset UTF8MB4;

