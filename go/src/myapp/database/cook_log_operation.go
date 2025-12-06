package database

import (
	"errors"
	"fmt"
	"myapp/model"
	"sort"

	"gorm.io/gorm"
)

// CookLogResponse データベースからのレスポンス用
type CookLogResponse struct {
	CookLogID  uint   `json:"cook_log_id"`
	RecipeID   uint   `json:"recipe_id"`
	RecipeName string `json:"recipe_name"`
}

// ユーザーに紐づく CookLog 一覧を取得
func FindCookLogsByUserLinkID(userLinkID uint) ([]CookLogResponse, error) {
	var results []CookLogResponse
	if err := db.
		Table("cook_logs").
		Select("cook_logs.id as cook_log_id, recipes.id as recipe_id, recipes.recipe_name as recipe_name").
		Joins("JOIN recipes ON cook_logs.recipe_id = recipes.id").
		Where("cook_logs.user_link_id = ?", userLinkID).
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// CookLogDetailResponse: APIで返す形のレスポンス
type CookLogDetailResponse struct {
	CookLogID  uint                `json:"cook_log_id"`
	RecipeName string              `json:"recipe_name"`
	Steps      []CookLogDetailStep `json:"steps"`
}

type CookLogDetailStep struct {
	Sequence int                       `json:"sequence"`
	Note     string                    `json:"note"`
	Evals    []CookLogDetailEvaluation `json:"evaluations"`
}

type CookLogDetailEvaluation struct {
	EvaluationItemID uint     `json:"evaluation_item_id"`
	Label            string   `json:"label"`
	Unit             string   `json:"unit"`
	Average          *float64 `json:"average"`
	Value            float64  `json:"value"`
}

// CookLog詳細取得
func FindCookLogDetail(userLinkID, cookLogID uint) (*CookLogDetailResponse, error) {
	fmt.Println("🔍 FindCookLogDetail 開始")
	fmt.Printf("➡️ userLinkID=%d, cookLogID=%d\n", userLinkID, cookLogID)

	// まずCookLogをJOINしてレシピを取得
	var cookLog model.CookLog
	if err := db.Preload("Recipe").Where("id = ? AND user_link_id = ?", cookLogID, userLinkID).First(&cookLog).Error; err != nil {
		fmt.Printf("❌ CookLog取得エラー: %v\n", err)
		return nil, err
	}
	fmt.Printf("✅ CookLog取得成功: CookLogID=%d, RecipeName=%s\n", cookLog.ID, cookLog.Recipe.RecipeName)

	// RecipeStepsを取得
	var steps []model.RecipeStep
	if err := db.Where("recipe_id = ?", cookLog.RecipeID).Order("sequence asc").Find(&steps).Error; err != nil {
		fmt.Printf("❌ RecipeStep取得エラー: %v\n", err)
		return nil, err
	}
	fmt.Printf("✅ RecipeSteps取得成功: 件数=%d\n", len(steps))

	// 各StepごとにEvaluationをJOIN
	var responseSteps []CookLogDetailStep
	for _, step := range steps {
		fmt.Printf("➡️ StepID=%d, Sequence=%d, Note=%v\n", step.ID, step.Sequence, step.Note)

		// FeatureDataLink JOIN EvaluationItem
		var rows []struct {
			EvaluationItemID   uint
			EvaluationItemName string
			Unit               string
			AverageData        *float64
			Value              float64
		}

		err := db.Table("feature_data_links").
			Select("evaluation_items.id as evaluation_item_id, evaluation_items.evaluation_item_name, evaluation_items.unit, evaluation_items.average_data, feature_data_links.data as value").
			Joins("JOIN evaluation_items ON feature_data_links.evaluation_item_id = evaluation_items.id").
			Where("feature_data_links.cook_log_id = ? AND feature_data_links.recipe_step_id = ?", cookLogID, step.ID).
			Scan(&rows).Error
		if err != nil {
			fmt.Printf("❌ Evaluation取得エラー StepID=%d: %v\n", step.ID, err)
			return nil, err
		}
		fmt.Printf("✅ StepID=%d のEvaluation取得: 件数=%d\n", step.ID, len(rows))

		// Evalを整形
		var evals []CookLogDetailEvaluation
		for _, r := range rows {
			evals = append(evals, CookLogDetailEvaluation{
				EvaluationItemID: r.EvaluationItemID,
				Label:            r.EvaluationItemName,
				Unit:             r.Unit,
				Average:          r.AverageData,
				Value:            r.Value,
			})
		}

		noteText := ""
		if step.Note != nil {
			noteText = *step.Note
		}

		responseSteps = append(responseSteps, CookLogDetailStep{
			Sequence: step.Sequence,
			Note:     noteText,
			Evals:    evals,
		})
	}

	// 最終レスポンス構築
	resp := &CookLogDetailResponse{
		CookLogID:  cookLog.ID,
		RecipeName: cookLog.Recipe.RecipeName,
		Steps:      responseSteps,
	}
	fmt.Println("✅ FindCookLogDetail 完了")
	return resp, nil
}

// // package database

// func FindCookLogDetail(userID, cookLogID uint) ([]model.RecipeDetailResponse, error) {
// 	var cookLog model.CookLog
// 	if err := db.Preload("Recipe").
// 		Preload("Recipe.Steps.CookActionIngredients.EvaluationItems").
// 		Preload("Recipe.Steps.FeatureDataLinks").
// 		Where("cook_logs.id = ? AND cook_logs.user_link_id = ?", cookLogID, userID).
// 		First(&cookLog).Error; err != nil {
// 		return nil, err
// 	}

// 	// 整形
// 	var steps []model.StepResponse
// 	for _, step := range cookLog.Recipe.Steps {
// 		var evals []model.EvaluationResponse
// 		for _, eval := range step.CookActionIngredients.EvaluationItems {
// 			// 対応する FeatureDataLink を探す
// 			var value float64
// 			for _, fdl := range step.FeatureDataLinks {
// 				if fdl.EvaluationItemID == eval.ID {
// 					value = fdl.Data
// 					break
// 				}
// 			}

// 			evals = append(evals, model.EvaluationResponse{
// 				Average:          eval.AverageData,
// 				EvaluationItemID: eval.ID,
// 				Label:            eval.EvaluationItemName,
// 				Unit:             eval.Unit,
// 				Value:            value,
// 			})
// 		}
// 		steps = append(steps, model.StepResponse{Evaluations: evals})
// 	}

// 	return []model.RecipeDetailResponse{
// 		{
// 			CookLogID:  cookLog.ID,
// 			RecipeName: cookLog.Recipe.RecipeName,
// 			Steps:      steps,
// 		},
// 	}, nil
// }

// CookCategoryAverageResponse APIレスポンス用
type CookCategoryAverageResponse struct {
	CookCategoryName    string  `json:"cook_category_name"`
	CookCategoryAverage float64 `json:"cook_category_average"`
}

// FindCookCategoryAveragesByUserLinkID 総合評価のみカテゴリ平均を算出
func FindCookCategoryAveragesByUserLinkID(userLinkID uint) ([]CookCategoryAverageResponse, error) {
	var results []CookCategoryAverageResponse

	if err := db.
		Table("feature_data_links").
		Select(`
			cook_categories.cook_category_name,
			AVG(feature_data_links.data) as cook_category_average
		`).
		Joins(`JOIN recipe_steps 
		           ON feature_data_links.recipe_step_id = recipe_steps.id`).
		Joins(`JOIN cook_action_ingredients 
		           ON recipe_steps.cook_action_ingredients_id = cook_action_ingredients.id`).
		Joins(`JOIN cook_actions 
		           ON cook_action_ingredients.cook_action_id = cook_actions.id`).
		Joins(`JOIN cook_categories 
		           ON cook_actions.cook_category_id = cook_categories.id`).
		// 🔑 総合評価だけを JOIN 時点で限定する
		Joins(`JOIN evaluation_items 
		           ON feature_data_links.evaluation_item_id = evaluation_items.id 
		          AND evaluation_items.evaluation_item_name = '総合評価'`).
		Joins(`JOIN cook_logs 
		           ON feature_data_links.cook_log_id = cook_logs.id`).
		Where("cook_logs.user_link_id = ?", userLinkID).
		Group("cook_categories.id, cook_categories.cook_category_name").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// 追加🟡🔴

// コントローラと同じ構造体（ここに置く派）
type LogValue struct {
	CookLogID  uint
	RecipeName string
	Value      float64
}
type EvaluationWithValues struct {
	EvaluationItemID uint
	Label            string
	Unit             string
	Average          *float64
	Values           []LogValue
}
type CommonStepResponse struct {
	Note        string
	Evaluations []EvaluationWithValues
}

// 共通 CAI を SQL で抽出するために CookLog 件数を取得
func countCookLogsByUser(userID uint) (int, error) {
	var n int64
	if err := db.Table("cook_logs").Where("user_link_id = ?", userID).Count(&n).Error; err != nil {
		return 0, err
	}
	return int(n), nil
}

// メイン：共通 CAI を求め、各評価×各ログの値をまとめて返す
func FindCommonStepsWithFeatureValues(userID uint) ([]CommonStepResponse, error) {
	totalLogs, err := countCookLogsByUser(userID)
	if err != nil {
		return nil, err
	}
	if totalLogs == 0 {
		return []CommonStepResponse{}, nil
	}

	// 1) 共通 CAI を求める（このユーザの全CookLogに出現する CAI）
	//   ※ 各 CookLog の recipe_id から recipe_steps を辿り、CAI を抽出
	var caiIDs []uint
	commonCaiSQL := `
        SELECT rs.cook_action_ingredients_id AS cai_id
        FROM cook_logs cl
        JOIN recipe_steps rs ON rs.recipe_id = cl.recipe_id
        WHERE cl.user_link_id = ?
        GROUP BY rs.cook_action_ingredients_id
        HAVING COUNT(DISTINCT cl.id) = ?
    `
	if err := db.Raw(commonCaiSQL, userID, totalLogs).Scan(&caiIDs).Error; err != nil {
		return nil, fmt.Errorf("common cai query: %w", err)
	}
	if len(caiIDs) == 0 {
		return []CommonStepResponse{}, nil
	}

	// 2) 共通 CAI について、各評価項目 × 各 CookLog の値を一気に取得
	//    - note はその CAI を持つ RecipeStep の最小 note（代表値）を採用
	//    - value は feature_data_links.data の AVG（万一複数行あっても頑健）
	type row struct {
		CaiID            uint
		Note             *string
		EvaluationItemID uint
		Label            string
		Unit             string
		Average          *float64
		CookLogID        uint
		RecipeName       string
		Value            *float64
	}

	rows := []row{}
	dataSQL := `
        SELECT
          rs.cook_action_ingredients_id              AS cai_id,
          MIN(rs.note)                                AS note,
          ei.id                                       AS evaluation_item_id,
          ei.evaluation_item_name                     AS label,
          ei.unit                                     AS unit,
          ei.average_data                             AS average,
          cl.id                                       AS cook_log_id,
          r.recipe_name                               AS recipe_name,
          AVG(fdl.data)                               AS value
        FROM cook_logs cl
        JOIN recipes r ON r.id = cl.recipe_id
        JOIN recipe_steps rs ON rs.recipe_id = cl.recipe_id
        JOIN evaluation_items ei ON ei.cook_action_ingredients_id = rs.cook_action_ingredients_id
        LEFT JOIN feature_data_links fdl
          ON  fdl.cook_log_id = cl.id
          AND fdl.recipe_step_id = rs.id
          AND fdl.evaluation_item_id = ei.id
        WHERE cl.user_link_id = ?
          AND rs.cook_action_ingredients_id IN (?)
        GROUP BY rs.cook_action_ingredients_id, ei.id, cl.id, r.recipe_name
        ORDER BY rs.cook_action_ingredients_id, ei.id, cl.id
    `
	// GORM の IN(?) プレースホルダ配列展開
	if err := db.Raw(dataSQL, userID, caiIDs).Scan(&rows).Error; err != nil {
		return nil, fmt.Errorf("data query: %w", err)
	}

	// 3) rows を整形（cai ごと -> evaluation ごと -> values 配列）
	//    cai_id -> (note, map[evaluation_item_id]Eval, order用スライス)
	type evalAgg struct {
		Label   string
		Unit    string
		Average *float64
		Values  []LogValue
	}
	type caiAgg struct {
		Note string
		E    map[uint]*evalAgg
		Keys []uint
	}

	ag := map[uint]*caiAgg{}

	for _, r := range rows {
		// CAI 初期化
		if ag[r.CaiID] == nil {
			note := ""
			if r.Note != nil {
				note = *r.Note
			}
			ag[r.CaiID] = &caiAgg{
				Note: note,
				E:    map[uint]*evalAgg{},
			}
		}
		// 評価項目 初期化
		if ag[r.CaiID].E[r.EvaluationItemID] == nil {
			ag[r.CaiID].E[r.EvaluationItemID] = &evalAgg{
				Label:   r.Label,
				Unit:    r.Unit,
				Average: r.Average,
				Values:  []LogValue{},
			}
			ag[r.CaiID].Keys = append(ag[r.CaiID].Keys, r.EvaluationItemID)
		}
		// 値を追加（NULL は 0 にしない。前段でAVGしているので nil の可能性は低いが安全側）
		v := 0.0
		if r.Value != nil {
			v = *r.Value
		}
		ag[r.CaiID].E[r.EvaluationItemID].Values = append(
			ag[r.CaiID].E[r.EvaluationItemID].Values,
			LogValue{CookLogID: r.CookLogID, RecipeName: r.RecipeName, Value: v},
		)
	}

	// 4) 出力整形（cai の順番は ID 昇順、評価も ID 昇順）
	caiIDsSorted := make([]int, 0, len(ag))
	for id := range ag {
		caiIDsSorted = append(caiIDsSorted, int(id))
	}
	sort.Ints(caiIDsSorted)

	out := make([]CommonStepResponse, 0, len(caiIDsSorted))
	for _, id := range caiIDsSorted {
		cai := ag[uint(id)]
		// 評価 ID 並べ替え
		sort.Slice(cai.Keys, func(i, j int) bool { return cai.Keys[i] < cai.Keys[j] })
		evals := make([]EvaluationWithValues, 0, len(cai.Keys))
		for _, eid := range cai.Keys {
			e := cai.E[eid]
			evals = append(evals, EvaluationWithValues{
				EvaluationItemID: eid,
				Label:            e.Label,
				Unit:             e.Unit,
				Average:          e.Average,
				Values:           e.Values,
			})
		}
		out = append(out, CommonStepResponse{
			Note:        cai.Note,
			Evaluations: evals,
		})
	}

	return out, nil
}

// 追加🟡🔴

type ValueRow struct {
	CookLogID  uint    `json:"cook_log"`
	RecipeName string  `json:"recipe_name"`
	Value      float64 `json:"value"`
}

type EvalRow struct {
	EvaluationItemID uint       `json:"evaluation_item_id"`
	Label            string     `json:"label"`
	Unit             string     `json:"unit"`
	Average          *float64   `json:"average"`
	Values           []ValueRow `json:"values"`
}

type StepRow struct {
	Note     *string `json:"note,omitempty"`
	Sequence *int    `json:"sequence,omitempty"`
	// CAI（CookActionIngredientsID）は内部で使うが返却はしない
	Evaluations []EvalRow `json:"evaluations"`
}

// 指定 cook_log 群が同一ユーザーに属するか検証
func validateCookLogsBelongToUser(db *gorm.DB, userID uint, cookLogIDs []uint) error {
	var cnt int64
	if err := db.Table("cook_logs").
		Where("user_link_id = ?", userID).
		Where("id IN ?", cookLogIDs).
		Count(&cnt).Error; err != nil {
		return err
	}
	if cnt != int64(len(cookLogIDs)) {
		return errors.New("some cook_logs do not belong to the user")
	}
	return nil
}

// 共通 CAI を抽出
func findCommonCAIByCookLogs(db *gorm.DB, cookLogIDs []uint) ([]uint, error) {
	var caiIDs []uint
	// cook_logs.id ごとにレシピが異なっても、同じ CAI を持つ step が各レシピに存在するものを抽出
	// recipe_steps を cook_logs.recipe_id にぶら下げて件数一致で共通化
	q := db.Table("recipe_steps").
		Select("recipe_steps.cook_action_ingredients_id AS cai").
		Joins("JOIN cook_logs ON cook_logs.recipe_id = recipe_steps.recipe_id").
		Where("cook_logs.id IN ?", cookLogIDs).
		Group("recipe_steps.cook_action_ingredients_id").
		Having("COUNT(DISTINCT cook_logs.id) = ?", len(cookLogIDs))

	if err := q.Pluck("cai", &caiIDs).Error; err != nil {
		return nil, err
	}
	return caiIDs, nil
}

// 代表の note/sequence（最初の cook_log のレシピから）
func fetchRepresentativeNoteSeq(db *gorm.DB, firstCookLogID uint, cai uint) (*string, *int, error) {
	var row struct {
		Note     *string
		Sequence *int
	}
	err := db.Table("recipe_steps").
		Select("recipe_steps.note, recipe_steps.sequence").
		Joins("JOIN cook_logs ON cook_logs.recipe_id = recipe_steps.recipe_id").
		Where("cook_logs.id = ?", firstCookLogID).
		Where("recipe_steps.cook_action_ingredients_id = ?", cai).
		// 1件あればOK。複数あれば sequence が若いもの優先にしておく
		Order("recipe_steps.sequence ASC").
		Limit(1).
		Scan(&row).Error
	if err != nil {
		return nil, nil, err
	}
	return row.Note, row.Sequence, nil
}

// ある CAI のすべての評価項目を取得
func fetchEvaluationItemsByCAI(db *gorm.DB, cai uint) ([]model.EvaluationItem, error) {
	var items []model.EvaluationItem
	if err := db.Where("cook_action_ingredients_id = ?", cai).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// 指定 cook_log 群 + CAI + 評価項目ID で値をまとめて取る
func fetchValuesForEval(db *gorm.DB, cookLogIDs []uint, cai uint, evaluationItemID uint) ([]ValueRow, error) {
	var rows []ValueRow
	err := db.Table("feature_data_links AS f").
		Select("f.cook_log_id, r.recipe_name, f.data AS value").
		Joins("JOIN recipe_steps rs ON rs.id = f.recipe_step_id").
		Joins("JOIN cook_logs cl ON cl.id = f.cook_log_id").
		Joins("JOIN recipes r ON r.id = cl.recipe_id").
		Where("f.cook_log_id IN ?", cookLogIDs).
		Where("rs.cook_action_ingredients_id = ?", cai).
		Where("f.evaluation_item_id = ?", evaluationItemID).
		Order("f.cook_log_id ASC").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// 公開API用：/cook_log/common_steps の実体
func FindCommonStepsWithValuesByCookLogs(userID uint, cookLogIDs []uint) ([]StepRow, error) {
	// 1) 所有者チェック
	if err := validateCookLogsBelongToUser(db, userID, cookLogIDs); err != nil {
		return nil, fmt.Errorf("ownership validation failed: %w", err)
	}

	// 2) 共通 CAI を見つける
	caiIDs, err := findCommonCAIByCookLogs(db, cookLogIDs)
	if err != nil {
		return nil, err
	}
	if len(caiIDs) == 0 {
		return []StepRow{}, nil
	}

	// 3) 各 CAI ごとに note/sequence（代表）、評価項目、値を組み立て
	firstLogID := cookLogIDs[0]
	resp := make([]StepRow, 0, len(caiIDs))

	for _, cai := range caiIDs {
		note, seq, err := fetchRepresentativeNoteSeq(db, firstLogID, cai)
		if err != nil {
			return nil, err
		}

		items, err := fetchEvaluationItemsByCAI(db, cai)
		if err != nil {
			return nil, err
		}

		evals := make([]EvalRow, 0, len(items))
		for _, it := range items {
			vals, err := fetchValuesForEval(db, cookLogIDs, cai, it.ID)
			if err != nil {
				return nil, err
			}
			evals = append(evals, EvalRow{
				EvaluationItemID: it.ID,
				Label:            it.EvaluationItemName,
				Unit:             it.Unit,
				Average:          it.AverageData,
				Values:           vals,
			})
		}

		resp = append(resp, StepRow{
			Note:        note,
			Sequence:    seq,
			Evaluations: evals,
		})
	}

	// 代表のレシピ手順順（sequence）で並べ替えたい場合はここでソート
	// sort.Slice(resp, func(i, j int) bool {
	// 	if resp[i].Sequence == nil || resp[j].Sequence == nil { return false }
	// 	return *resp[i].Sequence < *resp[j].Sequence
	// })

	return resp, nil
}

// controller.GetMasterAll
type MasterAllResponse struct {
	CookCategories        []model.CookCategory          `json:"cook_categories"`
	CookActions           []model.CookAction            `json:"cook_actions"`
	Ingredients           []model.Ingredients           `json:"ingredients"`
	CookActionIngredients []model.CookActionIngredients `json:"cook_action_ingredients"`
}

func FindMasterAll() (*MasterAllResponse, error) {
	var (
		cats  []model.CookCategory
		acts  []model.CookAction
		ingrs []model.Ingredients
		cais  []model.CookActionIngredients
	)

	// 順に取得（依存関係が無いので順不同でOK）
	if err := db.Find(&cats).Error; err != nil {
		return nil, err
	}
	if err := db.Find(&acts).Error; err != nil {
		return nil, err
	}
	if err := db.Find(&ingrs).Error; err != nil {
		return nil, err
	}
	if err := db.Find(&cais).Error; err != nil {
		return nil, err
	}

	return &MasterAllResponse{
		CookCategories:        cats,
		CookActions:           acts,
		Ingredients:           ingrs,
		CookActionIngredients: cais,
	}, nil
}

// POST /cook_log/common_action_values

type CommonActionValue struct {
	CookLogID  uint    `json:"cook_log"`
	RecipeName string  `json:"recipe_name"`
	Value      float64 `json:"value"`
}

type CommonActionEvaluation struct {
	EvaluationItemID uint                `json:"evaluation_item_id"`
	Label            string              `json:"label"`
	Unit             string              `json:"unit"`
	Average          *float64            `json:"average"`
	Values           []CommonActionValue `json:"values"`
}

type CommonActionStep struct {
	Note        *string                  `json:"note,omitempty"`
	Sequence    *int                     `json:"sequence,omitempty"`
	Evaluations []CommonActionEvaluation `json:"evaluations"`
}

func FindFeatureDataGroupedByEvaluation(userID, cai uint) ([]CommonActionStep, error) {
	var cookLogs []model.CookLog
	if err := db.Where("user_link_id = ?", userID).Find(&cookLogs).Error; err != nil {
		return nil, err
	}
	if len(cookLogs) == 0 {
		return nil, nil
	}

	var step model.RecipeStep
	if err := db.Where("cook_action_ingredients_id = ?", cai).First(&step).Error; err != nil {
		return nil, err
	}

	var evals []model.EvaluationItem
	if err := db.Where("cook_action_ingredients_id = ?", cai).Find(&evals).Error; err != nil {
		return nil, err
	}

	stepRow := CommonActionStep{
		Note:     step.Note,
		Sequence: &step.Sequence, // ← ポインタを渡す
	}

	for _, e := range evals {
		var values []CommonActionValue
		err := db.Table("feature_data_links AS f").
			Select("f.cook_log_id, r.recipe_name, f.data AS value").
			Joins("JOIN cook_logs cl ON cl.id = f.cook_log_id").
			Joins("JOIN recipes r ON r.id = cl.recipe_id").
			Joins("JOIN recipe_steps rs ON rs.id = f.recipe_step_id").
			Where("cl.user_link_id = ?", userID).
			Where("rs.cook_action_ingredients_id = ?", cai).
			Where("f.evaluation_item_id = ?", e.ID).
			Order("f.cook_log_id ASC").
			Scan(&values).Error
		if err != nil {
			return nil, err
		}

		stepRow.Evaluations = append(stepRow.Evaluations, CommonActionEvaluation{
			EvaluationItemID: e.ID,
			Label:            e.EvaluationItemName,
			Unit:             e.Unit,
			Average:          e.AverageData,
			Values:           values,
		})
	}

	return []CommonActionStep{stepRow}, nil
}
