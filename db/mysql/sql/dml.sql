INSERT INTO users(id,username) VALUES ("user1","yaga");

INSERT INTO user_auths(user_id,email,hash) VALUES ("user1","yaga@example.com","password");

INSERT INTO auth_tokens(id,user_id,token,expiry) VALUES ("auth_token1","user1","token","2020-05-02 21:00:00");
