package incr

import (
 	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncr(t *testing.T){
	Convey("Given a valid int argument (6)", t, func(){
		argument := 6

		Convey("Creating a new Incr struct should work", func (){
			incrStruct, err := New(argument)
			So(err, ShouldBeNil)

			Convey("Given a valid int receiver (11)" , func (){
				receiver := 11

				Convey("Applying should work, and the result should be as expected (an int of value 17)", func () {
					result, err := incrStruct.Apply(receiver)
					So(err, ShouldBeNil)
					So(result,ShouldEqual,17)
				})
			})

			Convey ("Given a valid float receiver (3.14)", func (){
				receiver := 3.14

				Convey("Applying should work, and the result should be as expected (an int of value 9.14)", func () {
					result, err := incrStruct.Apply(receiver)
					So(err, ShouldBeNil)
					So(result,ShouldEqual,9.14)
				})
			})

			Convey("Given an invalid receiver (a string)", func () {
				receiver := "applesauce"

				Convey ("Applying should cause an error", func (){
					_, err := incrStruct.Apply(receiver)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})


	Convey("Given a valid float argument (1.45)", t, func(){
		argument := 1.45

		Convey("Creating a new Incr struct should work", func (){
			incrStruct, err := New(argument)
			So(err, ShouldBeNil)

			Convey("Given a valid int receiver (5)" , func (){
				receiver := 5

				Convey("Applying should work, and the result should be as expected (an int of value 17)", func () {
					result, err := incrStruct.Apply(receiver)
					So(err, ShouldBeNil)
					So(result,ShouldEqual,6.45)
				})
			})

			Convey ("Given a valid float receiver (12.3)", func (){
				receiver := 12.3

				Convey("Applying should work, and the result should be as expected (an int of value 9.14)", func () {
					result, err := incrStruct.Apply(receiver)
					So(err, ShouldBeNil)
					So(result,ShouldEqual,13.75)
				})
			})

			Convey("Given an invalid receiver (a string)", func () {
				receiver := "applesauce"

				Convey ("Applying should cause an error", func (){
					_, err := incrStruct.Apply(receiver)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})

}