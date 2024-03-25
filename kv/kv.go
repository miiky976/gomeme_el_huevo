package kv

import "fmt"

var KV []map[string]interface{}

func Test() {
	KV = append(KV, map[string]interface{}{"name": "Miguel"})
}

func SET(key string, value interface{}) {
	for i, kv := range KV {
		if _, ok := kv[key]; ok {
			fmt.Println(i)
			KV[i] = map[string]interface{}{key: value}
			return
		}
	}
	KV = append(KV, map[string]interface{}{key: value})
}

func GET(key string) interface{} {
	for _, kv := range KV {
		if val, ok := kv[key]; ok {
			return val
		}
	}
	return ""
}

func DEL(key string) {
	for i, kv := range KV {
		if _, ok := kv[key]; ok {
			KV = append(KV[:i], KV[i+1:]...)
		}
	}
}

func INCR(key string, newVal int) {
	for i, kv := range KV {
		if val, ok := kv[key]; ok {
			if _, intOK := val.(int); intOK {
				KV[i] = map[string]interface{}{key: val.(int) + newVal}
			}
		}
	}
}
