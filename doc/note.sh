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

# 主键与排序键的关系：
# 1. 主键，如果不指定的话就是排序键
# 2. 如果跟排序键不一样，此时排序键用于"在数据片段中进行排序"，主键用于"在索引文件中进行标记的写入"
# 3. 这种情况下，主键表达式元组必须是排序键表达式元组的前缀(即主键为(a,b)，排序列必须为(a,b,**))

# 如果当前主键是 (a, b) ，在下列情况下添加另一个 c 列会提升性能：
# 1. 查询会使用 c 列作为条件
# 2. 很长的数据范围（ index_granularity 的数倍）里 (a, b) 都是相同的值，并且这样的情况很普遍。换言之，就是加入另一列后，可以让您的查询略过很长的数据范围。
# 3. 改善数据压缩；ClickHouse 以主键排序片段数据，所以，数据的一致性越高，压缩越好。
# 4. 在 CollapsingMergeTree 和 SummingMergeTree 引擎里进行数据合并时会提供额外的处理逻辑。
[PRIMARY KEY expr]

[SAMPLE BY expr]


[TTL expr [DELETE|TO DISK 'xxx'|TO VOLUME 'xxx'], ...]


[SETTINGS name=value, ...]

# SummingMergeTree 和 AggregatingMergeTree 会对排序键相同的行进行聚合，
# 所以 "把所有的维度" 放进 "排序键" 是很自然的做法。
# 但这将导致"排序键"中包含大量的"列"，并且排序键会伴随着新添加的维度不断的更新。