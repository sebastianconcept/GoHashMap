package main

import "C"
import (
	"github.com/cornelk/hashmap"
)

// var store = make(map[string]string)
var storage = &hashmap.HashMap{}

//export size
func size() C.uint {
	// return C.uint(len(store))
	return C.uint(storage.Len())
}

//export includes
// func includes(key *C.char) *C.bool {
// 	var localKey = C.GoString(key)
// 	// var _, ok = store[localKey]
// 	var _, ok = storage.Get(localKey)
// 	return C.CBOOL(ok)
// }

//export get
func get(key *C.char) *C.char {
	var localKey = C.GoString(key)
	// var found, ok = store[localKey]
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
	// store[local_key] = local_value
	storage.Set(local_key, local_value)
	return k
}

//export remove
func remove(k *C.char) {
	var local_key = C.GoString(k)
	// delete(store, local_key)
	storage.Del(local_key)
}

func main() {

}
