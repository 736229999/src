SET client_min_messages = warning;
SET client_encoding = 'UTF8';

CREATE TABLE ssq (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);

ALTER TABLE ONLY ssq ADD CONSTRAINT ssq_issue_key UNIQUE (issue);

CREATE TABLE dlt (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY dlt ADD CONSTRAINT dlt_issue_key UNIQUE (issue);

CREATE TABLE fc3d (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY fc3d ADD CONSTRAINT fc3d_issue_key UNIQUE (issue);

CREATE TABLE pl3 (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY pl3 ADD CONSTRAINT pl3_issue_key UNIQUE (issue);

CREATE TABLE pl5 (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY pl5 ADD CONSTRAINT pl5_issue_key UNIQUE (issue);

CREATE TABLE gd11x5 (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY gd11x5 ADD CONSTRAINT gd11x5_issue_key UNIQUE (issue);

CREATE TABLE bjpk10 (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY bjpk10 ADD CONSTRAINT bjpk10_issue_key UNIQUE (issue);

CREATE TABLE cqssc (
    id SERIAL primary key not null,
    issue text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    open_time timestamp with time zone NOT NULL,
    open_balls text
);
ALTER TABLE ONLY cqssc ADD CONSTRAINT cqssc_issue_key UNIQUE (issue);

CREATE TABLE play_time_settings (
	 id SERIAL primary key not null,
	 lottery_id int unique not null,
	 start_time int not null,
	 end_time int not null,
	 chase_start_time int not null
)