package models

type UserKey int
const (
    UserReqKey UserKey = iota
)

type AuthenticatedUser struct {
    Email    string
    LoggedIn bool
}
