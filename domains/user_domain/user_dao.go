package user_domain

import (
	"golang-final-project4-team2/db"
	"golang-final-project4-team2/resources/user_resources"
	"golang-final-project4-team2/utils/error_formats"
	"golang-final-project4-team2/utils/error_utils"
)

var UserDomain userDomainRepo = &userDomain{}

const (
	queryCreateUser = `INSERT INTO users ( full_name, email, password, role, balance) 
	VALUES($1, $2, $3, $4, $5) RETURNING id, full_name, email,password,balance, created_at`
	queryUserLogin    = `SELECT * from users where email = $1`
	queryUserUpdate   = `UPDATE users set updated_at = now(), email = $1,full_name = $2 where id = $3 RETURNING id,full_name,email, password,role,created_at, updated_at`
	queryUserDelete   = `UPDATE users SET  deleted_at = now() where id = $1`
	queryUserById     = `SELECT * from users where id = $1 and deleted_at is NULL`
	queryGetBalance   = `SELECT balance from users where id = $1`
	queryTopupBalance = `UPDATE users set balance = balance + $1 where id = $2 RETURNING balance`
)

type userDomainRepo interface {
	UserRegister(*user_resources.UserRegisterRequest) (*User, error_utils.MessageErr)
	UserLogin(*user_resources.UserLoginRequest) (*User, error_utils.MessageErr)
	UserUpdate(string, *user_resources.UserUpdateRequest) (*User, error_utils.MessageErr)
	UserDelete(string) error_utils.MessageErr
	UserCheckIsExists(int64) bool
	UserGetBalance(string) (int64, error_utils.MessageErr)
	UserTopupBalance(string, *user_resources.UserTopupBalanceRequest) (int64, error_utils.MessageErr)
}

type userDomain struct {
}

func (u *userDomain) UserRegister(userReq *user_resources.UserRegisterRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateUser, userReq.FullName, userReq.Email, userReq.Password, user_resources.RoleCustomer, 0)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Balance, &user.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &user, nil
}

func (u *userDomain) UserCheckIsExists(id int64) bool {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserById, id)
	var user User
	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if row.Err() != nil || err != nil {
		return false
	}
	return true
}

func (u *userDomain) UserLogin(userReq *user_resources.UserLoginRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserLogin, userReq.Email)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	return &user, nil
}

func (u *userDomain) UserUpdate(id string, userReq *user_resources.UserUpdateRequest) (*User, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserUpdate, userReq.Email, userReq.FullName, id)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var user User

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	return &user, nil
}
func (u *userDomain) UserDelete(id string) error_utils.MessageErr {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryUserDelete, id)
	if row.Err() != nil {
		return error_utils.NewBadRequest(row.Err().Error())
	}
	return nil
}

func (u *userDomain) UserGetBalance(id string) (int64, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryGetBalance, id)
	var balance int64

	err := row.Scan(&balance)
	if err != nil {
		return 0, error_utils.NewBadRequest(err.Error())
	}
	return balance, nil
}

func (u *userDomain) UserTopupBalance(id string, topupBalance *user_resources.UserTopupBalanceRequest) (int64, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryTopupBalance, topupBalance.Balance, id)

	var balance int64
	err := row.Scan(&balance)

	if err != nil {
		return 0, error_utils.NewBadRequest(err.Error())
	}
	return balance, nil
}
