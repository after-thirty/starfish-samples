# Nacos 注册中心版本
## 一. 启动 MYSQL
```bash
$ docker run -itd --name starfish-mysql -p 3306:3306 -e MYSQL_DATABASE=starfish -e MYSQL_ROOT_PASSWORD=123456 mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## 二. 导入 MYSQL脚本
通过数据库连接工具，如DBeaver 导入 starfish仓库
```sql
-- starfish.branch_table definition

CREATE TABLE `branch_table` (
  `branch_id` bigint NOT NULL,
  `xid` varchar(128) CHARACTER SET utf8 NOT NULL,
  `transaction_id` bigint DEFAULT NULL,
  `resource_group_id` varchar(32) CHARACTER SET utf8 DEFAULT NULL,
  `resource_id` varchar(256) CHARACTER SET utf8 DEFAULT NULL,
  `branch_type` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `client_id` varchar(64) CHARACTER SET utf8 DEFAULT NULL,
  `application_data` varchar(2000) CHARACTER SET utf8 DEFAULT NULL,
  `gmt_create` datetime(6) DEFAULT NULL,
  `gmt_modified` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`branch_id`),
  KEY `idx_xid` (`xid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- starfish.global_table definition

CREATE TABLE `global_table` (
  `xid` varchar(128) CHARACTER SET utf8 NOT NULL,
  `transaction_id` bigint DEFAULT NULL,
  `status` tinyint NOT NULL,
  `application_id` varchar(32) CHARACTER SET utf8 DEFAULT NULL,
  `transaction_service_group` varchar(32) CHARACTER SET utf8 DEFAULT NULL,
  `transaction_name` varchar(128) CHARACTER SET utf8 DEFAULT NULL,
  `timeout` int DEFAULT NULL,
  `begin_time` bigint DEFAULT NULL,
  `application_data` varchar(2000) CHARACTER SET utf8 DEFAULT NULL,
  `gmt_create` datetime DEFAULT NULL,
  `gmt_modified` datetime DEFAULT NULL,
  PRIMARY KEY (`xid`),
  KEY `idx_gmt_modified_status` (`gmt_modified`,`status`),
  KEY `idx_transaction_id` (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- starfish.lock_table definition

CREATE TABLE `lock_table` (
  `row_key` varchar(128) CHARACTER SET utf8 NOT NULL,
  `xid` varchar(96) CHARACTER SET utf8 DEFAULT NULL,
  `transaction_id` bigint DEFAULT NULL,
  `branch_id` bigint NOT NULL,
  `resource_id` varchar(256) CHARACTER SET utf8 DEFAULT NULL,
  `table_name` varchar(32) CHARACTER SET utf8 DEFAULT NULL,
  `pk` varchar(36) CHARACTER SET utf8 DEFAULT NULL,
  `gmt_create` datetime DEFAULT NULL,
  `gmt_modified` datetime DEFAULT NULL,
  PRIMARY KEY (`row_key`),
  KEY `idx_branch_id` (`branch_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
## 启动Naocs

## 编译启动 server
```bash
$ ./tc start -config ../profiles/dev/config.yml   
```
可以看到输出
```
2022-03-17T10:30:12.426+0800    DEBUG   server/server.go:86     s bind addr{:8091} ok!
2022-03-17T10:30:12.428+0800    INFO    file/registry.go:24     file register
```

