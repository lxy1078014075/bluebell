# 分区
1. MySQL的分区是局部分区索引，一个分区中既存了数据，又放了索引。每个区的聚集索引和非聚集索引都放在各自区(不同的物理文件)。
2. 如果表存在主键或者唯一索引时，分区列必须是唯一索引的一个组成部分。
3. 分区类型：range、list、hash、key

# 分表
## 纵向分表的好处
1. 大表拆小表，更便于开发与维护，也能避免跨页问题。跨页会造成额外的开销。
2. 数据库以行为单位将数据加载到内存中，这样表中字段长度越短且访问频次较高，内存能加载更多的数据，命中率更高，减少磁盘IO，从而提升数据库性能。