DROP TABLE IF EXISTS wellbe_common.c_language cascade;
CREATE TABLE IF NOT EXISTS wellbe_common.c_language (
  language_cd integer not null
  , language_char_cd varchar(10) not null
  , language_name text not null
  , sort_number integer not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_language_PKC primary key (language_cd)
) ;
GRANT ALL ON wellbe_common.c_language TO wellbe;