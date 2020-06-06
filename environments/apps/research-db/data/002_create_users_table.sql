\c "research" dolphin;

create schema "research";

create table "research"."user" (
  id uuid primary key
  ,username text not null unique
  ,password text not null
  ,nickname text not null
);

create table "research"."editor" (
  id serial primary key
  ,name text not null
  ,icon bytea
);

create table "research"."article" (
  id text primary key
  ,title text not null
  ,description text not null
  ,editor int not null
  ,last_modified timestamp not null
  ,thumbnail text not null
  ,foreign key (editor) references research.editor(id)
);

create table "research"."article_content" (
  id text primary key
  ,body text not null
  ,foreign key (id) references research.article(id)
);

create table "research"."article_tag" (
  id serial primary key
  ,article_id text not null
  ,tag text not null
  ,foreign key (article_id) references research.article(id)
);

\c "history" dolphin;

create schema history;

create table "history"."user_history_log" (
  id serial primary key
  ,user_id text not null
  ,action_type int not null
  ,page_id text
)
