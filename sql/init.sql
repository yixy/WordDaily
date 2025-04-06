-- sqlite3 ecards.db < init.sql

-- Create the word table
CREATE TABLE IF NOT EXISTS words_t (
    user_name TEXT,
    word TEXT,
    meaning TEXT NOT NULL,
    example_sentence TEXT,
    word_status TEXT CHECK(word_status IN ('未学习', '保留', '击破')) NOT NULL,
    last_studied_date DATE,
    last_studied_time TIME,
    PRIMARY KEY (user_name, word)
);

CREATE INDEX IF NOT EXISTS index_user_word_idx ON words_t (user_name, word_status, last_studied_date);

-- Create the word_tag table
CREATE TABLE IF NOT EXISTS user_words_tag_t (
    user_name TEXT,
    word TEXT,
    tag TEXT,
    PRIMARY KEY (user_name, word, tag)
);

-- Create the user table
CREATE TABLE IF NOT EXISTS user_t (
    username TEXT PRIMARY KEY,
    public_key TEXT NOT NULL,
    headshot BLOB
);