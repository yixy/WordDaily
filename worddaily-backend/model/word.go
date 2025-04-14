package model

import (
	"time"
)

type Word struct {
	UserName        string    `json:"user_name"`
	Word            string    `json:"word"`
	WordTag         string    `json:"word_tag"`
	Example         string    `json:"example_sentence"`
	Meaning         string    `json:"meaning"`
	WordStatus      int       `json:"word_status"`
	LastStudiedDate string    `json:"last_studied_date"`
	LastStudiedTime time.Time `json:"last_studied_time"`
}

// InsertWord inserts a new word into the user_words_t table.
func InsertWord(word Word) (int64, error) {
	result, err := DB.Exec("INSERT INTO user_words_t (user_name, word, word_tag, example_sentence, meaning, word_status,last_studied_date,last_studied_time ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		word.UserName, word.Word, word.WordTag, word.Example, word.Meaning, word.WordStatus, word.LastStudiedDate, word.LastStudiedTime)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func DeleteWord(word Word) error {
	_, err := DB.Exec("DELETE FROM user_words_t WHERE user_name = ? AND word = ? AND word_tag = ?", word.UserName, word.Word, word.WordTag)
	if err != nil {
		return err
	}

	return nil
}

// FetchWordTranslation fetches the translation of a word from the en_words table.
func FetchWordTranslation(word string) (string, error) {
	var translation string
	err := DB.QueryRow("SELECT translation FROM en_words WHERE word = ?", word).Scan(&translation)
	if err != nil {
		return "", err
	}
	return translation, nil
}

// FetchWordMeaningAndExample fetches the meaning and example of a word from an external API.
func FetchWordMeaningAndExample(word string) (string, string, error) {
	translation, err := FetchWordTranslation(word)
	if err != nil {
		return "", "", err
	}

	// 这里可以继续调用外部API获取例句
	example := "This is an example sentence." // 假设从API获取的例句

	return translation, example, nil
}