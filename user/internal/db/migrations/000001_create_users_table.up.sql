CREATE TABLE IF NOT EXISTS users
(
    uuid       varchar(50) UNIQUE,
    username text NOT NULL,
    email    text UNIQUE,
    password text NOT NULL,
    refresh text NOT NULL,
    isverified bool NOT NULL,
    joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
CREATE TABLE IF NOT EXISTS user_profiles
(
    uuid varchar(50) UNIQUE,
    avatar text NOT NULL,
    username text NOT NULL,
    email text UNIQUE,
    birthday date NOT NULL,
    bio text NOT NULL,
    lastonline TIMESTAMP
);
CREATE TABLE IF NOT EXISTS codes
(
    code text,
    email text UNIQUE,
    isverified bool
);
CREATE TABLE IF NOT EXISTS friend_reqs
(
    receiverid text NOT NULL,
    requestid text UNIQUE,
    avatar text NOT NULL,
    username text NOT NULL
);
CREATE TABLE IF NOT EXISTS friends
(
  uuid text NOT NULL,
  friendid text NOT NULL,
  avatar text NOT NULL,
  username text NOT NULL
);
