CREATE TABLE "builders" (
  "id"   SERIAL,
  "name" VARCHAR(255) NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("name")
);

CREATE TABLE "models" (
  "id"         SERIAL,
  "name"       VARCHAR(255) NOT NULL,
  "builder_id" INT REFERENCES builders (id),
  PRIMARY KEY ("id")
);

CREATE TABLE "companies" (
  "id"   SERIAL,
  "name" VARCHAR(255) NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "yachts" (
  "id"         SERIAL,
  "model_id"   INT REFERENCES models (id),
  "company_id" INT REFERENCES companies (id),
  "name"       VARCHAR(255)                                       NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp NOT NULL,
  "deleted_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "reservations" (
  "id"          SERIAL,
  "yacht_id"    INT REFERENCES yachts (id) ON DELETE CASCADE,
  "period_from" TIMESTAMP WITH TIME ZONE NOT NULL,
  "period_to"   TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id")
);

CREATE INDEX yachts_deleted_at_idx
  ON yachts (deleted_at);

CREATE INDEX models_name_idx
  ON models (name);


CREATE TABLE "gds" (
  "id"   SERIAL,
  "name" VARCHAR(255) NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("name")
);

CREATE TABLE "updaters" (
  "id"       SERIAL,
  "gds_id"   INT REFERENCES gds (id) ON DELETE CASCADE,
  "start_at" TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp NOT NULL,
  "stop_at"  TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  "status"   boolean DEFAULT FALSE                              NOT NULL
);

INSERT INTO "gds" (name) VALUES ('nausys');