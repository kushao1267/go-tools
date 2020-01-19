package funcs

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy 深拷贝（golang是值传递，如果被拷贝对象内有指针变量时，将只会拷贝指针值）
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
