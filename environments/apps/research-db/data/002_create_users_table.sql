\c "research" dolphin;

create schema "research";

create table "research"."user" (
  id uuid primary key
  ,username text not null unique
  ,password text not null
  ,name text not null
);

create table "research"."article_overview" (
  id text primary key
  ,name text not null
  ,description text not null
  ,category text not null
  ,url text not null unique
);
