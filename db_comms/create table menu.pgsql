create table menu
(
    id serial constraint menu_pk primary key,
    restaurant_id int constraint restaurant_id references restaurants,
    name varchar,
    photo uuid,
    price int
);
