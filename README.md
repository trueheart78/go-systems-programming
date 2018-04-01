# Go Systems Programming

From the [Go Systems Programming book][book-site], and [my book notes][book-notes].

## Examples

### Basic Go Features

1. [Commandline Param `-i`](parameter.go)
1. [Illustrative Functions](functions.go)
1. [Deferred Functions](defer.go)
1. [Pointer Basics](pointers.go)
1. [Converting an Array to a Map](array2map.go)
1. [Data Structures](dataStructures.go)
1. [Interface / Duck Typing](interfaces.go)
1. [Random Numbers](random.go)
1. [Add Commandline Args](addCLA.go)

### Advanced Go Features

1. [Function Error Handling](funErr.go)
1. [Logging](logging.go)
1. [Add Commandline Args (Improved!)](addCLAImproved.go)
1. [Regular Expressions](regExp.go)
1. [Read Column in Data](readColumn.go)
1. [Column Summary](summary.go)
1. [Occurrences](occurrences.go)
1. [Find and Replace](findReplace.go)
1. [Reflection](reflection.go)
1. [Unreachable Code](unreachable.go)
   * `go tool vet unreachable.go`

### Packages, Algorithms & Data Structures

1. [Sorting Slices](sortSlice.go)
1. [Runtime](runtime.go)
1. [A Simple Package](aSimplePackage.go)
   * `mkdir -p ~/go/src/aSimplePackage`
   * `cp aSimplePackage.go ~/go/src/aSimplePackage/`
   * `go install aSimplePackage`
1. [Using A Simple Package](usePackage.go)
1. [Using The MySQL Package](useMySQL.go)
   * `go get github.com/go-sql-driver/mysql`
1. [Garbage Collection](garbageCol.go) 

[book-site]: https://www.packtpub.com/networking-and-servers/go-systems-programming
[book-notes]: https://github.com/trueheart78/book-notes/tree/master/go-systems-programming
