package entity

type LoginSuccessResponseEntity struct {
	EmployeeDetail EmployeeDetailEntity `json:"employee_detail"`
	AccessToken    string               `json:"access_token"`
}
