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
	head := "Traceback (most recent call last)"
	src := fmt.Sprintf("\tFile \"%s.py\", line %d, in _execute_context",
		URI(),
		gofakeit.Number(0, 1000),
	)
	bla := fmt.Sprintf("\t\t%s", gofakeit.HackerPhrase())
	tail := "python.some.exception: m out of mind"

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
			res["traceback"] += 1
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
