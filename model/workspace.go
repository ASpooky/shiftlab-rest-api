package model

import "time"

/* 職場テーブル(workspaces)
   店舗名(name)
   時給(salary)
   早朝手当(morning_salary)
   深夜手当(night_salary)
   表示色(color) */

type Workspace struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"not null"`
	Salary        uint      `json:"salary" gorm:"not null"`
	MorningSalary uint      `json:"morning_salary" gorm:"not null"`
	NightSalary   uint      `json:"night_salary" gorm:"not null"`
	Color         string    `json:"color" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
	User          User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId        uint      `json:"user_id" gorm:"not null"`
}

type WorkspaceResponse struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"not null"`
	Salary        uint      `json:"salary" gorm:"not null"`
	MorningSalary uint      `json:"morning_salary" gorm:"not null"`
	NightSalary   uint      `json:"night_salary" gorm:"not null"`
	Color         string    `json:"color" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
}
