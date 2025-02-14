/*
Copyright 2018 The KubeSphere Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reflectutils

import (
	"reflect"
)

func In(value interface{}, container interface{}) bool {
	containerValue := reflect.ValueOf(container)
	switch reflect.TypeOf(container).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < containerValue.Len(); i++ {
			if containerValue.Index(i).Interface() == value {
				return true
			}
		}
	case reflect.Map:
		if containerValue.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true
		}
	default:
		return false
	}
	return false
}

func Override(left interface{}, right interface{}) {
	if left == nil || right == nil {
		return
	}

	if reflect.ValueOf(left).Type().Kind() != reflect.Ptr ||
		reflect.ValueOf(right).Type().Kind() != reflect.Ptr {
		return
	}

	old := reflect.ValueOf(left).Elem()
	new := reflect.ValueOf(right).Elem()

	for i := 0; i < old.NumField(); i++ {
		val := new.Field(i).Interface()
		if !reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface()) {
			old.Field(i).Set(reflect.ValueOf(val))
		}
	}
}
