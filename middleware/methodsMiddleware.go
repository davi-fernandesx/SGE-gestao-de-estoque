package middleware

import "slices"

import "net/http"


// Como usar:
// http.Handle("/minha-rota", MethodMiddleware([]string{http.MethodGet}, http.HandlerFunc(meuHandler)))
func MethodMiddleware(metados []string, next http.HandlerFunc) http.HandlerFunc {

	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		permitido:= slices.Contains(metados, r.Method)

		if !permitido {

			http.Error(w, "request não aceito", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w,r)
	})

	
}
