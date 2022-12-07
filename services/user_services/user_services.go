package user_services

import (
	"golang-final-project3-team2/domains/user_domain"
	"golang-final-project3-team2/resources/user_resources"
	"golang-final-project3-team2/utils/error_utils"
	"golang-final-project3-team2/utils/helpers"
)

var UserService userServiceRepo = &userService{}

type userServiceRepo interface {
	UserRegister(*user_resources.UserRegisterRequest) (*user_resources.UserRegisterResponse, error_utils.MessageErr)
	UserLogin(*user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr)
	UserUpdate(string, *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr)
	UserTopup(string, *user_resources.UserTopupBalanceRequest) (int64, error_utils.MessageErr)
	UserDelete(string) error_utils.MessageErr
	GenerateAdminData() error_utils.MessageErr
}

type userService struct{}

func (u *userService) UserRegister(userReq *user_resources.UserRegisterRequest) (*user_resources.UserRegisterResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	newPass, err := helpers.HashPass(userReq.Password)
	if err != nil {
		return nil, err
	}

	userReq.Password = *newPass

	user, err := user_domain.UserDomain.UserRegister(userReq)

	if err != nil {
		return nil, err
	}

	return &user_resources.UserRegisterResponse{
		Id:        user.Id,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *userService) UserLogin(userReq *user_resources.UserLoginRequest) (*user_resources.UserLoginResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	user, err := user_domain.UserDomain.UserLogin(userReq)

	if err != nil {
		return nil, err
	}

	if valid := helpers.ComparePass([]byte(user.Password), []byte(userReq.Password)); !valid {
		return nil, error_utils.NewBadRequest("invalid credential")
	}

	token, err := helpers.GenerateToken(&user_resources.UserGenerateTokenParam{
		Id:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	})

	if err != nil {
		return nil, err
	}

	return &user_resources.UserLoginResponse{
		Token: *token,
	}, nil
}

func (u *userService) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*user_resources.UserUpdateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return nil, err
	}
	user, err := user_domain.UserDomain.UserUpdate(id, userReq)

	if err != nil {
		return nil, err
	}

	return &user_resources.UserUpdateResponse{
		Id:        user.Id,
		FullName:  user.FullName,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *userService) UserTopup(id string, userReq *user_resources.UserTopupBalanceRequest) (int64, error_utils.MessageErr) {
	err := helpers.ValidateRequest(userReq)

	if err != nil {
		return 0, err
	}
	balance, err := user_domain.UserDomain.UserTopupBalance(id, userReq)

	if err != nil {
		return balance, err
	}

	return balance, nil
}

func (u *userService) UserDelete(id string) error_utils.MessageErr {
	err := user_domain.UserDomain.UserDelete(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GenerateAdminData() error_utils.MessageErr {
	return nil
}
