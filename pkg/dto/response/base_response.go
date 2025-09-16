package response

type BaseResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Errors  ArrErrorResponse `json:"errors"`
	Data    interface{}      `json:"data"`
}

func (r *BaseResponse) DefaultOK() {
	r.Code = 200
	r.Message = "success"
}

func (r *BaseResponse) DefaultCreated() {
	r.Code = 201
	r.Message = "created"
}

func (r *BaseResponse) DefaultBadRequest() {
	r.Code = 400
	r.Message = "Request body not match"
}
func (r *BaseResponse) DefaultUnauthorized() {
	r.Code = 401
	r.Message = "Unauthorized access"
}

func (r *BaseResponse) DefaultForbidden() {
	r.Code = 403
	r.Message = "Forbidden access"
}

func (r *BaseResponse) DefaultNotFound() {
	r.Code = 404
	r.Message = "Record not found"
}

func (r *BaseResponse) DefaultConflict() {
	r.Code = 409
	r.Message = "New data already exists"
}

func (r *BaseResponse) DefaultInternalError() {
	r.Code = 500
	r.Message = "Request failed, server error"
}
