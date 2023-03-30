# service-graph

## Running service

To run this service export two env vars, and start Neo4j Instance:

```bash
# HTTP PORT (for health checks)
export PORT=8000
# GRPC PORT
export GRPC_PORT=50051
```

then you can run start service by using this command:

```bash
go run main.go --neo4j=<NEO4J_URI> --neo4jUsername=<NEO4J_USERNAME> --neo4jPassword=<NEO4J_PASSWORD>
```

You can either run instance of neo4j locally or use docker.

```bash
docker run \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data \
    --name=knowledge_graph
    neo4j
```

## TODO

I'm trying to avoid premature optimization; but if the performance is not sufficiant, we can improve it by implementing at least one of the following:

* [ ] In memory LFU cache with a cap (env var).
* [ ] Multi-threaded permutations builder.
* [ ] using a memo (dynamic programming) internally inside `GraphService.MedecinePermutation` to avoid overlapped computation.
