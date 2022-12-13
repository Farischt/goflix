package models

// Postgresql schema for movies table
const Schema = `
CREATE TABLE IF NOT EXISTS movies (
	id SERIAL PRIMARY KEY,
	title TEXT,
	description TEXT,
	release_date TEXT,
	duration INT,
	trailer_url TEXT
);
`
