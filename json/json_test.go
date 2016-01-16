package jsonp

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golangplus/testing/assert"
)

type T struct {
	S string
}

func TestMarshalToFile(t *testing.T) {
	fn := path.Join(os.TempDir(), "TestMarshalToFile.json")
	s := T{
		S: "abc",
	}
	assert.NoError(t, MarshalToFile(s, fn))
	bs, err := ioutil.ReadFile(fn)
	assert.NoError(t, err)
	assert.Equal(t, "bs", strings.TrimSpace(string(bs)), `{"S":"abc"}`)
}

func TestMarshalToFile_Failed(t *testing.T) {
	fn := path.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 16), "TestMarshalToFile.json")
	s := T{
		S: "abc",
	}
	assert.Error(t, MarshalToFile(s, fn))
}

func TestUnmarshalFile(t *testing.T) {
	fn := path.Join(os.TempDir(), "TestMarshalToFile.json")
	assert.NoError(t, ioutil.WriteFile(fn, []byte(`{"S":"abc"}`), 0644))
	var s T
	assert.NoError(t, UnmarshalFile(fn, &s))
	assert.Equal(t, "s.S", s.S, "abc")
}

func TestUnmarshalFile_Failed(t *testing.T) {
	fn := path.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 16), "TestMarshalToFile.json")
	var s T
	assert.Error(t, UnmarshalFile(fn, &s))
}

func TestMarshalIgnoreError(t *testing.T) {
	assert.Equal(t, "MarshalIgnoreError", string(MarshalIgnoreError(T{S: "abc"})), `{"S":"abc"}`)
}
