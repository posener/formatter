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
// Create a new formatter.
f := Formatter{
    // For formatter wraps a writer. The writer is where the output will be written to.
    Writer: os.Stdout,
    // Define text indentation. This could be spaces, tabs or any other character set. It will
    // be inserted in the beginning of the text and after every new line. Leave it nil for not
    // indenting the text.
    Indent: []byte("  "),
    // Width defines line width for applying line wrap. Any non-positive number will be ignored.
    Width: 100,
}
// Writing into the formatter will result in the formatted text being written into the defined
// `Formatter.Writer`.
_, err := f.Write([]byte(loremIpsum))
if err != nil {
    // Handle error :-)
    panic("failed writing Lorem Ipsum")
}

// The result is the same text formatted with indentation of two spaces and 100 characters line
// width:
```


---

Created by [goreadme](https://github.com/apps/goreadme)
