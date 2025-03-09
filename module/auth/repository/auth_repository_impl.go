package repository

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/auth/data/data_source"
	"asset_management_backend/module/auth/domain/entity"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authRepository struct {
	authRemoteService datasource.AuthServiceClient
}

// Login implements AuthRepository.
func (a *authRepository) Login(context context.Context, username string, password string) (*entity.LoginSuccessResponseEntity, error) {
	loginRequest := datasource.LoginRequestRPC{
		Username: username,
		Password: password,
	}

	res, err := a.authRemoteService.Login(context, &loginRequest)
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		switch grpcStatus.Code() {
		case codes.InvalidArgument:
			return nil, error_app.ErrBadRequest
		case codes.Unauthenticated:
			return nil, error_app.ErrUnauthorized
		case codes.NotFound:
			return nil, error_app.ErrDocumentNotFound
		default:
			return nil, err
		}
	}

	employee := res.EmployeeDetail

	loginResponseEntity := entity.LoginSuccessResponseEntity{
		EmployeeDetail: entity.EmployeeDetailEntity{
			ID:             int(employee.Id),
			Name:           employee.Name,
			Username:       employee.Username,
			Code:           employee.Code,
			Email:          employee.Email,
			PhoneNumber:    employee.PhoneNumber,
			AvatarUrl:      employee.AvatarUrl,
			HireDate:       employee.HireDate.AsTime(),
			Birthday:       employee.Birthday.AsTime(),
			CreatedAt:      employee.CreatedAt.AsTime(),
			DepartmentCode: employee.DepartmentCode,
			DepartmentName: employee.DepartmentName,
			PositionCode:   employee.PositionCode,
			PositionName:   employee.PositionName,
		},
		AccessToken: res.AccessToken,
	}

	return &loginResponseEntity, nil

}

func NewAuthRepository() AuthRepository {
	return &authRepository{
		authRemoteService: datasource.NewAuthServiceClient(infras.GetEmployeeServiceConn()),
	}
}
