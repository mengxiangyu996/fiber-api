create table if not exists `admin` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `username` varchar(128) comment '用户名',
    `nickname` varchar(128) comment '昵称',
    `gender` tinyint(3) comment '性别：1-男；2-女',
    `email` varchar(128) comment '邮箱',
    `phone` varchar(128) comment '电话',
    `password` varchar(256) comment '密码',
    `avatar` varchar(256) comment '头像',
    `status` tinyint(3) default 1 comment '状态：1-启用；2-禁用'
) comment '管理员表';

create table if not exists `role` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `name` varchar(128) comment '角色名',
    `status` int(3) default 1 comment '状态：1-启用；2-禁用'
) comment '角色表';

create table if not exists `menu` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `parent_id` int(10) comment '父级id',
    `name` varchar(128) comment '菜单名',
    `type` tinyint(3) comment '类型：1-目录；2-菜单；3-按钮',
    `sort` int(10) default 0 comment '排序',
    `path` varchar(128) comment '路由地址',
    `component` varchar(128) comment '组件',
    `icon` varchar(128) comment '图标',
    `redirect` varchar(128) comment '重定向',
    `status` tinyint(3) default 1 comment '状态：1-启用；2-禁用'
) comment '菜单表';

create table if not exists `permission` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `name` varchar(128) comment '权限名',
    `group_name` varchar(128) comment '组名',
    `path` varchar(128) default '1' comment '权限地址',
    `method` varchar(128) comment '请求方式',
    `status` tinyint(3) default 1 comment '状态：1-启用；2-禁用'
) comment '权限表';

create table if not exists `admin_role_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `admin_id` int(10) comment '管理员id',
    `role_id` int(10) comment '角色id'
) comment '用户角色关系表';

create table if not exists `role_menu_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `role_id` int(10) comment '角色id',
    `menu_id` int(10) comment '菜单id'
) comment '角色菜单关系表';

create table if not exists `role_permission_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `role_id` int(10) comment '角色id',
    `permission_id` int(10) comment '权限id'
) comment '角色权限关系表';
