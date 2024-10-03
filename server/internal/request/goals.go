package request

type CreateNewGoalRequest struct {
	Title string `json:"title"`
	DesiredWeeklyFrequency int `json:"desired_weekly_frequency"`
}
