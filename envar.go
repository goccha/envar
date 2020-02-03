package envar

import (
	"github.com/goccha/log"
	"os"
	"strconv"
	"time"
)

func Get(names ...string) Env {
	return Env{Names: names}
}

type Env struct {
	Names []string
}

func (e Env) lookup() (string, bool) {
	for _, name := range e.Names {
		v, ok := os.LookupEnv(name)
		if ok {
			return v, true
		}
	}
	return "", false
}

func (e Env) Has() bool {
	if _, ok := e.lookup(); ok {
		return true
	}
	return false
}

func (e Env) Bool(defaultValue bool) bool {
	if v, ok := e.lookup(); ok {
		if v == "true" {
			return true
		} else {
			return false
		}
	}
	return defaultValue
}
func (e Env) Int(defaultValue int) int {
	if v, ok := e.lookup(); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return i
	}
	return defaultValue
}
func (e Env) Int32(defaultValue int32) int32 {
	if v, ok := e.lookup(); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int32(i)
	}
	return defaultValue
}
func (e Env) Int64(defaultValue int64) int64 {
	if v, ok := e.lookup(); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return int64(i)
	}
	return defaultValue
}
func (e Env) Uint32(defaultValue uint32) uint32 {
	if v, ok := e.lookup(); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint32(i)
	}
	return defaultValue
}
func (e Env) Uint64(defaultValue uint64) uint64 {
	if v, ok := e.lookup(); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return uint64(i)
	}
	return defaultValue
}
func (e Env) Float32(defaultValue float32) float32 {
	if v, ok := e.lookup(); ok {
		value, err := strconv.ParseFloat(v, 32)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return float32(value)
	}
	return defaultValue
}
func (e Env) Float64(defaultValue float64) float64 {
	if v, ok := e.lookup(); ok {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Warn("%+v", err)
			return defaultValue
		}
		return float64(value)
	}
	return defaultValue
}
func (e Env) String(defaultValue string) string {
	if v, ok := e.lookup(); ok {
		return v
	}
	return defaultValue
}
func (e Env) Duration(defaultValue time.Duration) time.Duration {
	if v, ok := e.lookup(); ok {
		d, err := time.ParseDuration(v)
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

func Int(names ...string) int {
	return Get(names...).Int(0)
}

func Int32(names ...string) int32 {
	return Get(names...).Int32(0)
}

func Int64(names ...string) int64 {
	return Get(names...).Int64(0)
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
