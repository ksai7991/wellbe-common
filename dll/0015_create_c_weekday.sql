CREATE TABLE IF NOT EXISTS wellbe_common.c_weekday (
  weekday_cd integer not null
  , language_cd integer not null
  , weekday_name text not null
  , weekday_abbreviation text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_weekday_PKC primary key (weekday_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_weekday TO wellbe;