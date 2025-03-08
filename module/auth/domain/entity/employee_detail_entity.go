package entity

import (
	"time"
)

type EmployeeDetailEntity struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Code           string    `json:"code"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	AvatarUrl      string    `json:"avatar_url"`
	HireDate       time.Time `json:"hire_date"`
	Birthday       time.Time `json:"birthday"`
	CreatedAt      time.Time `json:"created_at"`
	DepartmentCode string    `json:"department_code"`
	DepartmentName string    `json:"department_name"`
	PositionCode   string    `json:"position_code"`
	PositionName   string    `json:"position_name"`
}

