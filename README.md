# rinha-backend-2024


## ScyllDB

- [x] Deploy
- [x] CQLSH
- [X] Keyspace
- [X] Table
- [X] Insert
- [X] Select

### Deploy

```bash
docker-compose up

docker exec -it scylla-node1 nodetool status
```

### CQLSH

```bash
docker exec -it scylla-node1 cqlsh
```

### Keyspace

```sql
DROP KEYSPACE rinha_db;

CREATE KEYSPACE rinha_db
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
       
USE rinha_db;
```
### Table

```sql

CREATE TABLE conta_lock (
   id INT,
   time TIMESTAMP,
   primary key (id)
) WITH default_time_to_live = 3;

CREATE TABLE conta (
    id INT,
    limite INT,
    saldo_inicial INT,
    version INT,
    primary key (id)                   
);

```

### Insert

```sql
INSERT INTO conta (id, limite, saldo_inicial, version) VALUES (1, 100000, 0, 1);
INSERT INTO conta (id, limite, saldo_inicial, version) VALUES (2, 80000, 0, 1);
INSERT INTO conta (id, limite, saldo_inicial, version) VALUES (3, 1000000, 0, 1);
INSERT INTO conta (id, limite, saldo_inicial, version) VALUES (4, 10000000, 0, 1);
INSERT INTO conta (id, limite, saldo_inicial, version) VALUES (5, 500000, 0, 1);


INSERT INTO conta_lock (id, time) VALUES (1, toTimestamp(now())) IF NOT EXISTS; 
```

### Select*

```sql
SELECT * FROM conta where id = 1;

SELECT * FROM conta;
```


## Api

- [x] config fiber rota clientes transacoes 
- [x] Usecase clientes transacoes 
- [x] Repository clientes transacoes scyllaDB
