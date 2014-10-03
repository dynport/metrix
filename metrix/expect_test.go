package metrix

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func New(t *testing.T) func(target interface{}) *expectation {
	return (&expector{T: t}).expect
}

type expector struct {
	T *testing.T
}

func (e *expector) expect(target interface{}) *expectation {
	return &expectation{T: e.T, target: target}
}

type expectation struct {
	T      *testing.T
	target interface{}
}

func (e *expectation) ToHaveLength(l int) {
	realLen, ok := haveLength(e.target, l)
	e.expect(ok, "expected %#v to have length %d but was %d", e.target, l, realLen)
}

func (e *expectation) ToContain(element interface{}) {
	e.expect(contains(e.target, element), "expected %#v to contain %#v", e.target, element)
}
func (e *expectation) ToHavePrefix(prefix string) {
	s := e.expectString()
	e.expect(strings.HasPrefix(s, prefix), "expected %q to have prefix %q", s, prefix)
}

func (e *expectation) ToHaveSuffix(prefix string) {
	s := e.expectString()
	e.expect(strings.HasSuffix(s, prefix), "expected %q to have suffix %q", s, prefix)
}

func (e *expectation) ToNotEqual(i interface{}) {
	e.expect(e.target != i, "expected %#v to equal %#v", e.target, i)
}

func (e *expectation) ToEqual(i interface{}) {
	e.expect(equal(e.target, i), "expected %#v to equal %#v", e.target, i)
}

func (e *expectation) ToNotBeNil() {
	e.expect(e.target != nil, "expected %q to not be nil", e.target)
}

func (e *expectation) ToBeNil() {
	e.expect(e.target == nil, "expected %q to be nil", e.target)
}

// helpers

func equal(target interface{}, other interface{}) bool {
	return fmt.Sprint(target) == fmt.Sprint(other)
}

func contains(in interface{}, target interface{}) bool {
	switch v := valueOf(in).(type) {
	case string:
		if s, ok := target.(string); ok {
			return strings.Contains(v, s)
		}
		return false
	case []interface{}:
		for _, el := range v {
			if el == target {
				return true
			}
		}
	}
	return false
}

func haveLength(in interface{}, l int) (int, bool) {
	switch v := valueOf(in).(type) {
	case string:
		return len(v), len(v) == l
	case []interface{}:
		return len(v), len(v) == l
	}
	return 0, false
}

func valueOf(list interface{}) interface{} {
	out := []interface{}{}
	v := reflect.ValueOf(list)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			out = append(out, v.Index(i).Interface())
		}
		return out
	}
	return nil
}

func (e *expectation) expect(check bool, format string, args ...interface{}) {
	if check {
		e.pass()
	} else {
		e.fail(format, args...)
	}
}

func (e *expectation) pass() {
	fmt.Print(".")
}

func (e *expectation) fail(format string, i ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	fmt.Printf("\n\t%s:%d: %s\n", path.Base(file), line, fmt.Sprintf(format, i...))
	e.T.FailNow()
}

func (e *expectation) expectString() string {
	s, ok := e.target.(string)
	if !ok {
		e.fail("expected target to be string but was %T", e.target)
	}
	return s
}
