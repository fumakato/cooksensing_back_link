package controller

import (
	"log"
	"myapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// リクエスト用
type CookLogRequest struct {
	UserLinkID uint `json:"user_link_id" binding:"required"`
}

// ハンドラー
func GetCookLogAllByUserLinkID(c *gin.Context) {
	var req CookLogRequest

	// JSONバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// DB検索
	results, err := database.FindCookLogsByUserLinkID(req.UserLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cook logs"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{"data": results})
}

// リクエスト用構造体
type CookLogDetailRequest struct {
	UserID    uint `json:"user_id" binding:"required"`
	CookLogID uint `json:"cook_log_id" binding:"required"`
}

// レスポンス用構造体
type EvaluationResponse struct {
	Average          float64 `json:"average"`
	EvaluationItemID uint    `json:"evaluation_item_id"`
	Label            string  `json:"label"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

type StepResponse struct {
	Evaluations []EvaluationResponse `json:"evaluations"`
}

type RecipeDetailResponse struct {
	CookLogID  uint           `json:"cookLogId"`
	RecipeName string         `json:"recipeName"`
	Steps      []StepResponse `json:"steps"`
}

// ハンドラー
func GetCookLogDetail(c *gin.Context) {
	var req CookLogDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// データベースからCookLog詳細を取得
	recipeDetail, err := database.FindCookLogDetail(req.UserID, req.CookLogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cook log detail"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{"recipes": recipeDetail})
}

type UserIDRequest struct {
	UserLinkID uint `json:"user_id" binding:"required"`
}

func GetCookCategoryAverages(c *gin.Context) {
	var req UserIDRequest

	// リクエストパース
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// DBから取得
	results, err := database.FindCookCategoryAveragesByUserLinkID(req.UserLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cook category averages"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{"data": results})
}

// 追加分
// リクエスト
type CommonStepsRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

// 値（各ログごと）
type LogValue struct {
	CookLogID  uint    `json:"cook_log"`
	RecipeName string  `json:"recipe_name"`
	Value      float64 `json:"value"`
}

// 評価項目 + 各ログの値
type EvaluationWithValues struct {
	EvaluationItemID uint       `json:"evaluation_item_id"`
	Label            string     `json:"label"`
	Unit             string     `json:"unit"`
	Average          *float64   `json:"average,omitempty"`
	Values           []LogValue `json:"values"`
}

// ステップ（共通 CAI ごと）
type CommonStepResponse struct {
	Note        string                 `json:"note"`
	Evaluations []EvaluationWithValues `json:"evaluations"`
}

func GetCommonStepsByUserID(c *gin.Context) {
	var req CommonStepsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	steps, err := database.FindCommonStepsWithFeatureValues(req.UserID)
	if err != nil {
		log.Printf("FindCommonStepsWithFeatureValues error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve common steps"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"steps": steps})
}

type CommonStepsByCookLogsRequest struct {
	UserID     uint   `json:"user_id" binding:"required"`
	CookLogIDs []uint `json:"cook_log_ids" binding:"required"` // 2個以上想定（バリデーションはコードで）
}

func GetCommonStepsByCookLogs(c *gin.Context) {
	var req CommonStepsByCookLogsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if len(req.CookLogIDs) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cook_log_ids must contain at least 2 ids"})
		return
	}

	resp, err := database.FindCommonStepsWithValuesByCookLogs(req.UserID, req.CookLogIDs)
	if err != nil {
		log.Printf("FindCommonStepsWithValuesByCookLogs error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve common steps"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"steps": resp})
}

// controller/cook_log.go

// /master/all
func GetMasterAll(c *gin.Context) {
	resp, err := database.FindMasterAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch masters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// /cook_log/common_action_values
type CommonActionValuesRequest struct {
	UserID                uint `json:"user_id" binding:"required"`
	CookActionIngredients uint `json:"cook_action_ingredients_id" binding:"required"`
}

// POST /cook_log/common_action_values
func GetCommonActionValues(c *gin.Context) {
	var req CommonActionValuesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := database.FindFeatureDataGroupedByEvaluation(req.UserID, req.CookActionIngredients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"steps": resp})
}
