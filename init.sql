create table if not exists crud.todos (
	id serial primary key,
 	ts timestamp default current_timestamp,
	todo varchar(255) not null
);
