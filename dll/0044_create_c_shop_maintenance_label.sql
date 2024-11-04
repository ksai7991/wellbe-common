create table IF NOT EXISTS wellbe_common.c_shop_maintenance_label (
  shop_maintenance_label_cd integer not null
  , language_cd integer not null
  , shop_maintenance_label_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_shop_maintenance_label_PKC primary key (shop_maintenance_label_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_shop_maintenance_label TO wellbe;