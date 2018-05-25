# aidb_portal

web portal for aidb


## Technology

* http/https server written in Go using Echo framework

## Design

* Only data that can be entered by human are location, room, and jump servers
* **HW data are read-only** and should be populated only by data collectors. (trust no human)
