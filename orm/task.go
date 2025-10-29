package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Task struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(20);unique"`
	model.GormModel

	RunningNo      int        `json:"running_no" query:"running_no" gorm:"primaryKey;autoIncrement;not null"`
	ProjectID      string     `json:"project_id" query:"project_id" gorm:"type:varchar(20)"`
	TaskTypeID     string     `json:"task_type_id" query:"task_type_id" gorm:"type:varchar(2)"`
	TaskPriorityID string     `json:"task_priority_id" query:"task_priority_id" gorm:"type:varchar(2)"`
	TaskStatusID   string     `json:"task_status_id" query:"task_status_id" gorm:"type:varchar(2)"`
	Subject        string     `json:"subject" query:"subject" gorm:"type:text"`
	SprintID       string     `json:"sprint_id" query:"sprint_id" gorm:"type:varchar(4)"`
	DueDate        *time.Time `json:"due_date" query:"due_date" gorm:"type:datetime"`
}
