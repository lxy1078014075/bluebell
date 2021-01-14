-- 为什么不使用自增id作为用户id
-- 1.别人可以通过注册用户来获取到数据库的用户量
-- 2.分库分表的时候用户id可能会重复
-- 使用分布式ID生成器
create table user (
    id bigint(20) not null auto_increment,
    user_id bigint(20) not null ,
    username varchar(64) not null,
    password varchar(64) not null,
    email varchar (64),
    gender tinyint(4) not null default 0,
    create_time timestamp null default current_timestamp ,
    update_time timestamp null  default current_timestamp on update current_timestamp ,
    primary key(id),
    unique key idx_username(username),
    unique key idx_user_id(user_id)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

create table community(
    id int(11) not null auto_increment,
    community_id int(10) unsigned not null ,
    community_name varchar(128) collate utf8mb4_general_ci not null ,
    introduction varchar(256) collate utf8mb4_general_ci not null ,
    create_time timestamp not null default current_timestamp ,
    update_time timestamp not null default current_timestamp on update current_timestamp ,
    primary key (id),
    unique key idx_community_id(community_id),
    unique key idx_community_name(community_name)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

insert into community(community_id,community_name,introduction) values (1,"GO","Golang");
insert into community(community_id,community_name,introduction) values (2,"leetcode","刷题刷题");
insert into community(community_id,community_name,introduction) values (3,"CS:GO","Rush B...");
insert into community(community_id,community_name,introduction) values (4,"LOL","欢迎来到英雄联盟");

create table post(
    id bigint(20) not null auto_increment,
    post_id bigint(20) not null comment '帖子id',
    title varchar(128) collate utf8mb4_general_ci not null comment '标题',
    content varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    author_id bigint(20) not null comment '作者的用户id',
    community_id bigint(10) not null comment '所属社区',
    status tinyint(4) not null default 1 comment '帖子状态',
    create_time timestamp not null default current_timestamp ,
    update_time timestamp not null default current_timestamp on update current_timestamp ,
    primary key (id),
    unique key idx_post_id(post_id),
    key idx_auther_id(author_id),
    key idx_community_id(community_id)
)engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;
