package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "time"
	u "github.com/ijasmoopan/usermanagement-api/api/users"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/gorilla/mux"
	// h "userregistration/common/helpers"
)

type Admin struct {
	Id        int `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"index"`
	Password  string `json:"password"`
}

func (h *BaseHandler) UserList(w http.ResponseWriter, r *http.Request){
	// Connecting to database
	var user []u.User
	if result := h.db.Find(&user); result.Error != nil {
		panic(result.Error)
	}

	// Response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}



func (h *BaseHandler) LogIn(w http.ResponseWriter, r *http.Request){

	var adminForm u.FormData
	if err := json.NewDecoder(r.Body).Decode(&adminForm); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Println("Admin: ", adminForm)
	var admin Admin

	if result := h.db.First(&admin, "username = ?", adminForm.Username); result.Error != nil {
		panic(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&admin)
}

func (h *BaseHandler) CreateUser (w http.ResponseWriter, r *http.Request){
	var user u.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	// ------check already exist or not--------------
	var userCheck u.User
	h.db.First(&userCheck, "username = ?", user.Username)
	
	if userCheck.Username == user.Username {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Username is already exist")
	} else {
		// Creating new user
		if result := h.db.Create(&user); result.Error != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("New user added")
	}	
}

func (h *BaseHandler) UpdateUser (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user u.User
	if result := h.db.First(&user, "username = ?", username); result.Error != nil {
		panic(result.Error)
	}
	userBytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	
	patchedJSON, err := jsonpatch.MergePatch(userBytes, body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(patchedJSON, &user)
	if err != nil {
		panic(err)
	}
	h.db.Save(&user)
	// patch, err := jsonpatch.DecodePatch(patchedJSON)
	// modified, err := patch.Apply(userBytes)
	// w.Write(modified)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
	json.NewEncoder(w).Encode(&user)
}

func (h *BaseHandler) BlockUser (w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	username := vars["username"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user u.User
	if result := h.db.First(&user, "username = ?", username); result.Error != nil {
		panic(result.Error)
	}
	userBytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	// var status u.User
	// if user.Status {
	// 	status.Status = false
	// } else {
	// 	status.Status = true
	// }
	// result, err := json.Marshal(&status)
	patchedJSON, _ := jsonpatch.MergePatch(userBytes, body)
	err = json.Unmarshal(patchedJSON, &user)
	h.db.Save(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
	json.NewEncoder(w).Encode(&user)
}

func (h *BaseHandler) DeleteUser (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]

	var user u.User
	if result := h.db.First(&user, "username = ?", username); result.Error != nil {
		panic(result.Error)
	}
	h.db.Delete(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}