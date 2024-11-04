drop table IF EXISTS wellbe_common.c_area;
create table IF NOT EXISTS wellbe_common.c_area (
  language_cd integer not null
  , area_cd integer not null
  , state_cd integer not null
  , area_name text not null
  , search_area_name_seo text not null
  , west_longitude numeric(20,16) not null
  , east_longitude numeric(20,16) not null
  , north_latitude numeric(20,16) not null
  , south_latitude numeric(20,16) not null
  , summary_area_flg boolean not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_area_PKC primary key (language_cd,area_cd)
) ;
GRANT ALL ON wellbe_common.c_area TO wellbe;