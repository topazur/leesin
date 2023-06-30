CREATE TABLE IF NOT EXISTS "public"."accounts" (
  "id" BIGSERIAL PRIMARY KEY,

  -- libphonenumber-js
  "phone" VARCHAR UNIQUE,
  "email" VARCHAR UNIQUE,
  "username" VARCHAR UNIQUE,
  "nickname" VARCHAR UNIQUE,

  -- 密码
  "password" VARCHAR,
  "password_salt" VARCHAR,
  "password_updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),

  -- 创建、更新、删除
  "created_by" BIGINT,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_by" BIGINT,
  "updated_at" TIMESTAMPTZ,
  "deleted_by" BIGINT,
  "deleted_at" TIMESTAMPTZ
);

-- 注释
COMMENT ON TABLE "public"."accounts" IS '用户表';
COMMENT ON COLUMN "public"."accounts"."id" IS '主键:用户ID';

CREATE TABLE IF NOT EXISTS "public"."sessions" (
  "id" UUID PRIMARY KEY,

  "uid" BIGINT NOT NULL,
  "client_ip" VARCHAR NOT NULL,
  "user_agent" VARCHAR NOT NULL,
  "is_blocked" BOOLEAN NOT NULL DEFAULT false,

  "refresh_token" VARCHAR NOT NULL,
  "expires_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

-- 注释
COMMENT ON TABLE "public"."sessions" IS '用户登录session记录表';
COMMENT ON COLUMN "public"."sessions"."id" IS '主键:session_id';
COMMENT ON COLUMN "public"."sessions"."uid" IS '外键:public.accounts.id';

ALTER TABLE "public"."sessions" ADD FOREIGN KEY ("uid") REFERENCES "public"."accounts" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
