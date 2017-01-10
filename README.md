# Learnalist - Education by one list at a time.

# Far from production ready

# Today
[Vaporware](https://en.wikipedia.org/wiki/Vaporware).
Check [status.json](./status.json) for current status. (Not very useful, yet...)

# Tomorrow

A way to learn via "alist". Made by you, another human or something else.
It will be a service, which will consume the Learnalist API. Hosted via learnalist.net or privately.


# Getting Started

* Make sure you have [govendor](https://github.com/kardianos/govendor) installed, it is used to manage dependencies.
* Grab the repo
```
git clone https://github.com/freshteapot/learnalist.git
cd learnalist/api
govendor sync
go run cmd/api/main.go --port=1234 --database=/tmp/api.db
```
Your server should now be running on port 1234 with the database created at /tmp/api.db

```
curl -i http://localhost:1234
```

Should produce something like
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 24 Sep 2016 14:24:46 GMT
Content-Length: 31

{"message":"1, 2, 3. Lets go!"}
```

## Play along.
When the database is created, it is empty.

### Add a list of type v1.

```
curl -XPOST http://localhost:1234/alist -d'
{
    "data": [
        "monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday",
        "saturday",
        "sunday"
    ],
    "info": {
        "title": "Days of the Week",
        "type": "v1"
    }
}
'
```

### Get all lists created by you.
```
curl http://localhost:1234/alist/by/me
```

### Add a list of type v2.

```
curl -XPOST http://localhost:1234/alist -d'
{
    "data": {
        "car": "bil",
        "water": "vann"
    },
    "info": {
        "title": "A few words from English to Norwegian.",
        "type": "v2"
    }
}
'
```

Again, query all the lists by you.
```
curl http://localhost:1234/alist/by/me
```

Or an individual list.
```
curl http://localhost:1234/alist/{uuid}
```

# Api

| Method | Uri | Description |
| --- | --- | --- |
| POST | /alist | Save a list. |
| PATCH | /alist/{uuid} | Update one or more fields to the list. |
| PUT | /alist/{uuid} | Update all fields allowed to a list. |
| GET | /alist/{uuid} | Get a list via uuid. |
| GET | /alist/by/{uuid} | Get lists by {uuid}. Allow for both public, private lists. |



# List types

| Type | Description |
| --- | --- |
| v1 | An array of a string.|
| v2 | An array of key:value objects.|

### V1

```
{
    "data": [
        "a",
        "b"
    ],
    "info": {
        "title": "A list of strings",
        "type": "v1"
    }
}
'
```

### V2

```
{
    "data": {
        "key": "value"
    },
    "info": {
        "title": "A list of key:value pairs.",
        "type": "v2"
    }
}
```

# References as this becomes more useful.

* https://echo.labstack.com/
* Managing dependencies with [govendor](https://github.com/kardianos/govendor)
* https://github.com/thewhitetulip/web-dev-golang-anti-textbook
* https://gobyexample.com/command-line-flags
* https://developer.github.com/v3/
* [Example that helped understand Unmarshall and Marshall 1](http://mattyjwilliams.blogspot.no/2013/01/using-go-to-unmarshal-json-lists-with.html)
* [Example that helped understand Unmarshall and Marshall 2](https://gist.github.com/mdwhatcott/8dd2eef0042f7f1c0cd8)

# References as I dive deeper into golang.
* https://gobyexample.com/json
* [Like casting but not](https://golang.org/ref/spec#Type_assertions)
* Interfaces http://go-book.appspot.com/interfaces.html


# Problems

* Slow to run via 'go run'
```
cd ./vendor/github.com/mattn/go-sqlite3/
go install
```

Thanks to http://stackoverflow.com/a/38296407.
