# schutterij
Backend and middleware component

## Local
### Build and run
```
go build -o ./build ./cmd/main.go
./build/main
```

### Testing with Postgres
```
# Create a postgres instance
docker run --rm -d -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=pgpassword \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /var/lib/postgresql/data:/var/lib/postgresql/data \
    postgres:13.0-alpine

# Get shell into the Postgres container
docker exec -it <mycontainer> bash
docker exec -it fde9238d0681 bash

# Start psql
psql -U postgres

# Create postgrest
docker run --rm --net=host -p 3000:3000 \
  -e PGRST_DB_URI="postgres://postgres:pgpassword@localhost:5432/postgres" \
  -e PGRST_DB_ANON_ROLE="none" \
  -e PGRST_DB_SCHEMA="api" \
  -e PGRST_JWT_SECRET="bh3lfEY6f0hQ7TxHv0n8zj6s76ubN1hK" \
  postgrest/postgrest:v7.0.1

go build -o ./build ./cmd/main.go
./build/main

sudo mv exampleproject-config ~/.kube/config
kubectl create -f test/test-deployment.yaml
```

## Production
### Build and push container image
```
docker build -t jordan396/schutterij .
docker push jordan396/schutterij
docker run jordan396/schutterij
```