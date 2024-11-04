create table IF NOT EXISTS wellbe_common.c_treatment_time_range (
  treatment_time_cd integer not null
  , language_cd integer not null
  , treatment_time_name text not null
  , min_time integer not null
  , max_time integer
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_treatment_time_range_PKC primary key (treatment_time_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_treatment_time_range TO wellbe;