// The main functions are:
//
//   - [Validate] validates the given JSON using the validator.
//   - [Unmarshal] validates the JSON and unmarshals it into the given type.
//   - [Schema] generates JSON Schema for the validator.
//
// # Types
//
// The top-level validator defines the type of the data.
// It can be a primitive type, a collection, or a composition
// of multiple types.
//
//   - Primitive types: [Bool], [Float64], [Int], [String], [Null], [Any].
//   - Collections: [Array], [Object], [Map]
//   - Composition: [AllOf], [Not]
//
// # Constraints
//
// The types also accept a number of constraints. Either
// as an argument of their constructor or as Constrain method.
//
//   - Numeric constraints: [ExclMax], [ExclMin], [Max], [Min], [MultipleOf]
//   - String constraints: [MaxLen], [MinLen], [Pattern]
//   - Object constraints: [MaxProperties], [MinProperties], [PropertyNames]
//   - Array constraints: [Contains], [MaxItems], [MinItems]
//
// # Errors
//
// [Validate] returns one of the following errors:
//
//   - [Errors]
//   - [ErrNoInput]
//   - [ErrProperty]
//   - [ErrIndex]
//   - [ErrType]
//   - [ErrRequired]
//   - [ErrUnexpected]
//   - [ErrNot]
//
// Or one of the constraint errors:
//
//   - [ErrMultipleOf]
//   - [ErrMin]
//   - [ErrExclMin]
//   - [ErrMax]
//   - [ErrExclMax]
//   - [ErrMinLen]
//   - [ErrMaxLen]
//   - [ErrPattern]
//   - [ErrContains]
//   - [ErrMinItems]
//   - [ErrMaxItems]
//   - [ErrPropertyNames]
//   - [ErrMinProperties]
//   - [ErrMaxProperties]
//
// The errors can be translated using [Locale].
// Multiple locales can be combined in a single registry
// using [Locales].
package valdo
