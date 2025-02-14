// CREDITS:
// go-httplogger project by Gleicon Moraes https://github.com/gleicon/go-httplogger/tree/master
// more on log formats: https://www.sumologic.com/blog/apache-access-log/

package telemetry

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type stResponseWriter struct {
	http.ResponseWriter
	HTTPStatus   int
	ResponseSize int
}

func (w *stResponseWriter) WriteHeader(status int) {
	w.HTTPStatus = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *stResponseWriter) Flush() {
	z := w.ResponseWriter
	if f, ok := z.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *stResponseWriter) CloseNotify() <-chan bool {
	z := w.ResponseWriter
	return z.(http.CloseNotifier).CloseNotify()
}

func (w *stResponseWriter) Write(b []byte) (int, error) {
	if w.HTTPStatus == 0 {
		w.HTTPStatus = 200
	}
	w.ResponseSize = len(b)
	return w.ResponseWriter.Write(b)
}

func WriteAccessLog(apacherow string) {

	logPath := os.Getenv("LOGPATH")

	logf, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to create access log file: %s", err)
		fmt.Println(logPath)
		return
	}

	_, err = fmt.Fprintln(logf, apacherow)

	if err != nil {
		log.Fatalf("Failed to write access log file. %s", err)
		return
	}

	err = logf.Close()

	if err != nil {
		log.Fatalf("Failed to close access log file. %s", err)
		return
	}

}

func HTTPLogger(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		interceptWriter := stResponseWriter{w, 0, 0}

		handler.ServeHTTP(&interceptWriter, r)

		// write access log row
		// Apache Combined Log Format
		// TODO Could also support the shorter format

		var s strings.Builder
		s.WriteString("HTTP - ")
		s.WriteString(r.RemoteAddr)
		s.WriteString(" - - ")
		s.WriteString(t.Format("02/Jan/2006:15:04:05 -0700"))
		s.WriteString(" \"")
		s.WriteString(r.Method)
		s.WriteString(" ")
		s.WriteString(r.URL.Path)
		s.WriteString(" ")
		s.WriteString(r.Proto)
		s.WriteString("\" ")
		s.WriteString(strconv.Itoa(interceptWriter.HTTPStatus))
		s.WriteString(" ")
		s.WriteString(strconv.Itoa(interceptWriter.ResponseSize))
		s.WriteString(" ")
		s.WriteString(r.UserAgent())
		s.WriteString(" ")
		s.WriteString(strconv.FormatInt(time.Since(t).Microseconds(), 10))
		// s.WriteString("us\n")

		fmt.Println(s.String())

		go WriteAccessLog(s.String())

	})
}
