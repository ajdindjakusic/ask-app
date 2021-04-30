CREATE TABLE IF NOT EXISTS user (
          user_id SMALLINT NOT NULL AUTO_INCREMENT,          
          username varchar(30) NULL,          
          email varchar(40) NOT NULL,          
          password varchar(30) NOT NULL,          
          created timestamp not null default current_timestamp,          
          PRIMARY KEY (user_id)        
        );

CREATE TABLE IF NOT EXISTS question (
          question_id SMALLINT NOT NULL AUTO_INCREMENT,
          text varchar(1000) NOT NULL,
          user_id smallint NOT NULL,
          created timestamp not null default current_timestamp,
          PRIMARY KEY (question_id)
        );

CREATE TABLE IF NOT EXISTS answer (
          answer_id SMALLINT NOT NULL AUTO_INCREMENT,
          text varchar(1000) NOT NULL,
          user_id smallint NOT NULL,
          question_id smallint NOT NULL,
          created timestamp not null default current_timestamp,
          modified timestamp not null default current_timestamp,
          visible boolean default true,
          PRIMARY KEY (answer_id)
        );

CREATE TABLE IF NOT EXISTS likes (
          entity_id SMALLINT NOT NULL AUTO_INCREMENT,
          user_id smallint not NULL,
          question_id smallint default NULL,
          answer_id smallint default NULL,
          created timestamp not null default current_timestamp,
          PRIMARY KEY (entity_id)
        );

CREATE TABLE IF NOT EXISTS dislikes (
            entity_id SMALLINT NOT NULL AUTO_INCREMENT,
            user_id smallint not NULL,
            question_id smallint default NULL,
            answer_id smallint default NULL,
            created timestamp not null default current_timestamp,
            PRIMARY KEY (entity_id)
        );
