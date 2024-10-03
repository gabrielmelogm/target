CREATE TABLE IF NOT EXISTS "goal_completions" (
	"id" text PRIMARY KEY NOT NULL,
	"goal_id" text NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE "goal_completions"
ADD CONSTRAINT "goal_completions_goal_id_goals_id_fk"
FOREIGN KEY ("goal_id") REFERENCES "public"."goals"("id");

