## This service require PostgreSQL database of version *9.5* or higher

User keys and secrets are symmetrically encrypted so that optional "SECRET" environment variable can be provided (otherwise default one would be used).
SECRET should be a random string of 16, 24 or 32 characters 

## You can "run.sh" shell script to execute full lifecycle build/test/run/cleanup

```shell
 ./run.sh
```

## Alternatively each step could be executed separately

### Build test image and execute unit tests

```shell
docker build --target builder -t blocksize-assignment-test .
docker run --rm --network bsm  -e "POSTGRES_URL=postgres://postgres:test@postgres_host/blocksize?sslmode=disable&pool_max_conns=10" blocksize-assignment-test
```

### Build and run service

```shell
docker build --target app -t blocksize-assignment .
docker run --rm -p 50051:50051 -e "POSTGRES_URL=postgres://postgres:test@postgres_host/blocksize?sslmode=disable&pool_max_conns=10" blocksize-assignment
```

## What is left out of scope... In a real-world scenario I wouldn't go to production without:
1. Proper Logging
2. Performance metrics
3. Reliable database migrations
