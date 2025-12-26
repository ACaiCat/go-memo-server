/*
 Navicat Premium Dump SQL

 Source Server         : test-pg
 Source Server Type    : PostgreSQL
 Source Server Version : 180000 (180000)
 Source Host           : localhost:15664
 Source Catalog        : memo
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 180000 (180000)
 File Encoding         : 65001

 Date: 26/12/2025 16:00:17
*/


-- ----------------------------
-- Table structure for memo
-- ----------------------------
CREATE TABLE "public"."memo" (
  "id" int8 NOT NULL DEFAULT nextval('memo_id_seq'::regclass),
  "created_at" timestamptz(6),
  "user_id" int8,
  "title" text COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "status" int8,
  "start_time" timestamptz(6),
  "end_time" timestamptz(6)
)
;

-- ----------------------------
-- Indexes structure for table memo
-- ----------------------------
CREATE INDEX "idx_memo_title" ON "public"."memo" USING btree (
  "title" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_memo_user_id" ON "public"."memo" USING btree (
  "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table memo
-- ----------------------------
ALTER TABLE "public"."memo" ADD CONSTRAINT "memo_pkey" PRIMARY KEY ("id");
