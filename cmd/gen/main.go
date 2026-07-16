package main

import (
	"Diggpher/internal/dao"

	"gorm.io/gen"
)

// 从 internal/dao 模型生成类型安全的 query API 到 internal/query。
// 模型变更后重新执行: go run ./cmd/gen
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// 不连库：以手写 dao 模型为源生成 Query
	g.ApplyBasic(
		dao.Admin{},
	)

	g.Execute()
}
