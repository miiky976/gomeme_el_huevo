package kv

import "fmt"

var KV []*kv

// why struct?
// because i want to store the type
// of the value, because it can be a string but also a file binary
// and for my use case i need to know the type

type kv struct {
	Key   string
	Value []byte
	Type  string
}

func SET(key string, value []byte, Type string) {
	for _, v := range KV {
		if v.Key == key {
			v.Value = value
			v.Type = Type
			return
		}
	}
	KV = append(KV, &kv{key, value, Type})
}

func GET(key string) ([]byte, string) {
	for _, v := range KV {
		if v.Key == key {
			return v.Value, v.Type
		}
	}
	return nil, ""
}

func DEL(key string) {
	for i, v := range KV {
		if v.Key == key {
			KV = append(KV[:i], KV[i+1:]...)
			return
		}
	}
}

func Test() {
	SET("key", []byte("value"), "string")
	fmt.Println(GET("key"))
	SET("key", []byte("XDXD"), "string")
	fmt.Println(GET("key"))
	SET("newkey", []byte("newvalue"), "string")
	SET("newone", []byte("newone"), "string")
	fmt.Println(GET("newkey"))
	fmt.Println(GET("newone"))
	DEL("newkey")
	fmt.Println(GET("newkey"))

}
