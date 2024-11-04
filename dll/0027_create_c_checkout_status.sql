create table IF NOT EXISTS wellbe_common.c_checkout_status (
  checkout_status_cd integer not null
  , language_cd integer not null
  , checkout_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_checkout_status_PKC primary key (checkout_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_checkout_status TO wellbe;