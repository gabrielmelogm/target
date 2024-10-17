-- name: GetPendingGoals :many
WITH goals_created_up_to_week AS (
    SELECT id, title, desired_weekly_frequency, goals.created_at 
    FROM goals 
    WHERE goals.created_at <= $2
),
goal_completion_counts AS (
    SELECT goal_id, COUNT(id) AS completion_count
    FROM goal_completions
    WHERE goal_completions.created_at >= $1
      AND goal_completions.created_at <= $2
    GROUP BY goal_completions.goal_id
)
select
	goals_created_up_to_week.id,
	goals_created_up_to_week.title,
	goals_created_up_to_week.desired_weekly_frequency,
	goals_created_up_to_week.created_at,
	goal_completion_counts.completion_count
FROM goals_created_up_to_week
LEFT JOIN goal_completion_counts ON goals_created_up_to_week.id = goal_completion_counts.goal_id;

-- name: GetGoalCompletionsCountById :one
WITH goals_created_up_to_week AS (
    SELECT id, title, desired_weekly_frequency, goals.created_at 
    FROM goals 
    WHERE goals.created_at <= $2
    and goals.id = $3
),
goal_completion_counts AS (
    SELECT goal_id, COUNT(id) AS completion_count
    FROM goal_completions
    WHERE goal_completions.created_at >= $1
    AND goal_completions.created_at <= $2
    GROUP BY goal_completions.goal_id
)
select
	goals_created_up_to_week.desired_weekly_frequency,
	COALESCE(goal_completion_counts.completion_count, 0) as completion_count
FROM goals_created_up_to_week
LEFT JOIN goal_completion_counts ON goals_created_up_to_week.id = goal_completion_counts.goal_id;

-- name: GetWeekSummary :one
WITH goals_created_up_to_week AS (
    SELECT id, title, desired_weekly_frequency, goals.created_at 
    FROM goals 
    WHERE goals.created_at <= $2 
),
goals_completed_in_week AS (
	select
		goals.id as id,
		goals.title as title,
		goal_completions.created_at as completed_at,
		DATE(goal_completions.created_at) as completed_at_date
	FROM goal_completions
	inner join goals on goals.id = goal_completions.goal_id 
	WHERE goal_completions.created_at >= $1 
	AND goal_completions.created_at <= $2 
	ORDER BY goal_completions.created_at DESC
),
goals_completed_by_week_day as (
	select 
		goals_completed_in_week.completed_at_date as completed_at_date,
		jsonb_agg(
			jsonb_build_object(
				'id', goals_completed_in_week.id,
				'title', goals_completed_in_week.title,
				'completed_at', goals_completed_in_week.completed_at
			)
		) as completions
	from goals_completed_in_week
	group by goals_completed_in_week.completed_at_date
)
select
	(select count(*) from goals_completed_in_week) as completed,
	(select SUM(goals_created_up_to_week.desired_weekly_frequency) from goals_created_up_to_week) as total,
	json_object_agg(
		goals_completed_by_week_day.completed_at_date,
		goals_completed_by_week_day.completions
	) as goals_per_day
FROM goals_completed_by_week_day;

-- name: CreateNewGoal :one
INSERT INTO public.goals
(id, title, desired_weekly_frequency, created_at)
VALUES($1, $2, $3, now()) RETURNING *;

-- name: CreateNewGoalCompletion :exec
INSERT INTO public.goal_completions
(id, goal_id, created_at)
VALUES($1, $2, $3);

-- name: DeleteAllGoals :exec
DELETE FROM public.goals;

-- name: DeleteAllGoalCompletions :exec
DELETE FROM public.goal_completions;
