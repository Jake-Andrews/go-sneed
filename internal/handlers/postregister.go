package handlers

import (
	"context"
	store "go-sneed/internal/db"
	"log"
	"net/http"
)

type PostRegisterHandler struct {
    userStore store.UserStore
}

func NewPostRegisterHandler(UserStore store.UserStore) *PostRegisterHandler {
    return &PostRegisterHandler{userStore: UserStore}
}


func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // register logic
    if err := r.ParseForm(); err != nil {
        log.Fatal(err)
    }
    email := r.FormValue("email")
    password := r.FormValue("password")
    //log.Printf("User, email: %v", r.Form)
    log.Printf("email %s, password %s", email, password)
    log.Println("PostRegisterHandler calling db")

    if err := h.userStore.CreateUser(context.Background(), email, password); err != nil {
        log.Fatal(err)
    }
    log.Printf("Inserted user into db: %s", email)
}
