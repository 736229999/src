SET client_min_messages = warning;

CREATE TABLE bjpk10 (
    issue text NOT NULL,                        /* 期号 */
    opentime timestamp with time zone NOT NULL, /* 开奖时间 */
    balls text NOT NULL,                        /* 开奖号码 */
    grabtime timestamp with time zone NOT NULL, /* 抓取时间 */
    grabsource text NOT NULL,                   /* 抓取数据源 */
    detail json                                 /* 开奖详情 */
);
ALTER TABLE ONLY bjpk10 ADD CONSTRAINT bjpk10_pkey PRIMARY KEY (issue);

CREATE TABLE cqssc (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY cqssc ADD CONSTRAINT cqssc_pkey PRIMARY KEY (issue);

CREATE TABLE gd11x5 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY gd11x5 ADD CONSTRAINT gd11x5_pkey PRIMARY KEY (issue);

CREATE TABLE dlt (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY dlt ADD CONSTRAINT dlt_pkey PRIMARY KEY (issue);

CREATE TABLE fc3d (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY fc3d ADD CONSTRAINT fc3d_pkey PRIMARY KEY (issue);

CREATE TABLE ssq (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY ssq ADD CONSTRAINT ssq_pkey PRIMARY KEY (issue);

CREATE TABLE pl3 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY pl3 ADD CONSTRAINT pl3_pkey PRIMARY KEY (issue);

CREATE TABLE pl5 (
    issue text NOT NULL,
    opentime timestamp with time zone NOT NULL,
    balls text NOT NULL,
    grabtime timestamp with time zone NOT NULL,
    grabsource text NOT NULL,
    detail json
);
ALTER TABLE ONLY pl5 ADD CONSTRAINT pl5_pkey PRIMARY KEY (issue);

