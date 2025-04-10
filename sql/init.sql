-- sqlite3 ecards.db < init.sql

-- Create the word table
CREATE TABLE IF NOT EXISTS user_words_t (
    user_name TEXT,
    word TEXT,
    word_tag TEXT,
    meaning TEXT NOT NULL,
    example_sentence TEXT,
    word_status TEXT CHECK(word_status IN ('0','1','2')) NOT NULL,--'未学习', '保留', '击破'
    last_studied_date TEXT, --YYYY-MM-DD
    last_studied_time TIME,
    PRIMARY KEY (user_name, word, word_tag)
);

CREATE INDEX IF NOT EXISTS index_user_word_idx ON user_words_t (user_name, word_status, last_studied_date);

-- Create the user table
CREATE TABLE IF NOT EXISTS user_t (
    username TEXT PRIMARY KEY,
    user_password TEXT NOT NULL,
    headshot BLOB
);

INSERT INTO user_t (username, user_password, headshot)
VALUES ('seven', 'password', '');
commit;