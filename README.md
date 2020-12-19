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
  ```
