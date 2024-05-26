package kv

import "errors"

var Master []*KV

// why struct?
// because i want to store the type
// of the value, because it can be a string but also a file binary
// and for my use case i need to know the type

type KV struct {
	Value []byte
	Type  string
}

func Create(value []byte, Type string) *KV {
	Master = append(Master, &KV{Value: value, Type: Type})
	return &KV{Value: value, Type: Type}
}

func Read(i int) (*KV, error) {
	elem := Master[i]
	if elem == nil {
		return nil, errors.New("Not found")
	}
	return elem, nil
}

func Update(i int, value []byte, typo string) (*KV, error) {
	elem := Master[i]
	if elem == nil {
		return nil, errors.New("Not found")
	}
	elem = &KV{value, typo}
	return elem, nil
}

func Delete(i int) error {
	elem := Master[i]
	if elem == nil {
		return errors.New("not exists")
	}

	Master = append(Master[:i], Master[i+1:]...)

	return nil
}

func Test() {
}
