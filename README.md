# WIP: Cars trace microservice

### Data example

| carId  |   latitude |  longitude | operator | stopped | lineId |
|--------|-----------:|-----------:|:--------:|:-------:|:------:|
| c7831s | -17.167034 | -40.390721 |   o152   |    0    |  l171  |
| c7832s | -17.509504 | -40.491718 |   o252   |    0    |  l172  |
| c7833s | -17.167034 | -40.390721 |   o352   |    1    |  l173  |


### Database (Postgres)

Using command line (may need admin privileges) run the follow commands:

1. Create volume

```shell
docker volume create --name=postgresdata
```

2. Pull/Run postgres pointing to the volume created before and exposing 5432 port

```shell
docker run --name postgres-db -v postgresdata:/data/db -e POSTGRES_PASSWORD=developer -p 5432:5432 -d postgres
```

3. (Optional) Create database

```sql
CREATE DATABASE cars;
GRANT ALL PRIVILEGES ON DATABASE cars TO postgres;
```

#### Requirements
Download and install required tools:
- Go 1.19
- Docker (can use [rancher-desktop](https://rancherdesktop.io/))
- IDE
- DBeaver (Optional)
