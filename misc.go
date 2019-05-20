package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// NginxPiped : '$remote_addr|$host|[$time_local]|$request_time|$upstream_response_time|"$request"|$status'|$body_bytes_sent|"$http_referer"|$upstream_addr'
	NginxPiped   = "%s|%s|[%s]|%f|%f|\"%s %s HTTP/1.1\"|%d|%d|\"%s\"|-\n"
	NginxTime    = "02/Jan/2006:15:04:05 +0000"
	NginxErrTime = "02/Jan/2006 15:04:05"
	NginxError   = "%s [%s] 1001#1001: *2222 %s, client: %s, server %s, request: \"%s %s HTTP/1.1\", upstream: \"uwsgi://127.0.0.1:8080\", host: \"%s\"\n"
)

var StatusCodes = []int{200, 301, 403, 400, 500, 503, 504}
var ErrorLevels = []string{"warn", "error", "crit", "alert", "emerg"}

func GetRespCode() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(StatusCodes)
	return StatusCodes[n]
}

func GetErr() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(ErrorLevels)
	return ErrorLevels[n]
}

// NewNginxPipedLog creates a nginx log string with pipe delimiter
func NewNginxPiped(resp int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(
		NginxPiped,
		gofakeit.IPv4Address(),
		gofakeit.DomainName(),
		GetTime("NginxTime"),
		rand.Float32(),
		rand.Float32(),
		gofakeit.HTTPMethod(),
		URI(),
		resp,
		gofakeit.Number(0, 20000),
		gofakeit.URL(),
	)
}

func NewNginxError(llevel string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(
		NginxError,
		GetTime("NginxErrTime"),
		llevel,
		gofakeit.HackerPhrase(),
		gofakeit.IPv4Address(),
		gofakeit.DomainName(),
		gofakeit.HTTPMethod(),
		URI(),
		gofakeit.DomainName(),
	)
}

func GetTime(format string) string {
	t := time.Now()
	var line string

	switch format {
	case "NginxTime":
		line = fmt.Sprintf(t.Format(NginxTime))
	case "NginxErrTime":
		line = fmt.Sprintf(t.Format(NginxErrTime))
	default:
		panic("PANIC! unknown time pattern: " + format)
	}
	return line
}

func URI() string {
	// Slugs
	num := Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = gofakeit.Word()
	}
	url := "/" + strings.ToLower(strings.Join(slug, "/"))

	return url
}

func Number(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}
