package category_domain

import (
	"golang-final-project3-team2/db"
	"golang-final-project3-team2/resources/category_resources"
	"golang-final-project3-team2/utils/error_formats"
	"golang-final-project3-team2/utils/error_utils"
)

var CategoryDomain categoryDomainRepo = &categoryDomain{}

const (
	queryCreateCategory = `INSERT INTO categories (type) 
	VALUES($1) RETURNING id,type, created_at`

	queryGetCategories = `
	select c.id as c_id, type, c.updated_at as c_updated_at, c.created_at as c_created_at,
	t.id as t_id, t.title , t.description, t.user_id, t.created_at as t_created_at, t.updated_at as t_updated_at
	from categories c left join tasks t on c.id = t.category_id
	ORDER BY c.id`

	queryCategoryUpdate          = `UPDATE categories set updated_at = now(), type = $1 where id = $2 RETURNING id,type, updated_at`
	queryDeleteTasksByCategoryId = `DELETE from tasks where category_id = $1`
	queryDeleteCategory          = `DELETE from categories where id = $1`
)

type categoryDomainRepo interface {
	CreateCategory(*category_resources.CategoryCreateRequest, string) (*category_resources.CategoryCreateResponse, error_utils.MessageErr)
	GetCategories() (*[]category_resources.CategoriesGetResponse, error_utils.MessageErr)
	UpdateCategory(*category_resources.CategoryUpdateRequest, string) (*category_resources.CategoryUpdateResponse, error_utils.MessageErr)
	DeleteCategory(string) error_utils.MessageErr
}

type categoryDomain struct {
}

func (u *categoryDomain) CreateCategory(req *category_resources.CategoryCreateRequest, userId string) (*category_resources.CategoryCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateCategory, req.Type)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var category category_resources.CategoryCreateResponse

	err := row.Scan(&category.Id, &category.Type, &category.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &category, nil
}

func (u *categoryDomain) GetCategories() (*[]category_resources.CategoriesGetResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetCategories)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}

	categories := make([]category_resources.CategoriesGetResponse, 0)
	categoryTasks := make([]category_resources.CategoriesTaskGetResponse, 0)
	var category category_resources.CategoriesGetResponse
	var lastCategoryId int64 = -1

	for rows.Next() {
		var categoryTask category_resources.CategoriesTaskGetResponse

		// Scan Data
		err = rows.Scan(&category.Id, &category.Type, &category.UpdatedAt, &category.CreatedAt, &categoryTask.Id, &categoryTask.Title, &categoryTask.Description, &categoryTask.UserId, &categoryTask.CreatedAt, &categoryTask.UpdatedAt)

		// Cek apakah ada task di category ini
		if categoryTask.Id != nil {
			var categoryId *int64
			categoryId = new(int64)
			*categoryId = category.Id
			categoryTask.CategoryId = categoryId
			categoryTasks = append(categoryTasks, categoryTask)
		}

		// Apabila category ini bukan category sebelumnya maka tambahkan kedalam array
		if lastCategoryId != category.Id {
			category.Tasks = categoryTasks
			categoryTasks = make([]category_resources.CategoriesTaskGetResponse, 0)
			categories = append(categories, category)
		} else {
			// Jika Iya maka update array task category terakhir
			categories[len(categories)-1].Tasks = append(categories[len(categories)-1].Tasks, categoryTask)
			categoryTasks = make([]category_resources.CategoriesTaskGetResponse, 0)
		}

		lastCategoryId = category.Id

		if err != nil {
			return nil, error_formats.ParseError(err)
		}
	}

	return &categories, nil
}

func (u *categoryDomain) UpdateCategory(request *category_resources.CategoryUpdateRequest, categoryId string) (*category_resources.CategoryUpdateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCategoryUpdate, request.Type, categoryId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var category category_resources.CategoryUpdateResponse

	err := row.Scan(&category.Id, &category.Type, &category.UpdatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &category, nil
}

func (u *categoryDomain) DeleteCategory(id string) (error error_utils.MessageErr) {
	dbInstance, err := db.GetDB().Begin()
	if err != nil {
		error = error_utils.NewInternalServerError(err.Error())
		return
	}

	defer func() {
		if error != nil {
			dbInstance.Rollback()
			return
		}
		err := dbInstance.Commit()
		if err != nil {
			error = error_utils.NewInternalServerError(err.Error())
		}
	}()

	rows, err := dbInstance.Query(queryDeleteTasksByCategoryId, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	rows, err = dbInstance.Query(queryDeleteCategory, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil
}
