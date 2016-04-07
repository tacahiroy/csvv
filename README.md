csvv
=====
A simple CSV extractor.

Motivation
==========
Sometimes, I need to process too large CSV files from command-line. Needless to say, we have great tools such as `cat`, `sed`, `grep` and `awk`.
```console
cat /path/to/csv-file | awk -F ',' '{ print $2 "," $1 }'
```
However, how I can specify columns by column name, instead of column number like `$1` or `$2`?
Although I didn't google it, I didn't have clear answer(s) to achieve.
So I created csvv.

Installation
============
```console
go get github.com/tacahiroy/csvv
```

Build
=====
```console
go build csvv.go
```

Run
===
```console
csvv /path/to/csv column1,column2[,column3...]
```

Example
=======
Here's a CSV file, say `users.csv`.
```console
id,name,email,github,twitter
1,Bob,bob@aaa.xyz,bobaaa,bob123
2,Jake,jake@aaa.xyz,jakeaaa,jake123
```

If you want to get fields named `id` and `email`, you may run like this:
```console
csvv users.csv id,email
```

Then you get:
```console
id,email
1,bob@aaa.xyz
2,jake@aaa.xyz
```
