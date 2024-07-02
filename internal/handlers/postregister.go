package handlers

import (
	"context"
	store "go-sneed/internal/db"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
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
    // not needed for formvalue just form, remove after testing
    if err := r.ParseForm(); err != nil {
        log.Fatal(err)
    }
    log.Printf("User: %+v\n", r.Form)
    email := r.FormValue("email")
    password := r.FormValue("password")
    username := r.FormValue("username")
    //log.Printf("User, email: %v", r.Form)
    log.Printf("email %q, password %q, username %q", email, password, username)
    log.Println("PostRegisterHandler calling db")

    user := store.User{
        Email: email,
        Password: password,
        Username: username,
    }
    if err := h.userStore.CreateUser(context.Background(), &user); err != nil {
        log.Printf("Error creating user %v", err)
        utils.RenderTemplWithLayout(templates.RegisterError(), r.Context(), w)
        return
    }
    log.Printf("Success creating user %q", username)
    utils.RenderTemplWithLayout(templates.RegisterSuccess(username), r.Context(), w)
}

