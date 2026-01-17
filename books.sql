CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    co_authors VARCHAR(255)[] NOT NULL DEFAULT '{}',
    CONSTRAINT unique_book UNIQUE (title, author, co_authors)
);

CREATE INDEX IF NOT EXISTS idx_books_title_lower ON books (LOWER(title));
CREATE INDEX IF NOT EXISTS idx_books_author_lower ON books (LOWER(author));
