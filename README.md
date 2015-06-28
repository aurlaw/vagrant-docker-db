# vagrant-docker-db
Sandbox for experimenting with RethinkDB, MongoDB and PostgreSQL.

Uses Vagrant and Docker along with examples written in Go.


Vagrant and Docker setup based on [https://github.com/DevBandit/vagrant-docker](https://github.com/DevBandit/vagrant-docker)

## Install

First, install the [vagrant-docker-compose](https://github.com/leighmcculloch/vagrant-docker-compose) plugin (as root).

```bash
vagrant plugin install vagrant-docker-compose
```

Then run `vagrant up`.


## Usage

VM uses private ip 192.168.20.100

# Go App
```
export GOPATH=/path/to/this/repo/app
```

# PostgreSQL
PG Studio: [http://192.168.20.100:8080/](http://192.168.20.100:8080/)
Database host: 192.168.20.100
Port: 5432
Username: postgres
Password: postgres

# Mongo
TODO

# RethinkDB
TODO
