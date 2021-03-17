# go

Go path on Mac:

```
$ cat /etc/paths.d/go
/usr/local/go/bin
```

Run a go application:

```
$ go build file.go # optional
$ go run helloworld.go
Hello from Go!
```

Check documentation for package:

```
$ go doc strings
```

Format source code:

```
$ gofmt file.go
$ gofmt -w file.go
```

Check Go enviromental variables:

```
$ go env
```

Installed package goes to the $GOPATH/bin directory:

```
$ go env GOPATH
/Users/gengwg/go
$ ls /Users/gengwg/go
pkg
$ go install
$ ll /Users/gengwg/go
total 0
drwxr-xr-x  3 gengwg  staff  96 Mar 10 08:57 bin
drwxr-xr-x  3 gengwg  staff  96 Feb  4 09:53 pkg
$ ll /Users/gengwg/go/bin/
total 43552
-rwxr-xr-x  1 gengwg  staff  22297084 Mar 10 08:57 prometheus-ods-adapter
```
