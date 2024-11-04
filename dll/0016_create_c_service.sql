CREATE TABLE IF NOT EXISTS wellbe_common.c_service (
  service_cd integer not null
  , language_cd integer not null
  , service_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_service_PKC primary key (service_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_service TO wellbe;