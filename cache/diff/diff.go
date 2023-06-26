package diff

import (
	"fmt"
	"reflect"

	"github.com/infinitybotlist/grevolt/types"
)

// Applies all non-nil fields from new to to, thus creating a partial update
//
// Note: PartialUpdate updates to as well in the process
//
// As such, this function *modifies* the to object itself
func PartialUpdate[T any](to *T, new *T) *T {
	return partialUpdateImpl(to, new).(*T)
}

func partialUpdateImpl(to any, new any) any {
	// We use reflection here to get every field
	// and update it if it's not nil

	toVal := reflect.ValueOf(to).Elem()
	newVal := reflect.ValueOf(new).Elem()

	// Check if newVal is zero value
	if newVal.IsZero() {
		return to
	}

	for i := 0; i < toVal.NumField(); i++ {
		// Get the to field
		toField := toVal.Field(i)

		// Get the new field
		newField := newVal.Field(i)

		var set reflect.Value

		switch newField.Kind() {
		case reflect.Ptr:
			// Pointers are painful, lets indirect, and recurse through this function
			// again

			// Check if it's nil
			if newField.IsNil() {
				continue
			}

			if newField.IsZero() {
				continue
			}

			switch newField.Elem().Kind() {
			case reflect.Struct:
				set = newField
			}
		case reflect.Struct:
			// Recursively apply a partial update here too

			// Get type of field
			f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

			// Set the field
			set = reflect.ValueOf(f)
		case reflect.Slice:
			// Ensure len of slice of new is greater than 0 and not nil
			if newField.Len() > 0 && !newField.IsNil() {
				// Get type of field
				f := partialUpdateImpl(toField.Interface(), newField.Interface())

				// Set the field
				set = reflect.ValueOf(f)
			}
		case reflect.Map:
			// Ensure len of map of new is greater than 0 and not nil
			if newField.Len() > 0 && !newField.IsNil() {
				// Get type of field
				f := partialUpdateImpl(toField.Addr().Interface(), newField.Addr().Interface())

				// Set the field
				set = reflect.ValueOf(f)
			}
		case reflect.Interface:
			// Get type of field
			interf := newField.Interface()

			// Check if it's nil
			if interf == nil {
				continue
			}

			// Set the field
			set = newField
		case reflect.String:
			// Check that string is not empty in new

			// As the string may be a type alias, we need to get the underlying type
			// and check if it's empty
			fieldType := newField.Type()

			str := newField.Convert(fieldType).Interface().(string)

			if str != "" {
				// Set the field
				set = newField
			}
		case reflect.Bool:
			// Only set field if its an update
			if newField.Interface().(bool) {
				// Set the field
				set = newField
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// Only set field if its nonzero
			if newField.Int() != 0 {
				// Set the field
				set = newField
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			// Only set field if its nonzero
			if newField.Uint() != 0 {
				// Set the field
				set = newField
			}
		case reflect.Float64, reflect.Float32:
			// Only set field if its nonzero
			if newField.Float() != 0 {
				// Set the field
				set = newField
			}
		default:
			panic("unknown type:" + toField.Kind().String() + fmt.Sprint(toField))
		}

		if set.IsValid() {
			if toField.CanSet() {
				toField.Set(set)
			} else {
				panic("can't set field")
			}
		}
	}

	return to
}

type testStruct struct {
	Str  string
	Meow string
	ABC  int
	DEF  *types.OverrideField
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
		DEF: &types.OverrideField{
			A: 1,
		},
	}

	fmt.Println(PartialUpdate(&to, &new))
}
