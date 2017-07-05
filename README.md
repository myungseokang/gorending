gorending
===

[![Go Report Card](https://goreportcard.com/badge/github.com/Leop0ld/gorending)](https://goreportcard.com/report/github.com/Leop0ld/gorending)
[![Build Status](https://travis-ci.org/Leop0ld/gorending.svg?branch=master)](https://travis-ci.org/Leop0ld/gorending)

Demo
---
<a href="https://asciinema.org/a/4lfp48hlt8vSikodruSKsnpK4" target="_blank"><img src="https://asciinema.org/a/4lfp48hlt8vSikodruSKsnpK4.png" /></a>

Installation
---
```shell
$ go get github.com/leop0ld/gorending
```

```shell
# in mac OS
$ cp ~/go/bin/gorending /usr/local/bin
```

#### Dependencies

- [github.com/urfave/cli](https://github.com/urfave/cli)
- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

Example
---
If you want to see 5 repositories about golang in Github trending, type like below.

```shell
$ gorending --lang go --count 5
```

or 

```shell
$ gorending -L go -C 5
```

If you want to see 15 repositories about all languages in Github trending, type this.

```shell
$ gorending --count 15
```

or

```shell
$ gorending -C 15
```


License
---
[MIT License](https://en.wikipedia.org/wiki/MIT_License)
