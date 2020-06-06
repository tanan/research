create user dolphin with password 'dolphin' superuser;
create database "research" owner dolphin template template0 encoding 'UTF8' lc_collate 'C' lc_ctype 'C';
grant all privileges on database "research" to dolphin;
create database "history" owner dolphin template template0 encoding 'UTF8' lc_collate 'C' lc_ctype 'C';
grant all privileges on database "history" to dolphin;

\c "research";

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

\c "history";

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;


ALTER SYSTEM SET max_connections = 5000;
ALTER SYSTEM RESET shared_buffers;

