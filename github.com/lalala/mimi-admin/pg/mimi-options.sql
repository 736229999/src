create table lottery_options (
	 id int primary key NOT NULL,
	 lottery_name text unique not null,
	 is_plus_award bool,
	 info text not null,
	 stop_sale bool not null,
	 create_time int not null,
	 update_time int not null
)

create table news (
	id BIGSERIAL PRIMARY KEY,
	content text,
	title text NOT NULL,
	description text,
	pageviews bigint default 0 NOT NULL,
	author text NOT NULL,
	updated int,
	html text,
	cover text,
	is_visible bool default true,
	news_class int default 0,
	created int
)

create table banner (
	id BIGSERIAL PRIMARY KEY,
	url text,
	target_type int,
	target_link text,
	is_visible bool default true,
	description text,
	created int,
	sort int,
	target_id bigint,
	location int,
	updated bigint
)

create table contact (
	id SERIAL primary key not null,
	qq text,
	wechat text,
	email text,
	telphone text
)
create table feedback (
	id SERIAL primary key not null,
	email text,
	name text not null,
	content text not null,
	create_time int not null
)

create table faq (
	id BIGSERIAL PRIMARY KEY,
	title text NOT NULL default '',
	content_url text default '',
	create_time int,
	update_time int,
	content text NOT NULL default '',
	is_visible bool default true,
	html text NOT NULL default '',
	sort int default 0
)
