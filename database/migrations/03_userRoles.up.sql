create table USERS_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);