# formatter

[![Build Status](https://travis-ci.org/posener/formatter.svg?branch=master)](https://travis-ci.org/posener/formatter)
[![codecov](https://codecov.io/gh/posener/formatter/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/formatter)
[![GoDoc](https://godoc.org/github.com/posener/formatter?status.svg)](http://godoc.org/github.com/posener/formatter)
[![goreadme](https://goreadme.herokuapp.com/badge/posener/formatter.svg)](https://goreadme.herokuapp.com)

formatter is a library for formatting text.

The `Formatter` wraps an `io.Writer` and formats text written to it. It enables text indenting
and line width wrapping.

#### Examples

```golang
f := Formatter{Writer: os.Stdout, Indent: []byte("  "), Width: 100}
_, err := f.Write([]byte(loremIpsum))
if err != nil {
    panic("failed writing Lorem Ipsum")
}
```


---

Created by [goreadme](https://github.com/apps/goreadme)
