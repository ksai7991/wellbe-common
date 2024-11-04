create table IF NOT EXISTS wellbe_common.c_contact_status (
  contact_status_cd integer not null
  , language_cd integer not null
  , contact_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_contact_status_PKC primary key (contact_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_contact_status TO wellbe;