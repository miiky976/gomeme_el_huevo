package kv

import "errors"

var kv = make(map[uint]*Data)

// why struct?
// because i want to store the type
// of the value, because it can be a string but also a file binary
// and for my use case i need to know the type

type Data struct {
	Value []byte
	Type  string
}

func Create(value []byte, Type string) {
	key := GetNewKey()
	kv[key] = &Data{Value: value, Type: Type}
}

func Read(key uint) *Data {
	return kv[key]
}

func Update(key uint, val []byte, typ string) (*Data, error) {
	if kv[key] == nil {
		return nil, errors.New("Not found")
	}
	kv[key] = &Data{val, typ}
	return kv[key], nil
}

func Delete(key uint) error {
	if kv[key] == nil {
		return errors.New("Not found")
	}
	delete(kv, key)
	return nil
}

func Test() {
	Create([]byte("test"), "string")
}

func GetNewKey() uint {
	if len(kv) == 0 {
		return 0
	}
	keys := make([]uint, 0, len(kv)-1)
	for k := range kv {
		keys = append(keys, k)
	}
	return keys[len(keys)-1] + 1
}
