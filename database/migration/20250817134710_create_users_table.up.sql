CREATE TABLE users(
                      id varchar(64) NOT NULL,
                      first_name varchar(255) NOT NULL,
                      last_name varchar(255) NULL,
                      email varchar(255) UNIQUE,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY (id)
)ENGINE = INNODB