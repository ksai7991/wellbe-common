create table IF NOT EXISTS wellbe_common.c_coupon_target_attr (
  coupon_target_attr_cd integer not null
  , language_cd integer not null
  , coupon_target_attr_name text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_coupon_target_attr_PKC primary key (coupon_target_attr_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_coupon_target_attr TO wellbe;