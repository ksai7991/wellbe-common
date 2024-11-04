CREATE TABLE IF NOT EXISTS wellbe_common.c_tell_country (
  language_cd integer not null
  , tell_country_cd integer not null
  , country_name text not null
  , country_no text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_tell_country_PKC primary key (language_cd,tell_country_cd)
) ;
GRANT ALL ON wellbe_common.c_tell_country TO wellbe;