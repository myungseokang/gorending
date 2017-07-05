gorending
===

[![Go Report Card](https://goreportcard.com/badge/github.com/Leop0ld/gorending)](https://goreportcard.com/report/github.com/Leop0ld/gorending)
<a href="https://godoc.org/github.com/Leop0ld/gorending" target="_blank"><img src="https://godoc.org/github.com/Leop0ld/gorending?status.svg" alt="GoDoc"></a>
[![Build Status](https://travis-ci.org/Leop0ld/gorending.svg?branch=master)](https://travis-ci.org/Leop0ld/gorending)

Demo
---
<script type="text/javascript" src="https://asciinema.org/a/ujAF6W6sHJMMMXCOmKqQUOzyR.js" id="asciicast-ujAF6W6sHJMMMXCOmKqQUOzyR" async></script>

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
