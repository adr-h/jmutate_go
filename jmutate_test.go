package jmutate_go

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJsonMutation(t *testing.T){

	type TestDoc struct {
		Person struct {
			Name string `json:"name"`
			Occupation struct {
				Title string `json:"title"`
				Years int `json:"years"`
				Company string `json:"company"`
				IsSellingSecretsToTheEnemy bool `json:"is_selling_secrets_to_the_enemy"`
			} `json:"occupation"`
			SuccessRate float64 `json:"success_rate"`
		} `json:"person"`
	}

	Convey("Given a valid mutation document and a valid JSON document to apply it to", t, func(){
		jsonString := `{
			"person" : {
				"name" : "Bob",
				"occupation" : {
					"title" : "Senior Developer",
					"years" : 5,
					"company" : "Gate Breakers Inc",
					"is_selling_secrets_to_the_enemy" : false
				},
				"success_rate" : 91.5
			}
		}`

		mutationDocument := `{
			"/person/occupation/years" : {
				"op" : "INCR",
				"arg" : 2
			},
			"/person/success_rate" : {
				"op" : "INCR",
				"arg" : 3.4
			},
			"/person/occupation/title" : {
				"op" : "SET",
				"arg" : "CTO"
			},
			"/person/occupation/is_selling_secrets_to_the_enemy" : {
				"op" : "DEL"
			}
		}`

		Convey("Creating a JsonMutation should work", func (){
			mutation, err := New([]byte(mutationDocument))
			So(err, ShouldBeNil)


			Convey("Applying the mutation on the valid JSON document should work and give the expected results", func (){
				result, err := mutation.Apply([]byte(jsonString))
				So(err, ShouldBeNil)


				//unmarshal to parse the results; need a map as well to confirm "DELETE" operation
				var testDoc TestDoc
				var testMap map[string]interface{}

				err = json.Unmarshal(result, &testDoc)
				So(err, ShouldBeNil)
				err = json.Unmarshal(result, &testMap)
				So(err, ShouldBeNil)

				So(testDoc.Person.Occupation.Years, ShouldEqual, 7)
				So(testDoc.Person.SuccessRate, ShouldEqual, 94.9)
				So(testDoc.Person.Occupation.Title, ShouldEqual, "CTO")

				personMap := testMap["person"].(map[string]interface{})
				occupationMap := personMap["occupation"].(map[string]interface{})
				_, deletedKeyIsPresent := occupationMap["is_selling_secrets_to_the_enemy"]
				So(deletedKeyIsPresent, ShouldBeFalse)
			})
		})
	})
}
