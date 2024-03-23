CREATE TABLE IF NOT EXISTS users (
    id          INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username    CHAR(64) NOT NULL,
    password    CHAR(255) NOT NULL,
    email       CHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS memories (
    id          INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title       CHAR(64) NOT NULL,
    summary     CHAR(255) NOT NULL,
    time_limit  DATETIME NOT NULL,
    status      enum('Planning', 'Running', 'Completion') NOT NULL,
    description CHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS pictures (
    id          INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    memory_id   INTEGER NOT NULL,
    picture     MEDIUMBLOB NOT NULL,
    FOREIGN KEY (memory_id) REFERENCES memories(id)
);

CREATE TABLE IF NOT EXISTS tags (
    id          INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    memory_id   INTEGER NOT NULL,
    label       MEDIUMBLOB NOT NULL,
    FOREIGN KEY (memory_id) REFERENCES memories(id)
);
