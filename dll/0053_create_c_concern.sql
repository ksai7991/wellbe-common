drop table IF EXISTS wellbe_common.c_concern;
create table IF NOT EXISTS wellbe_common.c_concern (
  concern_cd integer not null
  , language_cd integer not null
  , concern_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_concern_PKC primary key (concern_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_concern TO wellbe;