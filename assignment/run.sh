docker network create bsm
docker build --target builder -t blocksize-assignment-test .
docker run -d --network bsm -e "POSTGRES_PASSWORD=test" -e "POSTGRES_DB=blocksize" --name bs_postgres_test postgres:alpine
docker run --rm --network bsm  -e "POSTGRES_URL=postgres://postgres:test@bs_postgres_test/blocksize?sslmode=disable&pool_max_conns=10" blocksize-assignment-test
docker stop bs_postgres_test
docker rm bs_postgres_test

docker build --target app -t blocksize-assignment .
docker run -d --network bsm -e "POSTGRES_PASSWORD=test" -e "POSTGRES_DB=blocksize" --name bs_postgres postgres:alpine
docker run --rm -p 50051:50051 --network bsm -e "POSTGRES_URL=postgres://postgres:test@bs_postgres/blocksize?sslmode=disable&pool_max_conns=10" blocksize-assignment
docker stop bs_postgres
docker rm bs_postgres
docker network rm bsm
docker rmi blocksize-assignment