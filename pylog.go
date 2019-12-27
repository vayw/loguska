package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	PyLine = "[%s] %s %d %s: %s\n"
)

var PyLogLevels = []string{"WARNING", "INFO", "ERROR", "DEBUG", "CRITICAL"}
var PyLogger = []string{"app.system.worker",
	"app.logic.incapsulator", "app.logic.duplicator"}

func makeTraceback() string {
	var tb string
	t := rand.Float32()
	switch {
	case t > 0.6:
		tb = "    "
	case 0.6 >= t && t > 0.3:
		tb = "\t"
	default:
		tb = ""
	}
	head := fmt.Sprintf("%sTraceback (most recent call last)", tb)
	src := fmt.Sprintf("%s\tFile \"%s.py\", line %d, in _execute_context",
		tb, URI(),
		gofakeit.Number(0, 1000),
	)
	bla := fmt.Sprintf("%s\t\t%s", tb, gofakeit.HackerPhrase())
	tail := fmt.Sprintf("%spython.some.exception: m out of mind", tb)

	result := fmt.Sprintf("%s\n%s\n%s\n%s\n", head, src, bla, tail)
	return result
}

func pyLog(w *bufio.Writer, l int) string {
	res := make(map[string]int)
	for i := 0; i < l; i++ {
		var line string
		rand.Seed(time.Now().UnixNano())
		n := rand.Int() % 100
		if n > 10 {
			loglevel := pickLine(PyLogLevels)
			logger := pickLine(PyLogger)
			line = fmt.Sprintf(PyLine,
				loglevel,
				GetTime("PyLogTime"),
				gofakeit.Number(0, 2000),
				logger,
				gofakeit.HackerPhrase(),
			)
			key := fmt.Sprintf("%s:%s", logger, loglevel)
			res[key] += 1
		} else {
			line = makeTraceback()
			res["Traceback:Traceback"] += 1
		}
		_, err := w.WriteString(line)
		if err != nil {
			fmt.Println("error on write")
		}
	}
	b := new(bytes.Buffer)
	for k, v := range res {
		fmt.Fprintf(b, "%s=%d\n", k, v)
	}
	return b.String()
}

func pickLine(ar []string) string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(ar)
	return ar[n]
}
