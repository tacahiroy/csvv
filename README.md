csvv
=====
A simple CSV extractor. Intends to use as an AWK alternative.

Installation
============
```console
go get github.com/tacahiroy/csvv
```

Run
===
```console
csvv /path/to/csv column1,column2[,column3...]
```
Or you can run Go source w/o compile
```console
go run csvv.go /path/to/csv column1,column2[,column3...]
```

Known Issues
============
Currently, it doesn't care quoting for each column.
