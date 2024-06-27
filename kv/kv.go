package kv

import (
	"errors"
	"fmt"
	"slices"
)

var kv = make(map[uint]*Data)
var lastkey uint = 0

// why struct?
// because i want to store the type
// of the value, because it can be a string but also a file binary
// and for my use case i need to know the type

type Data struct {
	Value []byte
	Type  string
}

func Create(value []byte, Type string) uint {
	key := lastkey // GetNewKey()
	lastkey++
	fmt.Println("new:", key)
	kv[key] = &Data{Value: value, Type: Type}
	return key
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
	Create([]byte("test2"), "string")
	Create([]byte("test3"), "string")
	Create([]byte("test4"), "string")
}

func GetKeys() []uint {
	if len(kv) == 0 {
		return []uint{}
	}
	keys := make([]uint, 0, len(kv))
	for k := range kv {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	return keys
}
