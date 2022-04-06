package service

import (
	"context"
	"database/sql"
	"golang_restful_api/exception"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
	"golang_restful_api/model/web"
	"golang_restful_api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

// best practice nya yaitu membuat provider function tetap return ke struct nya bukan ke interface nya
func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// validasi request nya
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// mulai transaction
	tx, err := service.DB.Begin()
	defer helper.TxCommitOrRollback(tx, err)

	// bikin 'object' category
	category := domain.Category{
		Name:       request.Name,
		CreateTime: request.CreateTime,
		UpdateTime: request.UpdateTime,
	}
	category = service.CategoryRepository.Save(ctx, tx, category)
	categoryResponse := web.CategoryResponse{
		Id:         category.Id,
		Name:       category.Name,
		CreateTime: category.CreateTime,
		UpdateTime: category.UpdateTime,
	}

	return categoryResponse
}
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	defer helper.TxCommitOrRollback(tx, err)

	// cek dulu apakah category tersedia atau tidak, sekaligus untuk mendapatkan data lengkapnya
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// kemudian jika tersedia maka reset dengan yang terdapat pada request
	category.Name = request.Name
	category.UpdateTime = request.UpdateTime
	categoryUpdated := service.CategoryRepository.Update(ctx, tx, category)
	categoryResponse := web.CategoryResponse{
		Id:         categoryUpdated.Id,
		Name:       categoryUpdated.Name,
		CreateTime: categoryUpdated.CreateTime,
		UpdateTime: categoryUpdated.UpdateTime,
	}
	return categoryResponse
}
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	defer helper.TxCommitOrRollback(tx, err)

	// cek dulu apakah category tersedia atau tidak
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category.Id)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	defer helper.TxCommitOrRollback(tx, err)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	categoryResponse := web.CategoryResponse{
		Id:         category.Id,
		Name:       category.Name,
		CreateTime: category.CreateTime,
		UpdateTime: category.UpdateTime,
	}

	return categoryResponse
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	defer helper.TxCommitOrRollback(tx, err)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponse := web.CategoryResponse{
			Id:         category.Id,
			Name:       category.Name,
			CreateTime: category.CreateTime,
			UpdateTime: category.UpdateTime,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses
}
