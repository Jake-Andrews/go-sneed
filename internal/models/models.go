package models

type FormErrors struct {
    Username []string
    Email    []string
    Password []string
}

type FormData struct {
    Username string
    Email    string
    Password string
}

