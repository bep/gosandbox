package log4me

import (
	"fmt"
	"io/ioutil"
)

type LoggerI struct {
	Enabled bool
}

func (l LoggerI) Info(v ...interface{}) {
	if !l.Enabled {
		return
	}

	fmt.Fprintln(ioutil.Discard, v...)

}

func (l LoggerI) String() string {
	if l.Enabled {
		return "LoggerI: Enabled"
	}
	return "LoggerI: Disabled"
}

type LoggerF struct {
	Enabled bool
}

type F func()

func (l LoggerF) Info(f F) {
	if l.Enabled {
		f()
	}
}

func (l LoggerF) String() string {
	if l.Enabled {
		return "LoggerF: Enabled"
	}
	return "LoggerF: Disabled"
}
