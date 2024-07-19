--
-- PostgreSQL database dump
--

-- Dumped from database version 14.12 (Homebrew)
-- Dumped by pg_dump version 14.12 (Homebrew)

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
-- Name: users; Type: TABLE; Schema: public; Owner: ssanyoq
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying,
    password character varying,
    name character varying,
    updated_at integer,
    created_at integer
);


ALTER TABLE public.users OWNER TO ssanyoq;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ssanyoq
--

COPY public.users (id, email, password, name, updated_at, created_at) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

