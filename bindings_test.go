package envar

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type TestStruct struct {
	Name       string        `envar:" TEST_1 ,TEST_2; default=TEST"`
	Expiration time.Duration `envar:"EXPIRATION;default=10h"`
	Value      int64         `envar:"TEST_INT64;default=99"`
}

func Test_Bind(t *testing.T) {
	_ = os.Setenv("TEST_1", "testName")
	v := &TestStruct{}
	if err := Bind(v); err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(t, "testName", v.Name)
}
