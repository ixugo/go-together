package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"sync"
)

type BodyMap map[string]interface{}

var mu = new(sync.RWMutex)

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) BodyMap {
	mu.Lock()
	bm[key] = value
	mu.Unlock()
	return bm
}

func (bm BodyMap) SetBodyMap(key string, value func(bm BodyMap)) BodyMap {
	_bm := make(BodyMap)
	value(_bm)

	mu.Lock()
	bm[key] = _bm
	mu.Unlock()
	return bm
}

// 获取参数转换string
func (bm BodyMap) GetString(key string) string {
	if bm == nil {
		return ""
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := bm[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		// 不是string强行转json字符串
		return convertToString(value)
	}
	return v
}

// 获取原始参数
func (bm BodyMap) GetInterface(key string) interface{} {
	if bm == nil {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	return bm[key]
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	mu.Lock()
	delete(bm, key)
	mu.Unlock()
}

// 置空BodyMap
func (bm BodyMap) Reset() {
	mu.Lock()
	for k := range bm {
		delete(bm, k)
	}
	mu.Unlock()
}

func (bm BodyMap) JsonBody() (jb string) {
	mu.Lock()
	defer mu.Unlock()
	bs, err := json.Marshal(bm)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == "" {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return errors.New(strings.Join(emptyKeys, ", ") + " : cannot be empty")
	}
	return nil
}

func (bm BodyMap) BufferBody() *bytes.Buffer {
	mu.Lock()
	defer mu.Unlock()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(bm)
	if err != nil {
		panic(err)
	}
	return &buf
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
