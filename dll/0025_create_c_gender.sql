create table IF NOT EXISTS wellbe_common.c_gender (
  gender_cd integer not null
  , language_cd integer not null
  , gender_name text not null
  , gender_abbreviation text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_gender_PKC primary key (gender_cd,language_cd)
) ;

GRANT ALL ON wellbe_common.c_gender TO wellbe;