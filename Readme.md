# methods

[![Build Status](https://travis-ci.org/nowk/methods.svg?branch=master)](https://travis-ci.org/nowk/methods)
[![GoDoc](https://godoc.org/github.com/nowk/methods?status.svg)](http://godoc.org/github.com/nowk/methods)

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