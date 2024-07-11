package models

import (
    "time"
)

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

type VideoData struct {
    UserID         string
    Title          string
    Description    string
    Duration       time.Duration
    FilePath       string
    ThumbnailPath  string
    Quality        map[string]interface{}
}
