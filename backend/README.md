# business-dashboard BACKEND

Documenting difficulties in establishing a database connection using _"github.com/mattn/go-sqlite3"_ for my learning purposes.
Ultimately with the return on investment and the given time frame in which I want to "complete" this project, I decided to abandon using _"github.com/mattn/go-sqlite3"_ and opted for _"modernc.org/sqlite"_.

Sidenote: I try to avoid using "I" in documentation for formal or professional documentation. Using it here is a special case.

Following [the Go Wiki for SQL Drivers](https://github.com/golang/go/wiki/SQLDrivers) I found a list of Drivers that work with SQLite, _"github.com/mattn/go-sqlite3"_ being among them. After some research, I learned that _"github.com/mattn/go-sqlite3"_ is a fairly popular and well-documented choice, so I decided to use it. 

1. Added `_ "github.com/mattn/go-sqlite3"` as an `import` to `main.go`.
(The undersore indicates that it is a "blank identifier".)

2. Ran `go mod tidy` to update `go.mod`.

3. Followed various documentation and tutorials in an attempt to create a database connection.
```
var DB *sql.DB
...
db, err := sql.Open("sqlite", "business.db")
if err != nil {
    fmt.Println("An error in the db connection has occured")
    return
}
DB = db
```
4. Ran `go run main.go` to see if the application still runs after making those changes.

The application seemed to "hang" before starting, meaning the cursor in the cmd would continue to blink without the application starting up, but the run did not crash or error out either. There was no console output. Repeated reruns gave the same result.

I found that if I commented out the import with and without commenting out the code to create a db connection, I would get an error for a missing import as expected. Commenting out the db connection code and keeping the import caused the application to "hang", so it seemed the issue was with the import itself or some other kind of incompatibility. 

After lots of research, I set up and validated several parameters. 
1. Changed `CGO_ENABLED=0` to `CGO_ENABLED=1`
```
go help
go help env
go env CGO_ENABLED        // check value of CGO_ENABLED
go env -w CGO_ENABLED=1   // update CGO_ENABLED value to 1
go env CGO_ENABLED        // validate change 
```
2. Tried moving my database to the `backend` folder thinking I made a mistake in the `sql.Open("sqlite", "business.db")` statement.
3. Reran `go mod tidy` and attempted to validate dependencies.
4. Validated SQLite version is compatible with "go-sqlite3" by entering `sqlite` into the command line.
5. Downgraded to go1.20.11 --> [https://go.dev/dl/](https://go.dev/dl/)
6. Validated PATH variables for Go and SQLite. --> `echo %PATH%` as well as checking from the Environment Variables panel.
7. Followed this guide for using GCC --> [https://code.visualstudio.com/docs/cpp/config-mingw](https://code.visualstudio.com/docs/cpp/config-mingw)
8. Validated I have permission to read/write to the database with `icacls business.db`.

After some time troubleshooting, I decided to try a resource that did not use cgo, leading me to _"modernc.org/sqlite"_. After validating that _"modernc.org/sqlite"_ would work with my current SQLite version and the downgraded version of Go (wanted to save some time and avoid upgrading), I imported _"modernc.org/sqlite"_, ran `go mod tidy`, and then `go run main.go`.

For a few scary seconds, I thought the same issue would occur, but the application started up and I was able to attempt hitting my endpoint that would query the database. Now I can go and fix my "row mapper" aka the Scan!

----
Various commands:
* `go mod example.com/business-dashboard`
* `go run main.go`
* `go mod tidy`
* `curl localhost:8080/api/v1/all-active-employees`
