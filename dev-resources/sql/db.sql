/*Create database out of this file*/
-- CREATE DATABASE tech_review_test;
/*Needed to create UUIDs*/
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*Create article categories table*/
CREATE TABLE topics
(
    id   BIGSERIAL PRIMARY KEY NOT NULL,
    name TEXT
);

/*Create user table*/
CREATE TABLE users
(
    /* reconsider the id*/
    id         VARCHAR DEFAULT REPLACE(uuid_generate_v4() :: text, '-', '') PRIMARY KEY,
    username   VARCHAR(250) NOT NULL UNIQUE,
    first_name VARCHAR(200) NOT NULL,
    last_name  VARCHAR(200) NOT NULL,
    password   VARCHAR(250) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    role_id    VARCHAR      NOT NULL,
    interests  JSONB
);

/*Create Article table */
CREATE TABLE articles
(
    id                 VARCHAR      NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4() :: text, '-', ''),
    author_id          VARCHAR(250) NOT NULL,
    content            JSONB        NOT NULL,
    topics             JSONB        NOT NULL,
    average_rating     FLOAT,
    number_of_ratings  INTEGER,
    posted_at          TIMESTAMP,
    review             BOOL,
    number_of_comments INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id)
);

/*Create Comment table */
CREATE TABLE comments
(
    id         VARCHAR      NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4() :: text, '-', ''),
    writer     VARCHAR(250) NOT NULL,
    content    TEXT         NOT NULL,
    article_id VARCHAR      NOT NULL,
    posted_at    TIMESTAMP,
    likes      INTEGER,
    FOREIGN KEY (writer) REFERENCES users (id),
    FOREIGN KEY (article_id) REFERENCES articles (id)
);

/*Create Question table*/
CREATE TABLE questions
(
    id                VARCHAR      NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4() :: text, '-', ''),
    inquirer_id          VARCHAR(250) NOT NULL,
    inquiry           TEXT         NOT NULL,
    asked_at          TIMESTAMP,
    follows        INTEGER,
    number_of_answers INTEGER,
    topics            JSONB,
    FOREIGN KEY (inquirer_id) REFERENCES users (id)
);

-- check replier_name ?
/*Create Answer table*/
CREATE TABLE answers
(
    id      VARCHAR      NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4() :: text, '-', ''),
    question_id     VARCHAR      NOT NULL,
    replier_id VARCHAR(250) NOT NULL,
    answer  VARCHAR      NOT NULL,
    votes INTEGER,
    FOREIGN KEY (question_id) REFERENCES questions (id),
    FOREIGN KEY (replier_id) REFERENCES users (id)
);
-- does the name have to be username ?
CREATE TABLE article_ratings (
    article_id VARCHAR,
    user_id VARCHAR,
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (user_id)  REFERENCES users(id)
);

CREATE TABLE comment_likes(
    article_id VARCHAR,
    user_id VARCHAR,
    FOREIGN KEY (article_id) references articles(id),
    FOREIGN KEY (user_id) references users(id)
);

CREATE TABLE answer_upvotes(
    answer_id VARCHAR,
    user_id VARCHAR,
    FOREIGN KEY (answer_id) references answers(id),
    FOREIGN KEY (user_id) references users(id)
    );

CREATE TABLE question_follows
(
    question_id VARCHAR,
    user_id     VARCHAR,
    FOREIGN KEY (question_id) references questions (id),
    FOREIGN KEY (user_id) references users (id)

);
