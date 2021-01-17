package envar

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestStruct struct {
	Name string `envar:" TEST_1 ,TEST_2; default=TEST"`
}

func Test_Bind(t *testing.T) {
	_ = os.Setenv("TEST_1", "testName")
	v := &TestStruct{}
	if err := Bind(v); err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(t, "testName", v.Name)
}
