package sql

const (
	GetArea = `
  WITH PARAMS AS (
    SELECT
      CAST($1 AS INTEGER) AS LANGUAGE_CD
      , $2 AS COUNTRY_CD
      , $3 AS STATE_CD
      , $4 AS AREA_CD
  )
  SELECT
  A.area_cd
  ,A.area_name
  ,A.search_area_name_seo
  ,B.state_cd
  ,B.state_name
  ,B.state_cd_iso
  ,C.country_cd
  ,C.country_cd_iso
  ,C.country_name
 FROM wellbe_common.c_area AS A
 CROSS JOIN PARAMS AS P
 INNER JOIN wellbe_common.c_state AS B ON 1=1
   AND A.state_cd = B.state_cd
   AND B.language_cd = P.LANGUAGE_CD
 INNER JOIN wellbe_common.c_country AS C ON 1=1
   AND B.country_cd = C.country_cd
   AND C.language_cd = P.LANGUAGE_CD
 WHERE 1=1
   AND (P.AREA_CD = '' OR A.area_cd = CAST(P.AREA_CD AS INTEGER))
   AND (P.STATE_CD = '' OR B.state_cd = CAST(P.STATE_CD AS INTEGER))
   AND (P.COUNTRY_CD = '' OR C.country_cd = CAST(P.COUNTRY_CD AS INTEGER))
   AND A.language_cd = P.LANGUAGE_CD
`
)
