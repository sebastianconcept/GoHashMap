package main

/*
#include<stdbool.h>
*/
import "C"
import (
	"github.com/cornelk/hashmap"
)

var defaultSize uintptr = 50
var storage = hashmap.New(defaultSize)

//export reset
func reset() {
	storage = hashmap.New(defaultSize)
}

//export size
func size() C.uint {
	return C.uint(storage.Len())
}

//export includes
func includes(key *C.char) C.bool {
	var localKey = C.GoString(key)
	var _, ok = storage.Get(localKey)
	return C.bool(ok)
}

//export get
func get(key *C.char) *C.char {
	var localKey = C.GoString(key)
	var found, ok = storage.Get(localKey)
	if ok {
		return C.CString(found.(string))
	} else {
		return nil
	}
}

//export set
func set(k *C.char, v *C.char) *C.char {
	var local_key = C.GoString(k)
	var local_value = C.GoString(v)
	storage.Set(local_key, local_value)
	return k
}

//export remove
func remove(k *C.char) {
	var local_key = C.GoString(k)
	storage.Del(local_key)
}

func main() {
	reset()
}
