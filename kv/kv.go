package kv

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
	SET("key", []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris laoreet, nisi at sodales tristique, quam nibh cursus nisi, eu vestibulum orci dolor sit amet magna.\nNullam quis turpis nisi. Aliquam consequat, tortor sit amet tempor lobortis, sem nisl maximus nulla, id ornare dolor neque vel neque.\nMauris ut mollis dui, id pellentesque purus. Quisque pharetra pellentesque leo consequat consequat. Vestibulum dapibus lectus quam, vitae facilisis eros vulputate sed.\nSuspendisse viverra mi ac erat lobortis venenatis. Duis ultrices ipsum vitae mi finibus, a elementum felis lobortis.\nPhasellus id lacus orci. Phasellus imperdiet cursus dolor, eu mattis neque tristique eu. Donec non lorem justo."), "string")
	SET("newkey", []byte("newvalue"), "string")
	SET("newone", []byte("newone"), "string")
	SET("aaa", []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris laoreet, nisi at sodales tristique, quam nibh cursus nisi, eu vestibulum orci dolor sit amet magna.\nNullam quis turpis nisi. Aliquam consequat, tortor sit amet tempor lobortis, sem nisl maximus nulla, id ornare dolor neque vel neque.\nMauris ut mollis dui, id pellentesque purus. Quisque pharetra pellentesque leo consequat consequat. Vestibulum dapibus lectus quam, vitae facilisis eros vulputate sed.\nSuspendisse viverra mi ac erat lobortis venenatis. Duis ultrices ipsum vitae mi finibus, a elementum felis lobortis.\nPhasellus id lacus orci. Phasellus imperdiet cursus dolor, eu mattis neque tristique eu. Donec non lorem justo."), "string")
	SET("sam", []byte("jflaskd"), "string")
	SET("des", []byte("mck"), "string")
}
