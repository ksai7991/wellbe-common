CREATE TABLE IF NOT EXISTS wellbe_common.c_recommend_label (
  recommend_label_cd integer not null
  , language_cd integer not null
  , recommend_label_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_recommend_label_PKC primary key (recommend_label_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_recommend_label TO wellbe;