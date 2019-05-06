package views

import "net/http"

// true  :  has logged in
// false :  not logged in
func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] == "true" {
		_, _ = w.Write([]byte("true"))
		return
	}
	_, _ = w.Write([]byte("false"))
	return
}
