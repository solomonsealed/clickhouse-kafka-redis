create table regional_activity
(
    identity_id String,
    country String,
    events_count UInt32
) engine = MergeTree()
order by identity_id;

create table kafka_regional_activity
(
    identity_id String,
    country String,
    events_count UInt32
) engine = Kafka
settings
    kafka_broker_list = 'kafka:9092',
    kafka_topic_list = 'clickhouse_topic',
    kafka_group_name = 'clickhouse_group',
    kafka_format = 'JSONEachRow',
    kafka_num_consumers = 1;

create materialized view to_kafka_regional_activity
to kafka_regional_activity as
select * from regional_activity;

insert into regional_activity (identity_id, country, events_count) values ('abc', 'US', 12), ('abc', 'CA', 2), ('xyz', 'US', 3);
insert into regional_activity (identity_id, country, events_count) values ('xyz', 'CA', 4);
insert into regional_activity (identity_id, country, events_count) values ('xyz', 'MX', 1);
insert into regional_activity (identity_id, country, events_count) values ('abc', 'US', 13), ('abc', 'CA', 3), ('xyz', 'US', 4), ('xyz', 'CA', 5), ('xyz', 'MX', 2), ('efg', 'MX', 1), ('efg', 'CN', 1), ('efg', 'US', 1), ('efg', 'CA', 1), ('efg', 'UK', 1);