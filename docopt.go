package commander

import (
	"fmt"
	"strconv"
)

type DocoptMap map[string]interface{}

func (d DocoptMap) Get(key string) interface{} {
	if v, ok := d[key]; ok {
		return v
	}
	return nil
}

func (d DocoptMap) Contain(key string) bool {
	_, ok := d[key]
	return ok
}

func (d DocoptMap) GetString(key string) string {
	if v := d.Get(key); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func (d DocoptMap) GetBool(key string) (bool, bool) {
	b, err := strconv.ParseBool(d.GetString(key))
	return b, err == nil
}

func (d DocoptMap) GetInt64(key string) (int64, bool) {
	i, err := strconv.ParseInt(d.GetString(key), 10, 64)
	return i, err == nil
}

func (d DocoptMap) GetInt(key string) (int, bool) {
	i, err := strconv.ParseInt(d.GetString(key), 10, 32)
	return int(i), err == nil
}

func (d DocoptMap) GetFloat64(key string) (float64, bool) {
	f, err := strconv.ParseFloat(d.GetString(key), 64)
	return f, err == nil
}

func (d DocoptMap) GetFloat(key string) (float32, bool) {
	f, err := strconv.ParseFloat(d.GetString(key), 32)
	return float32(f), err == nil
}

func (d DocoptMap) String() string {
	return fmt.Sprintf("%#v", d)
}
