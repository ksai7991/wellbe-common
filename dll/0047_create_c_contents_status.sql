create table IF NOT EXISTS wellbe_common.c_contents_status (
  contents_status_cd integer not null
  , language_cd integer not null
  , contents_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_contents_status_PKC primary key (contents_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_contents_status TO wellbe;