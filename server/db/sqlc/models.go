// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"
)

type Goal struct {
	ID                     string    `json:"id"`
	Title                  string    `json:"title"`
	DesiredWeeklyFrequency int32     `json:"desired_weekly_frequency"`
	CreatedAt              time.Time `json:"created_at"`
}

type GoalCompletion struct {
	ID        string    `json:"id"`
	GoalID    string    `json:"goal_id"`
	CreatedAt time.Time `json:"created_at"`
}