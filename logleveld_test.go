package logfn_test

import (
	"fmt"
	"testing"

	. "github.com/noypi/logfn"

	assertpkg "github.com/stretchr/testify/assert"
)

func TestLogLeveled(t *testing.T) {
	assert := assertpkg.New(t)

	var result string
	fn := func(s string, as ...interface{}) {
		result = fmt.Sprintf(s, as...)
	}

	var level int
	lvl := LogLevel(level)

	critical := lvl.WrapFunc(0, fn, "[critical]: ")
	// tst reassign
	lvl = LogLevel(level)
	err := lvl.WrapFunc(3, fn, "[error]: ")
	// test reassign
	lvl = LogLevel(level)
	warn := lvl.WrapFunc(1, fn, "[warn]: ")
	// test reassign
	lvl = LogLevel(level)

	type testarr struct {
		fn       LogFunc
		level    int
		param    int
		expected string
	}

	arr := []testarr{
		testarr{critical, 0, 0, "[critical]: 0"},
		testarr{err, 3, 3, "[error]: 3"},
		testarr{warn, 3, 1, "[warn]: 1"},
		testarr{critical, 3, 3, "[critical]: 3"},
		testarr{err, 4, 3, "[error]: 3"},
		testarr{err, 2, 3, ""},
		testarr{warn, 2, 1, "[warn]: 1"},
		testarr{warn, 0, 3, ""},
	}

	testfn := func(v testarr) {
		fmt.Println("v=", v)
		result = ""
		lvl.SetLevel(v.level)
		assert.Equal(v.level, int(lvl))
		v.fn("%d", v.param)
		assert.Equal(v.expected, result)
	}

	for _, v := range arr {
		testfn(v)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		testfn(arr[i])
	}

}
