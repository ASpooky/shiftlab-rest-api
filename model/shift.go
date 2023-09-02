package model

import "time"

/* シフトテーブル(shifts)
   ユーザid(user_name:foreignkey)
   店舗id(workspace_name:foreignkey)
   開始時間(start_time)
   終了時間(end_time)
*/

type Shift struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	Workspace   Workspace `json:"workspace" gorm:"foreignKey:WorkspaceId; constraint:OnDelete:CASCADE"`
	WorkspaceId uint      `json:"workspace_id" gorm:"not null`
	User        User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint      `json:"user_id" gorm:"not null`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type ShiftResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	WorkspaceId uint      `json:"workspace_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}
