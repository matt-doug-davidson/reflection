package reflection

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Pointer struct {
	X  *float64 `json:"Xplus,omitempty"`
	Y  float64  `json:"Yplus,omitempty"`
	IP *int     `json:"Iintp,omitempty"`
	I  int      `json:"Iint,omitempty"`
}

// JsonifyPretty converts the object to a pretty JSON string
func JsonifyPretty(d interface{}) string {
	jsonData, err := json.MarshalIndent(d, "", "    ")
	//jsonData, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

// JsonifyPretty converts the object to a pretty JSON string
func Jsonify(d interface{}) string {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

func TestSetFlt2Ptr(t *testing.T) {
	x := 33.66
	y := SetFlt3Ptr(x)

	fmt.Println(y)
	fmt.Println(JsonifyPretty(y))

	x0 := 99.0
	y0 := SetFltPrecisionPtr(x0, 0)
	fmt.Println(JsonifyPretty(y0))
	y1 := SetFltPrecisionPtr(x0, 1)
	fmt.Println(JsonifyPretty(y1))
	y2 := SetFltPrecisionPtr(x0, 2)
	fmt.Println(JsonifyPretty(y2))

}

// func TestFloat(t *testing.T) {
// 	p := &Pointer{}

// 	err := SetField(p, "Y", 969.5)
// 	if err != nil {
// 		t.Errorf("setting field Y to 969.5 failed. Cause: %s", err.Error())
// 	}
// 	jsonString := Jsonify(p)
// 	if jsonString != "{\"Yplus\":969.5}" {
// 		t.Errorf("Failed.")
// 	}

// 	err = SetField(p, "Y", 55)
// 	if err != nil {
// 		if err.Error() != "cannot set Y field of type float64 with value of type int" {
// 			t.Errorf("Failed.")
// 		}
// 	}
// }

// func TestFloatPointer(t *testing.T) {
// 	p := &Pointer{}

// 	err := SetField(p, "X", 638.87)
// 	if err != nil {
// 		t.Errorf("setting field X to 638.87 failed. Cause: %s", err.Error())
// 	}
// 	jsonString := Jsonify(p)
// 	if jsonString != "{\"Xplus\":638.87}" {
// 		t.Errorf("Failed.")
// 	}

// 	err = SetField(p, "X", 554)
// 	if err != nil {
// 		if err.Error() != "field is *float64 but value is int" {
// 			t.Errorf("Failed.")
// 		}
// 	}
// }

// func TestInt(t *testing.T) {
// 	p := &Pointer{}

// 	err := SetField(p, "I", 55)
// 	if err != nil {
// 		t.Errorf("setting field I to 55 failed. Cause: %s", err.Error())
// 	}
// 	jsonString := Jsonify(p)
// 	if jsonString != "{\"Iint\":55}" {
// 		t.Errorf("Failed.")
// 	}

// 	err = SetField(p, "I", 55.)
// 	if err != nil {
// 		if err.Error() != "cannot set I field of type int with value of type float64" {
// 			t.Errorf("Failed.")
// 		}
// 	}
// }

// func TestIntPointer(t *testing.T) {
// 	p := &Pointer{}

// 	err := SetField(p, "IP", 55)
// 	if err != nil {
// 		t.Errorf("setting field IP to 55 failed. Cause: %s", err.Error())
// 	}
// 	jsonString := Jsonify(p)
// 	if jsonString != "{\"Iintp\":55}" {
// 		t.Errorf("Failed.")
// 	}

// 	err = SetField(p, "IP", 55.)
// 	if err != nil {
// 		if err.Error() != "field is *int but value is float64" {
// 			t.Errorf("Failed.")
// 		}
// 	}
// }

// func TestBadField(t *testing.T) {
// 	p := &Pointer{}

// 	err := SetField(p, "J", 55)
// 	if err != nil {
// 		if err.Error() != "no such field: J in structure" {
// 			t.Errorf("Failed")
// 		}
// 	}
// }

// func TestGetFieldInt(t *testing.T) {
// 	p := &Pointer{
// 		Y: 9654.364,
// 		I: 3762,
// 	}

// 	f, err := GetField(p, "Y")

// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}
// 	if f != 9654.364 {
// 		t.Errorf("floating value not 9654.346")
// 	}

// 	i, err := GetField(p, "I")
// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}
// 	if i != 3762 {
// 		t.Errorf("int value not 3762")
// 	}
// }

// func TestGetIntPointer(t *testing.T) {
// 	p := &Pointer{
// 		Y: 9654.364,
// 		I: 3762,
// 	}
// 	err := SetField(p, "IP", 66)
// 	if err != nil {
// 		t.Errorf("setting field IP to 55 failed. Cause: %s", err.Error())
// 	}

// 	i, err := GetField(p, "IP")
// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}
// 	expected := 66
// 	if i != expected {
// 		t.Errorf("int value, %d, not %d", i, expected)
// 	}
// }

// func TestGetFloatPointer(t *testing.T) {
// 	p := &Pointer{
// 		Y: 9654.364,
// 		I: 3762,
// 	}
// 	testValue := 66.777
// 	err := SetField(p, "X", testValue)
// 	if err != nil {
// 		t.Errorf("setting field IP to 66.666 failed. Cause: %s", err.Error())
// 	}

// 	f, err := GetField(p, "X")
// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}

// 	if f != testValue {
// 		t.Errorf("int value, %f, not %f", f, testValue)
// 	}
// }
