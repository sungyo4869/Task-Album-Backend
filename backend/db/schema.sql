CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS memories (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(64) NOT NULL,
    summary VARCHAR(255) NOT NULL,
    time_limit DATETIME NOT NULL,
    status ENUM('planning', 'running', 'completion') NOT NULL,
    description VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS pictures (
    id INT NOT NULL AUTO_INCREMENT,
    memory_id INT NOT NULL,
    picture MEDIUMBLOB NOT NULL,
    FOREIGN KEY (memory_id) REFERENCES memories(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tags (
    id INT NOT NULL AUTO_INCREMENT,
    memory_id INT NOT NULL,
    label MEDIUMBLOB NOT NULL,
    FOREIGN KEY (memory_id) REFERENCES memories(id),
    PRIMARY KEY (id)
);
