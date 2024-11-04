create table IF NOT EXISTS wellbe_common.c_type_of_contact (
  type_of_contact_cd integer not null
  , language_cd integer not null
  , type_of_contact_name text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_type_of_contact_PKC primary key (type_of_contact_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_type_of_contact TO wellbe;