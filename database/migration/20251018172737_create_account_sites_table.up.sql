CREATE TABLE account_sites(
                              id varchar(64) NOT NULL,
                              name varchar(255) NOT NULL,
                              description text NULL,
                              url varchar(255) NULL,
                              icon varchar(255) NULL,
                              created_by_id VARCHAR(64) NULL,

                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (id),
                              CONSTRAINT fk_account_sites_created_by
                                  FOREIGN KEY (created_by_id)
                                      REFERENCES users(id)
                                      ON UPDATE RESTRICT
                                      ON DELETE RESTRICT
)ENGINE = INNODB


