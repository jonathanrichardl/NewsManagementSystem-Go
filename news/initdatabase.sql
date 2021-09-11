USE testers;
drop table tags;
drop table news;

drop table deleted;
CREATE TABLE news(
id integer PRIMARY KEY AUTO_INCREMENT,
title varchar(45),
topic varchar(25),
status varchar(10)
);

CREATE TABLE tags(
tags varchar(20),
news_id integer,
FOREIGN KEY(news_id) REFERENCES news(id)
);

CREATE TABLE deleted(
    id integer,
    title varchar(45)
);

