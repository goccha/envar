package envar

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	_ = os.Setenv("TEST_BOOL_TRUE", "true")
	assert.Equal(t, true, Get("TEST_BOOL_TRUE").Bool(false))
	assert.NotEqual(t, true, Get("TEST_BOOL_FALSE").Bool(false))

	_ = os.Setenv("TEST_BOOL_TRUE", "1")
	assert.Equal(t, false, Get("TEST_BOOL_TRUE").Bool(false))
}

func TestInt(t *testing.T) {
	_ = os.Setenv("TEST_INT_1", "1")
	assert.Equal(t, 1, Get("TEST_INT_1").Int(2))
	assert.NotEqual(t, 1, Get("TEST_INT_2").Int(2))

	_ = os.Setenv("TEST_INT_1", "zero")
	assert.Equal(t, 2, Get("TEST_INT_1").Int(2))
}

func TestInt32(t *testing.T) {
	_ = os.Setenv("TEST_INT32_1", "1")
	var expected int32 = 1
	assert.Equal(t, expected, Get("TEST_INT32_1").Int32(2))
	assert.NotEqual(t, expected, Get("TEST_INT32_2").Int32(2))

	_ = os.Setenv("TEST_INT32_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_INT32_1").Int32(2))
}

func TestInt64(t *testing.T) {
	_ = os.Setenv("TEST_INT64_1", "1")
	var expected int64 = 1
	assert.Equal(t, expected, Get("TEST_INT64_1").Int64(2))
	assert.NotEqual(t, expected, Get("TEST_INT64_2").Int64(2))

	_ = os.Setenv("TEST_INT64_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_INT64_1").Int64(2))
}

func TestUint32(t *testing.T) {
	_ = os.Setenv("TEST_UINT32_1", "1")
	var expected uint32 = 1
	assert.Equal(t, expected, Get("TEST_UINT32_1").Uint32(2))
	assert.NotEqual(t, expected, Get("TEST_UINT32_2").Uint32(2))

	_ = os.Setenv("TEST_UINT32_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT32_1").Uint32(2))
}

func TestUint64(t *testing.T) {
	_ = os.Setenv("TEST_UINT64_1", "1")
	var expected uint64 = 1
	assert.Equal(t, expected, Get("TEST_UINT64_1").Uint64(2))
	assert.NotEqual(t, expected, Get("TEST_UINT64_2").Uint64(2))

	_ = os.Setenv("TEST_UINT64_1", "zero")
	expected = 2
	assert.Equal(t, expected, Get("TEST_UINT64_1").Uint64(2))
}

func TestFloat32(t *testing.T) {
	_ = os.Setenv("TEST_FLOAT32_1", "1.0")
	var expected float32 = 1.0
	assert.Equal(t, expected, Get("TEST_FLOAT32_1").Float32(2.0))
	assert.NotEqual(t, expected, Get("TEST_FLOAT32_2").Float32(2.0))

	_ = os.Setenv("TEST_FLOAT32_1", "zero")
	expected = 2.0
	v := Get("TEST_FLOAT32_1").Float32(2.0)
	assert.Equal(t, expected, v)
	assert.NotEqual(t, 2, v)
}

func TestFloat64(t *testing.T) {
	_ = os.Setenv("TEST_FLOAT64_1", "1.0")
	assert.Equal(t, 1.0, Get("TEST_FLOAT64_1").Float64(2.0))
	assert.NotEqual(t, 1.0, Get("TEST_FLOAT64_2").Float64(2.0))

	_ = os.Setenv("TEST_FLOAT64_1", "zero")
	v := Get("TEST_FLOAT64_1").Float64(2.0)
	assert.Equal(t, 2.0, v)
	assert.NotEqual(t, 2, v)
}

func TestString(t *testing.T) {
	_ = os.Setenv("TEST_STRING", "test")
	assert.Equal(t, "test", Get("TEST_STRING").String("test"))
	assert.NotEqual(t, "test", Get("TEST_STRING_PROD").String("prod"))
}

func TestDuration(t *testing.T) {
	_ = os.Setenv("TEST_DURATION_1", "1m")
	expected := 1 * time.Minute
	value := 2 * time.Minute
	assert.Equal(t, expected, Get("TEST_DURATION_1").Duration(value))
	assert.NotEqual(t, expected, Get("TEST_DURATION_2").Duration(value))

	_ = os.Setenv("TEST_DURATION_1", "zero")
	assert.Equal(t, value, Get("TEST_DURATION_1").Duration(value))
	assert.NotEqual(t, expected, Get("TEST_DURATION_1").Duration(value))
}

func TestGetHostname(t *testing.T) {
	assert.NotEmpty(t, GetHostname(""))
}
