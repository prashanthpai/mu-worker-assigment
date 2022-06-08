# mu-worker assignment

mu-worker reads feedback (upvote/downvote) on videos sent over
a RabbitMQ queue and persists them to the database.

## Pre-requisites

1. [Install docker](https://docs.docker.com/engine/install/)
2. [Install go](https://go.dev/doc/install)

## Libraries/packages to be used

#### Config managment

1. Parsing toml files: github.com/BurntSushi/toml
2. Validating config structs using struct tags: github.com/go-playground/validator/v10
3. Loading structs with environment variables: github.com/sethvargo/go-envconfig

#### Database operations

Please refrain from using third-party libraries and use the built-in
[database/sql](https://pkg.go.dev/database/sql) from standard library
for interacting with MySQL database.

#### RabbitMQ operations

Use github.com/streadway/amqp

You can either use MySQL and RabitMQ hosted on Unacademy infrastructure for
developing your assignment code or choose it to run locally on your machine
as described below.

If you wish to use MySQL and RabitMQ hosted by Unacademy, please contact
Campsite organizers for credentials.

## Local Setup: MySQL and RabbitMQ (Optional)

Bring up rabbitmq and mysql containers:

```sh
$ docker-compose up
```

Log into the rabbitmq container and reset guest password:

```
$ docker exec -it rabbitmq /bin/sh
# rabbitmqctl list_users
# rabbitmqctl change_password guest guest
```

After changing password, restart the container.

```
$ docker-compose down -v
$ docker-compose up
```

RabbitMQ web UI can now be opened in the browser at: http://localhost:15672/#/

MySQL can be accessed from local machine:

```
$ mysql -u root -h 127.0.0.1 --protocol=TCP -P 3306 -p
Enter password:
```

or if mysql client isn't installed on your machine, you can run the client
from inside the container:

```
$ docker exec -it mysql /bin/sh
# mysql -u root -p
Enter password:
```

Create required table:

```
mysql> USE unacademy;
mysql> CREATE TABLE video_votes
    -> (
    ->     video_id    varchar(6)    NOT NULL,
    ->     upvote      int           NOT NULL,
    ->     downvote    int           NOT NULL,
    ->     created_at  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ->     updated_at  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    ->     PRIMARY KEY (video_id)
    -> );
Query OK, 0 rows affected (0.12 sec)
```

## Run worker

Edit [config/config.toml](config/config.toml) file to set configuration for
RabbitMQ and MySQL.

Compile and run the worker process. Send a valid message via RabbitMQ
web management interface or via Django app integration.

```sh
$ make
$ ./mu-worker config/config.toml
2022/06/02 00:09:20  [*] Waiting for messages. To exit press CTRL+C
2022/06/02 00:09:25 Received a message: {
    "video_id": "MEOWS",
    "upvote": 9,
    "downvote": 3
}
```

#### References

* https://hub.docker.com/_/rabbitmq
* https://hub.docker.com/_/mysql
* https://www.rabbitmq.com/tutorials/tutorial-one-go.html
* https://access.redhat.com/solutions/2172871
* https://stackoverflow.com/questions/59838692/mysql-root-password-is-set-but-getting-access-denied-for-user-rootlocalhost
