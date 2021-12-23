## 排序键，分区键，主键

```sh
CREATE TABLE [IF NOT EXISTS] [db.]table_name [ON CLUSTER cluster]
# shellcheck disable=SC1073
(
    name1 [type1] [DEFAULT|MATERIALIZED|ALIAS expr1] [TTL expr1],
    name2 [type2] [DEFAULT|MATERIALIZED|ALIAS expr2] [TTL expr2],
    ...

    INDEX index_name1 expr1 TYPE type1(...) GRANULARITY value1,
    INDEX index_name2 expr2 TYPE type2(...) GRANULARITY value2
)

ENGINE = MergeTree()

# 排序键
# 按 xxx 字段进行排序（必要）
ORDER BY expr

# 分区键
# 按 xxx 进行分区；如：按每个月进行分区；按每个月 + 性别进行分区.....
# 不同分区的数据片段不会进行合并
[PARTITION BY expr]

# 主键
# 主键与排序键的关系：
# 1. 主键，如果不指定的话就是排序键
# 2. 如果跟排序键不一样，此时排序键用于"在数据片段中进行排序"，主键用于"在索引文件中进行标记的写入"
# 3. 这种情况下，主键表达式元组必须是排序键表达式元组的前缀(即主键为(a,b)，排序列必须为(a,b,**))

# 如果当前主键是 (a, b) ，在下列情况下添加另一个 c 列会提升性能：
# 1. 查询会使用 c 列作为条件
# 2. 很长的数据范围（ index_granularity 的数倍）里 (a, b) 都是相同的值，并且这样的情况很普遍。换言之，就是加入另一列后，可以让您的查询略过很长的数据范围。
# 3. 改善数据压缩；ClickHouse 以主键排序片段数据，所以，数据的一致性越高，压缩越好。
# 4. 在 CollapsingMergeTree 和 SummingMergeTree 引擎里进行数据合并时会提供额外的处理逻辑。
# 原因：
# SummingMergeTree 和 AggregatingMergeTree 会对排序键相同的行进行聚合，
# 所以 "把所有的维度" 放进 "排序键" 是很自然的做法。
# 但这将导致"排序键"中包含大量的"列"，并且排序键会伴随着新添加的维度不断的更新。
[PRIMARY KEY expr]

[SAMPLE BY expr]


[TTL expr [DELETE|TO DISK 'xxx'|TO VOLUME 'xxx'], ...]


[SETTINGS name=value, ...]
```

---

## 数据存储

### 表

表由按主键排序的***数据片段（DATA PART）***组成。
当数据被插入到表中时，会创建***多个数据片段***并按***主键***的***字典序***排序。
e.g. 主键是 `(CounterID, Date)` 时，片段中数据首先按 `CounterID` 排序，具有相同 `CounterID` 的部分按 `Date` 排序。

### 分区

指定分区键后形成分区
不同分区的数据会被分成不同的片段，ClickHouse 在后台合并数据片段以便更高效存储。
不同分区的数据片段***不会进行合并***。—— 同一分区的可能有很多个数据片段
合并机制并不保证***具有相同主键***的行全都合并到同一个数据片段中。

### 数据片段

数据片段可以以 `Wide` 或 `Compact` 格式存储。
 `Wide` 格式下，***每一列***都会在文件系统中存储为***单独的文件***。
 `Compact` 格式下***所有列***都存储在***一个文件***中，`Compact` 格式可以提高***插入量少***，***插入频率频繁***时的性能。
数据存储格式由 `min_bytes_for_wide_part` 和 `min_rows_for_wide_part` 表引擎参数控制，如果数据片段中的***字节数***或***行数***少于相应的设置值，数据片段会以 `Compact` 格式存储，否则会以 `Wide` 格式存储。
***同一分区***的各个片段，会在插入后10-15分钟***整合成一个***片段。

### 颗粒

每个数据片段被逻辑的分割成颗粒（granules）。
颗粒是 ClickHouse 中进行数据查询时的***最小不可分割数据集***。ClickHouse 不会对行或值进行拆分，所以每个颗粒总是包含整数个行。
每个颗粒的***第一行***通过该行的***主键值***进行标记，ClickHouse 会为每个***数据片段***创建一个***索引文件***来存储这些标记。
对于***每列***，无论它是否包含在主键当中，ClickHouse ***都会存储类似标记***。这些标记让您可以***在列文件中直接找到数据***。
颗粒的***大小***通过表引擎参数 `index_granularity` 和 `index_granularity_bytes` 控制。
颗粒的行数的在 `[1, index_granularity]` 范围中，这取决于行的大小。如果单行的大小超过了 `index_granularity_bytes` 设置的值，那么一个颗粒的大小会超过 `index_granularity_bytes`。在这种情况下，颗粒的大小等于该行的大小。

## ![image-20211216201728439](/Users/xl/Library/Application Support/typora-user-images/image-20211216201728439.png)

---

## 分片及副本

### 表级副本

在两个节点（实例）上同时建同一张表
填入相同的zk ip，指定相同路径
表引擎选择 `ENGINE = ReplicatedMergeTree('zk_path', 'replica_name')`
这样就创出了带副本的表

### 分片节点

每个实例都可以成为一个分片 shard，或一个副本 replica
在集群配置中指定节点为 分片/副本 关系
分片是为了将数据分散
分片表的引擎可以用任意引擎，但是如果使用非`ReplicatedMergeTree`引擎的话，副本的数据同步需要`Distributed`来负责（写分片和副本），加大了其压力。所以如果使用分片，并且需要副本，推荐使用`ReplicatedMergeTree`引擎，数据同步交由其处理，减轻Distributed压力（需要增加internal_replication参数：<internal_replication>true</internal_rep`lication>）。

分片的好处：解决了单节点达到瓶颈的问题，和通过分布式表`Distributed`引擎能本身实现数据的路由和聚合。
分片分布式表的不足：`Distributed`表在写入时会在本地节点生成临时数据，会产生写放大，所以会对CPU及内存造成一些额外消耗，也会增加merge负担。

```xml
<yandex>
    <ck_remote_servers>
        <test_cluster_1_repl>
            <shard>  -- 创建一个集群，有两个shard，每个shard有两个副本
                <internal_replication>true</internal_replication>  -- 有复制表引擎自己分发同步数据，减少Distributed压力。
                <weight>1</weight>
                <replica>
                    <host>12.16.20.12</host>
                    <port>9000</port>
                    <priority>1</priority>
                </replica>
                <replica>
                    <host>12.16.20.12</host>
                    <port>9010</port>  -- 同节点，不同端口亦可
                    <priority>1</priority>
                </replica>
            </shard>
            <shard>
                <internal_replication>true</internal_replication>
                <weight>1</weight>
                <replica>
                    <host>12.16.20.17</host>
                    <port>9000</port>
                    <priority>1</priority>
                </replica>
                <replica>
                    <host>12.16.20.17</host>
                    <port>9010</port>
                    <priority>1</priority>
                </replica>
            </shard>
        </test_cluster_1_repl>
    </ck_remote_servers>
</yandex>
```

