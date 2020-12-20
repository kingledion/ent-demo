# ent-demo
Demo of graph database with ent

## Setup

Clone repository in GOPATH. Dependencies should install via gomod, but you may wish to vendor them explicitly with `go mod vendor` depending on your GOFLAGS setup.

Initialize a local mysql database and grant a user access to it.

Enter user information in `internal/config/config.yaml`, in the following format:
```
database:
  user: dbuser
  pass: dbpass
  port: 3306
  dbname: ent_demo
http:
  port: 8080

  ```

## Running

Run the executable with `go run main.go`

To create an order, POST to `localhost:8080/v1/order` with

```
{
    "user": {
        "uuid": "123-456",
        "firstname": "Bilbo",
        "lastname": "Baggins"
    },
    "merchant": {
        "uuid": "980-733",
        "dba": "Bag End Pipes & Weed"
    }
}
```

Retrieve a recommendation by GET to `localhost:8080/v1/123-456/recommend`
