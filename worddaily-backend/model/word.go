package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

// FetchWordMeaningAndExample fetches the meaning and example of a word from an external API.
func FetchWordMeaningAndExample(word string) (string, string, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}

	if len(data) == 0 {
		return "", "", fmt.Errorf("word not found")
	}

	meaning := ""
	example := ""

	if meanings, ok := data[0]["meanings"].([]interface{}); ok {
		for _, m := range meanings {
			if partOfSpeech, ok := m.(map[string]interface{})["partOfSpeech"].(string); ok {
				meaning += partOfSpeech + ": "
			}
			if definitions, ok := m.(map[string]interface{})["definitions"].([]interface{}); ok {
				for _, d := range definitions {
					if def, ok := d.(map[string]interface{})["definition"].(string); ok {
						meaning += def + "; "
					}
					if examples, ok := d.(map[string]interface{})["examples"].([]interface{}); ok {
						for _, ex := range examples {
							if exText, ok := ex.(map[string]interface{})["text"].(string); ok {
								example += exText + "; "
							}
						}
					}
				}
			}
		}
	}

	return strings.TrimSpace(meaning), strings.TrimSpace(example), nil
}
