package envar

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goccha/envar/pkg/log"
)

func Get(names ...string) Env {
	for _, name := range names {
		name = strings.TrimSpace(name)
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
func (e Env) Writer(defaultValue string) *os.File {
	value := defaultValue
	if e.value != "" {
		value = e.value
	}
	if value != "" {
		if f, err := os.Create(value); err != nil {
			log.Warn("%+v", err)
		} else {
			return f
		}
	}
	return os.Stdout
}
func (e Env) Reader(defaultValue string) *os.File {
	value := defaultValue
	if e.value != "" {
		value = e.value
	}
	if value != "" {
		if f, err := os.Open(value); err != nil {
			log.Warn("%+v", err)
		} else {
			return f
		}
	}
	return os.Stdin
}
func (e Env) String(defaultValue string) string {
	if e.value != "" {
		return e.value
	}
	return defaultValue
}
func (e Env) Bytes(defaultValue string) []byte {
	if e.value != "" {
		return []byte(e.value)
	}
	return []byte(defaultValue)
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
func (e Env) ByteSlice(defaultValue []byte, sep string) []byte {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	e.String("")
	array := make([]byte, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseUint(s, 10, 8); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, byte(v))
		}
	}
	return array
}
func (e Env) IntSlice(defaultValue []int, sep string) []int {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]int, 0, len(a))
	for i, s := range a {
		if v, err := strconv.Atoi(s); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, v)
		}
	}
	return array
}
func (e Env) Int8Slice(defaultValue []int8, sep string) []int8 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]int8, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseInt(s, 10, 8); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, int8(v))
		}
	}
	return array
}
func (e Env) Int16Slice(defaultValue []int16, sep string) []int16 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]int16, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseInt(s, 10, 16); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, int16(v))
		}
	}
	return array
}
func (e Env) Int32Slice(defaultValue []int32, sep string) []int32 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]int32, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseInt(s, 10, 16); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, int32(v))
		}
	}
	return array
}
func (e Env) Int64Slice(defaultValue []int64, sep string) []int64 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]int64, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseInt(s, 10, 64); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, v)
		}
	}
	return array
}
func (e Env) UintSlice(defaultValue []uint, sep string) []uint {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]uint, 0, len(a))
	for i, s := range a {
		if v, err := strconv.Atoi(s); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, uint(v))
		}
	}
	return array
}
func (e Env) Uint8Slice(defaultValue []uint8, sep string) []uint8 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]uint8, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseUint(s, 10, 8); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, uint8(v))
		}
	}
	return array
}
func (e Env) Uint16Slice(defaultValue []uint16, sep string) []uint16 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]uint16, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseUint(s, 10, 16); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, uint16(v))
		}
	}
	return array
}
func (e Env) Uint32Slice(defaultValue []uint32, sep string) []uint32 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]uint32, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseUint(s, 10, 16); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, uint32(v))
		}
	}
	return array
}
func (e Env) Uint64Slice(defaultValue []uint64, sep string) []uint64 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]uint64, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseUint(s, 10, 64); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, v)
		}
	}
	return array
}
func (e Env) Float32Slice(defaultValue []float32, sep string) []float32 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]float32, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseFloat(s, 32); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, float32(v))
		}
	}
	return array
}
func (e Env) Float64Slice(defaultValue []float64, sep string) []float64 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]float64, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseFloat(s, 64); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, v)
		}
	}
	return array
}
func (e Env) Complex64Slice(defaultValue []complex64, sep string) []complex64 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]complex64, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseComplex(s, 64); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, complex64(v))
		}
	}
	return array
}
func (e Env) Complex128Slice(defaultValue []complex128, sep string) []complex128 {
	a := e.Split("", sep)
	if len(a) == 0 {
		return defaultValue
	}
	array := make([]complex128, 0, len(a))
	for i, s := range a {
		if v, err := strconv.ParseComplex(s, 128); err != nil {
			log.Warn("[%d]:%v", i, err)
		} else {
			array = append(array, v)
		}
	}
	return array
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

func Split(names ...string) []string {
	return Get(names...).Split("", ",")
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

func IntSlice(names ...string) []int {
	return Get(names...).IntSlice([]int{}, ",")
}

func Int8(names ...string) int8 {
	return Get(names...).Int8(0)
}

func Int8Slice(names ...string) []int8 {
	return Get(names...).Int8Slice([]int8{}, ",")
}

func Int16(names ...string) int16 {
	return Get(names...).Int16(0)
}

func Int16Slice(names ...string) []int16 {
	return Get(names...).Int16Slice([]int16{}, ",")
}

func Int32(names ...string) int32 {
	return Get(names...).Int32(0)
}

func Int32Slice(names ...string) []int32 {
	return Get(names...).Int32Slice([]int32{}, ",")
}

func Int64(names ...string) int64 {
	return Get(names...).Int64(0)
}

func Int64Slice(names ...string) []int64 {
	return Get(names...).Int64Slice([]int64{}, ",")
}

func Uint(names ...string) uint {
	return Get(names...).Uint(0)
}

func UintSlice(names ...string) []uint {
	return Get(names...).UintSlice([]uint{}, ",")
}

func Uint8(names ...string) uint8 {
	return Get(names...).Uint8(0)
}

func Uint8Slice(names ...string) []uint8 {
	return Get(names...).Uint8Slice([]uint8{}, ",")
}

func Uint16(names ...string) uint16 {
	return Get(names...).Uint16(0)
}

func Uint16Slice(names ...string) []uint16 {
	return Get(names...).Uint16Slice([]uint16{}, ",")
}

func Uint32(names ...string) uint32 {
	return Get(names...).Uint32(0)
}

func Uint32Slice(names ...string) []uint32 {
	return Get(names...).Uint32Slice([]uint32{}, ",")
}

func Uint64(names ...string) uint64 {
	return Get(names...).Uint64(0)
}

func Uint64Slice(names ...string) []uint64 {
	return Get(names...).Uint64Slice([]uint64{}, ",")
}

func Float32(names ...string) float32 {
	return Get(names...).Float32(0)
}

func Float32Slice(names ...string) []float32 {
	return Get(names...).Float32Slice([]float32{}, ",")
}

func Float64(names ...string) float64 {
	return Get(names...).Float64(0)
}

func Float64Slice(names ...string) []float64 {
	return Get(names...).Float64Slice([]float64{}, ",")
}

func Complex64(names ...string) complex64 {
	return Get(names...).Complex64(0)
}

func Complex64Slice(names ...string) []complex64 {
	return Get(names...).Complex64Slice([]complex64{}, ",")
}

func Complex128(names ...string) complex128 {
	return Get(names...).Complex128(0)
}

func Complex128Slice(names ...string) []complex128 {
	return Get(names...).Complex128Slice([]complex128{}, ",")
}

func Writer(names ...string) *os.File {
	return Get(names...).Writer("")
}

func Reader(names ...string) *os.File {
	return Get(names...).Reader("")
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
