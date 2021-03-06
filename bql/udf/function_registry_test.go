package udf

import (
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"testing"
)

func TestDefaultFunctionRegistry(t *testing.T) {
	Convey("Given a default function registry", t, func() {
		fr := CopyGlobalUDFRegistry(&core.Context{}) // context is not directly used now

		Convey("When asking for an unknown function", func() {
			_, err := fr.Lookup("hoge", 17)
			Convey("An error is returned", func() {
				So(core.IsNotExist(err), ShouldBeTrue)
			})
		})

		Convey("When adding a unary function via Func", func() {
			fun := func(ctx *core.Context, vs ...data.Value) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("Test", Func(fun, 1))

			Convey("Then it can be looked up as unary", func() {
				_, err := fr.Lookup("TEST", 1)
				So(err, ShouldBeNil)
			})

			Convey("And it won't be found as binary", func() {
				_, err := fr.Lookup("test", 2)
				So(err, ShouldNotBeNil)
				So(core.IsNotExist(err), ShouldBeFalse) // the function was found but its arity was wrong
			})
		})

		Convey("When adding a variadic function via VariadicFunc", func() {
			fun := func(ctx *core.Context, vs ...data.Value) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("hello", VariadicFunc(fun))

			Convey("Then it can be looked up as unary", func() {
				_, err := fr.Lookup("hello", 1)
				So(err, ShouldBeNil)
			})
			Convey("And it can be looked up as binary", func() {
				_, err := fr.Lookup("hello", 2)
				So(err, ShouldBeNil)
			})
		})

		Convey("When adding a nullary function via NullaryFunc", func() {
			fun := func(*core.Context) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("test0", NullaryFunc(fun))

			Convey("Then it can be looked up as nullary", func() {
				_, err := fr.Lookup("test0", 0)
				So(err, ShouldBeNil)
			})

			Convey("And it won't be found as binary", func() {
				_, err := fr.Lookup("test0", 2)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When adding a unary function via UnaryFunc", func() {
			fun := func(*core.Context, data.Value) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("test1", UnaryFunc(fun))

			Convey("Then it can be looked up as unary", func() {
				_, err := fr.Lookup("test1", 1)
				So(err, ShouldBeNil)
			})

			Convey("And it won't be found as binary", func() {
				_, err := fr.Lookup("test1", 2)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When adding a binary function via BinaryFunc", func() {
			fun := func(*core.Context, data.Value, data.Value) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("test2", BinaryFunc(fun))

			Convey("Then it can be looked up as binary", func() {
				_, err := fr.Lookup("test2", 2)
				So(err, ShouldBeNil)
			})

			Convey("And it won't be found as unary", func() {
				_, err := fr.Lookup("test2", 1)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When adding a ternary function via TernaryFunc", func() {
			fun := func(*core.Context, data.Value, data.Value, data.Value) (data.Value, error) {
				return data.Bool(true), nil
			}
			fr.Register("test3", TernaryFunc(fun))

			Convey("Then it can be looked up as ternary", func() {
				_, err := fr.Lookup("test3", 3)
				So(err, ShouldBeNil)
			})

			Convey("And it won't be found as unary", func() {
				_, err := fr.Lookup("test3", 1)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
