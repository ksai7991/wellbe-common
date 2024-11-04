CREATE TABLE IF NOT EXISTS wellbe_common.c_country (
  country_cd integer not null
  , language_cd integer not null
  , country_name text not null
  , country_cd_iso varchar(2) not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_country_PKC primary key (country_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_country TO wellbe;