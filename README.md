# Go Systems Programming

From the [Go Systems Programming book][book-site], and [my book notes][book-notes].

## Examples

### Basic Go Features

1. [Commandline Param `-i`](parameter/parameter.go)
1. [Illustrative Functions](functions/functions.go)
1. [Deferred Functions](defer/defer.go)
1. [Pointer Basics](pointers/pointers.go)
1. [Converting an Array to a Map](array2map/array2map.go)
1. [Data Structures](dataStructures/dataStructures.go)
1. [Interface / Duck Typing](interfaces/interfaces.go)
1. [Random Numbers](random/random.go)
1. [Add Commandline Args](addCLA/addCLA.go)

### Advanced Go Features

1. [Function Error Handling](funErr/funErr.go)
1. [Logging](logging/logging.go)
1. [Add Commandline Args (Improved!)](addCLAImproved/addCLAImproved.go)
1. [Regular Expressions](regExp/regExp.go)
1. [Read Column in Data](readColumn/readColumn.go)
1. [Column Summary](summary/summary.go)
1. [Occurrences](occurrences/occurrences.go)
1. [Find and Replace](findReplace/findReplace.go)
1. [Reflection](reflection/reflection.go)
1. [Unreachable Code](unreachable/unreachable.go)
   * `go tool vet unreachable.go`

### Packages, Algorithms & Data Structures

1. [Sorting Slices](sortSlice/sortSlice.go)
1. [Runtime](runtime/runtime.go)
1. [A Simple Package](aSimplePackage/aSimplePackage.go)
   * `mkdir -p ~/go/src/aSimplePackage`
   * `cp aSimplePackage.go ~/go/src/aSimplePackage/`
   * `go install aSimplePackage`
1. [Using A Simple Package](usePackage/usePackage.go)
1. [Using The MySQL Package](useMySQL/useMySQL.go)
   * `go get github.com/go-sql-driver/mysql`
1. [Garbage Collection](garbageCol/garbageCol.go) 

### Files and Directories

1. [Using Command-Line Flags](usingFlag/usingFlag.go)
   * `go run usingFlag.go -o -c -k 5`
1. [Symbolic Links](symbLink/symbLink.go)
   * `go run symbLink.go /etc`
1. [pwd(1)](pwd/pwd.go)
   * `go run pwd.go -P`
1. [which(1)](which/which.go)
   * `go run which.go ls` to show the first
   * `go run which.go -a ls` to show all
   * `go run which.go -s ls && echo $?` to print the exit code
   * `go run which.go -s butts && echo $?` to print the exit code
1. [permissions](permissions/permissions.go)
   * `go run permissions.go ~`

[book-site]: https://www.packtpub.com/networking-and-servers/go-systems-programming
[book-notes]: https://github.com/trueheart78/book-notes/tree/master/go-systems-programming
