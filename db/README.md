

```
docker run --name example-driver-db \
           -p 5432:5432 \
           -e POSTGRES_USER=cycle \
           -e POSTGRES_PASSWORD=labs1 \
           -e POSTGRES_DB=accountdb \
           -v $PWD/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
