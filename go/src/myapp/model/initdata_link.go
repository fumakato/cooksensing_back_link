package model

import "time"

// ==================== UserLink InitData ====================
var UserLink_InitData = []UserLink{
	{ID: 1, UserName: "fuma", IsIdentification: false},
}

// ==================== CookCategory InitData ====================
var CookCategory_InitData = []CookCategory{
	{ID: 1, CookCategoryName: "洗う"},
	{ID: 2, CookCategoryName: "量る"},
	{ID: 3, CookCategoryName: "切る"},
	{ID: 4, CookCategoryName: "混ぜる・こねる"},
	{ID: 5, CookCategoryName: "入れる・取り出す"},
	{ID: 6, CookCategoryName: "火を入れる"},
	{ID: 7, CookCategoryName: "盛り付け・片付け"},
}

var CookAction_InitData = []CookAction{
	{ID: 1, CookActionName: "洗う", CookCategoryID: 1},
	{ID: 2, CookActionName: "量る", CookCategoryID: 2},
	{ID: 3, CookActionName: "乱切り", CookCategoryID: 3},
	{ID: 4, CookActionName: "混ぜる", CookCategoryID: 4},
	{ID: 5, CookActionName: "入れる", CookCategoryID: 5},
	{ID: 6, CookActionName: "炒める", CookCategoryID: 6},
	{ID: 7, CookActionName: "煮込む/煮詰める/煮立てる", CookCategoryID: 6},
	{ID: 8, CookActionName: "蒸らす", CookCategoryID: 6},
	{ID: 9, CookActionName: "灰汁を取る", CookCategoryID: 5},
	{ID: 10, CookActionName: "盛り付け・片付け", CookCategoryID: 7},
	{ID: 11, CookActionName: "くし切り", CookCategoryID: 3},
	{ID: 12, CookActionName: "粗みじん切り", CookCategoryID: 3},
	{ID: 13, CookActionName: "みじん切り", CookCategoryID: 3},
	{ID: 14, CookActionName: "割る", CookCategoryID: 5}, //「卵を割る」に使用
	{ID: 15, CookActionName: "火を通す", CookCategoryID: 6},
	{ID: 16, CookActionName: "皮を剥く", CookCategoryID: 3},
}

var Ingredients_InitData = []Ingredients{
	{ID: 1, IngredientsName: "ジャガイモ"},
	{ID: 2, IngredientsName: "ニンジン"},
	{ID: 3, IngredientsName: "玉ねぎ"},
	{ID: 4, IngredientsName: "白滝"},
	{ID: 5, IngredientsName: "調味料"},
	{ID: 6, IngredientsName: "サラダ油"},
	{ID: 7, IngredientsName: "牛肉"},
	{ID: 8, IngredientsName: "バター"},
	{ID: 9, IngredientsName: "おろしニンニク"},
	{ID: 10, IngredientsName: "ルー"},
	{ID: 11, IngredientsName: "塩胡椒"},
	{ID: 12, IngredientsName: "水"},
	{ID: 13, IngredientsName: "ローリエの葉"},
	{ID: 14, IngredientsName: "複合"},
	{ID: 15, IngredientsName: "豚バラ肉"},
	{ID: 16, IngredientsName: "長ねぎ"},
	{ID: 17, IngredientsName: "生姜"},
	{ID: 18, IngredientsName: "卵"},
	{ID: 19, IngredientsName: "ご飯"},
	{ID: 20, IngredientsName: "醤油"},
	{ID: 21, IngredientsName: "塩"},
	{ID: 22, IngredientsName: "ごま油"},
}

var Recipe_InitData = []Recipe{
	{ID: 1, RecipeName: "肉じゃが"},
	{ID: 2, RecipeName: "カレー"},
	{ID: 3, RecipeName: "チャーハン"},
}

var CookActionIngredients_InitData = []CookActionIngredients{
	// ==== 肉じゃが ====
	{ID: 1, CookActionID: 3, IngredientsID: 1},   // ジャガイモを乱切り
	{ID: 2, CookActionID: 3, IngredientsID: 2},   // ニンジンを乱切り
	{ID: 3, CookActionID: 11, IngredientsID: 3},  // 玉ねぎをくし切り
	{ID: 4, CookActionID: 1, IngredientsID: 4},   // 白滝を洗う
	{ID: 5, CookActionID: 3, IngredientsID: 4},   // 白滝を切る
	{ID: 6, CookActionID: 2, IngredientsID: 5},   // 調味料を測る
	{ID: 7, CookActionID: 5, IngredientsID: 6},   // 鍋にサラダ油を入れる
	{ID: 8, CookActionID: 5, IngredientsID: 7},   // 肉を入れる
	{ID: 9, CookActionID: 6, IngredientsID: 7},   // 牛肉を炒める
	{ID: 10, CookActionID: 5, IngredientsID: 1},  // ジャガイモを入れる
	{ID: 11, CookActionID: 5, IngredientsID: 2},  // ニンジンを入れる
	{ID: 12, CookActionID: 5, IngredientsID: 3},  // 玉ねぎを入れる
	{ID: 13, CookActionID: 6, IngredientsID: 14}, // 複合を炒める（※便宜的に調味料ID）
	{ID: 14, CookActionID: 5, IngredientsID: 12}, // 水を入れる
	{ID: 15, CookActionID: 5, IngredientsID: 5},  // 調味料を入れる
	{ID: 16, CookActionID: 9, IngredientsID: 12}, // 灰汁を取る（水から）
	{ID: 17, CookActionID: 5, IngredientsID: 4},  // 白滝を入れる
	{ID: 18, CookActionID: 5, IngredientsID: 14}, // 落とし蓋をする
	{ID: 19, CookActionID: 7, IngredientsID: 14}, // 複合を煮詰める（※調味料ID）
	{ID: 20, CookActionID: 4, IngredientsID: 14}, // 複合を混ぜる（※調味料ID）
	{ID: 21, CookActionID: 7, IngredientsID: 14}, // 複合を煮詰める（再度）
	{ID: 22, CookActionID: 8, IngredientsID: 14}, // 複合を蒸らす（※調味料ID）

	// ==== カレー ====
	// {CookActionID:3, IngredientsID:2} → 肉じゃがで既に使用済み（ID:2）、削除
	// {CookActionID:11, IngredientsID:3} → 肉じゃがで既に使用済み（ID:3）、削除
	// {CookActionID:2, IngredientsID:5} → 肉じゃがで既に使用済み（ID:6）、削除
	{ID: 23, CookActionID: 5, IngredientsID: 8}, // バターを入れる
	{ID: 24, CookActionID: 5, IngredientsID: 9}, // おろしニンニクを入れる
	// {CookActionID:5, IngredientsID:3} → 玉ねぎを入れる（肉じゃがで既に使用済み、ID:12）、削除
	{ID: 25, CookActionID: 6, IngredientsID: 3}, // 玉ねぎを炒める
	{ID: 26, CookActionID: 5, IngredientsID: 3}, // 玉ねぎを取り出す（同じ組み合わせでも意味が違うため残す）
	// {CookActionID:5, IngredientsID:6} → サラダ油を入れる（肉じゃがで既に使用済み、ID:7）、削除
	// {CookActionID:5, IngredientsID:7} → 牛肉を入れる（肉じゃがで既に使用済み、ID:8）、削除
	// {CookActionID:6, IngredientsID:7} → 牛肉を炒める（肉じゃがで既に使用済み、ID:9）、削除
	{ID: 27, CookActionID: 5, IngredientsID: 11}, // 塩胡椒を入れる
	// {CookActionID:5, IngredientsID:2} → ニンジンを入れる（肉じゃがで既に使用済み、ID:11）、削除
	// {CookActionID:5, IngredientsID:3} → 玉ねぎを入れる（肉じゃがで既に使用済み、ID:12）、削除
	// {CookActionID:6, IngredientsID:5} → 複合を炒める（肉じゃがで既に使用済み、ID:13）、削除
	// {CookActionID:5, IngredientsID:12} → 水を入れる（肉じゃがで既に使用済み、ID:14）、削除
	{ID: 28, CookActionID: 7, IngredientsID: 12}, // 水を煮立てる
	// {CookActionID:9, IngredientsID:12} → 灰汁を取る（水から）（肉じゃがで既に使用済み、ID:16）、削除
	{ID: 29, CookActionID: 5, IngredientsID: 13}, // ローリエの葉を入れる
	{ID: 30, CookActionID: 7, IngredientsID: 14}, // 複合を煮込む（※調味料ID）
	{ID: 31, CookActionID: 5, IngredientsID: 10}, // ルーを溶かす
	// {CookActionID:7, IngredientsID:5} → 複合を煮込む（肉じゃがで既に使用済み、ID:19,30）、削除
	// {CookActionID:5, IngredientsID:5} → 調味料を入れる（肉じゃがで既に使用済み、ID:15）、削除
	// {CookActionID:7, IngredientsID:5} → 複合を煮込む（三度目）（肉じゃがで既に使用済み、ID:19,30）、削除

	//追加
	{ID: 32, CookActionID: 10, IngredientsID: 14}, //盛り付ける

	// ==== チャーハン（新規追加）====
	{ID: 33, CookActionID: 12, IngredientsID: 15}, // 豚バラ肉を粗みじん切り
	{ID: 34, CookActionID: 5, IngredientsID: 21},  // 塩を入れる
	{ID: 35, CookActionID: 12, IngredientsID: 16}, // 長ねぎを粗みじん切り
	{ID: 36, CookActionID: 13, IngredientsID: 17}, // 生姜をみじん切り
	{ID: 37, CookActionID: 14, IngredientsID: 18}, // 卵を割る
	{ID: 38, CookActionID: 4, IngredientsID: 18},  // 卵を混ぜる
	{ID: 39, CookActionID: 5, IngredientsID: 15},  // 豚バラ肉を入れる
	{ID: 40, CookActionID: 5, IngredientsID: 16},  // 長ねぎを入れる
	{ID: 41, CookActionID: 5, IngredientsID: 17},  // 生姜を入れる
	{ID: 42, CookActionID: 15, IngredientsID: 14}, // 複合に火を通す
	{ID: 43, CookActionID: 5, IngredientsID: 18},  // 卵を入れる
	{ID: 44, CookActionID: 5, IngredientsID: 19},  // ご飯を入れる
	{ID: 45, CookActionID: 5, IngredientsID: 20},  // 醤油を入れる
	{ID: 46, CookActionID: 5, IngredientsID: 22},  // ごま油を入れる
	// 既存再利用（定義済みID）
	// 調味料を測る: ID 6 (CookActionID:2, IngredientsID:5)
	// サラダ油を入れる: ID 7 (5,6)
	// 複合を炒める: ID 13 (6,14)
	// 調味料を入れる: ID 15 (5,5)
	// 複合を混ぜる: ID 20 (4,14)
	// 複合を盛り付ける: ID 32 (10,14)

	//さらに追加
	{ID: 47, CookActionID: 1, IngredientsID: 1},   // ジャガイモを洗う
	{ID: 48, CookActionID: 16, IngredientsID: 1},  // ジャガイモの皮を剥く
	{ID: 49, CookActionID: 1, IngredientsID: 2},   // ニンジンを洗う
	{ID: 50, CookActionID: 16, IngredientsID: 3},  // 玉ねぎの皮を剥く
	{ID: 51, CookActionID: 1, IngredientsID: 16},  // 長ねぎのを洗う
	{ID: 52, CookActionID: 16, IngredientsID: 17}, // 生姜の皮を剥く

}

func ptr(s string) *string { return &s }

var RecipeStep_InitData = []RecipeStep{
	// ==== 肉じゃが ====
	{ID: 1, RecipeID: 1, CookActionIngredientsID: 1, Sequence: 1, Note: ptr("ジャガイモを乱切りする")},
	{ID: 2, RecipeID: 1, CookActionIngredientsID: 2, Sequence: 2, Note: ptr("ニンジンを乱切りする")},
	{ID: 3, RecipeID: 1, CookActionIngredientsID: 3, Sequence: 3, Note: ptr("玉ねぎをくし切りする")},
	{ID: 4, RecipeID: 1, CookActionIngredientsID: 4, Sequence: 4, Note: ptr("白滝を洗う")},
	{ID: 5, RecipeID: 1, CookActionIngredientsID: 5, Sequence: 5, Note: ptr("白滝を切る")},
	{ID: 6, RecipeID: 1, CookActionIngredientsID: 6, Sequence: 6, Note: ptr("調味料を量る")},
	{ID: 7, RecipeID: 1, CookActionIngredientsID: 7, Sequence: 7, Note: ptr("サラダ油を入れる")},
	{ID: 8, RecipeID: 1, CookActionIngredientsID: 8, Sequence: 8, Note: ptr("牛肉を入れる")},
	{ID: 9, RecipeID: 1, CookActionIngredientsID: 9, Sequence: 9, Note: ptr("牛肉を炒める")},
	{ID: 10, RecipeID: 1, CookActionIngredientsID: 10, Sequence: 10, Note: ptr("ジャガイモを入れる")},
	{ID: 11, RecipeID: 1, CookActionIngredientsID: 11, Sequence: 11, Note: ptr("ニンジンを入れる")},
	{ID: 12, RecipeID: 1, CookActionIngredientsID: 12, Sequence: 12, Note: ptr("玉ねぎを入れる")},
	{ID: 13, RecipeID: 1, CookActionIngredientsID: 13, Sequence: 13, Note: ptr("複合を炒める")},
	{ID: 14, RecipeID: 1, CookActionIngredientsID: 14, Sequence: 14, Note: ptr("水を入れる")},
	{ID: 15, RecipeID: 1, CookActionIngredientsID: 15, Sequence: 15, Note: ptr("調味料を入れる")},
	{ID: 16, RecipeID: 1, CookActionIngredientsID: 16, Sequence: 16, Note: ptr("灰汁を取る")},
	{ID: 17, RecipeID: 1, CookActionIngredientsID: 17, Sequence: 17, Note: ptr("白滝を入れる")},
	{ID: 18, RecipeID: 1, CookActionIngredientsID: 18, Sequence: 18, Note: ptr("落とし蓋をする")},
	{ID: 19, RecipeID: 1, CookActionIngredientsID: 19, Sequence: 19, Note: ptr("複合を煮詰める")},
	{ID: 20, RecipeID: 1, CookActionIngredientsID: 20, Sequence: 20, Note: ptr("複合を混ぜる")},
	{ID: 21, RecipeID: 1, CookActionIngredientsID: 21, Sequence: 21, Note: ptr("複合を煮詰める")},
	{ID: 22, RecipeID: 1, CookActionIngredientsID: 22, Sequence: 22, Note: ptr("複合を蒸らす")},
	{ID: 23, RecipeID: 1, CookActionIngredientsID: 32, Sequence: 23, Note: ptr("複合を盛り付ける")},

	// ==== カレー ====
	{ID: 24, RecipeID: 2, CookActionIngredientsID: 2, Sequence: 1, Note: ptr("ニンジンを乱切りする")},
	{ID: 25, RecipeID: 2, CookActionIngredientsID: 3, Sequence: 2, Note: ptr("玉ねぎをくし切りする")},
	{ID: 26, RecipeID: 2, CookActionIngredientsID: 6, Sequence: 3, Note: ptr("調味料を量る")},
	{ID: 27, RecipeID: 2, CookActionIngredientsID: 23, Sequence: 4, Note: ptr("バターを入れる")},
	{ID: 28, RecipeID: 2, CookActionIngredientsID: 24, Sequence: 5, Note: ptr("おろしニンニクを入れる")},
	{ID: 29, RecipeID: 2, CookActionIngredientsID: 12, Sequence: 6, Note: ptr("玉ねぎを入れる")},
	{ID: 30, RecipeID: 2, CookActionIngredientsID: 25, Sequence: 7, Note: ptr("玉ねぎを炒める")},
	{ID: 31, RecipeID: 2, CookActionIngredientsID: 26, Sequence: 8, Note: ptr("玉ねぎを取り出す")},
	{ID: 32, RecipeID: 2, CookActionIngredientsID: 7, Sequence: 9, Note: ptr("サラダ油を入れる")},
	{ID: 33, RecipeID: 2, CookActionIngredientsID: 8, Sequence: 10, Note: ptr("牛肉を入れる")},
	{ID: 34, RecipeID: 2, CookActionIngredientsID: 9, Sequence: 11, Note: ptr("牛肉を炒める")},
	{ID: 35, RecipeID: 2, CookActionIngredientsID: 27, Sequence: 12, Note: ptr("塩胡椒を入れる")},
	{ID: 36, RecipeID: 2, CookActionIngredientsID: 11, Sequence: 13, Note: ptr("ニンジンを入れる")},
	{ID: 37, RecipeID: 2, CookActionIngredientsID: 12, Sequence: 14, Note: ptr("玉ねぎを入れる")},
	{ID: 38, RecipeID: 2, CookActionIngredientsID: 13, Sequence: 15, Note: ptr("複合を炒める")},
	{ID: 39, RecipeID: 2, CookActionIngredientsID: 14, Sequence: 16, Note: ptr("水を入れる")},
	{ID: 40, RecipeID: 2, CookActionIngredientsID: 28, Sequence: 17, Note: ptr("水を煮立てる")},
	{ID: 41, RecipeID: 2, CookActionIngredientsID: 16, Sequence: 18, Note: ptr("灰汁を取る")},
	{ID: 42, RecipeID: 2, CookActionIngredientsID: 29, Sequence: 19, Note: ptr("ローリエの葉を入れる")},
	{ID: 43, RecipeID: 2, CookActionIngredientsID: 30, Sequence: 20, Note: ptr("複合を煮込む")},
	{ID: 44, RecipeID: 2, CookActionIngredientsID: 31, Sequence: 21, Note: ptr("ルーを溶かす")},
	{ID: 45, RecipeID: 2, CookActionIngredientsID: 30, Sequence: 22, Note: ptr("複合を煮込む")},
	{ID: 46, RecipeID: 2, CookActionIngredientsID: 15, Sequence: 23, Note: ptr("調味料を入れる")},
	{ID: 47, RecipeID: 2, CookActionIngredientsID: 30, Sequence: 24, Note: ptr("複合を煮込む")},
	{ID: 48, RecipeID: 2, CookActionIngredientsID: 32, Sequence: 25, Note: ptr("複合を盛り付ける")}, // 盛り付け・片付けの代用

	// ==== チャーハン ====
	{ID: 49, RecipeID: 3, CookActionIngredientsID: 33, Sequence: 1, Note: ptr("豚バラ肉を粗みじん切りする")},
	{ID: 50, RecipeID: 3, CookActionIngredientsID: 34, Sequence: 2, Note: ptr("塩を入れる")},
	{ID: 51, RecipeID: 3, CookActionIngredientsID: 35, Sequence: 3, Note: ptr("長ねぎを粗みじん切りする")},
	{ID: 52, RecipeID: 3, CookActionIngredientsID: 36, Sequence: 4, Note: ptr("生姜をみじん切りする")},
	{ID: 53, RecipeID: 3, CookActionIngredientsID: 37, Sequence: 5, Note: ptr("卵を割る")},
	{ID: 54, RecipeID: 3, CookActionIngredientsID: 38, Sequence: 6, Note: ptr("卵を混ぜる")},
	{ID: 55, RecipeID: 3, CookActionIngredientsID: 6, Sequence: 7, Note: ptr("調味料を量る")}, // 既存
	{ID: 56, RecipeID: 3, CookActionIngredientsID: 39, Sequence: 8, Note: ptr("豚バラ肉を入れる")},
	{ID: 57, RecipeID: 3, CookActionIngredientsID: 40, Sequence: 9, Note: ptr("長ねぎを入れる")},
	{ID: 58, RecipeID: 3, CookActionIngredientsID: 13, Sequence: 10, Note: ptr("複合を炒める")},  // 既存
	{ID: 59, RecipeID: 3, CookActionIngredientsID: 7, Sequence: 11, Note: ptr("サラダ油を入れる")}, // 既存
	{ID: 60, RecipeID: 3, CookActionIngredientsID: 41, Sequence: 12, Note: ptr("生姜を入れる")},
	{ID: 61, RecipeID: 3, CookActionIngredientsID: 42, Sequence: 13, Note: ptr("複合に火を通す")},
	{ID: 62, RecipeID: 3, CookActionIngredientsID: 43, Sequence: 14, Note: ptr("卵を入れる")},
	{ID: 63, RecipeID: 3, CookActionIngredientsID: 44, Sequence: 15, Note: ptr("ご飯を入れる")},
	{ID: 64, RecipeID: 3, CookActionIngredientsID: 20, Sequence: 16, Note: ptr("複合を混ぜる")},  // 既存
	{ID: 65, RecipeID: 3, CookActionIngredientsID: 15, Sequence: 17, Note: ptr("調味料を入れる")}, // 既存
	{ID: 66, RecipeID: 3, CookActionIngredientsID: 13, Sequence: 18, Note: ptr("複合を炒める")},  // 既存
	{ID: 67, RecipeID: 3, CookActionIngredientsID: 40, Sequence: 19, Note: ptr("長ねぎを入れる")},
	{ID: 68, RecipeID: 3, CookActionIngredientsID: 45, Sequence: 20, Note: ptr("醤油を入れる")},
	{ID: 69, RecipeID: 3, CookActionIngredientsID: 13, Sequence: 21, Note: ptr("複合を炒める")}, // 既存
	{ID: 70, RecipeID: 3, CookActionIngredientsID: 34, Sequence: 22, Note: ptr("塩を入れる")},  // 再利用
	{ID: 71, RecipeID: 3, CookActionIngredientsID: 46, Sequence: 23, Note: ptr("ごま油を入れる")},
	{ID: 72, RecipeID: 3, CookActionIngredientsID: 20, Sequence: 24, Note: ptr("複合を混ぜる")},   // 既存
	{ID: 73, RecipeID: 3, CookActionIngredientsID: 32, Sequence: 25, Note: ptr("複合を盛り付ける")}, // 既存

}

func floatPtr(f float64) *float64 { return &f }

var EvaluationItem_InitData = []EvaluationItem{
	// ==== 肉じゃが ====
	// 1. ジャガイモを乱切り
	{ID: 1, CookActionIngredientsID: 1, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 2, CookActionIngredientsID: 1, EvaluationItemName: "安定性", Unit: "", AverageData: floatPtr(4.3)},
	{ID: 3, CookActionIngredientsID: 1, EvaluationItemName: "綺麗さ", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 4, CookActionIngredientsID: 1, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 2. ニンジンを乱切り
	{ID: 5, CookActionIngredientsID: 2, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 6, CookActionIngredientsID: 2, EvaluationItemName: "安定性", Unit: "", AverageData: floatPtr(4.1)},
	{ID: 7, CookActionIngredientsID: 2, EvaluationItemName: "綺麗さ", Unit: "点", AverageData: floatPtr(4.0)},
	{ID: 8, CookActionIngredientsID: 2, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.1)},

	// 3. 玉ねぎをくし切り
	{ID: 9, CookActionIngredientsID: 3, EvaluationItemName: "速さ", Unit: "回/s", AverageData: floatPtr(0.8)},
	{ID: 10, CookActionIngredientsID: 3, EvaluationItemName: "安定性", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 11, CookActionIngredientsID: 3, EvaluationItemName: "綺麗さ", Unit: "点", AverageData: floatPtr(4.1)},
	{ID: 12, CookActionIngredientsID: 3, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 4. 白滝を洗う
	{ID: 13, CookActionIngredientsID: 4, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.5)},
	{ID: 14, CookActionIngredientsID: 4, EvaluationItemName: "安定性", Unit: "", AverageData: floatPtr(4.5)},
	{ID: 15, CookActionIngredientsID: 4, EvaluationItemName: "綺麗さ", Unit: "点", AverageData: floatPtr(4.6)},
	{ID: 16, CookActionIngredientsID: 4, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.5)},

	// 5. 白滝を切る
	{ID: 17, CookActionIngredientsID: 5, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 18, CookActionIngredientsID: 5, EvaluationItemName: "安定性", Unit: "", AverageData: floatPtr(4.2)},
	{ID: 19, CookActionIngredientsID: 5, EvaluationItemName: "綺麗さ", Unit: "点", AverageData: floatPtr(4.3)},
	{ID: 20, CookActionIngredientsID: 5, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 6. 調味料を測る
	{ID: 21, CookActionIngredientsID: 6, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 22, CookActionIngredientsID: 6, EvaluationItemName: "手際の良さ", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 23, CookActionIngredientsID: 6, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 7. サラダ油を入れる
	{ID: 24, CookActionIngredientsID: 7, EvaluationItemName: "速さ", Unit: "ml/s", AverageData: floatPtr(2.2)},
	{ID: 25, CookActionIngredientsID: 7, EvaluationItemName: "正確性", Unit: "点", AverageData: floatPtr(3.5)},
	{ID: 26, CookActionIngredientsID: 7, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 8. 牛肉を入れる
	{ID: 27, CookActionIngredientsID: 8, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.8)},
	{ID: 28, CookActionIngredientsID: 8, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(94.0)},
	{ID: 29, CookActionIngredientsID: 8, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 9. 牛肉を炒める
	{ID: 30, CookActionIngredientsID: 9, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 31, CookActionIngredientsID: 9, EvaluationItemName: "時間", Unit: "秒", AverageData: floatPtr(120)},
	{ID: 32, CookActionIngredientsID: 9, EvaluationItemName: "焦げ", Unit: "点", AverageData: floatPtr(0.4)},
	{ID: 33, CookActionIngredientsID: 9, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 10. ジャガイモを入れる
	{ID: 34, CookActionIngredientsID: 10, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 35, CookActionIngredientsID: 10, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 36, CookActionIngredientsID: 10, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 11. ニンジンを入れる
	{ID: 37, CookActionIngredientsID: 11, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 38, CookActionIngredientsID: 11, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(94.0)},
	{ID: 39, CookActionIngredientsID: 11, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 12. 玉ねぎを入れる
	{ID: 40, CookActionIngredientsID: 12, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.8)},
	{ID: 41, CookActionIngredientsID: 12, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(93.0)},
	{ID: 42, CookActionIngredientsID: 12, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.1)},

	// 13. 複合を炒める
	{ID: 43, CookActionIngredientsID: 13, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 44, CookActionIngredientsID: 13, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.0)},
	{ID: 45, CookActionIngredientsID: 13, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 14. 水を入れる
	{ID: 46, CookActionIngredientsID: 14, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.9)},
	{ID: 47, CookActionIngredientsID: 14, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(97.0)},
	{ID: 48, CookActionIngredientsID: 14, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.5)},

	// 15. 調味料を入れる
	{ID: 49, CookActionIngredientsID: 15, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 50, CookActionIngredientsID: 15, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 51, CookActionIngredientsID: 15, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 16. 灰汁を取る
	{ID: 52, CookActionIngredientsID: 16, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.5)},
	{ID: 53, CookActionIngredientsID: 16, EvaluationItemName: "丁寧さ", Unit: "点", AverageData: floatPtr(4.5)},
	{ID: 54, CookActionIngredientsID: 16, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 17. 白滝を入れる
	{ID: 55, CookActionIngredientsID: 17, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 56, CookActionIngredientsID: 17, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(94.0)},
	{ID: 57, CookActionIngredientsID: 17, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 18. 落とし蓋をする
	{ID: 58, CookActionIngredientsID: 18, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 59, CookActionIngredientsID: 18, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 60, CookActionIngredientsID: 18, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 19. 複合を煮詰める
	{ID: 61, CookActionIngredientsID: 19, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.5)},
	{ID: 62, CookActionIngredientsID: 19, EvaluationItemName: "火加減", Unit: "点", AverageData: floatPtr(4.1)},
	{ID: 63, CookActionIngredientsID: 19, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 20. 複合を混ぜる
	{ID: 64, CookActionIngredientsID: 20, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 65, CookActionIngredientsID: 20, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.3)},
	{ID: 66, CookActionIngredientsID: 20, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 21. 複合を煮詰める（再度）
	{ID: 67, CookActionIngredientsID: 21, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 68, CookActionIngredientsID: 21, EvaluationItemName: "火加減", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 69, CookActionIngredientsID: 21, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 22. 複合を蒸らす
	{ID: 70, CookActionIngredientsID: 22, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.5)},
	{ID: 71, CookActionIngredientsID: 22, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.4)},
	{ID: 72, CookActionIngredientsID: 22, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// ==== カレー ====
	// 23. バターを入れる
	{ID: 73, CookActionIngredientsID: 23, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 74, CookActionIngredientsID: 23, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 75, CookActionIngredientsID: 23, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 24. おろしニンニクを入れる
	{ID: 76, CookActionIngredientsID: 24, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 77, CookActionIngredientsID: 24, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 78, CookActionIngredientsID: 24, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 25. 玉ねぎを炒める
	{ID: 79, CookActionIngredientsID: 25, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 80, CookActionIngredientsID: 25, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.0)},
	{ID: 81, CookActionIngredientsID: 25, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 26. 玉ねぎを取り出す
	{ID: 82, CookActionIngredientsID: 26, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.8)},
	{ID: 83, CookActionIngredientsID: 26, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(94.0)},
	{ID: 84, CookActionIngredientsID: 26, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 27. 塩胡椒を入れる
	{ID: 85, CookActionIngredientsID: 27, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.9)},
	{ID: 86, CookActionIngredientsID: 27, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(96.0)},
	{ID: 87, CookActionIngredientsID: 27, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.5)},

	// 28. 水を煮立てる
	{ID: 88, CookActionIngredientsID: 28, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 89, CookActionIngredientsID: 28, EvaluationItemName: "火加減", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 90, CookActionIngredientsID: 28, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 29. ローリエの葉を入れる
	{ID: 91, CookActionIngredientsID: 29, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.8)},
	{ID: 92, CookActionIngredientsID: 29, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 93, CookActionIngredientsID: 29, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 30. 複合を煮込む
	{ID: 94, CookActionIngredientsID: 30, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.5)},
	{ID: 95, CookActionIngredientsID: 30, EvaluationItemName: "火加減", Unit: "点", AverageData: floatPtr(4.0)},
	{ID: 96, CookActionIngredientsID: 30, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 31. ルーを溶かす
	{ID: 97, CookActionIngredientsID: 31, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 98, CookActionIngredientsID: 31, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(94.0)},
	{ID: 99, CookActionIngredientsID: 31, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 32. 複合を盛り付ける
	{ID: 100, CookActionIngredientsID: 32, EvaluationItemName: "正確性", Unit: "点", AverageData: floatPtr(4.4)},
	{ID: 101, CookActionIngredientsID: 32, EvaluationItemName: "見栄え", Unit: "点", AverageData: floatPtr(4.5)},
	{ID: 102, CookActionIngredientsID: 32, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.6)},

	// ==== チャーハン（EvaluationItem：新規追加）====
	// 既存の最大ID=102 の続きから採番

	// 33: 豚バラ肉を粗みじん切り
	{ID: 103, CookActionIngredientsID: 33, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 104, CookActionIngredientsID: 33, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.1)},
	{ID: 105, CookActionIngredientsID: 33, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 34: 塩を入れる
	{ID: 106, CookActionIngredientsID: 34, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(96.0)},
	{ID: 107, CookActionIngredientsID: 34, EvaluationItemName: "手際の良さ", Unit: "点", AverageData: floatPtr(4.3)},
	{ID: 108, CookActionIngredientsID: 34, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 35: 長ねぎを粗みじん切り
	{ID: 109, CookActionIngredientsID: 35, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.65)},
	{ID: 110, CookActionIngredientsID: 35, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 111, CookActionIngredientsID: 35, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 36: 生姜をみじん切り
	{ID: 112, CookActionIngredientsID: 36, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.6)},
	{ID: 113, CookActionIngredientsID: 36, EvaluationItemName: "細かさ", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 114, CookActionIngredientsID: 36, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 37: 卵を割る
	{ID: 115, CookActionIngredientsID: 37, EvaluationItemName: "殻混入率", Unit: "%", AverageData: floatPtr(1.5)}, // 低いほど良い
	{ID: 116, CookActionIngredientsID: 37, EvaluationItemName: "手際の良さ", Unit: "点", AverageData: floatPtr(4.3)},
	{ID: 117, CookActionIngredientsID: 37, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 38: 卵を混ぜる
	{ID: 118, CookActionIngredientsID: 38, EvaluationItemName: "均一性", Unit: "点", AverageData: floatPtr(4.4)},
	{ID: 119, CookActionIngredientsID: 38, EvaluationItemName: "泡立ち", Unit: "点", AverageData: floatPtr(4.0)},
	{ID: 120, CookActionIngredientsID: 38, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 39: 豚バラ肉を入れる
	{ID: 121, CookActionIngredientsID: 39, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 122, CookActionIngredientsID: 39, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 123, CookActionIngredientsID: 39, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 40: 長ねぎを入れる
	{ID: 124, CookActionIngredientsID: 40, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 125, CookActionIngredientsID: 40, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 126, CookActionIngredientsID: 40, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.2)},

	// 41: 生姜を入れる
	{ID: 127, CookActionIngredientsID: 41, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 128, CookActionIngredientsID: 41, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 129, CookActionIngredientsID: 41, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 42: 複合に火を通す
	{ID: 130, CookActionIngredientsID: 42, EvaluationItemName: "火加減", Unit: "点", AverageData: floatPtr(4.2)},
	{ID: 131, CookActionIngredientsID: 42, EvaluationItemName: "加熱ムラ", Unit: "点", AverageData: floatPtr(4.1)},
	{ID: 132, CookActionIngredientsID: 42, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 43: 卵を入れる
	{ID: 133, CookActionIngredientsID: 43, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 134, CookActionIngredientsID: 43, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 135, CookActionIngredientsID: 43, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 44: ご飯を入れる
	{ID: 136, CookActionIngredientsID: 44, EvaluationItemName: "速さ", Unit: "m/s", AverageData: floatPtr(0.7)},
	{ID: 137, CookActionIngredientsID: 44, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 138, CookActionIngredientsID: 44, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.3)},

	// 45: 醤油を入れる
	{ID: 139, CookActionIngredientsID: 45, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(95.0)},
	{ID: 140, CookActionIngredientsID: 45, EvaluationItemName: "香り", Unit: "点", AverageData: floatPtr(4.4)},
	{ID: 141, CookActionIngredientsID: 45, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.4)},

	// 46: ごま油を入れる
	{ID: 142, CookActionIngredientsID: 46, EvaluationItemName: "正確性", Unit: "%", AverageData: floatPtr(96.0)},
	{ID: 143, CookActionIngredientsID: 46, EvaluationItemName: "香り", Unit: "点", AverageData: floatPtr(4.5)},
	{ID: 144, CookActionIngredientsID: 46, EvaluationItemName: "総合評価", Unit: "点", AverageData: floatPtr(4.5)},
}

// ==================== CookLog ====================
var CookLog_InitData = []CookLog{
	{
		ID:         1,
		UserLinkID: 1, // fuma
		RecipeID:   1, // 肉じゃが
		CookedAt:   time.Date(2025, 9, 15, 18, 30, 0, 0, time.Local),
	},
	{
		ID:         2,
		UserLinkID: 1, // fuma
		RecipeID:   2, // カレー
		CookedAt:   time.Date(2025, 9, 18, 19, 00, 0, 0, time.Local),
	},
	{
		ID:         3,
		UserLinkID: 1, // fuma
		RecipeID:   3, // チャーハン
		CookedAt:   time.Date(2025, 10, 20, 18, 00, 0, 0, time.Local),
	},
}

var FeatureDataLink_InitData = []FeatureDataLink{
	// 1. ジャガイモを乱切り (RecipeStepID:1)
	{ID: 1, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 1, Data: 0.63},
	{ID: 2, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 2, Data: 4.042},
	{ID: 3, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 3, Data: 4.536},
	{ID: 4, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 4, Data: 4.257},

	// 2. ニンジンを乱切り (RecipeStepID:2)
	{ID: 5, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 5, Data: 0.714},
	{ID: 6, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 6, Data: 3.977},
	{ID: 7, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 7, Data: 4.4},
	{ID: 8, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 8, Data: 3.813},

	// 3. 玉ねぎをくし切り (RecipeStepID:3)
	{ID: 9, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 9, Data: 0.848},
	{ID: 10, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 10, Data: 3.822},
	{ID: 11, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 11, Data: 4.305},
	{ID: 12, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 12, Data: 3.948},

	// 4. 白滝を洗う (RecipeStepID:4)
	{ID: 13, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 13, Data: 0.54},
	{ID: 14, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 14, Data: 4.455},
	{ID: 15, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 15, Data: 4.692},
	{ID: 16, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 16, Data: 4.365},

	// 5. 白滝を切る (RecipeStepID:5)
	{ID: 17, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 17, Data: 0.66},
	{ID: 18, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 18, Data: 3.906},
	{ID: 19, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 19, Data: 4.558},
	{ID: 20, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 20, Data: 3.913},

	// 6. 調味料を測る (RecipeStepID:6)
	{ID: 21, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 21, Data: 99.75},
	{ID: 22, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 22, Data: 3.948},
	{ID: 23, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 23, Data: 4.644},

	// 7. サラダ油を入れる (RecipeStepID:7)
	{ID: 24, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 24, Data: 2.12},
	{ID: 25, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 25, Data: 3.214},
	{ID: 26, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 26, Data: 4.268},

	// 8. 牛肉を入れる (RecipeStepID:8)
	{ID: 27, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 27, Data: 0.88},
	{ID: 28, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 28, Data: 87.42},
	{ID: 29, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 29, Data: 4.452},

	// 9. 牛肉を炒める (RecipeStepID:9)
	{ID: 30, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 30, Data: 0.546},
	{ID: 31, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 31, Data: 126},
	{ID: 32, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 32, Data: 0.376},
	{ID: 33, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 33, Data: 4.644},

	// 10. ジャガイモを入れる (RecipeStepID:10)
	{ID: 34, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 34, Data: 0.693},
	{ID: 35, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 35, Data: 96.9},
	{ID: 36, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 36, Data: 4.171},

	// 11. ニンジンを入れる (RecipeStepID:11)
	{ID: 37, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 37, Data: 0.77},
	{ID: 38, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 38, Data: 87.42},
	{ID: 39, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 39, Data: 4.452},

	// 12. 玉ねぎを入れる (RecipeStepID:12)
	{ID: 40, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 40, Data: 0.728},
	{ID: 41, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 41, Data: 97.65},
	{ID: 42, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 42, Data: 3.854},

	// 13. 複合を炒める (RecipeStepID:13)
	{ID: 43, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 43, Data: 0.648},
	{ID: 44, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 44, Data: 3.96},
	{ID: 45, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 45, Data: 4.284},

	// 14. 水を入れる (RecipeStepID:14)
	{ID: 46, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 46, Data: 0.873},
	{ID: 47, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 47, Data: 106.7},
	{ID: 48, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 48, Data: 4.185},

	// 15. 調味料を入れる (RecipeStepID:15)
	{ID: 49, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 49, Data: 0.742},
	{ID: 50, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 50, Data: 86.45},
	{ID: 51, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 51, Data: 4.515},

	// 16. 灰汁を取る (RecipeStepID:16)
	{ID: 52, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 52, Data: 0.47},
	{ID: 53, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 53, Data: 4.86},
	{ID: 54, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 54, Data: 4.356},

	// 17. 白滝を入れる (RecipeStepID:17)
	{ID: 55, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 55, Data: 0.612},
	{ID: 56, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 56, Data: 91.18},
	{ID: 57, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 57, Data: 4.62},

	// 18. 落とし蓋をする (RecipeStepID:18)
	{ID: 58, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 58, Data: 0.558},
	{ID: 59, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 59, Data: 100.7},
	{ID: 60, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 60, Data: 3.913},

	// 19. 複合を煮詰める (RecipeStepID:19)
	{ID: 61, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 61, Data: 0.525},
	{ID: 62, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 62, Data: 3.854},
	{ID: 63, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 63, Data: 4.536},

	// 20. 複合を混ぜる (RecipeStepID:20)
	{ID: 64, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 64, Data: 0.693},
	{ID: 65, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 65, Data: 4.386},
	{ID: 66, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 66, Data: 4.268},

	// 21. 複合を煮詰める（再度） (RecipeStepID:21)
	{ID: 67, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 67, Data: 0.66},
	{ID: 68, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 68, Data: 3.906},
	{ID: 69, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 69, Data: 4.558},

	// 22. 複合を蒸らす (RecipeStepID:22)
	{ID: 70, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 70, Data: 0.455},
	{ID: 71, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 71, Data: 4.62},
	{ID: 72, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 72, Data: 4.136},

	// 23. 複合を盛り付ける (RecipeStepID:23)
	{ID: 73, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 100, Data: 4.32},
	{ID: 74, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 101, Data: 4.455},
	{ID: 75, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 102, Data: 4.692},

	// 24. ニンジンを乱切り (RecipeStepID:24)
	{ID: 76, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 5, Data: 0.679},
	{ID: 77, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 6, Data: 4.51},
	{ID: 78, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 7, Data: 3.72},
	{ID: 79, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 8, Data: 4.346},

	// 25. 玉ねぎをくし切り (RecipeStepID:25)
	{ID: 80, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 9, Data: 0.728},
	{ID: 81, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 10, Data: 4.41},
	{ID: 82, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 11, Data: 3.854},
	{ID: 83, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 12, Data: 4.536},

	// 26. 調味料を測る (RecipeStepID:26)
	{ID: 84, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 21, Data: 94.05},
	{ID: 85, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 22, Data: 4.284},
	{ID: 86, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 23, Data: 4.171},

	// 27. バターを入れる (RecipeStepID:27)
	{ID: 87, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 73, Data: 0.77},
	{ID: 88, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 74, Data: 88.35},
	{ID: 89, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 75, Data: 4.558},

	// 28. おろしニンニクを入れる (RecipeStepID:28)
	{ID: 90, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 76, Data: 0.637},
	{ID: 91, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 77, Data: 99.75},
	{ID: 92, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 78, Data: 4.042},

	// 29. 玉ねぎを入れる (RecipeStepID:29)
	{ID: 93, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 40, Data: 0.864},
	{ID: 94, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 41, Data: 92.07},
	{ID: 95, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 42, Data: 4.182},

	// 30. 玉ねぎを炒める (RecipeStepID:30)
	{ID: 96, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 79, Data: 0.582},
	{ID: 97, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 80, Data: 4.4},
	{ID: 98, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 81, Data: 3.906},

	// 31. 玉ねぎを取り出す (RecipeStepID:31)
	{ID: 99, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 82, Data: 0.848},
	{ID: 100, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 83, Data: 85.54},
	{ID: 101, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 84, Data: 4.515},

	// 32. サラダ油を入れる (RecipeStepID:32)
	{ID: 102, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 24, Data: 2.3},
	{ID: 103, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 25, Data: 2.563},
	{ID: 104, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 26, Data: 4.356},

	// 33. 牛肉を入れる (RecipeStepID:33)
	{ID: 105, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 27, Data: 0.816},
	{ID: 106, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 28, Data: 91.18},
	{ID: 107, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 29, Data: 4.62},

	// 34. 牛肉を炒める (RecipeStepID:34)
	{ID: 108, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 30, Data: 0.558},
	{ID: 109, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 31, Data: 127.2},
	{ID: 110, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 32, Data: 0.364},
	{ID: 111, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 33, Data: 4.515},

	// 35. 塩胡椒を入れる (RecipeStepID:35)
	{ID: 112, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 86, Data: 90.24},
	{ID: 113, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 87, Data: 4.86},
	{ID: 114, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 88, Data: 0.594},

	// 36. ニンジンを入れる (RecipeStepID:36)
	{ID: 115, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 37, Data: 0.714},
	{ID: 116, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 38, Data: 91.18},
	{ID: 117, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 39, Data: 4.62},

	// 37. 玉ねぎを入れる (RecipeStepID:37)
	{ID: 118, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 40, Data: 0.744},
	{ID: 119, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 41, Data: 98.58},
	{ID: 120, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 42, Data: 3.731},

	// 38. 複合を炒める (RecipeStepID:38)
	{ID: 121, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 43, Data: 0.63},
	{ID: 122, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 44, Data: 3.76},
	{ID: 123, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 45, Data: 4.536},

	// 39. 水を入れる (RecipeStepID:39)
	{ID: 124, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 46, Data: 0.891},
	{ID: 125, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 47, Data: 98.94},
	{ID: 126, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 48, Data: 4.365},

	// 40. 水を煮立てる (RecipeStepID:40)
	{ID: 127, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 89, Data: 4.62},
	{ID: 128, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 90, Data: 4.092},
	{ID: 129, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 91, Data: 0.848},

	// 41. 灰汁を取る (RecipeStepID:41)
	{ID: 130, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 52, Data: 0.455},
	{ID: 131, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 53, Data: 4.725},
	{ID: 132, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 54, Data: 4.136},

	// 42. ローリエの葉を入れる (RecipeStepID:42)
	{ID: 133, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 92, Data: 102.6},
	{ID: 134, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 93, Data: 4.257},
	{ID: 135, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 94, Data: 0.51},

	// 43. 複合を煮込む (RecipeStepID:43)
	{ID: 136, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 95, Data: 3.88},
	{ID: 137, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 96, Data: 4.62},
	{ID: 138, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 97, Data: 0.651},

	// 44. ルーを溶かす (RecipeStepID:44)
	{ID: 139, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 97, Data: 0.742},
	{ID: 140, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 98, Data: 85.54},
	{ID: 141, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 99, Data: 4.62},

	// 45. 複合を煮込む (RecipeStepID:45)
	{ID: 142, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 95, Data: 3.76},
	{ID: 143, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 96, Data: 4.536},
	{ID: 144, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 97, Data: 0.693},

	// 46. 調味料を入れる (RecipeStepID:46)
	{ID: 145, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 49, Data: 0.714},
	{ID: 146, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 50, Data: 92.15},
	{ID: 147, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 51, Data: 4.73},

	// 47. 複合を煮込む (RecipeStepID:47)
	{ID: 148, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 95, Data: 3.72},
	{ID: 149, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 96, Data: 4.452},
	{ID: 150, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 97, Data: 0.637},

	// 48. 複合を盛り付ける (RecipeStepID:48)
	{ID: 151, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 100, Data: 4.58},
	{ID: 152, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 101, Data: 4.23},
	{ID: 153, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 102, Data: 4.968},

	// ==== チャーハン (RecipeStepID:49..73, CookLogID:3) ====
	// 49. 豚バラ肉を粗みじん切り (CAI:33 -> Eval:103,104,105)
	{ID: 154, CookLogID: 3, RecipeStepID: 49, EvaluationItemID: 103, Data: 0.58}, // 0.6 * 0.96
	{ID: 155, CookLogID: 3, RecipeStepID: 49, EvaluationItemID: 104, Data: 4.26}, // 4.1 * 1.04
	{ID: 156, CookLogID: 3, RecipeStepID: 49, EvaluationItemID: 105, Data: 4.28}, // 4.2 * 1.02

	// 50. 塩を入れる (CAI:34 -> Eval:106,107,108)
	{ID: 157, CookLogID: 3, RecipeStepID: 50, EvaluationItemID: 106, Data: 92.16}, // 96.0 * 0.96
	{ID: 158, CookLogID: 3, RecipeStepID: 50, EvaluationItemID: 107, Data: 4.47},  // 4.3 * 1.04
	{ID: 159, CookLogID: 3, RecipeStepID: 50, EvaluationItemID: 108, Data: 4.49},  // 4.4 * 1.02

	// 51. 長ねぎを粗みじん切り (CAI:35 -> Eval:109,110,111)
	{ID: 160, CookLogID: 3, RecipeStepID: 51, EvaluationItemID: 109, Data: 0.62}, // 0.65 * 0.96
	{ID: 161, CookLogID: 3, RecipeStepID: 51, EvaluationItemID: 110, Data: 4.37}, // 4.2 * 1.04
	{ID: 162, CookLogID: 3, RecipeStepID: 51, EvaluationItemID: 111, Data: 4.39}, // 4.3 * 1.02

	// 52. 生姜をみじん切り (CAI:36 -> Eval:112,113,114)
	{ID: 163, CookLogID: 3, RecipeStepID: 52, EvaluationItemID: 112, Data: 0.58}, // 0.6 * 0.96
	{ID: 164, CookLogID: 3, RecipeStepID: 52, EvaluationItemID: 113, Data: 4.37}, // 4.2 * 1.04
	{ID: 165, CookLogID: 3, RecipeStepID: 52, EvaluationItemID: 114, Data: 4.39}, // 4.3 * 1.02

	// 53. 卵を割る (CAI:37 -> Eval:115,116,117)
	{ID: 166, CookLogID: 3, RecipeStepID: 53, EvaluationItemID: 115, Data: 1.44}, // 1.5 * 0.96
	{ID: 167, CookLogID: 3, RecipeStepID: 53, EvaluationItemID: 116, Data: 4.47}, // 4.3 * 1.04
	{ID: 168, CookLogID: 3, RecipeStepID: 53, EvaluationItemID: 117, Data: 4.49}, // 4.4 * 1.02

	// 54. 卵を混ぜる (CAI:38 -> Eval:118,119,120)
	{ID: 169, CookLogID: 3, RecipeStepID: 54, EvaluationItemID: 118, Data: 4.22}, // 4.4 * 0.96
	{ID: 170, CookLogID: 3, RecipeStepID: 54, EvaluationItemID: 119, Data: 4.16}, // 4.0 * 1.04
	{ID: 171, CookLogID: 3, RecipeStepID: 54, EvaluationItemID: 120, Data: 4.49}, // 4.4 * 1.02

	// 55. 調味料を測る (CAI:6 -> Eval:21,22,23)
	{ID: 172, CookLogID: 3, RecipeStepID: 55, EvaluationItemID: 21, Data: 91.20}, // 95.0 * 0.96
	{ID: 173, CookLogID: 3, RecipeStepID: 55, EvaluationItemID: 22, Data: 4.37},  // 4.2 * 1.04
	{ID: 174, CookLogID: 3, RecipeStepID: 55, EvaluationItemID: 23, Data: 4.39},  // 4.3 * 1.02

	// 56. 豚バラ肉を入れる (CAI:39 -> Eval:121,122,123)
	{ID: 175, CookLogID: 3, RecipeStepID: 56, EvaluationItemID: 121, Data: 0.67},  // 0.7 * 0.96
	{ID: 176, CookLogID: 3, RecipeStepID: 56, EvaluationItemID: 122, Data: 98.80}, // 95.0 * 1.04
	{ID: 177, CookLogID: 3, RecipeStepID: 56, EvaluationItemID: 123, Data: 4.39},  // 4.3 * 1.02

	// 57. 長ねぎを入れる (CAI:40 -> Eval:124,125,126)
	{ID: 178, CookLogID: 3, RecipeStepID: 57, EvaluationItemID: 124, Data: 0.67},  // 0.7 * 0.96
	{ID: 179, CookLogID: 3, RecipeStepID: 57, EvaluationItemID: 125, Data: 98.80}, // 95.0 * 1.04
	{ID: 180, CookLogID: 3, RecipeStepID: 57, EvaluationItemID: 126, Data: 4.28},  // 4.2 * 1.02

	// 58. 複合を炒める (CAI:13 -> Eval:43,44,45)
	{ID: 181, CookLogID: 3, RecipeStepID: 58, EvaluationItemID: 43, Data: 0.58}, // 0.6 * 0.96
	{ID: 182, CookLogID: 3, RecipeStepID: 58, EvaluationItemID: 44, Data: 4.16}, // 4.0 * 1.04
	{ID: 183, CookLogID: 3, RecipeStepID: 58, EvaluationItemID: 45, Data: 4.28}, // 4.2 * 1.02

	// 59. サラダ油を入れる (CAI:7 -> Eval:24,25,26)
	{ID: 184, CookLogID: 3, RecipeStepID: 59, EvaluationItemID: 24, Data: 2.01}, // 0.7 * 0.96
	{ID: 185, CookLogID: 3, RecipeStepID: 59, EvaluationItemID: 25, Data: 3.89}, // 96.0 * 1.04
	{ID: 186, CookLogID: 3, RecipeStepID: 59, EvaluationItemID: 26, Data: 4.49}, // 4.4 * 1.02

	// 60. 生姜を入れる (CAI:41 -> Eval:127,128,129)
	{ID: 187, CookLogID: 3, RecipeStepID: 60, EvaluationItemID: 127, Data: 0.67},  // 0.7 * 0.96
	{ID: 188, CookLogID: 3, RecipeStepID: 60, EvaluationItemID: 128, Data: 98.80}, // 95.0 * 1.04
	{ID: 189, CookLogID: 3, RecipeStepID: 60, EvaluationItemID: 129, Data: 4.39},  // 4.3 * 1.02

	// 61. 複合に火を通す (CAI:42 -> Eval:130,131,132)
	{ID: 190, CookLogID: 3, RecipeStepID: 61, EvaluationItemID: 130, Data: 4.03}, // 4.2 * 0.96
	{ID: 191, CookLogID: 3, RecipeStepID: 61, EvaluationItemID: 131, Data: 4.26}, // 4.1 * 1.04
	{ID: 192, CookLogID: 3, RecipeStepID: 61, EvaluationItemID: 132, Data: 4.39}, // 4.3 * 1.02

	// 62. 卵を入れる (CAI:43 -> Eval:133,134,135)
	{ID: 193, CookLogID: 3, RecipeStepID: 62, EvaluationItemID: 133, Data: 0.67},  // 0.7 * 0.96
	{ID: 194, CookLogID: 3, RecipeStepID: 62, EvaluationItemID: 134, Data: 98.80}, // 95.0 * 1.04
	{ID: 195, CookLogID: 3, RecipeStepID: 62, EvaluationItemID: 135, Data: 4.49},  // 4.4 * 1.02

	// 63. ご飯を入れる (CAI:44 -> Eval:136,137,138)
	{ID: 196, CookLogID: 3, RecipeStepID: 63, EvaluationItemID: 136, Data: 0.67},  // 0.7 * 0.96
	{ID: 197, CookLogID: 3, RecipeStepID: 63, EvaluationItemID: 137, Data: 98.80}, // 95.0 * 1.04
	{ID: 198, CookLogID: 3, RecipeStepID: 63, EvaluationItemID: 138, Data: 4.39},  // 4.3 * 1.02

	// 64. 複合を混ぜる (CAI:20 -> Eval:64,65,66)
	{ID: 199, CookLogID: 3, RecipeStepID: 64, EvaluationItemID: 64, Data: 0.67}, // 0.7 * 0.96
	{ID: 200, CookLogID: 3, RecipeStepID: 64, EvaluationItemID: 65, Data: 4.47}, // 4.3 * 1.04
	{ID: 201, CookLogID: 3, RecipeStepID: 64, EvaluationItemID: 66, Data: 4.49}, // 4.4 * 1.02

	// 65. 調味料を入れる (CAI:15 -> Eval:49,50,51)
	{ID: 202, CookLogID: 3, RecipeStepID: 65, EvaluationItemID: 49, Data: 0.67},  // 0.7 * 0.96
	{ID: 203, CookLogID: 3, RecipeStepID: 65, EvaluationItemID: 50, Data: 98.80}, // 95.0 * 1.04
	{ID: 204, CookLogID: 3, RecipeStepID: 65, EvaluationItemID: 51, Data: 4.39},  // 4.3 * 1.02

	// 66. 複合を炒める (CAI:13 -> Eval:43,44,45)
	{ID: 205, CookLogID: 3, RecipeStepID: 66, EvaluationItemID: 43, Data: 0.58},
	{ID: 206, CookLogID: 3, RecipeStepID: 66, EvaluationItemID: 44, Data: 4.16},
	{ID: 207, CookLogID: 3, RecipeStepID: 66, EvaluationItemID: 45, Data: 4.28},

	// 67. 長ねぎを入れる (CAI:40 -> Eval:124,125,126)
	{ID: 208, CookLogID: 3, RecipeStepID: 67, EvaluationItemID: 124, Data: 0.67},
	{ID: 209, CookLogID: 3, RecipeStepID: 67, EvaluationItemID: 125, Data: 98.80},
	{ID: 210, CookLogID: 3, RecipeStepID: 67, EvaluationItemID: 126, Data: 4.28},

	// 68. 醤油を入れる (CAI:45 -> Eval:139,140,141)
	{ID: 211, CookLogID: 3, RecipeStepID: 68, EvaluationItemID: 139, Data: 91.20}, // 95.0 * 0.96
	{ID: 212, CookLogID: 3, RecipeStepID: 68, EvaluationItemID: 140, Data: 4.58},  // 4.4 * 1.04
	{ID: 213, CookLogID: 3, RecipeStepID: 68, EvaluationItemID: 141, Data: 4.49},  // 4.4 * 1.02

	// 69. 複合を炒める (CAI:13 -> Eval:43,44,45)
	{ID: 214, CookLogID: 3, RecipeStepID: 69, EvaluationItemID: 43, Data: 0.58},
	{ID: 215, CookLogID: 3, RecipeStepID: 69, EvaluationItemID: 44, Data: 4.16},
	{ID: 216, CookLogID: 3, RecipeStepID: 69, EvaluationItemID: 45, Data: 4.28},

	// 70. 塩を入れる (CAI:34 -> Eval:106,107,108)
	{ID: 217, CookLogID: 3, RecipeStepID: 70, EvaluationItemID: 106, Data: 92.16},
	{ID: 218, CookLogID: 3, RecipeStepID: 70, EvaluationItemID: 107, Data: 4.47},
	{ID: 219, CookLogID: 3, RecipeStepID: 70, EvaluationItemID: 108, Data: 4.49},

	// 71. ごま油を入れる (CAI:46 -> Eval:142,143,144)
	{ID: 220, CookLogID: 3, RecipeStepID: 71, EvaluationItemID: 142, Data: 92.16}, // 96.0 * 0.96
	{ID: 221, CookLogID: 3, RecipeStepID: 71, EvaluationItemID: 143, Data: 4.68},  // 4.5 * 1.04
	{ID: 222, CookLogID: 3, RecipeStepID: 71, EvaluationItemID: 144, Data: 4.59},  // 4.5 * 1.02

	// 72. 複合を混ぜる (CAI:20 -> Eval:64,65,66)
	{ID: 223, CookLogID: 3, RecipeStepID: 72, EvaluationItemID: 64, Data: 0.67},
	{ID: 224, CookLogID: 3, RecipeStepID: 72, EvaluationItemID: 65, Data: 4.47},
	{ID: 225, CookLogID: 3, RecipeStepID: 72, EvaluationItemID: 66, Data: 4.49},

	// 73. 複合を盛り付ける (CAI:32 -> Eval:100,101,102)
	{ID: 226, CookLogID: 3, RecipeStepID: 73, EvaluationItemID: 100, Data: 4.07}, // 0.8 * 0.96
	{ID: 227, CookLogID: 3, RecipeStepID: 73, EvaluationItemID: 101, Data: 4.68}, // 4.5 * 1.04
	{ID: 228, CookLogID: 3, RecipeStepID: 73, EvaluationItemID: 102, Data: 4.69}, // 4.6 * 1.02

}

// var FeatureDataLink_InitData = []FeatureDataLink{

// 	// RecipeStepID:1
// 	{ID: 1, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 1, Data: 0.62},
// 	{ID: 2, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 2, Data: 4.08},
// 	{ID: 3, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 3, Data: 4.54},
// 	{ID: 4, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 4, Data: 4.3},

// 	// RecipeStepID:2
// 	{ID: 5, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 5, Data: 0.64},
// 	{ID: 6, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 6, Data: 4.3},
// 	{ID: 7, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 7, Data: 3.88},
// 	{ID: 8, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 8, Data: 4.51},

// 	// RecipeStepID:3
// 	{ID: 9, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 9, Data: 0.82},
// 	{ID: 10, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 10, Data: 3.95},
// 	{ID: 11, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 11, Data: 4.39},
// 	{ID: 12, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 12, Data: 4.16},

// 	// RecipeStepID:4
// 	{ID: 13, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 13, Data: 0.46},
// 	{ID: 14, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 14, Data: 4.86},
// 	{ID: 15, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 15, Data: 4.37},
// 	{ID: 16, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 16, Data: 4.95},

// 	// RecipeStepID:5
// 	{ID: 17, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 17, Data: 0.66},
// 	{ID: 18, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 18, Data: 3.78},
// 	{ID: 19, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 19, Data: 4.42},
// 	{ID: 20, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 20, Data: 4.3},

// 	// RecipeStepID:6
// 	{ID: 21, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 21, Data: 89.3},
// 	{ID: 22, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 22, Data: 4.41},
// 	{ID: 23, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 23, Data: 4.21},

// 	// RecipeStepID:7
// 	{ID: 24, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 24, Data: 0.77},
// 	{ID: 25, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 25, Data: 103.68},
// 	{ID: 26, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 26, Data: 4.49},

// 	// RecipeStepID:8
// 	{ID: 27, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 27, Data: 0.9},
// 	{ID: 28, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 28, Data: 84.6},
// 	{ID: 29, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 29, Data: 4.49},

// 	// RecipeStepID:9
// 	{ID: 30, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 30, Data: 0.66},
// 	{ID: 31, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 31, Data: 109.2},
// 	{ID: 32, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 32, Data: 0.44},
// 	{ID: 33, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 33, Data: 4.43},

// 	// RecipeStepID:10
// 	{ID: 34, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 34, Data: 0.77},
// 	{ID: 35, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 35, Data: 103.55},
// 	{ID: 36, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 36, Data: 4.09},

// 	// RecipeStepID:11
// 	{ID: 37, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 37, Data: 0.77},
// 	{ID: 38, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 38, Data: 84.6},
// 	{ID: 39, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 39, Data: 4.62},

// 	// RecipeStepID:12
// 	{ID: 40, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 40, Data: 0.88},
// 	{ID: 41, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 41, Data: 96.9},
// 	{ID: 42, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 42, Data: 3.89},

// 	// RecipeStepID:13
// 	{ID: 43, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 43, Data: 0.66},
// 	{ID: 44, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 44, Data: 4.36},
// 	{ID: 45, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 45, Data: 4.49},

// 	// RecipeStepID:14
// 	{ID: 46, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 46, Data: 0.99},
// 	{ID: 47, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 47, Data: 106.67},
// 	{ID: 48, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 48, Data: 4.05},

// 	// RecipeStepID:15
// 	{ID: 49, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 49, Data: 0.77},
// 	{ID: 50, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 50, Data: 85.5},
// 	{ID: 51, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 51, Data: 4.73},

// 	// RecipeStepID:16
// 	{ID: 52, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 52, Data: 0.45},
// 	{ID: 53, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 53, Data: 4.27},
// 	{ID: 54, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 54, Data: 4.84},

// 	// RecipeStepID:17
// 	{ID: 55, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 55, Data: 0.66},
// 	{ID: 56, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 56, Data: 103.34},
// 	{ID: 57, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 57, Data: 3.78},

// 	// RecipeStepID:18
// 	{ID: 58, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 58, Data: 0.66},
// 	{ID: 59, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 59, Data: 85.5},
// 	{ID: 60, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 60, Data: 4.73},

// 	// RecipeStepID:19
// 	{ID: 61, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 61, Data: 0.45},
// 	{ID: 62, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 62, Data: 4.48},
// 	{ID: 63, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 63, Data: 3.78},

// 	// RecipeStepID:20
// 	{ID: 64, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 64, Data: 0.77},
// 	{ID: 65, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 65, Data: 4.04},
// 	{ID: 66, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 66, Data: 4.84},

// 	// RecipeStepID:21
// 	{ID: 67, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 67, Data: 0.66},
// 	{ID: 68, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 68, Data: 4.62},
// 	{ID: 69, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 69, Data: 3.87},

// 	// RecipeStepID:22
// 	{ID: 70, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 70, Data: 0.55},
// 	{ID: 71, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 71, Data: 4.18},
// 	{ID: 72, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 72, Data: 4.84},

// 	// RecipeStepID:23
// 	{ID: 73, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 100, Data: 0.72},
// 	{ID: 74, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 101, Data: 4.77},
// 	{ID: 75, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 102, Data: 4.14},

// 	// RecipeStepID:24
// 	{ID: 76, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 5, Data: 0.76},
// 	{ID: 77, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 6, Data: 3.89},
// 	{ID: 78, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 7, Data: 4.32},
// 	{ID: 79, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 8, Data: 3.89},

// 	// RecipeStepID:25
// 	{ID: 80, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 9, Data: 0.88},
// 	{ID: 81, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 10, Data: 4.62},
// 	{ID: 82, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 11, Data: 3.69},
// 	{ID: 83, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 12, Data: 4.62},

// 	// RecipeStepID:26
// 	{ID: 84, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 21, Data: 85.5},
// 	{ID: 85, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 22, Data: 4.41},
// 	{ID: 86, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 23, Data: 4.51},

// 	// RecipeStepID:27
// 	{ID: 87, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 73, Data: 0.63},
// 	{ID: 88, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 74, Data: 94.05},
// 	{ID: 89, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 75, Data: 4.6},

// 	// RecipeStepID:28
// 	{ID: 90, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 76, Data: 0.66},
// 	{ID: 91, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 77, Data: 85.5},
// 	{ID: 92, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 78, Data: 4.73},

// 	// RecipeStepID:29
// 	{ID: 93, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 40, Data: 0.66},
// 	{ID: 94, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 41, Data: 103.65},
// 	{ID: 95, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 42, Data: 4.31},

// 	// RecipeStepID:30
// 	{ID: 96, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 79, Data: 0.66},
// 	{ID: 97, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 80, Data: 3.8},
// 	{ID: 98, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 81, Data: 4.43},

// 	// RecipeStepID:31
// 	{ID: 99, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 82, Data: 0.88},
// 	{ID: 100, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 83, Data: 84.6},
// 	{ID: 101, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 84, Data: 4.64},

// 	// RecipeStepID:32
// 	{ID: 102, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 24, Data: 0.7},
// 	{ID: 103, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 25, Data: 103.68},
// 	{ID: 104, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 26, Data: 3.96},

// 	// RecipeStepID:33
// 	{ID: 105, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 27, Data: 0.72},
// 	{ID: 106, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 28, Data: 99.64},
// 	{ID: 107, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 29, Data: 3.78},

// 	// RecipeStepID:34
// 	{ID: 108, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 30, Data: 0.6},
// 	{ID: 109, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 31, Data: 132.0},
// 	{ID: 110, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 32, Data: 0.36},
// 	{ID: 111, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 33, Data: 4.73},

// 	// RecipeStepID:35
// 	{ID: 112, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 86, Data: 0.78},
// 	{ID: 113, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 87, Data: 84.48},
// 	{ID: 114, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 88, Data: 4.18},

// 	// RecipeStepID:36
// 	{ID: 115, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 37, Data: 0.63},
// 	{ID: 116, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 38, Data: 103.34},
// 	{ID: 117, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 39, Data: 3.78},

// 	// RecipeStepID:37
// 	{ID: 118, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 40, Data: 0.72},
// 	{ID: 119, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 41, Data: 85.26},
// 	{ID: 120, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 42, Data: 4.23},

// 	// RecipeStepID:38
// 	{ID: 121, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 43, Data: 0.66},
// 	{ID: 122, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 44, Data: 3.8},
// 	{ID: 123, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 45, Data: 4.62},

// 	// RecipeStepID:39
// 	{ID: 124, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 46, Data: 0.94},
// 	{ID: 125, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 47, Data: 94.09},
// 	{ID: 126, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 48, Data: 4.73},

// 	// RecipeStepID:40
// 	{ID: 127, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 89, Data: 0.72},
// 	{ID: 128, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 90, Data: 162.0},
// 	{ID: 129, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 91, Data: 4.13},

// 	// RecipeStepID:41
// 	{ID: 130, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 52, Data: 0.55},
// 	{ID: 131, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 53, Data: 4.95},
// 	{ID: 132, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 54, Data: 3.96},

// 	// RecipeStepID:42
// 	{ID: 133, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 92, Data: 0.55},
// 	{ID: 134, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 93, Data: 85.5},
// 	{ID: 135, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 94, Data: 3.98},

// 	// RecipeStepID:43
// 	{ID: 136, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 95, Data: 0.55},
// 	{ID: 137, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 96, Data: 4.4},
// 	{ID: 138, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 97, Data: 4.09},

// 	// RecipeStepID:44
// 	{ID: 139, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 97, Data: 4.84},
// 	{ID: 140, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 98, Data: 85.5},
// 	{ID: 141, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 99, Data: 4.0},

// 	// RecipeStepID:45
// 	{ID: 142, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 95, Data: 0.55},
// 	{ID: 143, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 96, Data: 4.4},
// 	{ID: 144, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 97, Data: 4.09},

// 	// RecipeStepID:46
// 	{ID: 145, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 49, Data: 0.63},
// 	{ID: 146, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 50, Data: 104.5},
// 	{ID: 147, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 51, Data: 3.87},

// 	// RecipeStepID:47
// 	{ID: 148, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 95, Data: 0.55},
// 	{ID: 149, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 96, Data: 4.4},
// 	{ID: 150, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 97, Data: 4.09},

// 	// RecipeStepID:48
// 	{ID: 151, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 100, Data: 0.88},
// 	{ID: 152, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 101, Data: 4.05},
// 	{ID: 153, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 102, Data: 4.73},
// }

// // // ==================== FeatureDataLink InitData (grouped by RecipeStepID, with comments) ====================
// // var FeatureDataLink_InitData = []FeatureDataLink{
// // 	// 1. ジャガイモを乱切り (RecipeStepID:1)
// // 	{ID: 1, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 1, Data: 0.6},
// // 	{ID: 2, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 2, Data: 4.3},
// // 	{ID: 3, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 3, Data: 4.2},
// // 	{ID: 4, CookLogID: 1, RecipeStepID: 1, EvaluationItemID: 4, Data: 4.3},

// // 	// 2. ニンジンを乱切り (RecipeStepID:2)
// // 	{ID: 5, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 5, Data: 0.7},
// // 	{ID: 6, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 6, Data: 4.1},
// // 	{ID: 7, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 7, Data: 4.0},
// // 	{ID: 8, CookLogID: 1, RecipeStepID: 2, EvaluationItemID: 8, Data: 4.1},

// // 	// 3. 玉ねぎをくし切り (RecipeStepID:3)
// // 	{ID: 9, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 9, Data: 0.8},
// // 	{ID: 10, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 10, Data: 4.2},
// // 	{ID: 11, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 11, Data: 4.1},
// // 	{ID: 12, CookLogID: 1, RecipeStepID: 3, EvaluationItemID: 12, Data: 4.2},

// // 	// 4. 白滝を洗う (RecipeStepID:4)
// // 	{ID: 13, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 13, Data: 0.5},
// // 	{ID: 14, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 14, Data: 4.5},
// // 	{ID: 15, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 15, Data: 4.6},
// // 	{ID: 16, CookLogID: 1, RecipeStepID: 4, EvaluationItemID: 16, Data: 4.5},

// // 	// 5. 白滝を切る (RecipeStepID:5)
// // 	{ID: 17, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 17, Data: 0.6},
// // 	{ID: 18, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 18, Data: 4.2},
// // 	{ID: 19, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 19, Data: 4.3},
// // 	{ID: 20, CookLogID: 1, RecipeStepID: 5, EvaluationItemID: 20, Data: 4.3},

// // 	// 6. 調味料を測る (RecipeStepID:6)
// // 	{ID: 21, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 21, Data: 95.0},
// // 	{ID: 22, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 22, Data: 4.2},
// // 	{ID: 23, CookLogID: 1, RecipeStepID: 6, EvaluationItemID: 23, Data: 4.3},

// // 	// 7. サラダ油を入れる (RecipeStepID:7)
// // 	{ID: 24, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 24, Data: 0.7},
// // 	{ID: 25, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 25, Data: 96.0},
// // 	{ID: 26, CookLogID: 1, RecipeStepID: 7, EvaluationItemID: 26, Data: 4.4},

// // 	// 8. 牛肉を入れる (RecipeStepID:8)
// // 	{ID: 27, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 27, Data: 0.8},
// // 	{ID: 28, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 28, Data: 94.0},
// // 	{ID: 29, CookLogID: 1, RecipeStepID: 8, EvaluationItemID: 29, Data: 4.2},

// // 	// 9. 牛肉を炒める (RecipeStepID:9)
// // 	{ID: 30, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 30, Data: 0.6},
// // 	{ID: 31, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 31, Data: 120},
// // 	{ID: 32, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 32, Data: 0.4},
// // 	{ID: 33, CookLogID: 1, RecipeStepID: 9, EvaluationItemID: 33, Data: 4.3},

// // 	// 10. ジャガイモを入れる (RecipeStepID:10)
// // 	{ID: 34, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 34, Data: 0.7},
// // 	{ID: 35, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 35, Data: 95.0},
// // 	{ID: 36, CookLogID: 1, RecipeStepID: 10, EvaluationItemID: 36, Data: 4.3},

// // 	// 11. ニンジンを入れる (RecipeStepID:11)
// // 	{ID: 37, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 37, Data: 0.7},
// // 	{ID: 38, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 38, Data: 94.0},
// // 	{ID: 39, CookLogID: 1, RecipeStepID: 11, EvaluationItemID: 39, Data: 4.2},

// // 	// 12. 玉ねぎを入れる (RecipeStepID:12)
// // 	{ID: 40, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 40, Data: 0.8},
// // 	{ID: 41, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 41, Data: 93.0},
// // 	{ID: 42, CookLogID: 1, RecipeStepID: 12, EvaluationItemID: 42, Data: 4.1},

// // 	// 13. 複合を炒める (RecipeStepID:13)
// // 	{ID: 43, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 43, Data: 0.6},
// // 	{ID: 44, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 44, Data: 4.0},
// // 	{ID: 45, CookLogID: 1, RecipeStepID: 13, EvaluationItemID: 45, Data: 4.2},

// // 	// 14. 水を入れる (RecipeStepID:14)
// // 	{ID: 46, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 46, Data: 0.9},
// // 	{ID: 47, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 47, Data: 97.0},
// // 	{ID: 48, CookLogID: 1, RecipeStepID: 14, EvaluationItemID: 48, Data: 4.5},

// // 	// 15. 調味料を入れる (RecipeStepID:15)
// // 	{ID: 49, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 49, Data: 0.7},
// // 	{ID: 50, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 50, Data: 95.0},
// // 	{ID: 51, CookLogID: 1, RecipeStepID: 15, EvaluationItemID: 51, Data: 4.3},

// // 	// 16. 灰汁を取る (RecipeStepID:16)
// // 	{ID: 52, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 52, Data: 0.5},
// // 	{ID: 53, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 53, Data: 4.5},
// // 	{ID: 54, CookLogID: 1, RecipeStepID: 16, EvaluationItemID: 54, Data: 4.4},

// // 	// 17. 白滝を入れる (RecipeStepID:17)
// // 	{ID: 55, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 55, Data: 0.6},
// // 	{ID: 56, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 56, Data: 94.0},
// // 	{ID: 57, CookLogID: 1, RecipeStepID: 17, EvaluationItemID: 57, Data: 4.2},

// // 	// 18. 落とし蓋をする (RecipeStepID:18)
// // 	{ID: 58, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 58, Data: 0.6},
// // 	{ID: 59, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 59, Data: 95.0},
// // 	{ID: 60, CookLogID: 1, RecipeStepID: 18, EvaluationItemID: 60, Data: 4.3},

// // 	// 19. 複合を煮詰める (RecipeStepID:19)
// // 	{ID: 61, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 61, Data: 0.5},
// // 	{ID: 62, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 62, Data: 4.1},
// // 	{ID: 63, CookLogID: 1, RecipeStepID: 19, EvaluationItemID: 63, Data: 4.2},

// // 	// 20. 複合を混ぜる (RecipeStepID:20)
// // 	{ID: 64, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 64, Data: 0.7},
// // 	{ID: 65, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 65, Data: 4.3},
// // 	{ID: 66, CookLogID: 1, RecipeStepID: 20, EvaluationItemID: 66, Data: 4.4},

// // 	// 21. 複合を煮詰める（再度） (RecipeStepID:21)
// // 	{ID: 67, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 67, Data: 0.6},
// // 	{ID: 68, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 68, Data: 4.2},
// // 	{ID: 69, CookLogID: 1, RecipeStepID: 21, EvaluationItemID: 69, Data: 4.3},

// // 	// 22. 複合を蒸らす (RecipeStepID:22)
// // 	{ID: 70, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 70, Data: 0.5},
// // 	{ID: 71, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 71, Data: 4.4},
// // 	{ID: 72, CookLogID: 1, RecipeStepID: 22, EvaluationItemID: 72, Data: 4.4},

// // 	// 23. 複合を盛り付ける (RecipeStepID:23)  ※CookActionIngredientsID=32 に合わせて EvaluationItemID=100..102
// // 	{ID: 73, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 100, Data: 0.8}, // 速さ
// // 	{ID: 74, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 101, Data: 4.5}, // 見栄え
// // 	{ID: 75, CookLogID: 1, RecipeStepID: 23, EvaluationItemID: 102, Data: 4.6}, // 総合評価

// // 	// 24. ニンジンを乱切り (RecipeStepID:24)
// // 	{ID: 76, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 5, Data: 0.7},
// // 	{ID: 77, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 6, Data: 4.1},
// // 	{ID: 78, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 7, Data: 4.0},
// // 	{ID: 79, CookLogID: 2, RecipeStepID: 24, EvaluationItemID: 8, Data: 4.1},

// // 	// 25. 玉ねぎをくし切り (RecipeStepID:25)
// // 	{ID: 80, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 9, Data: 0.8},
// // 	{ID: 81, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 10, Data: 4.2},
// // 	{ID: 82, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 11, Data: 4.1},
// // 	{ID: 83, CookLogID: 2, RecipeStepID: 25, EvaluationItemID: 12, Data: 4.2},

// // 	// 26. 調味料を測る (RecipeStepID:26)
// // 	{ID: 84, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 21, Data: 95.0},
// // 	{ID: 85, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 22, Data: 4.2},
// // 	{ID: 86, CookLogID: 2, RecipeStepID: 26, EvaluationItemID: 23, Data: 4.3},

// // 	// 27. バターを入れる (RecipeStepID:27)
// // 	{ID: 87, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 73, Data: 0.7},
// // 	{ID: 88, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 74, Data: 95.0},
// // 	{ID: 89, CookLogID: 2, RecipeStepID: 27, EvaluationItemID: 75, Data: 4.3},

// // 	// 28. おろしニンニクを入れる (RecipeStepID:28)
// // 	{ID: 90, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 76, Data: 0.6},
// // 	{ID: 91, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 77, Data: 96.0},
// // 	{ID: 92, CookLogID: 2, RecipeStepID: 28, EvaluationItemID: 78, Data: 4.4},

// // 	// 29. 玉ねぎを入れる (RecipeStepID:29)
// // 	{ID: 93, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 40, Data: 0.8},
// // 	{ID: 94, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 41, Data: 93.0},
// // 	{ID: 95, CookLogID: 2, RecipeStepID: 29, EvaluationItemID: 42, Data: 4.1},

// // 	// 30. 玉ねぎを炒める (RecipeStepID:30) ※EvaluationItemID=79..81 に修正（3件）
// // 	{ID: 96, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 79, Data: 0.7}, // 速さ
// // 	{ID: 97, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 80, Data: 0.5}, // 均一性（※元データの 0.5 を採用）
// // 	{ID: 98, CookLogID: 2, RecipeStepID: 30, EvaluationItemID: 81, Data: 4.3}, // 総合評価

// // 	// 31. 玉ねぎを取り出す (RecipeStepID:31) ※EvaluationItemID=82..84 に修正
// // 	{ID: 99, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 82, Data: 0.6},
// // 	{ID: 100, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 83, Data: 95.0},
// // 	{ID: 101, CookLogID: 2, RecipeStepID: 31, EvaluationItemID: 84, Data: 4.2},

// // 	// 32. サラダ油を入れる (RecipeStepID:32)
// // 	{ID: 102, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 24, Data: 0.7},
// // 	{ID: 103, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 25, Data: 96.0},
// // 	{ID: 104, CookLogID: 2, RecipeStepID: 32, EvaluationItemID: 26, Data: 4.4},

// // 	// 33. 牛肉を入れる (RecipeStepID:33)
// // 	{ID: 105, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 27, Data: 0.8},
// // 	{ID: 106, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 28, Data: 94.0},
// // 	{ID: 107, CookLogID: 2, RecipeStepID: 33, EvaluationItemID: 29, Data: 4.2},

// // 	// 34. 牛肉を炒める (RecipeStepID:34)
// // 	{ID: 108, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 30, Data: 0.6},
// // 	{ID: 109, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 31, Data: 120},
// // 	{ID: 110, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 32, Data: 0.4},
// // 	{ID: 111, CookLogID: 2, RecipeStepID: 34, EvaluationItemID: 33, Data: 4.3},

// // 	// 35. 塩胡椒を入れる (RecipeStepID:35)
// // 	{ID: 112, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 86, Data: 0.7},
// // 	{ID: 113, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 87, Data: 95.0},
// // 	{ID: 114, CookLogID: 2, RecipeStepID: 35, EvaluationItemID: 88, Data: 4.3},

// // 	// 36. ニンジンを入れる (RecipeStepID:36)
// // 	{ID: 115, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 37, Data: 0.7},
// // 	{ID: 116, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 38, Data: 94.0},
// // 	{ID: 117, CookLogID: 2, RecipeStepID: 36, EvaluationItemID: 39, Data: 4.2},

// // 	// 37. 玉ねぎを入れる (RecipeStepID:37)
// // 	{ID: 118, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 40, Data: 0.8},
// // 	{ID: 119, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 41, Data: 93.0},
// // 	{ID: 120, CookLogID: 2, RecipeStepID: 37, EvaluationItemID: 42, Data: 4.1},

// // 	// 38. 複合を炒める (RecipeStepID:38)
// // 	{ID: 121, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 43, Data: 0.6},
// // 	{ID: 122, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 44, Data: 4.0},
// // 	{ID: 123, CookLogID: 2, RecipeStepID: 38, EvaluationItemID: 45, Data: 4.2},

// // 	// 39. 水を入れる (RecipeStepID:39)
// // 	{ID: 124, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 46, Data: 0.9},
// // 	{ID: 125, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 47, Data: 97.0},
// // 	{ID: 126, CookLogID: 2, RecipeStepID: 39, EvaluationItemID: 48, Data: 4.5},

// // 	// 40. 水を煮立てる (RecipeStepID:40)
// // 	{ID: 127, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 89, Data: 0.8},
// // 	{ID: 128, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 90, Data: 180},
// // 	{ID: 129, CookLogID: 2, RecipeStepID: 40, EvaluationItemID: 91, Data: 4.3},

// // 	// 41. 灰汁を取る (RecipeStepID:41)
// // 	{ID: 130, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 52, Data: 0.5},
// // 	{ID: 131, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 53, Data: 4.5},
// // 	{ID: 132, CookLogID: 2, RecipeStepID: 41, EvaluationItemID: 54, Data: 4.4},

// // 	// 42. ローリエの葉を入れる (RecipeStepID:42)
// // 	{ID: 133, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 92, Data: 0.6},
// // 	{ID: 134, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 93, Data: 94.0},
// // 	{ID: 135, CookLogID: 2, RecipeStepID: 42, EvaluationItemID: 94, Data: 4.2},

// // 	// 43. 複合を煮込む (RecipeStepID:43)
// // 	{ID: 136, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 95, Data: 0.5},
// // 	{ID: 137, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 96, Data: 4.0},
// // 	{ID: 138, CookLogID: 2, RecipeStepID: 43, EvaluationItemID: 97, Data: 4.3},

// // 	// 44. ルーを溶かす (RecipeStepID:44)  ※EvaluationItemID=97..99 に修正
// // 	{ID: 139, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 97, Data: 0.7},
// // 	{ID: 140, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 98, Data: 95.0},
// // 	{ID: 141, CookLogID: 2, RecipeStepID: 44, EvaluationItemID: 99, Data: 4.4},

// // 	// 45. 複合を煮込む (RecipeStepID:45)
// // 	{ID: 142, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 95, Data: 0.5},
// // 	{ID: 143, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 96, Data: 4.0},
// // 	{ID: 144, CookLogID: 2, RecipeStepID: 45, EvaluationItemID: 97, Data: 4.3},

// // 	// 46. 調味料を入れる (RecipeStepID:46)
// // 	{ID: 145, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 49, Data: 0.7},
// // 	{ID: 146, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 50, Data: 95.0},
// // 	{ID: 147, CookLogID: 2, RecipeStepID: 46, EvaluationItemID: 51, Data: 4.3},

// // 	// 47. 複合を煮込む (RecipeStepID:47)
// // 	{ID: 148, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 95, Data: 0.5},
// // 	{ID: 149, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 96, Data: 4.0},
// // 	{ID: 150, CookLogID: 2, RecipeStepID: 47, EvaluationItemID: 97, Data: 4.3},

// // 	// 48. 複合を盛り付ける (RecipeStepID:48)  ※EvaluationItemID=100..102 に修正
// // 	{ID: 151, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 100, Data: 0.7}, // 速さ
// // 	{ID: 152, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 101, Data: 4.3}, // 見栄え
// // 	{ID: 153, CookLogID: 2, RecipeStepID: 48, EvaluationItemID: 102, Data: 4.3}, // 総合評価
// // }
