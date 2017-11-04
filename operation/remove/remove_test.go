package remove

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncr(t *testing.T){
	Convey("Given a parameter", t, func(){
		removeArgDefinition := map[string]interface{}{}

		Convey("Given a valid 'at' field", func (){
			removeArgDefinition["at"] = 3

			Convey("Given a valid 'count' field", func () {
				removeArgDefinition["count"] = 3
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should work", func (){
					removeStruct, err := New(constructorParam)
					So(err, ShouldBeNil)

					Convey("Given a valid document (has an 'array')", func (){
						arr := []interface{}{ "potato", "tomato", "shwarma", "samosa",
							"hellfire", "despair", "redemption", "mayonnaise" }
						document := map[string]interface{}{
							"favourites" : arr,
						}

						Convey("Applying the Remove struct should work", func (){
							mutatedDocument, err := removeStruct.Apply("/favourites", document)
							So(err, ShouldBeNil)

							Convey("The results should be as expected", func (){
								mutatedArr := mutatedDocument["favourites"].([]interface{})
								So(len(mutatedArr), ShouldEqual, 5)
								So (mutatedArr[3], ShouldNotEqual, "samosa")
								So (mutatedArr[3], ShouldEqual, "redemption")

								So (mutatedArr[4], ShouldNotEqual, "hellfire")
								So (mutatedArr[4], ShouldEqual, "mayonnaise")

							})

						})
					})

					Convey("Given an invalid document (has no 'array')", func (){
						document := map[string]interface{}{
							"favourites" : "blueberry",
						}

						Convey("Applying the Remove struct should cause an error", func (){
							_, err := removeStruct.Apply("/favourites", document)
							So(err, ShouldNotBeNil)
						})
					})




				})
			})

			Convey("Given no 'count' field", func () {
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should work", func (){
					removeStruct, err := New(constructorParam)
					So(err, ShouldBeNil)

					Convey("Given a valid document (has an 'array')", func (){
						arr := []interface{}{ "potato", "tomato", "shwarma", "samosa",
							"hellfire", "despair", "redemption", "mayonnaise" }
						document := map[string]interface{}{
							"favourites" : arr,
						}

						Convey("Applying the Remove struct should work", func (){
							mutatedDocument, err := removeStruct.Apply("/favourites", document)
							So(err, ShouldBeNil)

							Convey("The results should be as expected", func (){
								mutatedArr := mutatedDocument["favourites"].([]interface{})
								So(len(mutatedArr), ShouldEqual, 7)
								So (mutatedArr[3], ShouldNotEqual, "samosa")
								So (mutatedArr[3], ShouldEqual, "hellfire")
							})

						})
					})

					Convey("Given an invalid document (has no 'array')", func (){
						document := map[string]interface{}{
							"favourites" : "blueberry",
						}

						Convey("Applying the Remove struct should cause an error", func (){
							_, err := removeStruct.Apply("/favourites", document)
							So(err, ShouldNotBeNil)
						})
					})

				})
			})

			Convey("Given an invalid 'count' field (a string)", func () {
				removeArgDefinition["count"] = "applesauce"
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should cause an error", func (){
					_, err := New(constructorParam)
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("Given an invalid 'at' field (a string)", func (){
			removeArgDefinition["at"] = "chillisauce"

			Convey("Given no 'count' field", func () {
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should cause an error", func (){
					_, err := New(constructorParam)
					So(err, ShouldNotBeNil)
				})
			})

			Convey("Given a valid 'count' field", func () {
				removeArgDefinition["count"] = 2
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should cause an error", func (){
					_, err := New(constructorParam)
					So(err, ShouldNotBeNil)
				})
			})

			Convey("Given an invalid 'count' field (a string)", func () {
				removeArgDefinition["count"] = "applesauce"
				var constructorParam interface{} = removeArgDefinition
				Convey("Creating a Remove struct should cause an error", func (){
					_, err := New(constructorParam)
					So(err, ShouldNotBeNil)
				})
			})
		})

	})

}