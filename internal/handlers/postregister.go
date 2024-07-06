package handlers

import (
	"context"
	store "go-sneed/internal/db"
	"go-sneed/internal/models"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"log"
	"net/http"
	"net/mail"
	"reflect"
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
	user := models.FormData{
		Email:    email,
		Password: password,
		Username: username,
	}

	formErrors := validateFormData(user)
	if !StructOfStrSlicesEmpty(formErrors) {
		user.Password = ""
        w.WriteHeader(http.StatusBadRequest)
		utils.RenderTemplWithLayout(templates.RegisterPage(formErrors, user), r.Context(), w)
		return
	}

	if err := h.userStore.CreateUser(context.Background(), &user); err != nil {
		log.Printf("Error creating user %v", err)
        w.WriteHeader(http.StatusInternalServerError)
		user.Password = ""
        serverErrors := models.FormErrors{Email: []string{"Error registering"}, Password: []string{"Error registering"}, Username: []string{"Error registering"},}
		utils.RenderTemplWithLayout(templates.RegisterPage(serverErrors, user), r.Context(), w)
		return
	}

	log.Printf("Success creating user %q", username)
    //utils.RenderTemplWithLayout(templates.Index(username), r.Context(), w)
	user.Password = ""
    w.Header().Set("HX-Push-Url", "/")
    if err := templates.Index(user.Username).Render(r.Context(), w); err != nil {
        http.Error(w, "Error rendering template!", http.StatusInternalServerError)
    }
}

// If an err is encountered, "Issue validating x" will be added onto FormErrors
// for said property
func validateFormData(formData models.FormData) models.FormErrors {
	formErrors := models.FormErrors{}
	formErrors.Username = validateUsername(formData.Username)
	formErrors.Password = validatePassword(formData.Password)
	formErrors.Email = validateEmail(formData.Email)

	return formErrors
}

func validateUsername(userName string) []string {
	uErrors := make([]string, 0, 1)
	if len(userName) > 50 {
		uErrors = append(uErrors, "Error: username must be <= 50 characters!")
	}
	return uErrors
}

func validateEmail(email string) []string {
	eErrors := make([]string, 0, 2)
	if len(email) > 50 {
		eErrors = append(eErrors, "Email must be <= 50 characters!")
	}

	m, err := mail.ParseAddress(email)
	if err != nil {
		eErrors = append(eErrors, "Error: validating email!")
		log.Printf("Error validating email %q\nerr: %v", email, err)
		return eErrors
	}
	if m.Address != email {
		eErrors = append(eErrors, "Error: Invalid Email format!")
	}
	return eErrors
}

func validatePassword(password string) []string {
	pErrors := make([]string, 0, 1)
	if len(password) < 8 {
		pErrors = append(pErrors, "Error: Password must be >= 8 characters!")
	} else if len(password) > 50 {
		pErrors = append(pErrors, "Error: Password must be <= 50 characters!")
	}
	return pErrors
}

func StructOfStrSlicesEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.String {
			if field.Len() != 0 {
				return false
			}
		} else { // not a str slice
			return false
		}
	}
	return true
}
