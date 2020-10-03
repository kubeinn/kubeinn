# schutterij
Backend and middleware component

## Local
### Build
```
go build -o ./build ./cmd/main.go
```
### Run
```
./build/main.go
```

### Testing with Postgres
```
# Create a postgres instance
docker run -d -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=pgpassword \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /var/lib/postgresql/data:/var/lib/postgresql/data \
    postgres:13.0-alpine

# Get shell into the Postgres container
docker exec -it <mycontainer> bash
psql -U postgres
```

## Production
### Build and push container image
```
docker build -t jordan396/schutterij .
docker push jordan396/schutterij
docker run jordan396/schutterij
```