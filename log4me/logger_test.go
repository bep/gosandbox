package log4me

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	for _, enabled := range []bool{false, true} {
		li := LoggerI{Enabled: enabled}
		lf := LoggerF{Enabled: enabled}

		b.Run(li.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				li.Info(strings.Repeat("Log log ", 20), strings.Repeat("Log log ", 20), strings.Repeat("Log log ", 20))
			}
		})

		b.Run(lf.String(), func(b *testing.B) {
			lf.Info(
				func() {
					for i := 0; i < b.N; i++ {
						fmt.Fprintln(ioutil.Discard, strings.Repeat("Log log ", 20), strings.Repeat("Log log ", 20), strings.Repeat("Log log ", 20))
					}
				},
			)
		})
	}
}
