create table IF NOT EXISTS wellbe_common.c_invoice_status (
  invoice_status_cd integer not null
  , language_cd integer not null
  , invoice_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_invoice_status_PKC primary key (invoice_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_invoice_status TO wellbe;