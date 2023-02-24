create table restaurants
(
    id serial constraint restaurants_pk primary key,
    label varchar not null,
    category_type int not null constraint type references category
);
