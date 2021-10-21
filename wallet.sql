-- Adminer 4.8.1 PostgreSQL 13.3 (Debian 13.3-1.pgdg100+1) dump

DROP TABLE IF EXISTS "payments";
CREATE TABLE "public"."payments" (
    "id" character varying(60) NOT NULL,
    "amount" integer,
    "type" character varying(10) DEFAULT 'outcome',
    "wallet_id" character varying(60),
    "created_at" timestamp,
    "updated_at" timestamp,
    CONSTRAINT "payments_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "id" character varying(60) NOT NULL,
    "name" character varying(60),
    "surname" character varying(60),
    "username" character varying(60),
    "email" character varying(60),
    "password" character varying(60),
    "created_at" timestamp,
    "updated_at" timestamp,
    CONSTRAINT "users_email_key" UNIQUE ("email"),
    CONSTRAINT "users_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "users_username_key" UNIQUE ("username")
) WITH (oids = false);

INSERT INTO "users" ("id", "name", "surname", "username", "email", "password", "created_at", "updated_at") VALUES
('23f25f5c-2713-411a-82de-6f0f69e88232',	'Abrorbek',	'Ubaydullayev',	'aborbek_ubaydullayev',	'qwe@gmail.com',	NULL,	'2021-10-21 07:23:08.736572',	'2021-10-21 07:23:08.736572'),
('92b42359-d0d5-409d-a62f-fdad7a3f1a99',	'Mirza',	'Mirzayev',	'mirza',	'mirza@gmail.com',	NULL,	'2021-10-21 07:23:08.736572',	'2021-10-21 07:23:08.736572'),
('5e8d93be-6f87-4608-adc5-dcaa608be7e6',	'Bekzod',	'Hoshimov',	'bekzod',	'bekzod@gmail.com',	NULL,	'2021-10-21 07:23:08.736572',	'2021-10-21 07:23:08.736572'),
('a352140f-f439-4c97-89b9-ea6da4eae75b',	'Shaxa',	'Dalimov',	'shaxa',	'shaza@gmail.com',	NULL,	'2021-10-21 07:23:08.736572',	'2021-10-21 07:23:08.736572');

DROP TABLE IF EXISTS "wallets";
CREATE TABLE "public"."wallets" (
    "id" character varying(60) NOT NULL,
    "type" character varying(20) DEFAULT 'unidentified',
    "amount" integer,
    "user_id" character varying(60),
    "created_at" timestamp,
    "updated_at" timestamp,
    CONSTRAINT "wallets_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "wallets" ("id", "type", "amount", "user_id", "created_at", "updated_at") VALUES
('d8481cc1-2715-4ec7-b214-e51efe5ed574',	'unidentified',	10000,	'92b42359-d0d5-409d-a62f-fdad7a3f1a99',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('8e0eeffa-1731-4d69-8ab2-2f327b8d7b3a',	'unidentified',	10000,	'92b42359-d0d5-409d-a62f-fdad7a3f1a99',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('a179c140-290b-4d89-baa9-a725487c3b91',	'unidentified',	100000,	'5e8d93be-6f87-4608-adc5-dcaa608be7e6',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('30f7cd43-4c18-4825-8140-32fb2682bf8f',	'unidentified',	100000,	'5e8d93be-6f87-4608-adc5-dcaa608be7e6',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('c052cd2d-12ce-4409-bb81-8552445403cd',	'unidentified',	10000,	'23f25f5c-2713-411a-82de-6f0f69e88232',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('4d7116ed-79f1-4808-ac4b-b3e2441b18a6',	'unidentified',	10000,	'23f25f5c-2713-411a-82de-6f0f69e88232',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('bf2700a7-7172-47ed-8daf-e8f5df9c6800',	'unidentified',	10000,	'a352140f-f439-4c97-89b9-ea6da4eae75b',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347'),
('49c7e80f-2e72-4395-94e6-cc20db5f9813',	'unidentified',	100000,	'a352140f-f439-4c97-89b9-ea6da4eae75b',	'2021-10-21 07:23:08.814347',	'2021-10-21 07:23:08.814347');

ALTER TABLE ONLY "public"."payments" ADD CONSTRAINT "payments_wallet_id_fkey" FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE NOT DEFERRABLE;

ALTER TABLE ONLY "public"."wallets" ADD CONSTRAINT "wallets_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE NOT DEFERRABLE;

-- 2021-10-21 02:45:02.379835-05
