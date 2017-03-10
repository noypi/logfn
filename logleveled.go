package logfn

import (
	"strings"
)

type LogLevel int

func (this *LogLevel) WrapFunc(level int, fn LogFunc, prefix string) LogFunc {
	prefix = strings.Replace(prefix, "%", "", -1)
	return func(s string, as ...interface{}) {
		if level <= int(*this) {
			fn(prefix+s, as...)
		}
	}

}

func (this *LogLevel) SetLevel(n int) {
	*this = (LogLevel)(n)
}
