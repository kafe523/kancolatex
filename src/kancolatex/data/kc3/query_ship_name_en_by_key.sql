SELECT
  "value"
FROM
  (
    SELECT
      "data" AS "content"
    FROM 
      t_kc3_translation_json_ships
    ORDER BY 
      "timestamp" DESC
    LIMIT 1
  ),
  json_each("content")
WHERE
  "key" = ?
;
