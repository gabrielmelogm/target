-- name: CreateNewGoal :one
INSERT INTO public.goals
(id, title, desired_weekly_frequency, created_at)
VALUES($1, $2, $3, now()) RETURNING *;

-- name: DeleteAllGoals :exec
DELETE FROM public.goals;

-- name: DeleteAllGoalCompletions :exec
DELETE FROM public.goal_completions;
