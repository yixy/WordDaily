package model

import (
	"database/sql"
	"fmt"
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

const updateHeadshotQuery = `UPDATE user_t SET headshot = ? WHERE username = ?`

func UpdateUserHeadshot(headshotBase64, username string) error {
	// 执行更新操作
	result, err := DB.Exec(updateHeadshotQuery, headshotBase64, username)
	if err != nil {
		return err
	}

	// 验证是否有行受到影响
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// 如果没有行受到影响，可能是用户名不存在
		return fmt.Errorf("user %s not found", username)
	}

	return nil
}
