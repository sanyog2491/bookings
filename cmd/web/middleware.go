//writing the basic middleware

package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("hit the page")
// 		next.ServeHTTP(w, r)
// 	})
// }

//it adds CSRF protection to all POST requests

func Nosurf(next http.Handler) http.Handler {
	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfhandler
}

//it loads and saves the session data for the current request
func Sessionload(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
