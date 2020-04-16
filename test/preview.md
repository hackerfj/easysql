## 用户管理

#### createTableGoods
```sql
    CREATE TABLE `goods` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商品编号',
      
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品名称',
      `cover` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '封面',



      `stock` int(10) DEFAULT '0' COMMENT '库存数量',
      `price` decimal(10,2) DEFAULT NULL COMMENT '商品价格',
      `create_time` datetime DEFAULT NULL COMMENT '创建时间',
      `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
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
#### getCount
```sql
select count(id)as goodsCount from goods
```
#### txGetInfo
```sql
select * from goods where stock > 0 and id = ? for update
```
#### txUpdateInfo
```sql
update goods set stock = ? where id = ?
```