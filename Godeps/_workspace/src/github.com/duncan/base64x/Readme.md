# base64x

This tiny package provides a wrapper of Go’s `encoding/base64` package to fix up base64 encoded data before decoding. This includes:

* Ability to fixup missing padding in base64 encoded data
* Translation between base64 data in `StdEncoding` and `URLEncoding`

Both of these are intended to implement robust handling around base64 data that you need to process from sources that you don’t control. For example, according to the RFC, padding is optional for URL encoded data in cases where the data length is known.


## Usage

To use, first import this library:

```
import "github.com/duncan/base64x"
```

Then, replace calls like this, which will return only `abc` and a `CorruptInputError`:

```go
s, err := base64.URLEncoding.DecodeString("YWJjZGU")
```

with this:

```go
s, err := base64x.URLEncoding.DecodeString("YWJjZGU")
```

This correctly decodes `abcde` with no error.

In addition to `DecodeString`, an implementation that wraps the default `Decode` function is provided:

```go
i, err := StdEncoding.Decode(d, []byte(s))
```

If the behavior of Go’s `encoding/base64` package changes and becomes more accepting of unpadded content, you can easily move back to the default implementation by changing `base64x` to `base64` and dropping the import.

### AutoEncoding

While most of the time you’ll know if base64 encoded data is in either standard or URL encoding, there are times you might not know. The `AutoEncoding` helps with that by automatically standardizing encoded data before decoding.

```go
s, err := base64.AutoEncoding.DecodeString("YWJjZGU")
```

## Discussion of padding requirements

[RFC 4648](http://www.faqs.org/rfcs/rfc4648.html) has this to say in Section 3.2 about padding in base-encoded data:

> In some circumstances, the use of padding ("=") in base-encoded data
   is not required or used.  In the general case, when assumptions about
   the size of transported data cannot be made, padding is required to
   yield correct decoded data.

Furthermore, it says in Section 5:

> The pad character "=" is typically percent-encoded when used in an
   URI [9], but if the data length is known implicitly, this can be
   avoided by skipping the padding; see section 3.2.

In other words, it’s complicated. But, in the final analysis, getting things done in the real world when clients send URL encoded base64 data without padding means being robust and dealing with it accordingly.

## Questions

**What about fixing it in Go’s core library?**

The issue of the core library’s strict decoding has been discussed since at least 2012—see [issue #4237 on github.com/golang/go](https://github.com/golang/go/issues/4237). It has been postponed multiple times by the core maintainers and doesn’t look like it’ll be addressed anytime soon. I’ve considered submitting a pull request, but even if I do, I need the functionality in running applications today.

**What about encoding? Why not a wrapper to remove padding?**

Accepting base64 data without padding is in line with [Postel’s law](http://en.wikipedia.org/wiki/Robustness_principle), but so is being strict with output. Even in cases where padding is optional, there’s no need to remove it if it’s already been generated. Since this is the thinnest possible wrapper over the core library, there’s no need to address strict encoding.

## License

Licensed under the Apache License, Version 2.0
