package envar

import (
	"github.com/goccha/log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Get(names ...string) Env {
	for _, name := range names {
		name = strings.Trim(name, " ")
		v, ok := os.LookupEnv(name)
		if ok {
			return Env{Name: name, value: v}
		}
	}
	return Env{}
}

type Env struct {
	Name  string
	value string
}

func (e Env) Has() bool {
	return e.Name != ""
}

func (e Env) Bool(defaultValue bool) bool {
	if e.value != "" {
		if v, err := strconv.ParseBool(e.value); err != nil {
			log.Warn("%+v", err)
			return defaultValue
		} else {
			return v
		}
	}
	return defaultValue
}
func (e Env) Int(defaultValue int) int {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return i
	}
	return defaultValue
}
func (e Env) Int8(defaultValue int8) int8 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int8(i)
	}
	return defaultValue
}
func (e Env) Int16(defaultValue int16) int16 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int16(i)
	}
	return defaultValue
}
func (e Env) Int32(defaultValue int32) int32 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int32(i)
	}
	return defaultValue
}
func (e Env) Int64(defaultValue int64) int64 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int64(i)
	}
	return defaultValue
}
func (e Env) Uint(defaultValue uint) uint {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint(i)
	}
	return defaultValue
}
func (e Env) Uint8(defaultValue uint8) uint8 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint8(i)
	}
	return defaultValue
}
func (e Env) Uint16(defaultValue uint16) uint16 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint16(i)
	}
	return defaultValue
}
func (e Env) Uint32(defaultValue uint32) uint32 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint32(i)
	}
	return defaultValue
}
func (e Env) Uint64(defaultValue uint64) uint64 {
	if e.value != "" {
		i, err := strconv.Atoi(e.value)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint64(i)
	}
	return defaultValue
}
func (e Env) Float32(defaultValue float32) float32 {
	if e.value != "" {
		value, err := strconv.ParseFloat(e.value, 32)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return float32(value)
	}
	return defaultValue
}
func (e Env) Float64(defaultValue float64) float64 {
	if e.value != "" {
		value, err := strconv.ParseFloat(e.value, 64)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return value
	}
	return defaultValue
}
func (e Env) Complex64(defaultValue complex64) complex64 {
	if e.value != "" {
		value, err := strconv.ParseComplex(e.value, 64)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return complex64(value)
	}
	return defaultValue
}
func (e Env) Complex128(defaultValue complex128) complex128 {
	if e.value != "" {
		value, err := strconv.ParseComplex(e.value, 128)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return value
	}
	return defaultValue
}
func (e Env) String(defaultValue string) string {
	if e.value != "" {
		return e.value
	}
	return defaultValue
}
func (e Env) Split(defaultValue, sep string) []string {
	a := strings.Split(e.String(defaultValue), sep)
	if strings.Contains(sep, " ") {
		buf := make([]string, 0, len(a))
		for i := range a {
			v := strings.TrimSpace(a[i])
			if v != "" {
				buf = append(buf, v)
			}
		}
		return buf
	}
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}
func (e Env) ToUpper(defaultValue string) string {
	return strings.ToUpper(e.String(defaultValue))
}
func (e Env) ToLower(defaultValue string) string {
	return strings.ToLower(e.String(defaultValue))
}
func (e Env) Duration(defaultValue time.Duration) time.Duration {
	if e.value != "" {
		d, err := time.ParseDuration(e.value)
		if err == nil {
			return d
		}
	}
	return defaultValue
}

func Has(names ...string) bool {
	return Get(names...).Has()
}

func Bool(names ...string) bool {
	return Get(names...).Bool(false)
}

func String(names ...string) string {
	return Get(names...).String("")
}

func ToUpper(names ...string) string {
	return Get(names...).ToUpper("")
}

func ToLower(names ...string) string {
	return Get(names...).ToLower("")
}

func Int(names ...string) int {
	return Get(names...).Int(0)
}

func Int8(names ...string) int8 {
	return Get(names...).Int8(0)
}

func Int16(names ...string) int16 {
	return Get(names...).Int16(0)
}

func Int32(names ...string) int32 {
	return Get(names...).Int32(0)
}

func Int64(names ...string) int64 {
	return Get(names...).Int64(0)
}

func Uint(names ...string) uint {
	return Get(names...).Uint(0)
}

func Uint8(names ...string) uint8 {
	return Get(names...).Uint8(0)
}

func Uint16(names ...string) uint16 {
	return Get(names...).Uint16(0)
}

func Uint32(names ...string) uint32 {
	return Get(names...).Uint32(0)
}

func Uint64(names ...string) uint64 {
	return Get(names...).Uint64(0)
}

func Float32(names ...string) float32 {
	return Get(names...).Float32(0)
}

func Float64(names ...string) float64 {
	return Get(names...).Float64(0)
}

func Complex64(names ...string) complex64 {
	return Get(names...).Complex64(0)
}

func Complex128(names ...string) complex128 {
	return Get(names...).Complex128(0)
}

func Duration(names ...string) time.Duration {
	return Get(names...).Duration(0)
}

func GetHostname(defaultName string) string {
	name, err := os.Hostname()
	if err != nil {
		log.Error("%+v", err)
		return defaultName
	}
	return name
}
