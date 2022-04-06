package controller

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"golang_restful_api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// best practice nya yaitu membuat provider function tetap return ke struct nya bukan ke interface nya
func NewCategoryControllerImpl(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// ambil request
// decode (to Struct), masukin ke parameter Service
// setelah berhasil bikin struct model web response dan isi dengan return dari Service
// encode (to JSON) struct web response

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// request nya dijadiin object dari struct dengan menggunakan Decode
	categoryCreateRequest := web.CategoryCreateRequest{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	// ex :
	// decoder := json.NewDecoder(request.Body) // baca dari request
	// err := decoder.Decode(&categoryCreateRequest)
	// PanicIfError(err)

	// setelah jadi object kemudian masukan ke object struct web response
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	categoryWebResponse := web.CategoryWebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	// kemudian encode lagi object dari struct web response dan siap ditampilkan
	helper.WriteToResponseBody(writer, categoryWebResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// untuk dapatin request berupa Name dan UpdateTime, belum ada id nya
	categoryUpdateRequest := web.CategoryUpdateRequest{
		UpdateTime: time.Now(),
	}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// dapatkan id dan untuk menentukan id mana yang akan diupdate
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	categoryWebResponse := web.CategoryWebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, categoryWebResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	categoryWebResponse := web.CategoryWebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, categoryWebResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	categoryWebResponse := web.CategoryWebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, categoryWebResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	categoryWebResponse := web.CategoryWebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, categoryWebResponse)
}
