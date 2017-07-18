create table log(
	 id SERIAL PRIMARY KEY,
	 user_id int not null,
	 path text not null,
	 operating int not null,
	 params text not null,
	 message text not null,
	 create_time int not null
)

-- 礼包模板表
create table gift_template (
	 id SERIAL primary key NOT NULL,
	 title text not null,
	 content_desc text not null,
	 content json not null,
	 add_time int not null,
	 creator text not null
)

-- 用户礼包关系表
create table user_gift (
	 id SERIAL primary key NOT NULL,
	 gift_template_id int not null,
	 act_activity_id int not null,
	 account_id int not null,
	 content json not null,
	 status int not null,
	 create_time int not null, -- 礼包产生的时间
	 receive_time int          -- 领取礼包的时间
)
