# Change log

## [Unreleased]
### Changed
- `commands.Command.PostCondition` was called with the state before running the command. It makes
  much more sense to first do `commands.Command.NextState` and then `commands.Command.PostCondition`
- `commands.Commands.NewSystemUnderTest` now takes has an argument `initialState commands.State` to
  allow implementators to create/bootstrap a system under test based on an arbitrary initial state.
  So far examples were just using a constant initial state ... which is a bit boring.
- Fixed: Actually use `commands.Commands.InitialPreCondition` as sieve for
  `commands.Commands.GenInitialState`
- Gen.Map and Shrink.Map now accept `interface{}` instead of `func (interface{}) interface{}`

  This allows cleaner mapping functions without type conversion. E.g. instead of

  ```Go
  gen.AnyString().Map(function (v interface{}) interface{} {
     return strings.ToUpper(v.(string))
  })
  ```
  you can (and should) now write

  ```Go
  gen.AnyString().Map(function (v string) string {
     return strings.ToUpper(v)
  })
  ```
- Correspondingly Gen.SuchThat now also ccept `interface{}` instead of `func (interface{}) bool`

  This allows cleaner sieve functions without type conversion. E.g. instead of

  ```Go
  gen.AnyString().SuchThat(function (v interface{}) bool {
     return HasPrefix(v.(string), "P")
  })
  ```
  you can (and should) now write

  ```Go
  gen.AnyString().SuchThat(function (v string) bool {
     return HasPrefix(v, "P")
  })
  ```
- Gen.FlatMap now has a second parameter `resultType reflect.Type` defining the result type of the mapped generator
- Reason for these changes: The original `Map` and `FlatMap` had a recurring issue with empty results. If the original generator created an empty result there was no clean way to determine the result type of the mapped generator. The new version fixes this by extracting the return type of the mapping functions.

## [0.1] - 2016-04-30
### Added
- Initial implementation.

[Unreleased]: https://github.com/leanovate/gopter/compare/v0.1...HEAD
[0.1]: https://github.com/leanovate/gopter/tree/v0.1