CREATE TABLE IF NOT EXISTS wellbe_common.c_menu_label (
  menu_label_cd integer not null
  , language_cd integer not null
  , menu_label_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_menu_label_PKC primary key (menu_label_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_menu_label TO wellbe;