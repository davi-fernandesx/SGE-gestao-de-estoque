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
/*
Como funciona:
Recebe uma lista de métodos permitidos (metados []string) e o próximo handler (next http.HandlerFunc).
Quando a rota é chamada, verifica se o método da requisição está na lista usando slices.Contains.
Se não estiver, retorna erro 405 (Method Not Allowed).
Se estiver, chama o próximo handler normalmente.
*/