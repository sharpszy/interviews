# 数据类型
    String、List、Hash、Set，ZSet

# 原理
## hash
   - 使用ziplist场景（可配置）
     - 所有的键值对的健和值的字符串长度都小于等于64byte（一个英文字母一个字节）
     - 哈希对象保存的键值对数量小于512个
   - 有2个hash表，h[0]/h[1]
## rehash
  - 渐进式（扩、缩容）
  - rehashidx， 记录已完成的rehash进度，当为-1时，表示rehash完成
  - 步骤
    > 1. 为 ht[1] 分配指定空间，让字典同时持有 ht[0] 和 ht[1] 两个哈希表
    > 2. 将 rehashidx 设置为0，表示正式开始 rehash
    > 3. 每次对字典执行 增、删、改、查操作时，程序除了执行指定的操作外，还会将 哈希表 ht[0].table中下标为 rehashidx 位置上的所有的键值对 全部迁移到 ht[1].table 上
    > 4. 最后，当 ht[0].used 变为0时，代表所有的键值对都已经从 ht[0] 迁移到 ht[1] 了，释放 ht[0].table， 并且将 ht[0] 设置为 ht[1]，rehashidx 标记为 -1 代表 rehash 结束

## string

## ziplist
  - 使用更加 **紧凑的结构** 实现多个元素的 **连续存储**，所以在 **节省内存** 方面比 hashtable 更加优秀
  - 双向链表


# 集群

## 主从/哨兵

## redis cluster

# 持久化

## AOF

## RDB

# 其他
## 过期逐出策略