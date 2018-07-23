--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16387)
-- Name: account; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.account (
    id integer NOT NULL,
    auth_id character varying(40),
    username character varying(30)
);

--
-- TOC entry 197 (class 1259 OID 16390)
-- Name: account_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- TOC entry 3142 (class 0 OID 0)
-- Dependencies: 197
-- Name: account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.account_id_seq OWNED BY public.account.id;

--
-- TOC entry 198 (class 1259 OID 16392)
-- Name: phone_number; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.phone_number (
    id integer NOT NULL,
    number character varying(40),
    account_id integer
);

--
-- TOC entry 199 (class 1259 OID 16395)
-- Name: phone_number_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.phone_number_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- TOC entry 3143 (class 0 OID 0)
-- Dependencies: 199
-- Name: phone_number_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.phone_number_id_seq OWNED BY public.phone_number.id;

--
-- TOC entry 3004 (class 2604 OID 16399)
-- Name: account id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account ALTER COLUMN id SET DEFAULT nextval('public.account_id_seq'::regclass);

--
-- TOC entry 3005 (class 2604 OID 16400)
-- Name: phone_number id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.phone_number ALTER COLUMN id SET DEFAULT nextval('public.phone_number_id_seq'::regclass);

--
-- TOC entry 3132 (class 0 OID 16387)
-- Dependencies: 196
-- Data for Name: account; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.account VALUES (1, '20S0KPNOIM', 'plivo1');
INSERT INTO public.account VALUES (2, '54P2EOKQ47', 'plivo2');
INSERT INTO public.account VALUES (3, '9LLV6I4ZWI', 'plivo3');
INSERT INTO public.account VALUES (4, 'YHWE3HDLPQ', 'plivo4');
INSERT INTO public.account VALUES (5, '6DLH8A25XZ', 'plivo5');

--
-- TOC entry 3134 (class 0 OID 16392)
-- Dependencies: 198
-- Data for Name: phone_number; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.phone_number VALUES (1, '4924195509198', 1);
INSERT INTO public.phone_number VALUES (2, '4924195509196', 1);
INSERT INTO public.phone_number VALUES (3, '4924195509197', 1);
INSERT INTO public.phone_number VALUES (4, '4924195509195', 1);
INSERT INTO public.phone_number VALUES (5, '4924195509049', 1);
INSERT INTO public.phone_number VALUES (6, '4924195509012', 1);
INSERT INTO public.phone_number VALUES (7, '4924195509193', 1);
INSERT INTO public.phone_number VALUES (8, '4924195509029', 1);
INSERT INTO public.phone_number VALUES (9, '4924195509192', 1);
INSERT INTO public.phone_number VALUES (10, '4924195509194', 1);
INSERT INTO public.phone_number VALUES (11, '31297728125', 1);
INSERT INTO public.phone_number VALUES (12, '3253280312', 1);
INSERT INTO public.phone_number VALUES (13, '3253280311', 1);
INSERT INTO public.phone_number VALUES (14, '3253280315', 1);
INSERT INTO public.phone_number VALUES (15, '3253280313', 1);
INSERT INTO public.phone_number VALUES (16, '3253280329', 1);
INSERT INTO public.phone_number VALUES (17, '441224459508', 1);
INSERT INTO public.phone_number VALUES (18, '441224980086', 1);
INSERT INTO public.phone_number VALUES (19, '441224980087', 1);
INSERT INTO public.phone_number VALUES (20, '441224980096', 1);
INSERT INTO public.phone_number VALUES (21, '441224980098', 1);
INSERT INTO public.phone_number VALUES (22, '441224980099', 1);
INSERT INTO public.phone_number VALUES (23, '441224980100', 1);
INSERT INTO public.phone_number VALUES (24, '441224980094', 2);
INSERT INTO public.phone_number VALUES (25, '441224459426', 2);
INSERT INTO public.phone_number VALUES (26, '13605917249', 2);
INSERT INTO public.phone_number VALUES (27, '441224459548', 2);
INSERT INTO public.phone_number VALUES (28, '441224459571', 2);
INSERT INTO public.phone_number VALUES (29, '441224459598', 2);
INSERT INTO public.phone_number VALUES (30, '13605895047', 2);
INSERT INTO public.phone_number VALUES (31, '14433600975', 2);
INSERT INTO public.phone_number VALUES (32, '16052299352', 2);
INSERT INTO public.phone_number VALUES (33, '13602092244', 2);
INSERT INTO public.phone_number VALUES (34, '441224459590', 2);
INSERT INTO public.phone_number VALUES (35, '441224459620', 2);
INSERT INTO public.phone_number VALUES (36, '441224459660', 2);
INSERT INTO public.phone_number VALUES (37, '234568266473', 2);
INSERT INTO public.phone_number VALUES (38, '441224980091', 2);
INSERT INTO public.phone_number VALUES (39, '441224980092', 2);
INSERT INTO public.phone_number VALUES (40, '441224980089', 2);
INSERT INTO public.phone_number VALUES (41, '441224459482', 2);
INSERT INTO public.phone_number VALUES (42, '441224980093', 2);
INSERT INTO public.phone_number VALUES (43, '441887480051', 2);
INSERT INTO public.phone_number VALUES (44, '441873440028', 2);
INSERT INTO public.phone_number VALUES (45, '441873440017', 3);
INSERT INTO public.phone_number VALUES (46, '441970450009', 3);
INSERT INTO public.phone_number VALUES (47, '441235330075', 3);
INSERT INTO public.phone_number VALUES (48, '441235330053', 3);
INSERT INTO public.phone_number VALUES (49, '441235330044', 3);
INSERT INTO public.phone_number VALUES (50, '441235330078', 3);
INSERT INTO public.phone_number VALUES (51, '34881254103', 3);
INSERT INTO public.phone_number VALUES (52, '61871112946', 3);
INSERT INTO public.phone_number VALUES (53, '61871112915', 3);
INSERT INTO public.phone_number VALUES (54, '61881666904', 3);
INSERT INTO public.phone_number VALUES (55, '61881666939', 3);
INSERT INTO public.phone_number VALUES (56, '61871112913', 3);
INSERT INTO public.phone_number VALUES (57, '61871112901', 3);
INSERT INTO public.phone_number VALUES (58, '61871112938', 3);
INSERT INTO public.phone_number VALUES (59, '61871112934', 3);
INSERT INTO public.phone_number VALUES (60, '61871112902', 3);
INSERT INTO public.phone_number VALUES (61, '61881666926', 4);
INSERT INTO public.phone_number VALUES (62, '61871705936', 4);
INSERT INTO public.phone_number VALUES (63, '61871112920', 4);
INSERT INTO public.phone_number VALUES (64, '61881666923', 4);
INSERT INTO public.phone_number VALUES (65, '61871112947', 4);
INSERT INTO public.phone_number VALUES (66, '61871112948', 4);
INSERT INTO public.phone_number VALUES (67, '61871112921', 4);
INSERT INTO public.phone_number VALUES (68, '61881666914', 4);
INSERT INTO public.phone_number VALUES (69, '61881666942', 4);
INSERT INTO public.phone_number VALUES (70, '61871112922', 4);
INSERT INTO public.phone_number VALUES (71, '61871232393', 4);
INSERT INTO public.phone_number VALUES (72, '61871112916', 5);
INSERT INTO public.phone_number VALUES (73, '61881666921', 5);
INSERT INTO public.phone_number VALUES (74, '61871112905', 5);
INSERT INTO public.phone_number VALUES (75, '61871112937', 5);
INSERT INTO public.phone_number VALUES (76, '61361220301', 5);
INSERT INTO public.phone_number VALUES (77, '61871112931', 5);
INSERT INTO public.phone_number VALUES (78, '61871112939', 5);
INSERT INTO public.phone_number VALUES (79, '61871112940', 5);

--
-- TOC entry 3144 (class 0 OID 0)
-- Dependencies: 197
-- Name: account_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.account_id_seq', 5, true);

--
-- TOC entry 3145 (class 0 OID 0)
-- Dependencies: 199
-- Name: phone_number_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.phone_number_id_seq', 79, true);

--
-- TOC entry 3007 (class 2606 OID 16402)
-- Name: account account_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_pkey PRIMARY KEY (id);

--
-- TOC entry 3009 (class 2606 OID 16404)
-- Name: phone_number phone_number_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.phone_number
    ADD CONSTRAINT phone_number_pkey PRIMARY KEY (id);

--
-- TOC entry 3010 (class 2606 OID 16405)
-- Name: phone_number phone_number_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.phone_number
    ADD CONSTRAINT phone_number_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account(id);
