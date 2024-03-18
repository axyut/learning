.schema
.tables

CREATE TABLE user(
    id INTEGER NOT NULL UNIQUE,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE post(
	id INTEGER NOT NULL UNIQUE,
	user_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT,
	FOREIGN KEY(user_id) REFERENCES user(id),
	PRIMARY KEY(id)
);

CREATE TABLE comment (
	id INTEGER NOT NULL UNIQUE,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	FOREIGN KEY(post_id) REFERENCES post(id),
	FOREIGN KEY(user_id) REFERENCES user(id),
	PRIMARY KEY (id)
);

CREATE TABLE following(
	user_id INTEGER NOT NULL,
	following_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES user(id),
	FOREIGN KEY(following_id) REFERENCES user(id)
);

CREATE TABLE follower (
	user_id INTEGER NOT NULL,
	follower_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES user(id),
	FOREIGN KEY(follower_id) REFERENCES user(id)
);
