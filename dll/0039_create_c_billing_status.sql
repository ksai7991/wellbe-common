create table IF NOT EXISTS wellbe_common.c_billing_status (
  billing_status_cd integer not null
  , language_cd integer not null
  , billing_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_billing_status_PKC primary key (billing_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_billing_status TO wellbe;