package category_resources

type CategoryCreateRequest struct {
	Type string `json:"type" validate:"required"`
}

type CategoryUpdateRequest struct {
	Type string `json:"type" validate:"required"`
}
