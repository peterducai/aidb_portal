# aidb_portal

web portal for aidb


## ECHO web server

* http/https server written in Go using Echo framework

### GET

### POST

> curl -X POST -F 'jumpname=somejump' www.localhost/jumps

### UPDATE

> curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Joe"}' localhost:1323/users/1

### DELETE

> curl -X DELETE localhost:1323/users/1


## SSH

go.crypto/ssh is an implementation of SSH protocol in Go.

> go get golang.org/x/crypto/ssh

## Web server

Installation

>  go get -u github.com/labstack/echo/...

more can be found at [Echo guide](https://echo.labstack.com/guide)