package reflection

import (
	"fmt"
	"reflect"
	"strconv"
)

type Flt0 float64

func (f Flt0) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 0, 64)), nil
}

func SetFlt0Ptr(value float64) *Flt0 {
	ptr := new(Flt0)
	*ptr = Flt0(value)
	return ptr
}

type Flt1 float64

func (f Flt1) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 1, 64)), nil
}

func SetFlt1Ptr(value float64) *Flt1 {
	ptr := new(Flt1)
	*ptr = Flt1(value)
	return ptr
}

type Flt2 float64

func (f Flt2) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 2, 64)), nil
}

func SetFlt2Ptr(value float64) *Flt2 {
	ptr := new(Flt2)
	*ptr = Flt2(value)
	return ptr
}

type Flt3 float64

func (f Flt3) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 3, 64)), nil
}

func SetFlt3Ptr(value float64) *Flt3 {
	ptr := new(Flt3)
	*ptr = Flt3(value)
	return ptr
}

func SetFltPrecisionPtr(value float64, precision int) interface{} {

	switch precision {
	case 0:
		ptr := new(Flt0)
		*ptr = Flt0(value)
		return ptr
	case 1:
		ptr := new(Flt1)
		*ptr = Flt1(value)
		return ptr
	case 2:
		ptr := new(Flt2)
		*ptr = Flt2(value)
		return ptr
	}
	return 0.0
}

// GetField gets the value of a field (field) in a structure (pointer s).
func GetField(s interface{}, f string) (interface{}, error) {

	svoe := reflect.ValueOf(s).Elem()
	fbn := svoe.FieldByName(f)
	ko := fbn.Kind()
	if ko == reflect.Invalid {
		return nil, fmt.Errorf("field %s is not defined for structure", f)
	}

	if ko != reflect.Ptr {
		return fbn.Interface(), nil

	} else {
		fbnType := fbn.Type().String()
		if fbnType == "*int" {
			return int(fbn.Elem().Int()), nil
		} else if fbnType == "*int64" {
			return fbn.Elem().Int(), nil
		} else if fbnType == "*float64" {
			return fbn.Elem().Float(), nil
		}
	}

	return nil, fmt.Errorf("field %s was not set in structure", f)
}

// SetField sets field (f) in structure (pointer s) to value (v). Field
// type in structure (s."f") must match type of value (v)
func SetField(s interface{}, f string, v interface{}) error {

	svoe := reflect.ValueOf(s).Elem()
	fbn := svoe.FieldByName(f)

	// If field value is invalid return error
	if !fbn.IsValid() {
		return fmt.Errorf("no such field: %s in structure", f)
	}

	// If field value is not settable return error
	if !fbn.CanSet() {
		return fmt.Errorf("cannot set %s field value", f)
	}

	vvo := reflect.ValueOf(v)
	vko := vvo.Kind()
	fko := fbn.Kind()

	if fko == vko {
		// Matching types just set the field to the value
		val := reflect.ValueOf(v)
		fbn.Set(val)
	} else if fko == reflect.Ptr {
		fbnType := fbn.Type().String()
		// switch on the value kind
		switch vko {
		case reflect.Float64:
			// Value type is float64
			switch fbnType {
			case "*float64":
				ptr := new(float64)
				*ptr = v.(float64)
				ptrReflectValue := reflect.ValueOf(ptr)
				fbn.Set(ptrReflectValue)
			case "*reflection.Flt0":
				ptr := new(Flt0)
				*ptr = v.(Flt0)
				ptrReflectValue := reflect.ValueOf(ptr)
				fbn.Set(ptrReflectValue)
			case "*reflection.Flt1":
				ptr := new(Flt1)
				*ptr = v.(Flt1)
				ptrReflectValue := reflect.ValueOf(ptr)
				fbn.Set(ptrReflectValue)
			case "*reflection.Flt2":
				ptr := new(Flt2)
				*ptr = v.(Flt2)
				ptrReflectValue := reflect.ValueOf(ptr)
				fbn.Set(ptrReflectValue)
			case "*reflection.Flt3":
				ptr := new(Flt3)
				*ptr = v.(Flt3)
				ptrReflectValue := reflect.ValueOf(ptr)
				fbn.Set(ptrReflectValue)
			default:
				return fmt.Errorf("field is %s but value is float64", fbnType)
			}
		case reflect.Int:
			// Value type is int
			if fbnType != "*int" {
				return fmt.Errorf("field is %s but value is int", fbnType)
			}
			ptr := new(int)
			*ptr = v.(int)
			ptrReflectValue := reflect.ValueOf(ptr)
			fbn.Set(ptrReflectValue)
		default:
			return fmt.Errorf("field type of %s is not supported for value type %s", fbnType, vko.String())
		}
	} else {
		return fmt.Errorf("cannot set %s field of type %+v with value of type %+v", f, fko, vko)
	}

	return nil
}

func GetFieldNames(s interface{}) {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		fmt.Println(val.Type().Field(i).Name)
	}
}
