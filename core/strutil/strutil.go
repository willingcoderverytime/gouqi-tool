package strutil

import (
	"reflect"
)

func IsEmpty(str *string) bool {

	if nil == str || "" == *str {
		return true
	}

	return false
}
func IsNotEmpty(str *string) bool {
	return !IsEmpty(str)
}

//IsEmptyIFStr
func IsEmptyIFStr(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	if kind == reflect.String {
		value, ok := obj.(string)
		if ok {
			return IsEmpty(&value)
		} else {
			return false
		}

	} else {
		return false
	}
}

// IsBlank 不仅仅是判断是否为nil 还需要判断是不是空字符串以及空分割符号
func IsBlank(str string) bool {
	var length = len(str)

	if &str == nil || length == 0 {
		return true
	}

	for i := 0; i < length; i++ {
		
	
	}
	
	return false
}




func CharAt(str string,at int) byte {
	var code =	str[at]
	var value = reflect.ValueOf(code)

	switch(value.Kind()){




	}
}
