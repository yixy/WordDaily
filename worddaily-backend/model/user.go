package model

import (
    "database/sql"
)

type User struct {
    Username  string
    PublicKey string
    Headshot  []byte
}

func GetUserByUsername(username string) (*User, error) {
    var user User
    query := `SELECT username, public_key, headshot FROM user_t WHERE username = ?`
    err := DB.QueryRow(query, username).Scan(&user.Username, &user.PublicKey, &user.Headshot)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}