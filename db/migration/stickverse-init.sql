-- -------------------------------------------------------------
-- TablePlus 4.8.0(432)
--
-- https://tableplus.com/
--
-- Database: stickverse
-- Generation Time: 2022-07-12 20:23:40.2220
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS alliance_members_id_seq;

-- Table Definition

CREATE TABLE "public"."alliance_members" (
    "id" int8 NOT NULL DEFAULT nextval('alliance_members_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "alliance_id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "membership_type_id" int8 NOT NULL,
    "ended_at" timestamptz,
    PRIMARY KEY ("id"),
    CONSTRAINT "alliance_members_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE,
    CONSTRAINT "alliance_members_alliance_id_fkey" FOREIGN KEY ("alliance_id") REFERENCES "public"."alliances"("id") ON DELETE CASCADE,
    CONSTRAINT "alliance_members_membership_type_id_fkey" FOREIGN KEY ("membership_type_id") REFERENCES "public"."membership_types"("id") ON DELETE CASCADE
);


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS alliances_id_seq;

-- Table Definition
CREATE TABLE "public"."alliances" (
    "id" int8 NOT NULL DEFAULT nextval('alliances_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    "date_disbanded" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS allowed_actions_id_seq;

-- Table Definition
CREATE TABLE "public"."allowed_actions" (
    "id" int8 NOT NULL DEFAULT nextval('allowed_actions_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "membership_action_id" int8,
    "membership_type_id" int8,
    PRIMARY KEY ("id"),
    CONSTRAINT "allowed_actions_membership_action_id_fkey" FOREIGN KEY ("membership_action_id") REFERENCES "public"."membership_actions"("id") ON DELETE CASCADE,
    CONSTRAINT "allowed_actions_membership_type_id_fkey" FOREIGN KEY ("membership_type_id") REFERENCES "public"."membership_types"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS characteristics_id_seq;

-- Table Definition
CREATE TABLE "public"."characteristics" (
    "id" int8 NOT NULL DEFAULT nextval('characteristics_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS grid_types_id_seq;

-- Table Definition
CREATE TABLE "public"."grid_types" (
    "id" int8 NOT NULL DEFAULT nextval('grid_types_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "description" jsonb,
    "name_local" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS grids_id_seq;

-- Table Definition
CREATE TABLE "public"."grids" (
    "id" int8 NOT NULL DEFAULT nextval('grids_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "x" int8,
    "y" int8,
    "grid_type_id" int8,
    "capacity" int8 DEFAULT 10,
    "occupied" int8 DEFAULT 0,
    PRIMARY KEY ("id"),
    CONSTRAINT "grids_grid_type_id_fkey" FOREIGN KEY ("grid_type_id") REFERENCES "public"."grid_types"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS group_movement_units_id_seq;

-- Table Definition
CREATE TABLE "public"."group_movement_units" (
    "id" int8 NOT NULL DEFAULT nextval('group_movement_units_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "group_movement_id" int8,
    "unit_id" int8,
    "amount" int8,
    PRIMARY KEY ("id"),
    CONSTRAINT "group_movement_units_group_movement_id_fkey" FOREIGN KEY ("group_movement_id") REFERENCES "public"."group_movements"("id") ON DELETE CASCADE,
    CONSTRAINT "group_movement_units_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS group_movements_id_seq;

-- Table Definition
CREATE TABLE "public"."group_movements" (
    "id" int8 NOT NULL DEFAULT nextval('group_movements_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" int8 NOT NULL,
    "movement_type_id" int8 NOT NULL,
    "location_from_id" int8 NOT NULL,
    "location_to_id" int8 NOT NULL,
    "arrival_date" timestamptz NOT NULL,
    "return_date" timestamptz NOT NULL,
    "waiting_time" int8 NOT NULL,
    "comment" text NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "group_movements_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE,
    CONSTRAINT "group_movements_movement_type_id_fkey" FOREIGN KEY ("movement_type_id") REFERENCES "public"."movement_types"("id") ON DELETE CASCADE,
    CONSTRAINT "group_movements_location_from_id_fkey" FOREIGN KEY ("location_from_id") REFERENCES "public"."locations"("id") ON DELETE CASCADE,
    CONSTRAINT "group_movements_location_to_id_fkey" FOREIGN KEY ("location_to_id") REFERENCES "public"."locations"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS location_resources_id_seq;

-- Table Definition
CREATE TABLE "public"."location_resources" (
    "id" int8 NOT NULL DEFAULT nextval('location_resources_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "resource_id" int8 NOT NULL,
    "location_id" int8 NOT NULL,
    "quantity" int8 NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "location_resources_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "public"."locations"("id") ON DELETE CASCADE,
    CONSTRAINT "location_resources_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "public"."resources"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS location_types_id_seq;

-- Table Definition
CREATE TABLE "public"."location_types" (
    "id" int8 NOT NULL DEFAULT nextval('location_types_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS location_units_id_seq;

-- Table Definition
CREATE TABLE "public"."location_units" (
    "id" int8 NOT NULL DEFAULT nextval('location_units_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "location_id" int8 NOT NULL,
    "unit_id" int8 NOT NULL,
    "amount" int8 NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "location_units_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "public"."locations"("id") ON DELETE CASCADE,
    CONSTRAINT "location_units_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS locations_id_seq;

-- Table Definition
CREATE TABLE "public"."locations" (
    "id" int8 NOT NULL DEFAULT nextval('locations_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "location_type_id" int8 NOT NULL,
    "level" int8 NOT NULL,
    "user_id" int8,
    "grid_id" int8,
    "grid_index" int8 NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "locations_grid_id_fkey" FOREIGN KEY ("grid_id") REFERENCES "public"."grids"("id") ON DELETE CASCADE,
    CONSTRAINT "locations_location_type_id_fkey" FOREIGN KEY ("location_type_id") REFERENCES "public"."location_types"("id") ON DELETE CASCADE,
    CONSTRAINT "locations_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS login_histories_id_seq;

-- Table Definition
CREATE TABLE "public"."login_histories" (
    "id" int8 NOT NULL DEFAULT nextval('login_histories_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" int8 NOT NULL,
    "ip" text NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "login_histories_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS membership_actions_id_seq;

-- Table Definition
CREATE TABLE "public"."membership_actions" (
    "id" int8 NOT NULL DEFAULT nextval('membership_actions_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS membership_histories_id_seq;

-- Table Definition
CREATE TABLE "public"."membership_histories" (
    "id" int8 NOT NULL DEFAULT nextval('membership_histories_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "alliance_member_id" int8 NOT NULL,
    "membership_type_id" int8 NOT NULL,
    "date_from" timestamptz NOT NULL,
    "ended_at" timestamptz,
    PRIMARY KEY ("id"),
    CONSTRAINT "membership_histories_alliance_member_id_fkey" FOREIGN KEY ("alliance_member_id") REFERENCES "public"."alliance_members"("id") ON DELETE CASCADE,
    CONSTRAINT "membership_histories_membership_type_id_fkey" FOREIGN KEY ("membership_type_id") REFERENCES "public"."membership_types"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS membership_types_id_seq;

-- Table Definition
CREATE TABLE "public"."membership_types" (
    "id" int8 NOT NULL DEFAULT nextval('membership_types_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS movement_types_id_seq;

-- Table Definition
CREATE TABLE "public"."movement_types" (
    "id" int8 NOT NULL DEFAULT nextval('movement_types_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "allows_wait" bool NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS prerequisites_id_seq;

-- Table Definition
CREATE TABLE "public"."prerequisites" (
    "id" int8 NOT NULL DEFAULT nextval('prerequisites_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "type" varchar(255) NOT NULL,
    "prerequisite_formula" varchar(255),
    "structure_id" int8,
    "research_id" int8,
    "unit_id" int8,
    "required_structure_id" int8,
    "required_research_id" int8,
    "required_level" int8,
    PRIMARY KEY ("id"),
    CONSTRAINT "prerequisites_structure_id_fkey" FOREIGN KEY ("structure_id") REFERENCES "public"."structures"("id") ON DELETE CASCADE,
    CONSTRAINT "prerequisites_research_id_fkey" FOREIGN KEY ("research_id") REFERENCES "public"."researches"("id") ON DELETE CASCADE,
    CONSTRAINT "prerequisites_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units"("id") ON DELETE CASCADE,
    CONSTRAINT "prerequisites_required_structure_id_fkey" FOREIGN KEY ("required_structure_id") REFERENCES "public"."structures"("id") ON DELETE CASCADE,
    CONSTRAINT "prerequisites_required_research_id_fkey" FOREIGN KEY ("required_research_id") REFERENCES "public"."researches"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS research_levels_id_seq;

-- Table Definition
CREATE TABLE "public"."research_levels" (
    "id" int8 NOT NULL DEFAULT nextval('research_levels_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" int8 NOT NULL,
    "research_id" int8 NOT NULL,
    "level" int8 NOT NULL,
    "upgrade_ongoing" bool NOT NULL,
    "upgrade_started_at" timestamptz NOT NULL,
    "upgrade_ended_at" timestamptz NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "research_levels_research_id_fkey" FOREIGN KEY ("research_id") REFERENCES "public"."researches"("id") ON DELETE CASCADE,
    CONSTRAINT "research_levels_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS research_resources_id_seq;

-- Table Definition
CREATE TABLE "public"."research_resources" (
    "id" int8 NOT NULL DEFAULT nextval('research_resources_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "upgrade_formula" text,
    "research_id" int8 NOT NULL,
    "resource_id" int8 NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "research_resources_research_id_fkey" FOREIGN KEY ("research_id") REFERENCES "public"."researches"("id") ON DELETE CASCADE,
    CONSTRAINT "research_resources_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "public"."resources"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS researches_id_seq;

-- Table Definition
CREATE TABLE "public"."researches" (
    "id" int8 NOT NULL DEFAULT nextval('researches_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    "upgrade_time_formula" text NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS resources_id_seq;

-- Table Definition
CREATE TABLE "public"."resources" (
    "id" int8 NOT NULL DEFAULT nextval('resources_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS structure_builts_id_seq;

-- Table Definition
CREATE TABLE "public"."structure_builts" (
    "id" int8 NOT NULL DEFAULT nextval('structure_builts_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "location_id" int8 NOT NULL,
    "structure_id" int8 NOT NULL,
    "upgrade_ongoing" bool NOT NULL,
    "upgrade_started_at" timestamptz NOT NULL,
    "upgrade_ended_at" timestamptz NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "structure_builts_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "public"."locations"("id") ON DELETE CASCADE,
    CONSTRAINT "structure_builts_structure_id_fkey" FOREIGN KEY ("structure_id") REFERENCES "public"."structures"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS structure_resources_id_seq;

-- Table Definition
CREATE TABLE "public"."structure_resources" (
    "id" int8 NOT NULL DEFAULT nextval('structure_resources_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "structure_id" int8,
    "resource_id" int8,
    "upgrade_formula" text,
    "production_formula" text,
    PRIMARY KEY ("id"),
    CONSTRAINT "structure_resources_resource_id_fkey" FOREIGN KEY ("resource_id") REFERENCES "public"."resources"("id") ON DELETE CASCADE,
    CONSTRAINT "structure_resources_structure_id_fkey" FOREIGN KEY ("structure_id") REFERENCES "public"."structures"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS structures_id_seq;

-- Table Definition
CREATE TABLE "public"."structures" (
    "id" int8 NOT NULL DEFAULT nextval('structures_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    "upgrade_time_formula" text NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS unit_characteristics_id_seq;

-- Table Definition
CREATE TABLE "public"."unit_characteristics" (
    "id" int8 NOT NULL DEFAULT nextval('unit_characteristics_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "unit_id" int8 NOT NULL,
    "characteristic_id" int8 NOT NULL,
    "amount" numeric NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "unit_characteristics_characteristic_id_fkey" FOREIGN KEY ("characteristic_id") REFERENCES "public"."characteristics"("id") ON DELETE CASCADE,
    CONSTRAINT "unit_characteristics_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS unit_costs_id_seq;

-- Table Definition
CREATE TABLE "public"."unit_costs" (
    "id" int8 NOT NULL DEFAULT nextval('unit_costs_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "unit_id" int8 NOT NULL,
    "cost" numeric NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "unit_costs_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units"("id") ON DELETE CASCADE
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS units_id_seq;

-- Table Definition
CREATE TABLE "public"."units" (
    "id" int8 NOT NULL DEFAULT nextval('units_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "name_local" jsonb,
    "description" jsonb,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "username" text,
    "email" text,
    "hashed_password" text,
    "role" text,
    PRIMARY KEY ("id")
);

