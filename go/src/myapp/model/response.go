package model

type RecipeDetailResponse struct {
	CookLogID  uint           `json:"cookLogId"`
	RecipeName string         `json:"recipeName"`
	Steps      []StepResponse `json:"steps"`
}

type StepResponse struct {
	Sequence    int                  `json:"sequence"`
	Note        *string              `json:"note,omitempty"`
	Evaluations []EvaluationResponse `json:"evaluations"`
}

type EvaluationResponse struct {
	Average          float64 `json:"average"`
	EvaluationItemID uint    `json:"evaluation_item_id"`
	Label            string  `json:"label"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}
