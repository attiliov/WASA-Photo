CREATE TABLE IF NOT EXISTS User (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    signup_date DATETIME NOT NULL,
    last_seen DATETIME NOT NULL,
    email VARCHAR(255) DEFAULT NULL,
    bio TEXT DEFAULT "",
    prifile_image_id VARCHAR(36) DEFAULT NULL,
    followers_count INT NOT NULL DEFAULT 0,
    following_count INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Post (
    id VARCHAR(36) PRIMARY KEY,
    author_id VARCHAR(36) NOT NULL,
    author_username VARCHAR(255) NOT NULL,
    creation_date DATETIME NOT NULL,
    caption VARCHAR(5000) NOT NULL,
    image_id VARCHAR(36) NOT NULL DEFAULT "",
    like_count INT NOT NULL DEFAULT 0,
    comment_count INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Comment (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    username VARCHAR(255) NOT NULL,
    post_id VARCHAR(36) NOT NULL,
    caption VARCHAR(1000) NOT NULL,
    creation_date DATETIME NOT NULL,
    like_count INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS PostLike (
    user_id VARCHAR(36),
    post_id VARCHAR(36),
    username VARCHAR(255),
    PRIMARY KEY(user_id, post_id)
);

CREATE TABLE IF NOT EXISTS CommentLike (
    user_id VARCHAR(36) NOT NULL,
    comment_id VARCHAR(36) NOT NULL,
    username VARCHAR(255) NOT NULL,
    PRIMARY KEY(user_id, comment_id)
);

CREATE TABLE IF NOT EXISTS Follow (
    follower VARCHAR(36) NOT NULL,
    following VARCHAR(36) NOT NULL,
    PRIMARY KEY(follower, following)
);

CREATE TABLE IF NOT EXISTS Ban (
    user_id VARCHAR(36) NOT NULL,
    banned_user_id VARCHAR(36) NOT NULL,
    PRIMARY KEY(user_id, banned_user_id)
);

CREATE TABLE IF NOT EXISTS Photo (
    id VARCHAR(36) PRIMARY KEY,
    owner_id VARCHAR(36) NOT NULL
);

-- Add foreign key constraints
ALTER TABLE User ADD FOREIGN KEY(prifile_image_id) REFERENCES Photo(id);
ALTER TABLE Post ADD FOREIGN KEY(author_id) REFERENCES User(id);
ALTER TABLE Post ADD FOREIGN KEY(author_username) REFERENCES User(username);
ALTER TABLE Post ADD FOREIGN KEY(image_id) REFERENCES Photo(id);
ALTER TABLE Post ADD FOREIGN KEY(author_id) REFERENCES User(id);
ALTER TABLE Comment ADD FOREIGN KEY(user_id) REFERENCES User(id);
ALTER TABLE Comment ADD FOREIGN KEY(post_id) REFERENCES Post(id);
ALTER TABLE PostLike ADD FOREIGN KEY(user_id) REFERENCES User(id);
ALTER TABLE PostLike ADD FOREIGN KEY(post_id) REFERENCES Post(id);
ALTER TABLE CommentLike ADD FOREIGN KEY(user_id) REFERENCES User(id);
ALTER TABLE CommentLike ADD FOREIGN KEY(comment_id) REFERENCES Comment(id);
ALTER TABLE Ban ADD FOREIGN KEY(user_id) REFERENCES User(id);
ALTER TABLE Ban ADD FOREIGN KEY(banned_user_id) REFERENCES User(id);
ALTER TABLE Photo ADD FOREIGN KEY(owner_id) REFERENCES User(id);
ALTER TABLE Follow ADD FOREIGN KEY(follower) REFERENCES User(id);
ALTER TABLE Follow ADD FOREIGN KEY(following) REFERENCES User(id);