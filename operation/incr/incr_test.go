package incr

import (
 	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncr(t *testing.T){
	Convey("Given a valid (int) argument", t, func(){
		argument := 6

		Convey("Creating a new Incr struct should work", func (){
			incrStruct, err := New(argument)
			So(err, ShouldBeNil)

			Convey("Given a valid (int) receiver" , func (){
				receiver := 11

				Convey("Applying the Incr should work, and the result should be as expected", func (){
					result, err := incrStruct.Apply(receiver)
					So(err, ShouldBeNil)
					So(result,ShouldEqual,17)
				})
			})

		})
	})
}