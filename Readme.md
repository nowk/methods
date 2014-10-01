# go-methods

[![Build Status](https://travis-ci.org/nowk/go-methods.svg?branch=master)](https://travis-ci.org/nowk/go-methods)
[![GoDoc](https://godoc.org/github.com/nowk/go-methods?status.svg)](http://godoc.org/github.com/nowk/go-methods)

Method bouncer

## Example

    func(w http.ResponseWriter, req *http.Request) {
      m := methods.Allow("get", "post")
      if !m.Allowed(req) {
        // handle illegal methods
      }

      // OK!
    }

## License

MIT