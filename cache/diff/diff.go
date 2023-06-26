package diff

import (
	"fmt"
	"reflect"
)

// Note: PartialUpdate updates to
func PartialUpdate[T any](to *T, new *T) *T {
	return partialUpdateImpl(to, new).(*T)
}

func partialUpdateImpl(to any, new any) any {
	// We use reflection here to get every field
	// and update it if it's not nil

	toVal := reflect.ValueOf(to).Elem()
	newVal := reflect.ValueOf(new).Elem()

	for i := 0; i < toVal.NumField(); i++ {
		// Get the to field
		toField := toVal.Field(i)

		// Get the new field
		newField := newVal.Field(i)

		if toField.Kind() == reflect.Ptr {
			// Indirect the pointer if it's not nil
			if !toField.IsNil() {
				toField = toField.Elem()
			}
		}

		if newField.Kind() == reflect.Ptr {
			// Ensure we aren't nil?
			//
			// We can skip this field if it's nil
			if newField.IsNil() {
				continue
			}

			// Indirect the pointer
			newField = newField.Elem()
		}

		switch newField.Kind() {
		case reflect.Struct:
			// Recursively apply a partial update here too

			// Get type of field
			f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

			// Set the field
			if toField.CanSet() {
				toField.Set(reflect.ValueOf(f))
			} else {
				panic("can't set field")
			}
		case reflect.Slice:
			// Ensure len of slice of new is greater than 0 and not nil
			if newField.Len() > 0 && !newField.IsNil() {
				// Get type of field
				f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(f))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Map:
			// Ensure len of map of new is greater than 0 and not nil
			if newField.Len() > 0 && !newField.IsNil() {
				// Get type of field
				f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(f))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Interface:
			// Get type of field
			f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

			// Set the field
			if toField.CanSet() {
				toField.Set(reflect.ValueOf(f))
			} else {
				panic("can't set field")
			}
		case reflect.String:
			// Check that string is not empty in new

			// As the string may be a type alias, we need to get the underlying type
			// and check if it's empty

			// Get the string as a string, we can't do .Interface().(string) as may be a type alias
			str := fmt.Sprintf("%v", newField.Interface())

			if str != "" {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(str))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Bool:
			// Only set field if its an update
			if newField.Interface().(bool) {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(true))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Int:
			// Only set field if its nonzero
			if newField.Interface().(int) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(int)))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Int64:
			// Only set field if its nonzero
			if newField.Interface().(int64) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(int64)))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Uint:
			// Only set field if its nonzero
			if newField.Interface().(uint) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(uint)))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Uint64:
			// Only set field if its nonzero
			if newField.Interface().(uint64) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(uint64)))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Float64:
			// Only set field if its nonzero
			if newField.Interface().(float64) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(float64)))
				} else {
					panic("can't set field")
				}
			}
		case reflect.Float32:
			// Only set field if its nonzero
			if newField.Interface().(float32) != 0 {
				// Set the field
				if toField.CanSet() {
					toField.Set(reflect.ValueOf(newField.Interface().(float32)))
				} else {
					panic("can't set field")
				}
			}
		default:
			panic("unknown type:" + toField.Kind().String() + fmt.Sprint(toField))
		}
	}

	return to
}

type testStruct struct {
	Str  string
	Meow string
	ABC  int
}

func init() {
	var to = testStruct{
		Str:  "hello",
		Meow: "meow",
		ABC:  123,
	}

	var new = testStruct{
		Str: "world",
		ABC: 456,
	}

	fmt.Println(PartialUpdate(&to, &new))
}
