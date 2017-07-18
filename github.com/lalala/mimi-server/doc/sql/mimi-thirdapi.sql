SET client_min_messages = warning;

CREATE TABLE bankcard (
    id SERIAL primary key NOT NULL,
    idcard_no text NOT NULL,    /* 身份证号码 */
    bankcard_no text NOT NULL,  /* 银行卡号 */
    card_type text NOT NULL,    /* 银行卡类型(借记卡/储蓄卡/...)*/
    bankname text NOT NULL,     /* 开户行 */ 
    phone text NOT NULL,        /* 预留手机号 */
    add_time integer NOT NULL  /* 添加时间 */
);

CREATE TABLE idcard (
    id SERIAL primary key NOT NULL,
    cardno text NOT NULL,       /* 身份证号码 */
    realname text NOT NULL,     /* 真实姓名 */
    add_time integer NOT NULL  /* 添加时间 */
);

CREATE TABLE sms_detail (
    id SERIAL primary key NOT NULL,
    content text NOT NULL,              /* 短信内容 */
    vendor text NOT NULL,               /* 短信服务供应商 */
    send_time integer NOT NULL,         /* 发送时间 */
    is_success boolean NOT NULL,        /* 是否发送成功 */
    result text NOT NULL,               /* 发送结果 */
    sms_type integer NOT NULL,          /* 短信类型(消息/验证码) */
    sign text NOT NULL,                 /* 短信签名 */
    code text,                          /* 验证码 */
    expire_time integer                 /* 验证码过期时间 */
);

CREATE TABLE sms_history (
    id SERIAL primary key NOT NULL,
    phone text NOT NULL,                /* 手机号 */
    detail_id integer NOT NULL          /* sms_detail表id */
);

CREATE TABLE sms_stats (
    id SERIAL primary key NOT NULL,
    phone text NOT NULL,                /* 手机号 */
    count_time integer NOT NULL,       /* 最近一小时计数开始时间 */
    latest_time integer NOT NULL,       /* 最近一条短信发送时间 */
    hourly_count integer NOT NULL,      /* 一小时发送数量计数 */
    daily_count integer NOT NULL,       /* 每日发送数量计数 */
    total_count integer NOT NULL        /* 发送数量计数 */
);
ALTER TABLE ONLY sms_stats ADD CONSTRAINT sms_stats_phone_key UNIQUE (phone);

CREATE TABLE opencai_bjpk10 (
    issue text NOT NULL,                        /* 期号 */
    opentime timestamp with time zone NOT NULL, /* 开奖时间 */
    balls text NOT NULL,                        /* 开奖号码 */
    grabtime timestamp with time zone NOT NULL, /* 抓取时间 */
    grabsource text NOT NULL,                   /* 抓取数据源 */
    detail json                                 /* 开奖详情 */
);
ALTER TABLE ONLY opencai_bjpk10 ADD CONSTRAINT opencai_bjpk10_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_cqssc (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_cqssc ADD CONSTRAINT opencai_cqssc_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_gd11x5 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_gd11x5 ADD CONSTRAINT opencai_gd11x5_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_dlt (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_dlt ADD CONSTRAINT opencai_dlt_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_fc3d (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_fc3d ADD CONSTRAINT opencai_fc3d_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_ssq (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_ssq ADD CONSTRAINT opencai_ssq_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_pl3 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_pl3 ADD CONSTRAINT opencai_pl3_pkey PRIMARY KEY (issue);

CREATE TABLE opencai_pl5 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY opencai_pl5 ADD CONSTRAINT opencai_pl5_pkey PRIMARY KEY (issue);

