CREATE TABLE IF NOT EXISTS wellbe_common.c_shop_equipment (
  shop_equipment_cd integer not null
  , language_cd integer not null
  , shop_equipment_name text not null
  , unit_name text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_shop_equipment_PKC primary key (shop_equipment_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_shop_equipment TO wellbe;