package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An User"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}

func GetUserss(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}
func CreateUsesr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}
func GetUsersss(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An User"))
}
func UpdateUsersss(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
func DeleteUsersss(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
