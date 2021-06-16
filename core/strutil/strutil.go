package strutil

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/wang249639015/gouqi-tool/core/runeutil"
)

// 字符串是否为nil 或者 ''
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

/**
 * <p>如果对象是字符串是否为空白，空白的定义如下：</p>
 * <ol>
 *     <li>{@code null}</li>
 *     <li>空字符串：{@code ""}</li>
 *     <li>空格、全角空格、制表符、换行符，等不可见字符</li>
 * </ol>
 *
 * <p>例：</p>
 * <ul>
 *     <li>{@code StrUtil.isBlankIfStr(null)     // true}</li>
 *     <li>{@code StrUtil.isBlankIfStr("")       // true}</li>
 *     <li>{@code StrUtil.isBlankIfStr(" \t\n")  // true}</li>
 *     <li>{@code StrUtil.isBlankIfStr("abc")    // false}</li>
 * </ul>
 */
func IsBlank(str *string) bool {
	var length = len(*str)

	if &str == nil || length == 0 {
		return true
	}
	for v := range *str {
		if !runeutil.IsBlankRune(rune(v)) {
			return false
		}
	}
	return true
}

func IsBlankIfStr(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	if kind == reflect.String {
		value, ok := obj.(string)
		if ok {
			return IsBlank(&value)
		} else {
			return false
		}

	} else {
		return false
	}

}

func TrimArray(strs []string) {
	if strs == nil {
		return
	}
	for i, v := range strs {
		strs[i] = Trim(&v)
	}

}

func Trim(str *string) string {
	return strings.TrimSpace(*str)
}

//string 类型在golang中

func stringTest() {
	s := "测试"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	for _, v := range s {
		fmt.Println(v)
	}
}
