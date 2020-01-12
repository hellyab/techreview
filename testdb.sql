CREATE EXTENSION
IF NOT EXISTS "uuid-ossp";

/*Create article categories table*/
CREATE TABLE article_categories
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(150)
);

/*Create user table*/
CREATE TABLE user
(
    /* reconsider the id*/
    id VARCHAR DEFAULT REPLACE(uuid_generate_v4()
    :: text, '-', ''),
    username    VARCHAR
    (250) NOT NULL PRIMARY KEY,
    first_name  VARCHAR
    (200) NOT NULL,
    middle_name VARCHAR
    (15),
    last_name   VARCHAR
    (200) NOT NULL,
    password    VARCHAR
    (250) NOT NULL,
    email       VARCHAR
    (150) NOT NULL,
    interests   JSON,
    privileged  BOOL         NOT NULL
);

    /*Create Article table */
    CREATE TABLE article
    (
        id VARCHAR NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()
        :: text, '-', ''),
    author            VARCHAR
        (250) NOT NULL,
    content           JSON         NOT NULL,
    topics            JSON         NOT NULL,
    average_rating    FLOAT,
    number_of_ratings INTEGER,
    posted_at         TIMESTAMP,
    FOREIGN KEY
        (author) REFERENCES user
        (username)
);

        /*Create Comment table */
        CREATE TABLE comment
        (
            id VARCHAR NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()
            :: text, '-', ''),
    writer     VARCHAR
            (250) NOT NULL,
    content    JSON         NOT NULL,
    article_id VARCHAR      NOT NULL,
    post_at    TIMESTAMP,
    FOREIGN KEY
            (writer) REFERENCES user
            (username),
    FOREIGN KEY
            (article_id) REFERENCES article
            (id)
);

            /*Create Question table*/
            CREATE TABLE question
            (
                id VARCHAR NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()
                :: text, '-', ''),
    inquirer VARCHAR
                (250) NOT NULL,
    inquiry  TEXT         NOT NULL,
    FOREIGN KEY
                (inquirer) REFERENCES user
                (username)
);

                /*Create Answer table*/
                CREATE TABLE answer
                (
                    id VARCHAR NOT NULL PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()
                    :: text, '-', ''),
    qid     VARCHAR      NOT NULL,
    replier VARCHAR
                    (250) NOT NULL,
    answer  VARCHAR      NOT NULL,
    FOREIGN KEY
                    (qid) REFERENCES question
                    (id),
    FOREIGN KEY
                    (replier) REFERENCES user
                    (username)
);