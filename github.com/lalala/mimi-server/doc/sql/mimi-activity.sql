--活动表
DROP TABLE IF EXISTS "public"."act_activity";
create table act_activity (
	 id SERIAL PRIMARY KEY  NOT NULL ,
   title VARCHAR(50) NOT NULL DEFAULT '',
   des VARCHAR (255) NOT NULL DEFAULT '',
   logo VARCHAR (255) NOT NULL DEFAULT '',
   num int4 NOT NULL DEFAULT 0,
   starttime int4 NOT NULL DEFAULT 0,
   endtime int4 NOT NULL DEFAULT 0,
   left_num int4 NOT NULL DEFAULT 0,
   package_id int2 NOT NULL DEFAULT 0,
   is_expired int2 NOT NULL DEFAULT 0,
   create_time int4 NOT NULL DEFAULT 0
);
COMMENT ON COLUMN act_activity.title IS '活动标题';
COMMENT ON COLUMN act_activity.des IS '活动描述';
COMMENT ON COLUMN act_activity.logo IS '活动logo';
COMMENT ON COLUMN act_activity.num IS '限制人数';
COMMENT ON COLUMN act_activity.starttime IS '开始时间';
COMMENT ON COLUMN act_activity.endtime IS '结束时间';
COMMENT ON COLUMN act_activity.left_num IS '剩余份数';
COMMENT ON COLUMN act_activity.package_id IS '对应礼包';
COMMENT ON COLUMN act_activity.is_expired IS '是否过期  0：未过期，1：过期';
COMMENT ON COLUMN act_activity.create_time IS '创建时间';

--活动任务中间表
DROP TABLE IF EXISTS "public"."act_activity_task";
create table act_activity_task (
	 id SERIAL PRIMARY KEY  NOT NULL ,
   act_id int4 NOT NULL DEFAULT 0,
   task_id int4 NOT NULL DEFAULT 0,
   addtime int4 NOT NULL DEFAULT 0
);
COMMENT ON COLUMN act_activity_task.act_id IS '对应活动id';
COMMENT ON COLUMN act_activity_task.task_id IS '对应任务id';

--任务表
DROP TABLE IF EXISTS "public"."act_task";
create table act_task (
	 id SERIAL PRIMARY KEY  NOT NULL ,
   name VARCHAR (50) NOT NULL DEFAULT 0,
   des VARCHAR (255) NOT NULL DEFAULT '',
   addtime int4 NOT NULL DEFAULT 0,
   is_finish int2 NOT NULL DEFAULT 0,
   type VARCHAR (20) NOT NULL DEFAULT '',
   money float4 NOT NULL DEFAULT 0.00
);
COMMENT ON COLUMN act_task.name IS '任务名';
COMMENT ON COLUMN act_task.is_finish IS '后端是否有这个任务的实现，0：没有，1：有';

DROP TABLE IF EXISTS "public"."act_user_activity";
create table act_user_activity (
	 id SERIAL PRIMARY KEY  NOT NULL ,
   act_id int4 NOT NULL DEFAULT 0,
   account_id int4 NOT NULL DEFAULT 0,
   addtime int4 NOT NULL DEFAULT 0,
   is_finish int2 NOT NULL DEFAULT 0
);
COMMENT ON COLUMN act_user_activity.act_id IS '活动id';
COMMENT ON COLUMN act_user_activity.account_id IS '账户id';

DROP TABLE IF EXISTS "public"."act_user_task";
create table act_user_task (
	 id SERIAL PRIMARY KEY  NOT NULL ,
   task_id int4 NOT NULL DEFAULT 0,
   account_id int4 NOT NULL DEFAULT 0,
   addtime int4 NOT NULL DEFAULT 0
);
COMMENT ON COLUMN act_user_task.task_id IS '任务id';
COMMENT ON COLUMN act_user_task.account_id IS '账户id';
