# scope-filter

Filter the contents arriving from stdin based on a target list. The output will contain URLs, IPs and domains defined in the target file.

```
mikey@dev:~/Sites/scope-filter$ cat test/input.txt | ./main test/target.txt 
bookie.dubell.io
shit.infd.pw
hidden.c.collab.dubell.io
https://test.dubell.io/this/is/a/path
https://shit.infd.pw/admin.panel.html
192.168.1.34
192.168.1.156
```

## Install

```
$ go install github.com/dubs3c/scope-filter@latest
```