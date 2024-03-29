package auth

type IAuthUseCases interface {
	ValidateRequest(request ValidateRequestDTO) error
	ValidateListPersonBelongingsRequest(request ValidateListPersonBelongingsRequestDTO) error
}

type ValidateRequestDTO struct {
	Action string `json:"action"`
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
