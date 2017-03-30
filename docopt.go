package commander

import (
	"fmt"
	"strconv"
)

// DocoptMap docopt returns a map of option names to the values
type DocoptMap map[string]interface{}

func newDocoptMap(m map[string]interface{}) DocoptMap {
	return DocoptMap(m)
}

func (d DocoptMap) Map() map[string]interface{} {
	return map[string]interface{}(d)
}

func (d DocoptMap) Get(key string) interface{} {
	if v, ok := d[key]; ok {
		return v
	}
	return nil
}

func (d DocoptMap) Contain(key string) bool {
	if v, ok := d[key]; ok && v != nil {
		switch i := v.(type) {
		case string:
			return len(i) != 0 &&
				i != "0" && i != "false"
		case []string:
			return len(i) != 0
		case bool:
			return i
		case int, int8, int16, int32, int64:
			return i != 0
		case float32, float64:
			return i != 0
		}
	}
	return false
}

func (d DocoptMap) GetString(key string) (string, bool) {
	if v := d.Get(key); v != nil {
		if s, ok := v.(string); ok {
			return s, true
		}
		return fmt.Sprintf("%v", v), true
	}
	return "", false
}

func (d DocoptMap) MustString(key string) string {
	if s, ok := d.GetString(key); ok {
		return s
	}
	return ""
}

func (d DocoptMap) GetStrings(key string) ([]string, bool) {
	if v := d.Get(key); v != nil {
		if s, ok := v.([]string); ok {
			return s, true
		}
	}
	return nil, false
}

func (d DocoptMap) MustStrings(key string) []string {
	if s, ok := d.GetStrings(key); ok {
		return s
	}
	return []string{}
}

func (d DocoptMap) GetBool(key string) (bool, bool) {
	b, err := strconv.ParseBool(d.MustString(key))
	return b, err == nil
}

func (d DocoptMap) MustBool(key string) bool {
	b, ok := d.GetBool(key)
	return b && ok
}

func (d DocoptMap) GetInt64(key string) (int64, bool) {
	i, err := strconv.ParseInt(d.MustString(key), 10, 64)
	return i, err == nil
}

func (d DocoptMap) MustInt64(key string) int64 {
	if i, ok := d.GetInt64(key); ok {
		return i
	}
	return 0
}

func (d DocoptMap) GetInt(key string) (int, bool) {
	i, err := strconv.ParseInt(d.MustString(key), 10, 32)
	return int(i), err == nil
}

func (d DocoptMap) MustInt(key string) int {
	if i, ok := d.GetInt(key); ok {
		return i
	}
	return 0
}

func (d DocoptMap) GetFloat64(key string) (float64, bool) {
	f, err := strconv.ParseFloat(d.MustString(key), 64)
	return f, err == nil
}

func (d DocoptMap) MustFloat64(key string) float64 {
	if f, ok := d.GetFloat64(key); ok {
		return f
	}
	return 0
}

func (d DocoptMap) GetFloat(key string) (float32, bool) {
	f, err := strconv.ParseFloat(d.MustString(key), 32)
	return float32(f), err == nil
}

func (d DocoptMap) MustFloat(key string) float32 {
	if f, ok := d.GetFloat(key); ok {
		return f
	}
	return 0
}
