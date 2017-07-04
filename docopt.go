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

// Map returns a map type of it.
func (d DocoptMap) Map() map[string]interface{} {
	return map[string]interface{}(d)
}

// Get returns the value by the key.
func (d DocoptMap) Get(key string) interface{} {
	if v, ok := d[key]; ok {
		return v
	}
	return nil
}

// Contain returns true if it contains the key.
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

// GetString returns a string type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetString(key string) (string, bool) {
	if v := d.Get(key); v != nil {
		if s, ok := v.(string); ok {
			return s, true
		}
		return fmt.Sprintf("%v", v), true
	}
	return "", false
}

// MustString returns a string type value by key.
// Returns empty string if not exist for the key.
func (d DocoptMap) MustString(key string) string {
	if s, ok := d.GetString(key); ok {
		return s
	}
	return ""
}

// GetStrings returns a slice of string type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetStrings(key string) ([]string, bool) {
	if v := d.Get(key); v != nil {
		if s, ok := v.([]string); ok {
			return s, true
		}
	}
	return nil, false
}

// MustStrings returns a slice of string type value by key.
// Returns empty slice if not exist for the key.
func (d DocoptMap) MustStrings(key string) []string {
	if s, ok := d.GetStrings(key); ok {
		return s
	}
	return []string{}
}

// GetBool returns a boolean type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetBool(key string) (bool, bool) {
	b, err := strconv.ParseBool(d.MustString(key))
	return b, err == nil
}

// MustBool returns a boolean type value by key.
// Returns false if not exist for the key.
func (d DocoptMap) MustBool(key string) bool {
	b, ok := d.GetBool(key)
	return b && ok
}

// GetInt64 returns a int64 type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetInt64(key string) (int64, bool) {
	i, err := strconv.ParseInt(d.MustString(key), 10, 64)
	return i, err == nil
}

// MustInt64 returns a int64 type value by key.
// Returns 0 if not exist for the key.
func (d DocoptMap) MustInt64(key string) int64 {
	if i, ok := d.GetInt64(key); ok {
		return i
	}
	return 0
}

// GetInt returns a int type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetInt(key string) (int, bool) {
	i, err := strconv.ParseInt(d.MustString(key), 10, 32)
	return int(i), err == nil
}

// MustInt returns a int type value by key.
// Returns 0 if not exist for the key.
func (d DocoptMap) MustInt(key string) int {
	if i, ok := d.GetInt(key); ok {
		return i
	}
	return 0
}

// GetFloat64 returns a float64 type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetFloat64(key string) (float64, bool) {
	f, err := strconv.ParseFloat(d.MustString(key), 64)
	return f, err == nil
}

// MustFloat64 returns a float64 type value by key.
// Returns 0 if not exist for the key.
func (d DocoptMap) MustFloat64(key string) float64 {
	if f, ok := d.GetFloat64(key); ok {
		return f
	}
	return 0
}

// GetFloat returns a float32 type value by the key.
// Second returns false if not exist for the key.
func (d DocoptMap) GetFloat(key string) (float32, bool) {
	f, err := strconv.ParseFloat(d.MustString(key), 32)
	return float32(f), err == nil
}

// MustFloat returns a float32 type value by key.
// Returns 0 if not exist for the key.
func (d DocoptMap) MustFloat(key string) float32 {
	if f, ok := d.GetFloat(key); ok {
		return f
	}
	return 0
}
