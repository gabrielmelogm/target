CREATE TABLE IF NOT EXISTS "goals" (
	"id" text PRIMARY KEY NOT NULL,
	"title" text NOT NULL,
	"desired_weekly_frequency" INTEGER NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
