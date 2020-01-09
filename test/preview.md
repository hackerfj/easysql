## 用户管理

#### createTableGoods
```sql
CREATE TABLE `goods` (
	`id` BIGINT (20) NOT NULL AUTO_INCREMENT COMMENT '商品编号',
	`name` VARCHAR (255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品名称',
	`cover` VARCHAR (1000) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '封面',
	`stock` INT (10) DEFAULT NULL COMMENT '库存数量',
	`price` DECIMAL (10, 2) DEFAULT NULL COMMENT '商品价格',
	`create_time` datetime DEFAULT NULL COMMENT '创建时间',
	`update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;
```
#### addGoods
```sql
INSERT INTO goods (`name`,`price`) VALUES ('飞行棋',99.99)
```
#### updateGoods
```sql
UPDATE goods SET `name`='飞行棋',`price`=99.88 WHERE `id`= ?
```
#### getOne
```sql
select * from goods limit ?
```
#### findAll
```sql
select * from goods
```
#### deleteById
```sql
delete from goods where id = ?
```