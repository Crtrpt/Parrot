package gen

import (
	"context"

	"github.com/parrot/internal/core"
	"gorm.io/gen"
)

func Start(ctx context.Context) (err error) {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/pkg/model",
	})

	g.UseDB(core.DB)
	g.GenerateAllTable()
	g.Execute()

	return nil
}
