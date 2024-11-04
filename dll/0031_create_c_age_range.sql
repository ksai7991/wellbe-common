drop table IF EXISTS wellbe_common.c_age_range CASCADE;
create table IF NOT EXISTS wellbe_common.c_age_range (
  age_range_cd integer not null
  , language_cd integer not null
  , age_range_gender text not null
  , age_range_from integer
  , age_range_to integer
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_age_range_PKC primary key (age_range_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_age_range TO wellbe;