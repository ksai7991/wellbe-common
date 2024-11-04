drop table IF EXISTS wellbe_common.c_contents_label;
create table IF NOT EXISTS wellbe_common.c_contents_label (
  contents_label_cd integer not null
  , language_cd integer not null
  , contents_category_cd integer not null
  , contents_label_name text not null
  , contents_label_url text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_contents_label_PKC primary key (contents_label_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_contents_label TO wellbe;