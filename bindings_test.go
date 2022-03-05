package envar

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
)

type TestStruct struct {
	Name        string        `envar:" TEST_1 ,TEST_2; default=TEST"`
	Expiration  time.Duration `envar:"EXPIRATION;default=10h"`
	Value       int64         `envar:"TEST_INT64;default=99"`
	Values      []string      `envar:"TEST_SLICE;default=test1,test2"`
	Nums        []int         `envar:"TEST_NUMS;default=1,2"`
	Int8s       []int8        `envar:"TEST_INT8_ARRAY;default=1,2"`
	Int16s      []int16       `envar:"TEST_INT16_ARRAY;default=1,2"`
	Int32s      []int32       `envar:"TEST_INT32_ARRAY;default=1,2"`
	Int64s      []int64       `envar:"TEST_INT64_ARRAY;default=1,2"`
	Uints       []uint        `envar:"TEST_UINT_ARRAY;default=1,2"`
	Uint8s      []uint8       `envar:"TEST_UINT8_ARRAY;default=1,2"`
	Uint16s     []uint16      `envar:"TEST_UINT16_ARRAY;default=1,2"`
	Uint32s     []uint32      `envar:"TEST_UINT32_ARRAY;default=1,2"`
	Uint64s     []uint64      `envar:"TEST_UINT64_ARRAY;default=1,2"`
	Complex64s  []complex64   `envar:"TEST_COMPLEX64_ARRAY;default=1,2"`
	Complex128s []complex128  `envar:"TEST_COMPLEX128_ARRAY;default=1,2"`
	WriteFile   *os.File      `envar:"TEST_FILE;default=test.txt"`
	ReadFile    io.Reader     `envar:"TEST_READER;default=test.txt"`
}

func Test_Bind(t *testing.T) {
	_ = os.Setenv("TEST_1", "testName")
	v := &TestStruct{}
	if err := Bind(v); err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(t, "testName", v.Name)
	_ = v.WriteFile.Close()
	if closer, ok := v.ReadFile.(io.Closer); ok {
		_ = closer.Close()
	}
	err := os.Remove("test.txt")
	assert.NoError(t, err)
}
