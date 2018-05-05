package goutils

import "reflect"

// Contains tells an object is in a list or not
func Contains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if obj == o {
			return true
		}
	}
	return false
}

// DeepContains tells an object is in a list or not, but uses reflect.DeepEqual to determine the equality
func DeepContains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if reflect.DeepEqual(o, obj) {
			return true
		}
	}
	return false
}
