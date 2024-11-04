create table IF NOT EXISTS wellbe_common.c_shop_contract_plan_item (
  shop_contract_plan_item_cd integer not null
  , language_cd integer not null
  , shop_contract_plan_name text not null
  , unit text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_shop_contract_plan_item_PKC primary key (shop_contract_plan_item_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_shop_contract_plan_item TO wellbe;