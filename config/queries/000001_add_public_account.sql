-- CreateAccount: 创建用户
-- table: accounts
-- name: CreateAccount :one
INSERT INTO "public"."accounts" (
  "phone",
  "email",
  "username",
  "nickname",
  "password",
  "password_salt",
  "created_by",
  "updated_by",
  "updated_at",
  "deleted_by",
  "deleted_at"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- UpdateAccount: 更新用户 (非空表达式实现动态传参更新数据)
-- table: accounts
--  NOTICE: COALESCE返回表达式中第一个非空表达式(例如: COALESCE(NULL, NULL, 1, 2) 返回 1) -  缺点就是更新时不能设置为null值
--  [nullable-parameters](https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters)
-- name: UpdateAccount :one
UPDATE "public"."accounts" 
SET
 "phone" = COALESCE(sqlc.narg('phone'), "phone"),
 "email" = COALESCE(sqlc.narg('email'), "email"),
 "username" = COALESCE(sqlc.narg('username'), "username"),
 "nickname" = COALESCE(sqlc.narg('nickname'), "nickname"),
 "password" = COALESCE(sqlc.narg('password'), "password"),
 "password_salt" = COALESCE(sqlc.narg('password_salt'), "password_salt"),
 "password_updated_at" = COALESCE(sqlc.narg('password_updated_at'), "password_updated_at"),
 "updated_by" = COALESCE(sqlc.narg('updated_by'), "updated_by"),
 "updated_at" = COALESCE(sqlc.narg('updated_at'), "updated_at"),
 "deleted_by" = COALESCE(sqlc.narg('deleted_by'), "deleted_by"),
 "deleted_at" = COALESCE(sqlc.narg('deleted_at'), "deleted_at")
WHERE "id" = sqlc.arg('id')
RETURNING *;
