package diff

import (
	"reflect"
)

// Applies all non-nil fields from new to to, thus creating a partial update
//
// Note: PartialUpdate updates to as well in the process
//
// As such, this function *modifies* the to object itself
func PartialUpdate[T any](to *T, new *T) *T {
	return partialUpdateImpl(to, new).(*T)
}

func partialUpdateImpl(base any, contrast any) any {
	// From https://github.com/sentinelb51/revoltgo/blob/main/util.go
	//
	// Their merge() implementation was better so I copied it here
	//
	// All credit goes to them for this much more performant implementation
	baseValues := reflect.ValueOf(base).Elem()
	contrastValues := reflect.ValueOf(contrast).Elem()

	for i := 0; i < baseValues.NumField(); i++ {
		baseValuesField := baseValues.Field(i)
		contrastValuesField := contrastValues.Field(i)

		shouldUpdate := false

		if contrastValuesField.Kind() == reflect.Ptr {
			shouldUpdate = !contrastValuesField.IsNil()
		} else {
			shouldUpdate = !contrastValuesField.IsZero()
		}

		if shouldUpdate {
			baseValuesField.Set(contrastValuesField)
		}
	}

	return base
}
