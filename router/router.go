package router

import (
	"net/http"

	// "github.com/ijasmoopan/usermanagement-api/common/helpers"
	d "github.com/ijasmoopan/usermanagement-api/common/db"
	a "github.com/ijasmoopan/usermanagement-api/api/admin"
	u "github.com/ijasmoopan/usermanagement-api/api/users"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	DB := d.ConnectDB()

	ha := a.NewBaseHandler(DB)
	hu := u.NewBaseHandler(DB)

	r := mux.NewRouter()

	// Admin routes
	r.HandleFunc("/admin", ha.LogIn).Methods(http.MethodPost)

	r.HandleFunc("/admin/userlist", ha.UserList).Methods(http.MethodGet)

	r.HandleFunc("/admin/userlist/create", ha.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/admin/userlist/{username}/update", ha.UpdateUser).Methods(http.MethodPatch)
	r.HandleFunc("/admin/userlist/{username}/delete", ha.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/admin/userlist/{username}/block", ha.BlockUser).Methods(http.MethodPatch)

	// User routes
	r.HandleFunc("/login", hu.LogIn).Methods(http.MethodPost)

	r.HandleFunc("/signup", hu.SignUp).Methods(http.MethodPost)

	r.HandleFunc("/logout", hu.LogOut).Methods(http.MethodGet)

	r.HandleFunc("/home/{username}", hu.HomePage).Methods(http.MethodGet)

	return r
}