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
	);

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
);

create table lottery_options (
	 id int primary key NOT NULL,
	 lottery_name text unique not null,
	 is_plus_award bool not null,
	 info text not null,
	 stop_sale bool not null,
	 create_time int not null,
	 update_time int not null
);

CREATE table feedback (
  id int primary key NOT NULL,
  email text,
  name text not null,
  content text not null,
  status integer not null,
  create_time int not null
);
