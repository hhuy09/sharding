## Set up Sharding using Docker Containers

### Config servers
Start config servers (3 member replica set)
```
docker-compose -f config/docker-compose.yaml up -d
```
Initiate replica set
```
mongosh mongodb://172.31.96.1:10001
```
```
rs.initiate(
  {
    _id: "cfgrs",
    configsvr: true,
    members: [
      { _id : 0, host : "172.31.96.1:10001" },
      { _id : 1, host : "172.31.96.1:10002" },
      { _id : 2, host : "172.31.96.1:10003" }
    ]
  }
)

rs.status()
```

### Shard servers
Start shard server
```
docker-compose -f shard/docker-compose.yaml up -d
```
Initiate replica set
```
mongosh mongodb://172.31.96.1:20001
```
```
rs.initiate(
  {
    _id: "shard1rs",
    members: [
      { _id : 0, host : "172.31.96.1:20001" },
      { _id : 1, host : "172.31.96.1:20002" },
      { _id : 2, host : "172.31.96.1:20003" }
    ]
  }
)

rs.status()
```

### Mongos Router
Start mongos query router
```
docker-compose -f mongos/docker-compose.yaml up -d
```

### Add shard to the cluster
Connect to mongos
```
mongosh mongodb://172.31.96.1:30000
```
Add shard
```
mongos> sh.addShard("shard1rs/172.31.96.1:20001,172.31.96.1:20002,172.31.96.1:20003")
mongos> sh.status()
```

```
mongos> sh.enableSharding("testdb")
mongos> sh.shardCollection("testdb.users", { key: 1 } )
mongos> sh.shardCollection("testdb.users", { key: "hashed" } )
```
