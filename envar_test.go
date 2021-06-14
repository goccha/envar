package envar

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	_ = os.Setenv("TEST_BOOL_TRUE", "true")
	assert.Equal(t, true, Bool("TEST_BOOL_TRUE"))
	assert.NotEqual(t, true, Bool("TEST_BOOL_FALSE"))

	_ = os.Setenv("TEST_BOOL_TRUE", "1")
	assert.Equal(t, true, Bool("TEST_BOOL_TRUE"))

	_ = os.Setenv("TEST_BOOL", "0")
	assert.Equal(t, false, Bool("TEST_BOOL"))
}

func TestInt(t *testing.T) {
	_ = os.Setenv("TEST_INT_1", "1")
	_ = os.Setenv("TEST_INT_2", "2")
	assert.Equal(t, 1, Get("TEST_INT_1", "TEST_INT_2").Int(2))
	assert.Equal(t, 2, Get("TEST_INT_0", "TEST_INT_2").Int(-1))
	assert.NotEqual(t, 1, Get("TEST_INT_0").Int(3))

	_ = os.Setenv("TEST_INT_1", "zero")
	assert.Equal(t, 2, Get("TEST_INT_1").Int(2))
}

func TestInt8(t *testing.T) {
	_ = os.Setenv("TEST_INT8_1", "1")
	_ = os.Setenv("TEST_INT8_2", "2")
	assert.Equal(t, int8(1), Int8("TEST_INT8_1", "TEST_INT_2"))
	assert.Equal(t, int8(2), Get("TEST_INT8_0", "TEST_INT_2").Int8(-1))
	assert.NotEqual(t, int8(1), Get("TEST_INT8_0").Int8(3))

	_ = os.Setenv("TEST_INT8_1", "zero")
	assert.Equal(t, int8(2), Get("TEST_INT8_1").Int8(2))
}

func TestInt16(t *testing.T) {
	_ = os.Setenv("TEST_INT16_1", "1")
	_ = os.Setenv("TEST_INT16_2", "2")
	assert.Equal(t, int16(1), Int16("TEST_INT16_1", "TEST_INT_2"))
	assert.Equal(t, int16(2), Get("TEST_INT16_0", "TEST_INT16_2").Int16(-1))
	assert.NotEqual(t, int16(1), Get("TEST_INT16_0").Int16(3))

	_ = os.Setenv("TEST_INT16_1", "zero")
	assert.Equal(t, int16(2), Get("TEST_INT16_1").Int16(2))
}

func TestInt32(t *testing.T) {
	_ = os.Setenv("TEST_INT32_1", "1")
	var expected int32 = 1
	assert.Equal(t, expected, Int32("TEST_INT32_1"))
	assert.NotEqual(t, expected, Get("TEST_INT32_2").Int32(2))

	_ = os.Setenv("TEST_INT32_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_INT32_1").Int32(2))
}

func TestInt64(t *testing.T) {
	_ = os.Setenv("TEST_INT64_1", "1")
	var expected int64 = 1
	assert.Equal(t, expected, Int64("TEST_INT64_1"))
	assert.NotEqual(t, expected, Get("TEST_INT64_2").Int64(2))

	_ = os.Setenv("TEST_INT64_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_INT64_1").Int64(2))
}

func TestUint(t *testing.T) {
	_ = os.Setenv("TEST_UINT_1", "1")
	var expected uint = 1
	assert.Equal(t, expected, Uint("TEST_UINT_1"))
	assert.NotEqual(t, expected, Get("TEST_UINT_2").Uint(2))

	_ = os.Setenv("TEST_UINT_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT_1").Uint(2))
}

func TestUint8(t *testing.T) {
	_ = os.Setenv("TEST_UINT8_1", "1")
	var expected uint8 = 1
	assert.Equal(t, expected, Uint8("TEST_UINT8_1"))
	assert.NotEqual(t, expected, Get("TEST_UINT8_2").Uint8(2))

	_ = os.Setenv("TEST_UINT8_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT8_1").Uint8(2))
}

func TestUint16(t *testing.T) {
	_ = os.Setenv("TEST_UINT16_1", "1")
	var expected uint16 = 1
	assert.Equal(t, expected, Uint16("TEST_UINT16_1"))
	assert.NotEqual(t, expected, Get("TEST_UINT16_2").Uint16(2))

	_ = os.Setenv("TEST_UINT16_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT16_1").Uint16(2))
}

func TestUint32(t *testing.T) {
	_ = os.Setenv("TEST_UINT32_1", "1")
	var expected uint32 = 1
	assert.Equal(t, expected, Uint32("TEST_UINT32_1"))
	assert.NotEqual(t, expected, Get("TEST_UINT32_2").Uint32(2))

	_ = os.Setenv("TEST_UINT32_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT32_1").Uint32(2))
}

func TestUint64(t *testing.T) {
	_ = os.Setenv("TEST_UINT64_1", "1")
	var expected uint64 = 1
	assert.Equal(t, expected, Uint64("TEST_UINT64_1"))
	assert.NotEqual(t, expected, Get("TEST_UINT64_2").Uint64(2))

	_ = os.Setenv("TEST_UINT64_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT64_1").Uint64(2))
}

func TestFloat32(t *testing.T) {
	_ = os.Setenv("TEST_FLOAT32_1", "1.0")
	var expected float32 = 1.0
	assert.Equal(t, expected, Float32("TEST_FLOAT32_1"))
	assert.NotEqual(t, expected, Get("TEST_FLOAT32_2").Float32(2.0))

	_ = os.Setenv("TEST_FLOAT32_1", "zero")
	expected = 2.0
	v := Get("TEST_FLOAT32_1").Float32(2.0)
	assert.Equal(t, expected, v)
	assert.NotEqual(t, 2, v)
}

func TestFloat64(t *testing.T) {
	_ = os.Setenv("TEST_FLOAT64_1", "1.0")
	assert.Equal(t, 1.0, Float64("TEST_FLOAT64_1"))
	assert.NotEqual(t, 1.0, Get("TEST_FLOAT64_2").Float64(2.0))

	_ = os.Setenv("TEST_FLOAT64_1", "zero")
	v := Get("TEST_FLOAT64_1").Float64(2.0)
	assert.Equal(t, 2.0, v)
	assert.NotEqual(t, 2, v)
}

func TestComplex64(t *testing.T) {
	_ = os.Setenv("TEST_COMPLEX64_1", "1.0")
	var expected complex64 = 1.0
	assert.Equal(t, expected, Complex64("TEST_COMPLEX64_1"))
	assert.NotEqual(t, expected, Get("TEST_COMPLEX64_2").Complex64(2.0))

	_ = os.Setenv("TEST_COMPLEX64_1", "zero")
	expected = 2.0
	v := Get("TEST_COMPLEX64_1").Complex64(2.0)
	assert.Equal(t, expected, v)
	assert.NotEqual(t, 2, v)
}

func TestComplex128(t *testing.T) {
	_ = os.Setenv("TEST_COMPLEX128_1", "1.0")
	var expected complex128 = 1.0
	assert.Equal(t, expected, Complex128("TEST_COMPLEX128_1"))
	assert.NotEqual(t, expected, Get("TEST_COMPLEX128_2").Complex128(2.0))

	_ = os.Setenv("TEST_COMPLEX128_1", "zero")
	expected = 2.0
	v := Get("TEST_COMPLEX128_1").Complex128(2.0)
	assert.Equal(t, expected, v)
	assert.NotEqual(t, 2, v)
}

func TestString(t *testing.T) {
	_ = os.Setenv("TEST_STRING", "test")
	assert.Equal(t, "test", String("TEST_STRING"))
	assert.NotEqual(t, "test", Get("TEST_STRING_PROD").String("prod"))
}

func TestDuration(t *testing.T) {
	_ = os.Setenv("TEST_DURATION_1", "1m")
	expected := 1 * time.Minute
	value := 2 * time.Minute
	assert.Equal(t, expected, Duration("TEST_DURATION_1"))
	assert.NotEqual(t, expected, Get("TEST_DURATION_2").Duration(value))

	_ = os.Setenv("TEST_DURATION_1", "zero")
	assert.Equal(t, value, Get("TEST_DURATION_1").Duration(value))
	assert.NotEqual(t, expected, Get("TEST_DURATION_1").Duration(value))
}

func TestGetHostname(t *testing.T) {
	assert.NotEmpty(t, GetHostname(""))
}
