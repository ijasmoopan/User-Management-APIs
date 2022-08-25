package users

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *BaseHandler) LogIn(w http.ResponseWriter, r *http.Request) {

	// Reading body
	var userForm FormData

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// json.Unmarshal(body, &userForm)

	// Instead of above code we can use Decode()....
	
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println("user form:", userForm)

	var user User

	
	// Find soft deleted or not
	if result := h.db.Unscoped().Find(&user, "username = ?", userForm.Username); result.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Deleted account")
	}

	if result := h.db.First(&user, "username = ?", userForm.Username); result.Error != nil {
		panic(result.Error)
	}
	// -----------check password---------
	if user.Password != userForm.Password {  // Hashing is remaining...........
		fmt.Println("Incorrect Password")
		panic(err)
	}
	if !user.Status {
		fmt.Println("User is blocked")
		panic(err)
	}
	// --------------generate token----------------

	
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func (h *BaseHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	defer r.Body.Close()

	fmt.Println("Registration: ", newUser)

	// ------------Check already exist or not----------------
	// var user User
	// if result := h.db.First(&user, "username = ?", newUser.Username); result.QueryFields || result.Error != nil {
	// 	fmt.Println("Error while checking already existing or not")
	// 	panic(result.Error)
	// }

	if result := h.db.FirstOrCreate(&newUser, User{Username: newUser.Username}); result.Error != nil {
		fmt.Println("FirstOrCreate")
		panic(result.Error)
	}


	// if result := h.db.Create(&newUser); result.Error != nil {
	// 	panic(result.Error)
	// }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Registered a new user")
}

func (h *BaseHandler) HomePage(w http.ResponseWriter, r *http.Request) {

	var user User
	vars := mux.Vars(r)
	username := vars["username"]

	if result := h.db.First(&user, "username = ?", username); result.Error != nil {
		panic(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

	// userInput := fmt.Sprintf(username, "")
	// h.db.Where("username = ?", userInput).First(&user)

}

func (h *BaseHandler) LogOut(w http.ResponseWriter, r *http.Request) {

	// ---------delete jwt-----------------
}