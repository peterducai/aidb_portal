# aidb_portal

web portal for aidb

## Technology

* http/https server written in Go using Echo framework
* everything else is handled by external scripts

## Design

### Main principles

**Only data that can be entered by human are location, room, rack and jump servers** as this is only data that cannot be scanned.

**HW data are read-only** and should be populated only by data collectors (programs). Experience tells us that humans makes mistakes and thus should not be allowed in certain data collecting.

# Dependencies

## SSH

go.crypto/ssh is an implementation of SSH protocol in Go.

> go get golang.org/x/crypto/ssh

## Web server

Installation

>  go get -u github.com/labstack/echo/...

more can be found at [Echo guide](https://echo.labstack.com/guide)

> go get -u github.com/lib/pq

for connecting to AIDB (postgresql 10)