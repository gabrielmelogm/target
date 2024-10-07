package request

type CreateNewGoalRequest struct {
	Title                  string `json:"title" validate:"required"`
	DesiredWeeklyFrequency int    `json:"desired_weekly_frequency" validate:"required,min=1,max=7"`
}
