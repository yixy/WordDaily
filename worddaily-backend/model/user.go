package model

import (
	"database/sql"
	"encoding/base64"
	"fmt"
)

type User struct {
	Username string
	UserPwd  string
	Headshot string
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	query := `SELECT username, user_password, headshot FROM user_t WHERE username = ?`
	err := DB.QueryRow(query, username).Scan(&user.Username, &user.UserPwd, &user.Headshot)
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
	// 检查 headshotBase64 是否为有效的 base64 编码
	_, err := base64.StdEncoding.DecodeString(headshotBase64)
	if err != nil {
		return fmt.Errorf("invalid base64 encoding in headshot data")
	}

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
