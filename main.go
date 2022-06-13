package main

import "C"

var store = make(map[string]string)

//export getVersion
func getVersion() *C.char {
	var version = C.CString("4.3")
	return version
}

//export size
func size() C.uint {
	return C.uint(len(store))
}

//export includes
// func includes(key *C.char) *C.BOOL {
// 	var localKey = C.GoString(key)
// 	var _, ok = store[localKey]
// 	return ok
// }

//export get
func get(key *C.char) *C.char {
	var localKey = C.GoString(key)
	var found, ok = store[localKey]
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
	store[local_key] = local_value
	return k
}

//export remove
func remove(k *C.char) {
	var local_key = C.GoString(k)
	delete(store, local_key)
}

func main() {

}
