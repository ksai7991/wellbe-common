package sql

const (
	GetCurrencyForPayment = `
  WITH PARAM AS(
    SELECT
      CAST($1 AS INTEGER) AS LANGUAGE_CD
  )
  SELECT
    B.currency_cd
    ,B.language_cd
    ,B.currency_name
    ,B.currency_cd_iso
  FROM wellbe_common.currency_for_payment AS A
  CROSS JOIN PARAM AS P
  INNER JOIN wellbe_common.c_currency AS B ON 1=1
    AND A.currency_cd = B.currency_cd
    AND B.language_cd = P.LANGUAGE_CD
  
`
)
