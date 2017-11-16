package set

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

//TODO : flesh out this test with more cases
func TestSet(t *testing.T){
	Convey("Given a SET mutation struct", t, func() {
		setValue := 50
		setStruct, err := New(setValue)
		So(err, ShouldBeNil)

		Convey("Given an empty document", func() {
			emptyDocString := "{}"
			emptyDoc := map[string]interface{}{}
			json.Unmarshal([]byte(emptyDocString), emptyDoc)

			Convey("Given a JSON pointer path that doesn't exist in the document", func (){
				keyName := "bait"
				jsonPointerPath := "/" + keyName

				Convey("Applying the mutation should work", func (){
					newDoc, err := setStruct.Apply(jsonPointerPath,emptyDoc)

					So(err, ShouldBeNil)
					So(newDoc[keyName], ShouldEqual, setValue)
				})
			})
		})
	})
}