## Dependencies:
Standalone tailwindcss
example uses linux, replace with your platforms (https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.6)
```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.6/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
```
Air - for live reloading
```
go install github.com/air-verse/air@latest
```


## Environment Variables

Set the following environment variables inside the `.envrc` folder located in the root of your project:

- `PG_URI`
- `DBNAME`
- `USER`

`PG_URI`, see the [pgxpool documentation](https://pkg.go.dev/github.com/jackc/pgx/v5@v5.6.0/pgxpool#ParseConfig).

#### Example `PG_URI`

```plaintext
user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
```

Docker secret:
- `db_password.txt`

## Running
Bring up the database
```
make local
```
Migrate the database
```
tern migrate
```
Insert video/user data into database for testing:
```
make gendb
```
Bring up the website with live reloading
```
air
```
Bring down the database and remove all data:
```
make down-local
```
