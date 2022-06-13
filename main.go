package main

/*
#include<stdbool.h>
*/
import "C"

var storage = make(map[string]string)

//export reset
func reset() {
	storage = make(map[string]string)
}

//export size
func size() C.uint {
	return C.uint(len(storage))
}

//export includes
func includes(key *C.char) C.bool {
	var localKey = C.GoString(key)
	var _, ok = storage[localKey]
	return C.bool(ok)
}

//export get
func get(key *C.char) *C.char {
	var localKey = C.GoString(key)
	var found, ok = storage[localKey]
	if ok {
		return C.CString(found)
	} else {
		return nil
	}
}

//export set
func set(k *C.char, v *C.char) *C.char {
	var local_key = C.GoString(k)
	var local_value = C.GoString(v)
	storage[local_key] = local_value
	return k
}

//export remove
func remove(k *C.char) {
	var local_key = C.GoString(k)
	delete(storage, local_key)
}

func main() {
	reset()
}
