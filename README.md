# url-shortener

URL shortener app, written in Golang, using PostgresDB as persistant storage.

## Setup
In order to run the app, docker and docker-compose are needed. To build and run the app, from root directory of project run:
```
$ docker-compose up
```
App will be available on port :8080.

## Documentation
API documentation is publicly available at `/swagger/index.html#/`

## Next steps
1. Implement the URL shortening algorithm without relying on external packages, such as the one used here (teris-io/shortid).
2. Perform retrieval of full link and counter increment in a transaction.
3. Add other storage options.
4. Full test coverage.

