csvv
=====


Setup
=====
You need to install Go to compile csvv.go.
```console
brew install go
```

Compile
=======
```console
go build csvv.go
```

Run
===
```console
./csvv /path/to/csv column1,column2[,column3...]
```
Or you can run Go source w/o compile
```console
go run csvv.go /path/to/csv column1,column2[,column3...]
```

Note: Columns not found in the passed CSV file will be ignored.
