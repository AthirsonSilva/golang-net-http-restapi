package usecases

import "net/http"

// Logout is responsible for logging out the user and cleaning the session
func (r *Repository) Logout(res http.ResponseWriter, req *http.Request) {
	r.Config.Session.Destroy(req.Context())
	r.Config.Session.RenewToken(req.Context())
	r.Config.Session.Put(req.Context(), "flash", "User logged out successfully!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
