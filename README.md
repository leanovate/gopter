# GOPTER

... the GOlang Property TestER
[![Build Status](https://travis-ci.org/leanovate/gopter.svg?branch=master)](https://travis-ci.org/leanovate/gopter)
[![codecov.io](https://codecov.io/github/leanovate/gopter/coverage.svg?branch=master)](https://codecov.io/github/leanovate/gopter?branch=master)
[![GoDoc](https://godoc.org/github.com/leanovate/gopter?status.png)](http://godoc.org/github.com/leanovate/gopter)

## Synopsis

Gopter tries to bring the goodness of [ScalaCheck](https://www.scalacheck.org/) (and impliticly the goodness of [QuickCheck](http://hackage.haskell.org/package/QuickCheck)) to Go.
It can be also seen as a more sophisticated version of the testing/quick package.

Main differences to ScalaCheck:

* It is Go ... duh
* ... nevertheless: Do not expect the same typesafety and elegance as in ScalaCheck.
* For simplicity [Shrink](https://www.scalacheck.org/files/scalacheck_2.11-1.12.5-api/index.html#org.scalacheck.Shrink) has become part of the generators. They can be still easily changed if necessary.
* There is no [Pretty](https://www.scalacheck.org/files/scalacheck_2.11-1.12.5-api/index.html#org.scalacheck.util.Pretty) ... so far gopter feels quiet comfortable being ugly.
* A generator for regex matches
* No parallel commands ... yet?

Main differences to the already testing/quick package:

* Much tighter control over generators
* Shrinkers, i.e. automatically find the minimum value falsifying a property
* A generator for regex matches (already mentioned that ... but it's cool)

## License

[MIT Licence](http://opensource.org/licenses/MIT)
