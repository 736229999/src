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

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: banner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE banner (
    id integer NOT NULL,
    url text,
    link_type text,
    link text,
    status boolean DEFAULT true,
    label text,
    created integer
);


ALTER TABLE banner OWNER TO postgres;

--
-- Name: COLUMN banner.url; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.url IS '图片地址';


--
-- Name: COLUMN banner.link_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.link_type IS '连接类型';


--
-- Name: COLUMN banner.link; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.link IS '点击链接';


--
-- Name: COLUMN banner.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.status IS '显示状态';


--
-- Name: COLUMN banner.label; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.label IS '描述';


--
-- Name: COLUMN banner.created; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN banner.created IS '创建时间';


--
-- Name: banner_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE banner_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE banner_id_seq OWNER TO postgres;

--
-- Name: banner_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE banner_id_seq OWNED BY banner.id;


--
-- Name: news; Type: TABLE; Schema: public; Owner: yangzefeng
--

CREATE TABLE news (
    id bigint NOT NULL,
    content text,
    title text NOT NULL,
    description text,
    pageviews bigint DEFAULT '0'::bigint NOT NULL,
    author text NOT NULL,
    updated integer,
    html text,
    cover text,
    is_visible boolean DEFAULT true,
    news_calss integer DEFAULT 0,
    created integer
);


ALTER TABLE news OWNER TO yangzefeng;

--
-- Name: COLUMN news.content; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.content IS '内容';


--
-- Name: COLUMN news.title; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.title IS '标题';


--
-- Name: COLUMN news.description; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.description IS '描述';


--
-- Name: COLUMN news.pageviews; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.pageviews IS '浏览量';


--
-- Name: COLUMN news.author; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.author IS '作者';


--
-- Name: COLUMN news.updated; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.updated IS '更新时间';


--
-- Name: COLUMN news.html; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.html IS '编译生成的html';


--
-- Name: COLUMN news.cover; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.cover IS '封面';


--
-- Name: COLUMN news.is_visible; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.is_visible IS '可见性';


--
-- Name: COLUMN news.news_calss; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.news_calss IS '新闻类别';


--
-- Name: COLUMN news.created; Type: COMMENT; Schema: public; Owner: yangzefeng
--

COMMENT ON COLUMN news.created IS '创建时间';


--
-- Name: news_id_seq; Type: SEQUENCE; Schema: public; Owner: yangzefeng
--

CREATE SEQUENCE news_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE news_id_seq OWNER TO yangzefeng;

--
-- Name: news_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: yangzefeng
--

ALTER SEQUENCE news_id_seq OWNED BY news.id;


--
-- Name: banner id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY banner ALTER COLUMN id SET DEFAULT nextval('banner_id_seq'::regclass);


--
-- Name: news id; Type: DEFAULT; Schema: public; Owner: yangzefeng
--

ALTER TABLE ONLY news ALTER COLUMN id SET DEFAULT nextval('news_id_seq'::regclass);


--
-- Data for Name: banner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY banner (id, url, link_type, link, status, label, created) FROM stdin;
\.


--
-- Name: banner_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('banner_id_seq', 1, false);


--
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: yangzefeng
--

COPY news (id, content, title, description, pageviews, author, updated, html, cover, is_visible, news_calss, created) FROM stdin;
1	<div align="center">\n![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。\n# 等待二十年终于盼来百万大奖 \n\n20年是什么概念？相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497077425	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n<img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。</p>\n<h1 id="-">等待二十年终于盼来百万大奖</h1>\n<p>20年是什么概念？相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497077425
2	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/d13016449cb1460bf8141081d388bb87)\n</div>\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n\n# 这是一个测试标题\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n	阿森纳0-0贝西克塔斯，拉姆塞染...	北京时间今天凌晨，2014-15赛季欧冠(微博 专题) ...	0	亨通彩	1497086529	<!DOCTYPE html>\n<html>\n<head>\n\t<title>阿森纳0-0贝西克塔斯，拉姆塞染...</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">阿森纳0-0贝西克塔斯，拉姆塞染...</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/d13016449cb1460bf8141081d388bb87" alt="">\n</div>\n\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n<h1 id="-">这是一个测试标题</h1>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/d13016449cb1460bf8141081d388bb87	t	0	1497086529
3	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0)\n</div>\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n\n# 这是一个测试标题\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\n	关于重庆时时彩派奖活动公告	为庆祝体育彩票超级重庆时时彩上市十周年...	0	亨通彩	1497087184	<!DOCTYPE html>\n<html>\n<head>\n\t<title>关于重庆时时彩派奖活动公告</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">关于重庆时时彩派奖活动公告</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0" alt="">\n</div>\n\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n<h1 id="-">这是一个测试标题</h1>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/0e8814fe61bfdc3309848febc879ccc6	t	0	1497087184
4	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087327	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087327
5	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087329	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087329
6	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087331	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087331
7	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087338	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087338
8	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087339	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087339
9	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087341	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087341
10	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	阿森纳0-0贝西克塔斯，拉姆塞染...	为庆祝体育彩票超级重庆时时彩上市十周年...	0	亨通彩	1497087446	<!DOCTYPE html>\n<html>\n<head>\n\t<title>阿森纳0-0贝西克塔斯，拉姆塞染...</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">阿森纳0-0贝西克塔斯，拉姆塞染...</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0	t	0	1497087446
11	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	关于重庆时时彩派奖活动公告	为庆祝体育彩票超级重庆时时彩上市十周年...	0	亨通彩	1497087481	<!DOCTYPE html>\n<html>\n<head>\n\t<title>关于重庆时时彩派奖活动公告</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">关于重庆时时彩派奖活动公告</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0	t	0	1497087481
12	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0)\n</div>\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。\n\n# 这是一个测试标题\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？	关于重庆时时彩派奖活动公告	为庆祝体育彩票超级重庆时时彩上市十周年...	0	亨通彩	1497087556	<!DOCTYPE html>\n<html>\n<head>\n\t<title>关于重庆时时彩派奖活动公告</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">关于重庆时时彩派奖活动公告</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0" alt="">\n</div>\n\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。</p>\n<h1 id="-">这是一个测试标题</h1>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/d6df9b6082bb59df85c863b5c9bc31d0	t	0	1497087556
13	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/f85197351b3b569c11fc5022995340b1)\n</div>\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。\n\n# 这是一个测试标题\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？\n\n测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？	阿森纳0-0贝西克塔斯，拉姆塞染...	北京时间今天凌晨，2014-15赛季欧冠(微博 专题) ...	0	亨通彩	1497087615	<!DOCTYPE html>\n<html>\n<head>\n\t<title>阿森纳0-0贝西克塔斯，拉姆塞染...</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">阿森纳0-0贝西克塔斯，拉姆塞染...</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/f85197351b3b569c11fc5022995340b1" alt="">\n</div>\n\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。</p>\n<h1 id="-">这是一个测试标题</h1>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？</p>\n<p>测试测试测试测试测试测试测试测试测试测试测试测试测试，测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测，试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试。测试测试出测试测试？</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/f85197351b3b569c11fc5022995340b1	t	0	1497087615
14	<div align="center">\n                ![](http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8)\n</div>\n\n5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 \n\n辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 \n\n# 等待二十年终于盼来百万大奖 20年是什么概念？\n\n相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 \n\n以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”	二十年老站喜迎双色球头奖	5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖...	0	亨通彩	1497087682	<!DOCTYPE html>\n<html>\n<head>\n\t<title>二十年老站喜迎双色球头奖</title>\n</head>\n<body>\n\t<div class="article">\n\t\t<div class="article-header">\n\t\t\t<div>\n\t\t\t\t<h1 class="article-title">二十年老站喜迎双色球头奖</h1>\n\t\t\t</div>\n\t\t\t<div>\n\t\t\t\t<p class="article-author">2017-06-10 &nbsp亨通彩</p>\n\t\t\t</div>\n\t\t</div>\n\t\t<div class="article-content">\n\t\t<div align="center">\n                <img src="http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8" alt="">\n</div>\n\n<p>5月14日晚，中国福利彩票双色球游戏进行第2017055期开奖。当期双色球红球号码为07、12、13、20、24、31，蓝球号码为05。 </p>\n<p>辽宁中出的一等奖出自沈阳福彩007号投注站，彩站位于皇姑区黄河南大街62号，彩票是一张6+4蓝号复式票，共计花费8元钱，中奖者除中得1注一等奖外，还中得3注二等奖，总奖金6,212,320元，这也是沈阳市今年中出的第7注双色球头奖。 </p>\n<h1 id="-20-">等待二十年终于盼来百万大奖 20年是什么概念？</h1>\n<p>相信有很多刚接触彩票的年轻人在20年前才刚刚出生，而007号投注站的站主董大哥，经营他这家投注站已经整整20年的时间了。 </p>\n<p>以后我的手也可以称为幸运之手了，因为这票就是我打出来的，也希望我能为更多的彩民朋友带去好运，也祝愿沈阳百姓都能中得大奖。”</p>\n\n\t\t</div>\n\t\t<div class="article-footer">\n\t\t\t<p class="article-pageviews">{{.PageViews}}人浏览 &nbsp&nbsp&nbsp2017-06-10</p>\n\t\t</div>\n\t</div>\n</body>\n    <style type="text/css">\n\t.article {\n\t\tmargin: 12px 28px;\n\t\tfont-size: 2.0em;\n\t\tfont-family: "MicrosoftYaHei";\n\t}\t\n\t.article-title {\n\t\tfont-size: 2.0em;\n\t\tfont-weight: normal;\n\t\tmargin-bottom: 0;\n\t}\t\n\t.article-author {\n\t\tcolor: #BBBBBB;\n\t}\n\t.article-pageviews {\n\t\tcolor: #BBBBBB;\n\t\tmargin-top: 10%;\n\t}\n\t.article img {\n\t\tmargin: 3% auto;\n\t\twidth: 100%;\n\t}\n    .article-content p {\n        margin: 5% auto;\n        font-size: 1.1em;\n        line-height:160%;\n    }\n    .article-footer {\n        margin-bottom: 5%\n    }\n</style>\n</html>	http://cptest.kxkr.com:8088/assets/download/news/fe2b776ea2dbefeeaecd7dd03d1bf2f8	t	0	1497087682
\.


--
-- Name: news_id_seq; Type: SEQUENCE SET; Schema: public; Owner: yangzefeng
--

SELECT pg_catalog.setval('news_id_seq', 14, true);


--
-- Name: banner banner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY banner
    ADD CONSTRAINT banner_pkey PRIMARY KEY (id);


--
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: yangzefeng
--

ALTER TABLE ONLY news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

