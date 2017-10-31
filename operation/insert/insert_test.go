package insert

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncr(t *testing.T){
	Convey("Given a parameter", t, func(){
		insertDefinition := map[string]interface{}{}


		Convey("Given an 'at' field and multiple 'values'", func (){
			insertDefinition["at"] = 2
			insertDefinition["values"] = []interface{}{"tomato","potato","carrots"}
			var constructorParam interface{} = insertDefinition

			Convey("Creating a new Insert struct should work", func (){
				insertStruct, err := New(constructorParam)
				So(err, ShouldBeNil)

				Convey("Given a valid document with an array field" , func (){
					document := make(map[string]interface{})
					document["favourites"] = []interface{}{"soup", "stew", "steak", "salmon", "ramen"}


					Convey("Given a valid jsonpointer", func (){
						pointer := "/favourites"

						Convey("Applying should work, and the result should be as expected", func () {
							result, err := insertStruct.Apply(pointer,document)
							So(err, ShouldBeNil)

							favourites := result["favourites"].([]interface{})
							So(len(favourites), ShouldEqual, 8)
							So(favourites[0],ShouldEqual,"soup")
							So(favourites[1],ShouldEqual,"stew")
							So(favourites[2],ShouldEqual,"tomato")
							So(favourites[3],ShouldEqual,"potato")
							So(favourites[4],ShouldEqual,"carrots")
							So(favourites[5],ShouldEqual,"steak")
							So(favourites[6],ShouldEqual,"salmon")
							So(favourites[7],ShouldEqual,"ramen")
						})
					})

					Convey("Given an invalid jsonpointer", func (){
						pointer := "/notfavourites"

						Convey("Applying should cause an error", func () {
							_, err := insertStruct.Apply(pointer,document)
							So(err, ShouldNotBeNil)
						})
					})
				})

				Convey("Given an invalid document (field isn't an array)" , func (){
					document := make(map[string]interface{})
					document["favourites"] = "15"

					Convey("Given a valid jsonpointer", func (){
						pointer := "/favourites"

						Convey("Applying should cause an error", func () {
							_, err := insertStruct.Apply(pointer,document)
							So(err, ShouldNotBeNil)
						})
					})
				})
			})
		})

		Convey("Given an 'at' field and a single 'value'", func (){
			insertDefinition["at"] = 2
			insertDefinition["values"] = []interface{}{"tomato"}
			var constructorParam interface{} = insertDefinition

			Convey("Creating a new Insert struct should work", func (){
				insertStruct, err := New(constructorParam)
				So(err, ShouldBeNil)

				Convey("Given a valid document with an array field" , func (){
					document := make(map[string]interface{})
					document["favourites"] = []interface{}{"soup", "stew", "steak", "salmon", "ramen"}


					Convey("Given a valid jsonpointer", func (){
						pointer := "/favourites"

						Convey("Applying should work, and the result should be as expected", func () {
							result, err := insertStruct.Apply(pointer,document)
							So(err, ShouldBeNil)

							favourites := result["favourites"].([]interface{})
							So(len(favourites), ShouldEqual, 6)
							So(favourites[0],ShouldEqual,"soup")
							So(favourites[1],ShouldEqual,"stew")
							So(favourites[2],ShouldEqual,"tomato")
							So(favourites[3],ShouldEqual,"steak")
							So(favourites[4],ShouldEqual,"salmon")
							So(favourites[5],ShouldEqual,"ramen")
						})
					})
				})
			})
		})

		Convey("Given no 'at' field and a single 'value'", func (){
			insertDefinition["values"] = []interface{}{"tomato"}
			var constructorParam interface{} = insertDefinition

			Convey("Creating a new Insert struct should work", func (){
				insertStruct, err := New(constructorParam)
				So(err, ShouldBeNil)

				Convey("Given a valid document with an array field" , func (){
					document := make(map[string]interface{})
					document["favourites"] = []interface{}{"soup", "stew", "steak", "salmon", "ramen"}


					Convey("Given a valid jsonpointer", func (){
						pointer := "/favourites"

						Convey("Applying should work, and the result should be as expected", func () {
							result, err := insertStruct.Apply(pointer,document)
							So(err, ShouldBeNil)

							favourites := result["favourites"].([]interface{})
							So(len(favourites), ShouldEqual, 6)
							So(favourites[0],ShouldEqual,"soup")
							So(favourites[1],ShouldEqual,"stew")
							So(favourites[2],ShouldEqual,"steak")
							So(favourites[3],ShouldEqual,"salmon")
							So(favourites[4],ShouldEqual,"ramen")
							So(favourites[5],ShouldEqual,"tomato")
						})
					})
				})
			})
		})

		Convey("Given no 'at' or 'values' fields", func (){
			var constructorParam interface{} = insertDefinition

			Convey("Creating a new Insert struct should cause an error", func (){
				_, err := New(constructorParam)
				So(err, ShouldNotBeNil)
			})
		})
	})

}