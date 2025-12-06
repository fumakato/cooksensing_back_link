package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ==================== UserLink ====================
type UserLink struct {
	ID                       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName                 string  `gorm:"size:255;not null" json:"user_name"`
	Mail                     *string `gorm:"size:255;unique" json:"mail,omitempty"`
	Pass                     *string `gorm:"size:255" json:"-"`
	CookpadID                *string `gorm:"size:255" json:"cookpad_id,omitempty"`
	FirebaseAuthUUID         *string `gorm:"size:255;unique" json:"firebaseauth_uuid,omitempty"`
	IsIdentification         bool    `gorm:"default:false" json:"is_identification"`
	IdentityVerificationText *string `gorm:"size:255" json:"identity_verification_text,omitempty"`

	CookLogs []CookLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	BaseModel
}

// ==================== Recipe ====================
type Recipe struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	RecipeName string  `gorm:"size:255;not null" json:"recipe_name"`
	URL        *string `json:"url,omitempty"`

	Steps    []RecipeStep `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CookLogs []CookLog    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	BaseModel
}

// ==================== CookCategory ====================
type CookCategory struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CookCategoryName string `gorm:"size:255;unique;not null" json:"cook_category_name"`

	Actions []CookAction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	BaseModel
}

// ==================== CookAction ====================
type CookAction struct {
	ID             uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	CookActionName string       `gorm:"size:255;not null" json:"cook_action_name"`
	CookCategoryID uint         `json:"cook_category_id"`
	CookCategory   CookCategory `gorm:"foreignKey:CookCategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`

	ActionIngredients []CookActionIngredients `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	BaseModel
}

// ==================== Ingredients ====================
type Ingredients struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	IngredientsName string `gorm:"size:255;unique;not null" json:"ingredients_name"`

	ActionIngredients []CookActionIngredients `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	BaseModel
}

// ==================== CookActionIngredients ====================
type CookActionIngredients struct {
	ID            uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CookActionID  uint `json:"cook_action_id"`
	IngredientsID uint `json:"ingredients_id"`

	CookAction  CookAction  `gorm:"foreignKey:CookActionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	Ingredients Ingredients `gorm:"foreignKey:IngredientsID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`

	RecipeSteps     []RecipeStep     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	EvaluationItems []EvaluationItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	BaseModel
}

// ==================== RecipeStep ====================
type RecipeStep struct {
	ID                      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	RecipeID                uint    `json:"recipe_id"`
	CookActionIngredientsID uint    `json:"cook_action_ingredients_id"`
	Sequence                int     `gorm:"not null" json:"sequence"`
	Note                    *string `json:"note,omitempty"`

	Recipe                Recipe                `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CookActionIngredients CookActionIngredients `gorm:"foreignKey:CookActionIngredientsID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	FeatureDataLinks      []FeatureDataLink     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	BaseModel
}

// ==================== EvaluationItem ====================
type EvaluationItem struct {
	ID                      uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	CookActionIngredientsID uint     `json:"cook_action_ingredients_id"`
	EvaluationItemName      string   `gorm:"size:255;not null" json:"evaluation_item_name"`
	Unit                    string   `gorm:"size:64;not null" json:"unit"`
	AverageData             *float64 `json:"average_data,omitempty"`

	CookActionIngredients CookActionIngredients `gorm:"foreignKey:CookActionIngredientsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	FeatureDataLinks      []FeatureDataLink     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	BaseModel
}

// ==================== CookLog ====================
type CookLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserLinkID uint      `json:"user_link_id"`
	RecipeID   uint      `json:"recipe_id"`
	CookedAt   time.Time `json:"cooked_at"`

	// UserLinkID を外部キーとして指定
	UserLink UserLink `gorm:"foreignKey:UserLinkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	// RecipeID を外部キーとして指定
	Recipe Recipe `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	// CookLog に属する FeatureDataLink を定義
	Features []FeatureDataLink `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	BaseModel
}

// ==================== FeatureDataLink ====================
type FeatureDataLink struct {
	ID               uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	CookLogID        uint    `json:"cook_log_id"`
	RecipeStepID     uint    `json:"recipe_step_id"`
	EvaluationItemID uint    `json:"evaluation_item_id"`
	Data             float64 `json:"data"`

	CookLog        CookLog        `gorm:"foreignKey:CookLogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	RecipeStep     RecipeStep     `gorm:"foreignKey:RecipeStepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	EvaluationItem EvaluationItem `gorm:"foreignKey:EvaluationItemID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	BaseModel
}
