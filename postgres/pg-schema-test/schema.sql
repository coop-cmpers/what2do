--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.1

-- Started on 2024-05-18 10:03:10 AEST

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3354 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 16388)
-- Name: test_table; Type: TABLE; Schema: public; Owner: pgtestuser
--

CREATE TABLE public.test_table (
    id integer NOT NULL,
    random_string text,
    random_integer integer,
    random_datetime timestamp with time zone
);


ALTER TABLE public.test_table OWNER TO pgtestuser;

--
-- TOC entry 215 (class 1259 OID 16387)
-- Name: test_table_id_seq; Type: SEQUENCE; Schema: public; Owner: pgtestuser
--

CREATE SEQUENCE public.test_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.test_table_id_seq OWNER TO pgtestuser;

--
-- TOC entry 3355 (class 0 OID 0)
-- Dependencies: 215
-- Name: test_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pgtestuser
--

ALTER SEQUENCE public.test_table_id_seq OWNED BY public.test_table.id;


--
-- TOC entry 3203 (class 2604 OID 16391)
-- Name: test_table id; Type: DEFAULT; Schema: public; Owner: pgtestuser
--

ALTER TABLE ONLY public.test_table ALTER COLUMN id SET DEFAULT nextval('public.test_table_id_seq'::regclass);


--
-- TOC entry 3205 (class 2606 OID 16395)
-- Name: test_table test_table_pkey; Type: CONSTRAINT; Schema: public; Owner: pgtestuser
--

ALTER TABLE ONLY public.test_table
    ADD CONSTRAINT test_table_pkey PRIMARY KEY (id);


-- Completed on 2024-05-18 10:03:11 AEST

--
-- PostgreSQL database dump complete
--

