CREATE TABLE Video_table (
	video_id int GENERATED ALWAYS AS IDENTITY,
	video_link varchar(50),
	code jsonb,
	paid bool,
	video_topic varchar(50),
	book_id int,
	tools_id int,
	date_created TIMESTAMP,
	is_active bool,
	PRIMARY KEY(video_id),
	  CONSTRAINT fk_book
      FOREIGN KEY(book_id) 
	  REFERENCES book(book_id),
	  CONSTRAINT fk_tools
      FOREIGN KEY(tools_id) 
	  REFERENCES tools(tools_id)
);
select * from Video_table;

insert into Video_table (video_link,paid,video_topic,book_id,tools_id,date_created,is_active) VALUES
('test video link 1', TRUE, 'test video topic 1', (select book_id from book where book_name = 'book_name 1'),(select tools_id from tools where tool_name = 'tool_name 1'), (select current_timestamp), true),
('test video link 2', TRUE, 'test video topic 2', (select book_id from book where book_name = 'book_name 2'),(select tools_id from tools where tool_name = 'tool_name 2'), (select current_timestamp), true),
('test video link 3', TRUE, 'test video topic 3', (select book_id from book where book_name = 'book_name 3'),(select tools_id from tools where tool_name = 'tool_name 3'), (select current_timestamp), true);


drop table Video_table;

create table tools(
	tools_id int GENERATED ALWAYS AS IDENTITY,
	tool_name varchar(50),
	tool_link varchar(50),
	date_created TIMESTAMP,
	is_active bool,
	PRIMARY KEY(tools_id)
);
select * from tools;

insert into tools (tool_name, tool_link,date_created,is_active) VALUES 
('tool_name 1', 'tool_link 1', (select current_timestamp), true),
('tool_name 2', 'tool_link 2', (select current_timestamp), true),
('tool_name 3', 'tool_link 3', (select current_timestamp), true);


create table book(
	book_id int GENERATED ALWAYS AS IDENTITY,
	book_name varchar(50),
	book_link varchar(50),
	date_created TIMESTAMP,
	is_active bool,
	PRIMARY KEY(book_id)
);
select * from book;

insert into book(book_name, book_link, date_created, is_active) VALUES 
('book_name 1', 'book_link 1', (select current_timestamp), true),
('book_name 2', 'book_link 2', (select current_timestamp), true),
('book_name 3', 'book_link 3', (select current_timestamp), true);

create table Blog (
	Blog_id int GENERATED ALWAYS AS IDENTITY,
	blog_text text,
	video_id int,
	book_id int,
	tools_id int,
	reference_link varchar(50),
	date_created TIMESTAMP,
	is_active bool,
	  CONSTRAINT fk_video
      FOREIGN KEY(video_id) 
	  REFERENCES Video_table(video_id),
	  CONSTRAINT fk_book
      FOREIGN KEY(book_id) 
	  REFERENCES book(book_id),
	  CONSTRAINT fk_tools
      FOREIGN KEY(tools_id) 
	  REFERENCES tools(tools_id)
);

insert into Blog(blog_text,video_id,book_id,tools_id,reference_link,date_created,is_active) Values 
('blog text one', (select video_id from Video_table where video_topic = 'test video topic 1'), (select book_id from book where book_name = 'book_name 1'),(select tools_id from tools where tool_name = 'tool_name 1'),'reference link 1',(select current_timestamp), true),
('blog text two', (select video_id from Video_table where video_topic = 'test video topic 2'), (select book_id from book where book_name = 'book_name 2'),(select tools_id from tools where tool_name = 'tool_name 2'),'reference link 2',(select current_timestamp), true),
('blog text three', (select video_id from Video_table where video_topic = 'test video topic 3'), (select book_id from book where book_name = 'book_name 3'),(select tools_id from tools where tool_name = 'tool_name 3'),'reference link 3',(select current_timestamp), true);


drop table Blog;
select * from Blog;

create table user_role(
	role_id int GENERATED ALWAYS AS IDENTITY,
	user_role varchar,
	is_active bool,
	PRIMARY KEY(role_id)
)

insert into user_role(user_role, is_active)values
('admin', true),
('developers', true),
('tester', true),
('web_user', false);

create table user_table(
	user_id int GENERATED ALWAYS AS IDENTITY,
	user_name varchar(50),
	password varchar(50),
	email varchar(50),
	mob_no bigint,
	user_role int,
	date_created TIMESTAMP,
	is_active bool,
	PRIMARY KEY(user_id),
	  CONSTRAINT fk_user_role
      FOREIGN KEY(user_role) 
	  REFERENCES user_role(role_id)
);
select * from user_table;

drop table user_table;

insert into user_table(user_name,password,email,mob_no,user_role,date_created,is_active)values
('gaurav', 'Admin_123', 'gaurav@blog.com', 9988227472, (select role_id from user_role where user_role = 'admin' and is_active = true),(select current_timestamp), true),
('shubham', 'dev_123', 'shubham@blog.com', 1122334455, (select role_id from user_role where user_role = 'developers' and is_active = true),(select current_timestamp), true);

create table resources(
	resource_id int GENERATED ALWAYS AS IDENTITY,
	resource_link varchar(50),
	description varchar(50),
	pic varchar(50),
	date_created TIMESTAMP,
	is_active bool,
	PRIMARY KEY(resource_id)
);
drop table resources;

insert into resources(resource_link,description,pic,date_created,is_active)values
('resources link 1', 'description 1', 'pic link 1',(select current_timestamp), true),
('resources link 2', 'description 2', 'pic link 2',(select current_timestamp), true),
('resources link 3', 'description 3', 'pic link 3',(select current_timestamp), true);

