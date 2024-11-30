package dto

type ReactionRequest struct {
	Reaction     int8   `json:"reaction"`
	TargetUserID string `json:"target_user_id"`
}
