--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

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
-- Name: db_germentop; Type: DATABASE; Schema: -; Owner: arikarniawan
--

CREATE DATABASE db_germentop WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE db_germentop OWNER TO arikarniawan;

\connect db_germentop

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
-- Name: user_type; Type: TYPE; Schema: public; Owner: arikarniawan
--

CREATE TYPE public.user_type AS ENUM (
    'buyer',
    'saller',
    'admin',
    'root'
);


ALTER TYPE public.user_type OWNER TO arikarniawan;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: mt_address; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_address (
    id integer,
    user_id character varying(50),
    receiver_name character varying(30),
    receiver_nohp character varying(15),
    address text,
    subdistrict_id integer,
    address_name character varying(25),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.mt_address OWNER TO arikarniawan;

--
-- Name: mt_category; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_category (
    id integer,
    name character varying(50),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.mt_category OWNER TO arikarniawan;

--
-- Name: mt_order_state; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_order_state (
    id integer,
    name character varying(50),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    next_state json,
    final_state character varying(3)
);


ALTER TABLE public.mt_order_state OWNER TO arikarniawan;

--
-- Name: mt_product; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_product (
    id character varying,
    name character varying(60),
    category_id integer,
    supplier_id character varying(50),
    description json,
    variant json,
    gallery character varying,
    price numeric(12,2),
    disc_price numeric(9,2),
    disc_percent numeric(4,2),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.mt_product OWNER TO arikarniawan;

--
-- Name: mt_shipper; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_shipper (
    id integer NOT NULL,
    shipper_method character varying(25),
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.mt_shipper OWNER TO arikarniawan;

--
-- Name: mt_user; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.mt_user (
    id character varying(50) NOT NULL,
    name character varying(25),
    email character varying(50),
    password character varying(100),
    no_hp character varying(13),
    address character varying(250),
    subdistrict_id character varying(13),
    zip_code character varying(13),
    created_at date,
    updated_at date,
    deleted_at date,
    photo character varying(100),
    user_type public.user_type
);


ALTER TABLE public.mt_user OWNER TO arikarniawan;

--
-- Name: tr_payment_method; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.tr_payment_method (
    id integer,
    name integer,
    provider integer,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.tr_payment_method OWNER TO arikarniawan;

--
-- Name: tr_shoping_cart; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.tr_shoping_cart (
    id character varying(50),
    user_id character varying(50),
    items json,
    total money,
    created_at timestamp without time zone DEFAULT now(),
    meta_data text,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    status integer
);


ALTER TABLE public.tr_shopping_cart OWNER TO arikarniawan;

--
-- Name: tr_shoping_checkout; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.tr_shoping_checkout (
    id character varying(50) NOT NULL,
    cart_id character varying(50),
    user_id character varying(50),
    amount money,
    disc money,
    total money,
    promo_code integer,
    items json,
    address_id character varying(5),
    meta_data json,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    payment_method_id character varying(50),
    state_id integer
);


ALTER TABLE public.tr_shoping_checkout OWNER TO arikarniawan;

--
-- Name: tr_shopping_history; Type: TABLE; Schema: public; Owner: arikarniawan
--

CREATE TABLE public.tr_shopping_history (
    id character varying(50) NOT NULL,
    checkout_id character varying(5),
    state_id integer,
    meta_data json,
    shipper_id integer,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.tr_shopping_history OWNER TO arikarniawan;

--
-- Data for Name: mt_address; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Data for Name: mt_category; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--

INSERT INTO public.mt_category (id, name, created_at, updated_at, deleted_at) VALUES (1, 'Nasi Padang', '2021-02-04 22:32:06', NULL, NULL);


--
-- Data for Name: mt_order_state; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--

INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (11, 'Order Cancel', '2021-02-04 14:32:28.269112', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (14, 'Order Shipping Not Found', '2021-02-04 14:36:29.209815', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (3, 'Payment Canceled', '2021-02-04 14:20:33.742383', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (13, 'Order Rejected by Seller', '2021-02-04 14:36:29.209815', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (4, 'Payment Expired', '2021-02-04 14:20:47.231357', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (1, 'Open Payment', '2021-02-04 14:19:35.689414', NULL, NULL, '[2,3,4]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (9, 'Order On Way', '2021-02-04 14:28:08.814567', NULL, NULL, '[10]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (7, 'Order Ready', '2021-02-04 14:26:55.07174', NULL, NULL, '[8]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (2, 'Payment Completed', '2021-02-04 14:20:09.961286', NULL, NULL, '[5,11,13]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (5, 'Order Confirm', '2021-02-04 14:21:38.37357', NULL, NULL, '[6,11]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (6, 'Order Prepare', '2021-02-04 14:25:55.635975', NULL, NULL, '[7]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (8, 'Order Pickup', '2021-02-04 14:27:20.123392', NULL, NULL, '[9]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (10, 'Order Completed', '2021-02-04 14:32:28.269112', NULL, NULL, '[12]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (15, 'Order Return Process', '2021-02-04 15:00:42.355784', NULL, NULL, '[16]', 'no');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (16, 'Order Return Complated', '2021-02-04 15:00:42.355784', NULL, NULL, NULL, 'yes');
INSERT INTO public.mt_order_state (id, name, created_at, updated_at, deleted_at, next_state, final_state) VALUES (12, 'Order Return Request', '2021-02-04 14:32:28.269112', NULL, NULL, '[15]', 'no');


--
-- Data for Name: mt_product; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--

INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('965f4b67-055b-49b7-a5e5-71b5bc821173', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:56:18.267506', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('c7410763-a1e2-48dc-a281-b638dd35d547', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:45:43.348765', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('31ac8a01-e44f-4d2d-aeb0-9ad67ab744f2', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-02-02 00:23:20.722197', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('711c6211-0708-408f-ae5e-6186dd0068b6', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 00:27:13.316486', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('93f82efa-5406-4e57-b382-973cba12ce10', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-29 21:12:19.627397', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('cec7588e-f9d4-49e7-9672-bfdb9a823e1e', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:46:53.91179', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('2732fdaf-342c-451d-882b-65e3e5022313', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 0.00, 0.00, 0.00, '2021-01-29 21:05:34.725476', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('3ac85dfd-297b-41c7-aa97-a61f5ea5d8b7', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:52:56.759042', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('4f63f402-dd47-4ed3-8930-2be80499d244', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 0.00, 0.00, '2021-01-29 21:07:38.130295', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('1765f812-fb2f-49da-90c5-c9b36a5ffdce', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 21:23:18.733957', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('fdd6ec7e-a4cd-4407-a9fc-dbb5e08c1073', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 21:42:36.017367', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('6c218dd5-b900-4d15-8c6b-a1e16ca7c307', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:50:29.142656', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('9f184f60-6461-42df-a52e-7b5bf81654fc', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('791325c0-e4fc-49e3-99b1-f890280ac952', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac3854e0', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-29 21:12:19.348595', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('02458526-070c-40c9-b718-ba262bc11328', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 19:59:44.761826', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('8e9c7ef4-d310-45c2-ba37-e293d37d37c9', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:42:04.392635', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('28f6d334-2cae-4852-a523-9146ccd7bd8e', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:51:14.762681', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('6ab8a09e-fc21-4456-991a-c7c3779f184f', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:42:41.694581', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('7249e14e-2d2e-4799-a6b6-08efd87995cd', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:38:07.148363', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('565ec501-effd-4182-85ae-8a924daf215d', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-29 21:12:20.755376', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('b429c18a-390a-491b-82a9-95dfe8869cb9', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 0.00, 0.00, 0.00, '2021-01-29 21:04:25.678286', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('b609bad4-47c3-417c-b955-406636ac4e56', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:43:31.344539', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('da115187-c480-4b33-aed9-beed85c7850a', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:38:22.8715', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('1db0fe73-c2c5-4586-be8e-a5a0b3627380', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:01:10.272868', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('6c2b1be4-869b-4f49-b3fa-163702294369', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('7dd2260e-7f24-4c88-8cd0-30700d295b71', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:46:47.414561', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('bd326db6-f811-470a-a8f6-08c5503cda60', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-31 10:54:27.632588', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('b7c56422-f024-423f-81d0-2a246191430b', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-29 21:12:20.452111', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('0872fea5-fe3c-4682-bbde-9e97927d3dc0', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:54:52.031256', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('08828239-7c48-44f4-989e-94885af3ee71', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-02-02 00:19:56.258373', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('af88fd34-e616-4303-98f0-de2e5bf0e0b9', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 21:48:34.148078', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('ea07e694-7ddc-4b47-bc6c-89e9596a40b9', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 21:21:14.566569', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('dd8df9e8-dd87-4780-80ad-4b6632d8c94d', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-02-02 00:25:13.710347', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('40926816-046e-4aec-a34a-de76795b9ea9', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:43:47.332069', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('2f4bde35-2fe6-40c7-a318-0ed87e3fe99e', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:39:12.098163', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('30fb7ac0-718a-4f72-b9db-97488c4dae18', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('3c3cae5c-a4ee-4790-ac88-788819a4efff', 'AYE AYE', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 23000.00, 20000.00, 13.00, '2021-01-29 21:10:31.171496', NULL, NULL);
INSERT INTO public.mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent, created_at, updated_at, deleted_at) VALUES ('21c8efdf-14d0-40c8-9b48-0d7be9023041', 'Ahai Ahai', 1, '6993fad7-dc54-40ba-bca2-aaa2ac38', '[{"title":"Komposisi","content":"Minyak Sayur, Garam, bumbu rempah"}]', '[{"title":"Toping","type":"CheckList","options":[{"price":"5000"}]}]', '["a.jpg","b.jpg"]', 5000000.00, 20000.00, 13.00, '2021-02-02 20:17:15.554262', NULL, NULL);


--
-- Data for Name: mt_shipper; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Data for Name: mt_user; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--

INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('a1245', 'Ari Karniawan', 'arikarniawan@gmail.com', '123456', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'buyer');
INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('6993fad7-dc54-40ba-bca2-aaa2ac3854e0', NULL, 'abc@gmail.com', '$2a$04$4GNrI8A7eoru2q.7mRbBXOEAH7bN0bYbEV4ztuEDwuDDJQChYblh6', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'buyer');
INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('7cbdef1d-a64b-45f3-bf8d-43f11dd9b5bd', NULL, 'a@gmail.com', '$2a$04$7jD2i0PsTwI30uOmJtNIMOcyqKFfZMjVBDmIbtF8J.uCNxzS8DB7m', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'buyer');
INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('f61ecb42-1b64-4263-ba04-41c0f31dcdd1', NULL, 'abcterewrw@gmail.com', '$2a$04$MRVb8iBAdFF3c.FcUaxdPOJbpZyNcDE81dmvAEIiEo49z7PM1BlX2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'buyer');
INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('18e782c3-b800-40ea-8496-13d0d40c0df9', NULL, 'ab@gmail.com', '$2a$04$hI8RSasXMD/CN59TPxqqLu8w9RxXmfRQ0Q.qrj/hghYuf6oyEcDae', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'buyer');
INSERT INTO public.mt_user (id, name, email, password, no_hp, address, subdistrict_id, zip_code, created_at, updated_at, deleted_at, photo, user_type) VALUES ('a8311203-8b20-44ed-8f1f-c4b44d9e1838', NULL, 'ari.karniawan@onlineresto.id', '$2a$04$udUVC89AqVEKkIgtGWt9FeCw2/bCpPlZedKKs.gy2bdoX7BxIYNlu', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);


--
-- Data for Name: tr_payment_method; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Data for Name: tr_shoping_cart; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Data for Name: tr_shoping_checkout; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Data for Name: tr_shopping_history; Type: TABLE DATA; Schema: public; Owner: arikarniawan
--



--
-- Name: mt_address mt_address_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_address
    ADD CONSTRAINT mt_address_pk UNIQUE (id);


--
-- Name: mt_category mt_category_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_category
    ADD CONSTRAINT mt_category_pk UNIQUE (id);


--
-- Name: mt_product mt_product_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_product
    ADD CONSTRAINT mt_product_pk UNIQUE (id);


--
-- Name: mt_shipper mt_shipper_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_shipper
    ADD CONSTRAINT mt_shipper_pk PRIMARY KEY (id);


--
-- Name: mt_user table_name_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_user
    ADD CONSTRAINT table_name_pk PRIMARY KEY (id);


--
-- Name: tr_shoping_checkout tr_shoping_checkout_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.tr_shoping_checkout
    ADD CONSTRAINT tr_shoping_checkout_pk PRIMARY KEY (id);


--
-- Name: tr_shopping_history tr_shopping_history_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.tr_shopping_history
    ADD CONSTRAINT tr_shopping_history_pk PRIMARY KEY (id);


--
-- Name: mt_order_state tr_state_pk; Type: CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_order_state
    ADD CONSTRAINT tr_state_pk UNIQUE (id);


--
-- Name: table_name_id_users_uindex; Type: INDEX; Schema: public; Owner: arikarniawan
--

CREATE UNIQUE INDEX table_name_id_users_uindex ON public.mt_user USING btree (id);


--
-- Name: tr_shoping_checkout_state_id_cart_id_address_id_index; Type: INDEX; Schema: public; Owner: arikarniawan
--

CREATE INDEX tr_shoping_checkout_state_id_cart_id_address_id_index ON public.tr_shoping_checkout USING btree (state_id, cart_id, address_id);


--
-- Name: tr_shopping_history_courier_id_checkout_id_state_id_index; Type: INDEX; Schema: public; Owner: arikarniawan
--

CREATE INDEX tr_shopping_history_courier_id_checkout_id_state_id_index ON public.tr_shopping_history USING btree (shipper_id, checkout_id, state_id);


--
-- Name: mt_product mt_product_mt_category_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.mt_product
    ADD CONSTRAINT mt_product_mt_category_id_fk FOREIGN KEY (category_id) REFERENCES public.mt_category(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: tr_shopping_history tr_shopping_history_mt_shipper_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: arikarniawan
--

ALTER TABLE ONLY public.tr_shopping_history
    ADD CONSTRAINT tr_shopping_history_mt_shipper_id_fk FOREIGN KEY (shipper_id) REFERENCES public.mt_shipper(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

