
## Genji/Badger db golang example todoapp

storm db has all the capabilities of boltdb with the added query functionality 

### to run:
	$ dep ensure
	$ go clean
	$ go build

run the binary output
command input:
	"all"
	"new"
	"delete"
	"update"
	"find"

TODO: implement relational nested query (ie filter -> async map)
HACK: use different files as db