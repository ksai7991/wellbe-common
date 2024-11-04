create table IF NOT EXISTS wellbe_common.c_payout_item_category (
  payout_item_category_cd integer not null
  , language_cd integer not null
  , payout_item_category_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_payout_item_category_PKC primary key (payout_item_category_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_payout_item_category TO wellbe;