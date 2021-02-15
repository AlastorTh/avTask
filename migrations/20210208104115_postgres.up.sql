CREATE TABLE ads (
    id bigserial not null primary key,
    name varchar not null,
    descript varchar not null,
    price float not null default 0.00,
    mainpic varchar not null, 
    otherpics varchar[] not null,
    created_at TIMESTAMP NOT NULL  default NOW()
);