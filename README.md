# aidb_portal
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpeterducai%2Faidb_portal.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fpeterducai%2Faidb_portal?ref=badge_shield)


web portal for aidb

## Technology

* http/https server written in Go using Echo framework
* UI using [PatternFly](http://www.patternfly.org) (TODO)
* everything else is handled by external scripts


## Development

> go get github.com/peterducai/aidb_portal

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

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpeterducai%2Faidb_portal.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fpeterducai%2Faidb_portal?ref=badge_large)