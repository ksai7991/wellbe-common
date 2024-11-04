package sql

const (
	GetAgeRange = `
  WITH PARAMS AS(
    SELECT
      CAST($1 AS INTEGER) AS LANGUAGE_CD
      ,CAST($2 AS INTEGER) AS AGE
  )
  SELECT
    A.age_range_cd
    ,A.language_cd
    ,A.age_range_gender
  FROM wellbe_common.c_age_range AS A
  CROSS JOIN PARAMS AS P
  WHERE 1=1
    AND A.language_cd = P.LANGUAGE_CD
    AND P.AGE >= A.age_range_from
    AND P.AGE <= A.age_range_to
`
)
