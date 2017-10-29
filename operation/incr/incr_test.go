package incr

import (
 	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncr(t *testing.T){
	Convey("Given a valid constructor parameter (int, 6)", t, func(){
		constructorParam := 6

		Convey("Creating a new Incr struct should work", func (){
			incrStruct, err := New(constructorParam)
			So(err, ShouldBeNil)

			Convey("Given a document with a valid value (int, 11)" , func (){
				document := make(map[string]interface{})
				document["years"] = 11

				Convey("Given a valid jsonpointer", func (){
					pointer := "/years"

					Convey("Applying should work, and the result should be as expected (an int of value 17)", func () {
						result, err := incrStruct.Apply(pointer,document)
						So(err, ShouldBeNil)
						So(result["years"],ShouldEqual,17)
					})
				})

				Convey("Given an invalid valid jsonpointer", func (){
					pointer := "/tomatosauce"

					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})
			})

			Convey ("Given a document with a valid value (float, 3.14)", func (){
				document := make(map[string]interface{})
				document["years"] = 3.14

				Convey("Given a valid jsonpointer", func (){
					pointer := "/years"
					Convey("Applying should work, and the result should be as expected (an int of value 9.14)", func () {
						result, err := incrStruct.Apply(pointer,document)
						So(err, ShouldBeNil)
						So(result["years"],ShouldEqual,9.14)
					})
				})

				Convey("Given an invalid valid jsonpointer", func (){
					pointer := "/tomatosauce"

					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})
			})

			Convey("Given a document with an invalid value (string, 'applesauce')", func () {
				document := make(map[string]interface{})
				document["years"] = "applesauce"
				Convey("Given a valid jsonpointer", func () {
					pointer := "/years"

					Convey ("Applying should cause an error", func (){
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})

				Convey("Given an invalid valid jsonpointer", func (){
					pointer := "/tomatosauce"

					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})


	Convey("Given a valid constructor parameter (float, 1.45)", t, func(){
		constructorParam := 1.45

		Convey("Creating a new Incr struct should work", func (){
			incrStruct, err := New(constructorParam)
			So(err, ShouldBeNil)

			Convey("Given a document with a valid value (int, 5)" , func (){
				document := make(map[string]interface{})
				document["years"] = 5

				Convey("Given a valid pointer", func (){
					pointer := "/years"

					Convey("Applying should work, and the result should be as expected (a float of value 6.45)", func () {
						result, err := incrStruct.Apply(pointer,document)
						So(err, ShouldBeNil)
						So(result["years"],ShouldEqual,6.45)
					})
				})

				Convey("Given an invalid pointer", func (){
					pointer := "/tomatosauce"
					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})



			})

			Convey ("Given a document with a valid value (float, 12.3)", func (){
				document := make(map[string]interface{})
				document["years"] = 12.3

				Convey("Given a valid pointer", func (){
					pointer := "/years"
					Convey("Applying should work, and the result should be as expected (an int of value 9.14)", func () {
						result, err := incrStruct.Apply(pointer, document)
						So(err, ShouldBeNil)
						So(result["years"],ShouldEqual,13.75)
					})
				})

				Convey("Given an invalid pointer", func (){
					pointer := "/tomatosauce"
					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})

			})

			Convey("Given a document with an invalid value (string, 'applesauce')", func () {
				document := make(map[string]interface{})
				document["years"] = "applesauce"

				Convey("Given a valid pointer", func (){
					pointer := "/years"
					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})

				Convey("Given an invalid pointer", func (){
					pointer := "/tomatosauce"
					Convey("Applying should cause an error", func () {
						_, err := incrStruct.Apply(pointer,document)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})


	Convey("Given an invalid constructor parameter (string, 'Babylon')", t, func (){
		constructorParam := "Babylon"

		Convey("Creating a new Incr struct should cause an error", func (){
			_, err := New(constructorParam)
			So(err, ShouldNotBeNil)
		})
	})
}