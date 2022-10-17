# gRPC plants

CockroachDB single-node database, hosted in Docker. To retreive information within the database for gRPC client calls

---

### Database

Cockroachdb Single node database run in Docker container. Can be executed via the Makefile commands.

Database Name: ***native***

Table(s):

plants (***native.plants***):
```
column_name   | data_type | is_nullable |  column_default   | generation_expression |  indices  | is_hidden
+-------------+-----------+-------------+-------------------+-----------------------+-----------+-----------+
id            | UUID      |    false    | gen_random_uuid() |                       | {primary} |   false    
common_name   | STRING    |    false    | NULL              |                       | {}        |   false   
genus         | STRING    |    false    | NULL              |                       | {}        |   false    
species       | STRING    |    false    | NULL              |                       | {}        |   false    
```

## Makefile

<code>build_protoc</code>      build proto files<br/>
<code>build_server</code>      Builds server in bin/plants<br/>
<code>build_client</code>      Builds client in bin/plants<br/>
<code>docker_up</code>         Start up CockroachDB with docker-compose in detached mode<br/>
<code>docker_down</code>       Shut down CockroachDB with docker-compose<br/>
<code>about</code>             Display info related to the build<br/>

## Additional Commands:

To Get into the interactive terminal within the cockroachdb:

`sudo docker exec -it ${CONTAINER_ID} ./cockroach sql --insecure`

## To Run

Client and server executables will be written to the /bin/plants directory (<code>/bin/plants/(client/server)</code>)

Protocol files will the written into <code>/plants/proto</code> directory.

1) build protobuf stubs

    <code>make build_protoc</code>

2) build server code

    <code>make build_server</code>

3) build client code

    <code>make build_client</code>

4) Start Docker containers

    <code>make docker_up</code>

5) to run the actual code

    <code> ./bin/plants/server </code><br>
    <code> ./bin/plants/client </code>

6) To take docker containers down
    
    <code> make docker_down </code>

***Note:each docker command is prefixed with 'sudo' because the system is assumed to be a Linux distribution***

TODO: ... Create SQL dump 