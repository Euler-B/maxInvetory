create table PRODUCTS(
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_by int not null, 
    primary key (id),
    foreign key (created_by) references USERS(id)

);