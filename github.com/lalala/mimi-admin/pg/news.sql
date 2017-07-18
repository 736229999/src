--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.2
-- Dumped by pg_dump version 9.6.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: news; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE news (
    id bigint NOT NULL,
    content text,
    created integer,
    title text,
    description text,
    pageviews integer,
    author text,
    updated integer,
    html text,
    cover text,
    is_visible boolean DEFAULT true,
    news_class integer
);


ALTER TABLE news OWNER TO postgres;

--
-- Name: COLUMN news.content; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.content IS '内容';


--
-- Name: COLUMN news.created; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.created IS '创建时间';


--
-- Name: COLUMN news.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.title IS '标题';


--
-- Name: COLUMN news.description; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.description IS '描述';


--
-- Name: COLUMN news.pageviews; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.pageviews IS '浏览量';


--
-- Name: COLUMN news.author; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.author IS '作者';


--
-- Name: COLUMN news.updated; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.updated IS '更新时间';


--
-- Name: COLUMN news.html; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.html IS '编译生成的html';


--
-- Name: COLUMN news.cover; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.cover IS '封面url';


--
-- Name: COLUMN news.is_visible; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.is_visible IS '状态';


--
-- Name: COLUMN news.news_class; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN news.news_class IS '新闻类别';


--
-- Name: new_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE new_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE new_id_seq OWNER TO postgres;

--
-- Name: new_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE new_id_seq OWNED BY news.id;


--
-- Name: news id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY news ALTER COLUMN id SET DEFAULT nextval('new_id_seq'::regclass);


--
-- Name: news new_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY news
    ADD CONSTRAINT new_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

