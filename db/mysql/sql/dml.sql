INSERT INTO users(id,username) VALUES ("user1","yaga");

INSERT INTO user_auths(user_id,email,hash) VALUES ("user1","yaga@example.com","$2a$10$/z1SlkePlRKHfOOxb/w70.7B45svUsrqUq5kAFDUM/E4mjDMWvdwa");

INSERT INTO auth_tokens(id,user_id,token,expiry) VALUES ("auth_token1","user1","token","2020-05-02 21:00:00");

INSERT INTO documents(id,url,user_id) VALUES
("document1","https://go-tour-jp.appspot.com/welcome/1","user1"),
("document2","https://soudai.hatenablog.com/entry/2018/05/01/204442","user1"),
("document3","https://angular.jp/tutorial","user1");
