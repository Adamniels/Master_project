package middleware

import "net/http"


// Detta kan man ockåså göra med middleware. Detta godkänner godkända adresser
func WithCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lägg till CORS-headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Om det är en preflight (OPTIONS), svara direkt
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Gå vidare till riktiga handlern
		// TODO: hur funkar detta?
		next.ServeHTTP(w, r)
	}
}