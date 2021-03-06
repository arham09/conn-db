--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: invoice_groups; Type: TABLE; Schema: public; Owner: medea
--

CREATE TABLE public.invoice_groups (
    id bigint NOT NULL,
    code character varying(50) DEFAULT NULL::character varying,
    name character varying(100) NOT NULL,
    minimum_order double precision,
    status character varying(15) DEFAULT 'active'::character varying NOT NULL,
    external_id character varying(100) DEFAULT NULL::character varying,
    supplier_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.invoice_groups OWNER TO medea;

--
-- Name: invoice_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: medea
--

CREATE SEQUENCE public.invoice_groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.invoice_groups_id_seq OWNER TO medea;

--
-- Name: invoice_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: medea
--

ALTER SEQUENCE public.invoice_groups_id_seq OWNED BY public.invoice_groups.id;


--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: medea
--

CREATE TABLE public.suppliers (
    id bigint NOT NULL,
    code character varying(50) DEFAULT NULL::character varying,
    name character varying(100) NOT NULL,
    address text,
    longitude real,
    latitude real,
    phone_no character varying(20),
    status character varying(25) DEFAULT 'inactive'::character varying NOT NULL,
    urban_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.suppliers OWNER TO medea;

--
-- Name: suppliers_id_seq; Type: SEQUENCE; Schema: public; Owner: medea
--

CREATE SEQUENCE public.suppliers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.suppliers_id_seq OWNER TO medea;

--
-- Name: suppliers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: medea
--

ALTER SEQUENCE public.suppliers_id_seq OWNED BY public.suppliers.id;


--
-- Name: invoice_groups id; Type: DEFAULT; Schema: public; Owner: medea
--

ALTER TABLE ONLY public.invoice_groups ALTER COLUMN id SET DEFAULT nextval('public.invoice_groups_id_seq'::regclass);


--
-- Name: suppliers id; Type: DEFAULT; Schema: public; Owner: medea
--

ALTER TABLE ONLY public.suppliers ALTER COLUMN id SET DEFAULT nextval('public.suppliers_id_seq'::regclass);


--
-- Data for Name: invoice_groups; Type: TABLE DATA; Schema: public; Owner: medea
--

COPY public.invoice_groups (id, code, name, minimum_order, status, external_id, supplier_id, created_at, updated_at, deleted_at) FROM stdin;
1	\N	Lakmé	100000	active	\N	1	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
2	\N	Tigaraksa #1 (SGM)	100000	active	\N	2	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
3	\N	Tigaraksa #2 (Combined)	100000	active	\N	2	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
4	\N	Tigaraksa #3 (Combined)	100000	active	\N	2	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
5	\N	Tigaraksa #4 (Combined)	100000	active	\N	2	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
6	\N	Tigaraksa #5 (Combined)	100000	active	\N	2	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
7	\N	Lakmé #2	100000	active	\N	1	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
8	\N	So Good	100000	active	\N	3	2020-06-11 17:09:03.491569+07	2020-06-11 17:09:03.491569+07	\N
\.


--
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: medea
--

COPY public.suppliers (id, code, name, address, longitude, latitude, phone_no, status, urban_id, created_at, updated_at, deleted_at) FROM stdin;
1	SNA-LME	Lakmé Cosmetics	Apt. 709	3.3849	13.5561	0743 3303 583	active	1	2020-06-11 17:09:03.17348+07	2020-06-11 17:09:03.17348+07	\N
2	SNA-TRS	PT. Tigaraksa Satria Tbk	Apt. 777	10.9881	60.9387	(+62) 505 6983 863	active	1	2020-06-11 17:09:03.17348+07	2020-06-11 17:09:03.17348+07	\N
3	SNA-JPF	PT. Japfa Comfeed Indonesia	Apt. 066	44.9133	33.9131	0803 885 467	active	1	2020-06-11 17:09:03.17348+07	2020-06-11 17:09:03.17348+07	\N
\.


--
-- Name: invoice_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: medea
--

SELECT pg_catalog.setval('public.invoice_groups_id_seq', 8, true);


--
-- Name: suppliers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: medea
--

SELECT pg_catalog.setval('public.suppliers_id_seq', 3, true);


--
-- Name: invoice_groups invoice_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: medea
--

ALTER TABLE ONLY public.invoice_groups
    ADD CONSTRAINT invoice_groups_pkey PRIMARY KEY (id);


--
-- Name: suppliers suppliers_pkey; Type: CONSTRAINT; Schema: public; Owner: medea
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);


--
-- Name: supplierCodeIndex; Type: INDEX; Schema: public; Owner: medea
--

CREATE INDEX "supplierCodeIndex" ON public.suppliers USING btree (code);


--
-- Name: invoice_groups invoice_groups_supplier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: medea
--

ALTER TABLE ONLY public.invoice_groups
    ADD CONSTRAINT invoice_groups_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

