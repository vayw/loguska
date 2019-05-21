package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		logf     = flag.String("log", "piped", "log format; piped, err, py")
		logname  = flag.String("output", "log.log", "output file")
		loglines = flag.Int("lines", 100, "number of lines to generate")
	)
	flag.Parse()

	f, err := os.OpenFile(*logname, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	var res string

	switch *logf {
	case "piped":
		res = nginxPiped(w, *loglines)
	case "err":
		res = nginxError(w, *loglines)
	case "py":
		res = pyLog(w, *loglines)
	}

	w.Flush()
	fmt.Println(res)
}

func nginxPiped(w *bufio.Writer, l int) string {
	var errors = 0
	for i := 0; i < l; i++ {
		rc := GetRespCode()
		if rc >= 500 {
			errors += 1
		}
		_, err := w.WriteString(NewNginxPiped(rc))
		if err != nil {
			fmt.Println("error on write")
		}
	}

	return fmt.Sprintf("%d", errors)
}

func nginxError(w *bufio.Writer, l int) string {
	var err, crit, emerg = 0, 0, 0
	for i := 0; i < l; i++ {
		elevel := GetErr()
		switch elevel {
		case "error":
			err += 1
		case "crit":
			crit += 1
		case "emerg":
			emerg += 1
		}
		_, err := w.WriteString(NewNginxError(elevel))
		checke(err)
	}
	return fmt.Sprintf("%d;%d;%d", err, crit, emerg)
}

func checke(e error) {
	if e != nil {
		fmt.Println("error on write", e)
	}
}
