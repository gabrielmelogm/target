-- name: DeleteAllGoals :exec
DELETE FROM public.goals;

-- name: DeleteAllGoalCompletions :exec
DELETE FROM public.goal_completions;
