# clickhouse-kafka-redis
Example of how you can create a materialized view which sends the event to Redis on every insert

1. `docker up`
2. `docker ps` (note down the container ID for clickhouse and Redis)
3. `docker exec -it <clickhouse container ID here> clickhouse-client`
4. execute the commands in tables.sql
5. `docker exec -it "redis container ID here" redis-cli`
6. `keys *`
7. `hgetall efg` as you can see all the inserts to Clickhouse are now in Redis