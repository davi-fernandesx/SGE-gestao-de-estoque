package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggerResponseWriter struct {

	http.ResponseWriter
	statusCode int
} 

func (logg *loggerResponseWriter) WriteHeader(code int){


	logg.statusCode = code
	logg.ResponseWriter.WriteHeader(code)
}


func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()

		logger:= &loggerResponseWriter{
			ResponseWriter: w,
			statusCode: http.StatusOK,
		}

		next.ServeHTTP(logger,r)


		duration:= time.Since(start)
		log.Printf("\n Método:[%s]\n URL:%s\n Endereço IP:%s\n Status:%d\n Duração:%vs \n", r.Method, r.URL.Path, r.RemoteAddr, logger.statusCode, duration)

	})

}
