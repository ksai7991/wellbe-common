create table IF NOT EXISTS wellbe_common.c_checkout_timing (
  checkout_timing_cd integer not null
  , language_cd integer not null
  , checkout_timing_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_checkout_timing_PKC primary key (checkout_timing_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_checkout_timing TO wellbe;