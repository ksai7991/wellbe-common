CREATE TABLE IF NOT EXISTS wellbe_common.numbering_master_definition (
 numbering_key varchar(100) PRIMARY KEY
 , initial_value bigint
 , current_value bigint
 , max_value bigint
 , fix_length bigint -- 0 means no fixed
 , create_datetime varchar(100) not null
 , create_func varchar(100) not null
 , update_datetime varchar(100)
 , update_func varchar(100)
);

GRANT ALL ON wellbe_common.numbering_definition TO wellbe;