SET client_min_messages = warning;


-- 账户
CREATE TABLE account (
    id SERIAL primary key NOT NULL,     /* 账户id */
    phone_id integer,                   /* phone_user表id */
    qq_id integer,                      /* qq_user表id */
    weixin_id integer                   /* weixin_user表id */
);

-- 账户变动历史
CREATE TABLE account_history (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,        /* 账户id */
    change_type integer NOT NULL,       /* 变动类型 */
    user_type integer NOT NULL,         /* 用户类型(phone/qq/weixin) */
    user_id integer NOT NULL,           /* 用户id(phone_user/qq_user/weixin_user表id */ 
    change_time integer NOT NULL,       /* 变动时间 */
    ip text NOT NULL,                   /* 用户ip地址 */
    detail text                         /* 变动详情 */
);

CREATE TABLE bankcard (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,        /* 账户id */
    cardno text NOT NULL,               /* 银行卡号 */
    bankname integer NOT NULL,          /* 开户行枚举 */
    cardtype text NOT NULL,             /* 卡类型 */
    phone text NOT NULL,                /* 预留手机号 */
    idno text NOT NULL,                 /* 身份证号码 */
    realname text NOT NULL,             /* 真实姓名 */
    add_time integer NOT NULL           /* 添加时间 */
);
ALTER TABLE ONLY bankcard ADD CONSTRAINT bankcard_account_id_key UNIQUE (account_id);


CREATE TABLE buycai_user_order (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,
    lottery_id integer NOT NULL,
    issue_num integer DEFAULT 1 NOT NULL,
    chase_no integer DEFAULT 0 NOT NULL,
    cai double precision NOT NULL,
    balance double precision NOT NULL,
    sum_money double precision NOT NULL,
    issues json NOT NULL,
    scheme_list json NOT NULL,
    order_time integer NOT NULL,
    ticket_sub_money double precision NOT NULL,
    status integer NOT NULL,
    is_win_stop boolean DEFAULT false NOT NULL,
    cost_cai double precision NOT NULL,
    cost_balance double precision NOT NULL
);


CREATE TABLE buycai_vendor_order (
    id SERIAL primary key NOT NULL,
    user_order_id integer NOT NULL,
    account_id integer NOT NULL,
    lottery_id integer NOT NULL,
    issue text NOT NULL,
    sum_num integer NOT NULL,
    multiple integer NOT NULL,
    money double precision NOT NULL,
    cai double precision NOT NULL,
    balance double precision NOT NULL,
    chase_no integer NOT NULL,
    vendor text NOT NULL,
    scheme_list json NOT NULL,
    add_time integer NOT NULL,
    status integer NOT NULL,
    status_desc text NOT NULL,
    vendor_req_time integer NOT NULL,
    vendor_resp_time integer NOT NULL,
    vendor_resp_id text NOT NULL,
    win_money double precision NOT NULL
);

CREATE TABLE client_device (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,
    device_id text NOT NULL,
    device_os integer NOT NULL,
    device_os_version text NOT NULL,
    device_model text NOT NULL
);

CREATE TABLE credits_history (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,
    var integer NOT NULL,
    remain integer NOT NULL,
    reason integer NOT NULL,
    change_time integer NOT NULL,
    detail text
);

CREATE TABLE exchange_history (
    id SERIAL primary key not null,
    account_id integer NOT NULL,
    batch integer NOT NULL,
    cdkey integer NOT NULL,
    exchange_time integer NOT NULL
);


CREATE TABLE fund (
    account_id integer DEFAULT 0 NOT NULL,
    cai double precision DEFAULT 0 NOT NULL,
    balance double precision DEFAULT 0 NOT NULL,
    freeze_cai double precision DEFAULT 0 NOT NULL,
    freeze_balance double precision DEFAULT 0 NOT NULL,
    total_recharge double precision DEFAULT 0 NOT NULL,
    total_win double precision DEFAULT 0 NOT NULL,
    total_withdraw double precision DEFAULT 0 NOT NULL,
    total_cai double precision DEFAULT 0 NOT NULL,
    total_buycai double precision DEFAULT 0 NOT NULL
);
ALTER TABLE ONLY fund ADD CONSTRAINT fund_pkey PRIMARY KEY (account_id);

CREATE TABLE fund_history (
    id BIGSERIAL primary key NOT NULL,
    account_id bigint NOT NULL,
    change_type smallint NOT NULL,
    cai double precision NOT NULL,
    freeze_cai double precision NOT NULL,
    balance double precision NOT NULL,
    freeze_balance double precision NOT NULL,
    var_balance double precision DEFAULT 0 NOT NULL,
    var_freeze_balance double precision DEFAULT 0 NOT NULL,
    var_cai double precision DEFAULT 0 NOT NULL,
    var_freeze_cai double precision DEFAULT 0 NOT NULL,
    recharge_order_no text DEFAULT '' NOT NULL,
    vendor_order_id bigint DEFAULT 0 NOT NULL,
    user_order_id bigint DEFAULT 0 NOT NULL,
    withdraw_apply_id bigint DEFAULT 0 NOT NULL,
    change_time integer NOT NULL,
    change_comment text
);

CREATE TABLE gift_package (
    id SERIAL primary key NOT NULL,
    title text NOT NULL,
    content_desc text NOT NULL,
    content json NOT NULL,
    gift_type integer NOT NULL,
    add_time integer NOT NULL
);

CREATE TABLE idcard (
    account_id integer NOT NULL,        /* 账户id */
    realname text NOT NULL,             /* 真实姓名 */
    idno text NOT NULL,                 /* 身份证号码 */
    add_time integer NOT NULL           /* 添加时间 */
);
ALTER TABLE ONLY idcard ADD CONSTRAINT idcard_pkey PRIMARY KEY (account_id);


CREATE TABLE invite_history (
    id SERIAL primary key NOT NULL,
    inviter integer NOT NULL,
    invitee integer NOT NULL,
    accept_time integer NOT NULL,
    credits integer NOT NULL,
    tickets_num integer NOT NULL,
    tickets_money integer NOT NULL
);
CREATE UNIQUE INDEX invite_history_inviter_invitee_idx ON invite_history USING btree (inviter, invitee);


CREATE TABLE kxd_history (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,
    var integer NOT NULL,
    remain integer NOT NULL,
    reason integer NOT NULL,
    change_time integer NOT NULL,
    detail text
);

CREATE TABLE phone_user (
    id SERIAL primary key NOT NULL,
    account_id integer DEFAULT 0 NOT NULL,
    phone text NOT NULL,
    password text
);
ALTER TABLE ONLY phone_user ADD CONSTRAINT phone_user_phone_key UNIQUE (phone);


CREATE TABLE qq_user (
    id SERIAL primary key NOT NULL,
    account_id integer DEFAULT 0 NOT NULL,
    openid text NOT NULL
);
ALTER TABLE ONLY qq_user ADD CONSTRAINT qq_user_openid_key UNIQUE (openid);


CREATE TABLE ticket (
    id SERIAL primary key NOT NULL,
    account_id integer NOT NULL,
    use_base integer NOT NULL,
    use_sub integer NOT NULL,
    valid_start integer NOT NULL,
    valid_end integer NOT NULL,
    add_time integer NOT NULL,
    title text NOT NULL,
    restrict_id integer NOT NULL,
    restrict_type integer NOT NULL,
    order_id integer
);

CREATE TABLE userinfo (
    account_id integer NOT NULL,        /* 账户id */
    exp integer NOT NULL,               /* 当前成长值 */
    lvl integer NOT NULL,               /* 当前等级 */
    nickname text,                      /* 昵称 */
    icon text,                          /* 头像 */
    sex integer,                        /* 性别 */
    invitation_code text,               /* 邀请码 */
    pay_password text,                  /* 支付密码 */
    open_pay_password boolean,          /* 是否开启支付密码 */
    cont_check_days integer,            /* 连续签到天数 */
    daily_check_time integer            /* 每日签到时间 */
);
ALTER TABLE ONLY userinfo ADD CONSTRAINT user_info_pkey PRIMARY KEY (account_id);
ALTER TABLE ONLY userinfo ADD CONSTRAINT user_info_nickname_key UNIQUE (nickname);
ALTER TABLE ONLY userinfo ADD CONSTRAINT userinfo_invitation_code_key UNIQUE (invitation_code);


CREATE TABLE virtual_fund (
    account_id integer NOT NULL,
    credits integer NOT NULL,
    kxd integer NOT NULL
);
ALTER TABLE ONLY virtual_fund ADD CONSTRAINT virtual_fund_pkey PRIMARY KEY (account_id);


CREATE TABLE weixin_user (
    id SERIAL primary key NOT NULL,
    account_id integer DEFAULT 0 NOT NULL,
    openid text NOT NULL
);
ALTER TABLE ONLY weixin_user ADD CONSTRAINT weixin_user_openid_key UNIQUE (openid);


CREATE TABLE zf_balance_history (
    id SERIAL primary key NOT NULL,
    vendor_order_id integer NOT NULL,
    zf_order_id text NOT NULL,
    sum_money double precision NOT NULL,
    balance double precision NOT NULL,
    order_time integer NOT NULL
);
ALTER TABLE ONLY zf_balance_history ADD CONSTRAINT zf_balance_history_vendor_order_id_key UNIQUE (vendor_order_id);
ALTER TABLE ONLY zf_balance_history ADD CONSTRAINT zf_balance_history_zf_order_id_key UNIQUE (zf_order_id);

CREATE TABLE phone_regist_gift_history (
    account_id integer NOT NULL,
    receive_time integer NOT NULL
);
ALTER TABLE phone_regist_gift_history ADD CONSTRAINT phone_regist_gift_history_pkey PRIMARY KEY (account_id);


CREATE TABLE cdkey_batch (
    id SERIAL primary key NOT NULL,
    max_exchange integer NOT NULL,
    cdkey_desc text NOT NULL,
    title text NOT NULL,
    gift_package_id integer NOT NULL,
    valid_start integer NOT NULL,
    valid_end integer NOT NULL,
    add_time integer NOT NULL
);

-- 提现申请表
CREATE TABLE withdraw_apply (
    id BIGSERIAL primary key NOT NULL,
    account_id bigint NOT NULL,					/* 账户id */
    realname text NOT NULL,						/* 真实姓名 */
    create_time integer NOT NULL,				/* 申请时间 */
    amount float8 NOT NULL,						/* 提现金额 */
    in_bankname text NOT NULL DEFAULT '', 		/* 开户行 */
    in_no text NOT NULL,							/* 用户收款银行卡 */
    phone text NOT NULL,							/* 手机号码 */
    idcard_no text NOT NULL,						/* 身份证号码 */
    step int NOT NULL DEFAULT 1,					/* 当前步骤 1: 提现申请, 2: 平台审核, 3: 银行处理 */
    is_success boolean NOT NULL DEFAULT FALSE, 	/* 当前步骤是否成功 */
    status integer NOT NULL DEFAULT 0,			/* 提现状态 */
    auditor text DEFAULT '',						/* 审核人员 */
    audit_time integer DEFAULT 0,				/* 审核时间 */
    audit_comment text DEFAULT '',				/* 审核评论 */
    withdraw_type integer NOT NULL				/* 提现类型(支付宝/银行卡) */
);

