package backend

import (
	"context"
	"could-work/backend/core/define"
	"could-work/backend/event"
	"could-work/backend/util"
)

var log = define.Log

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	log.Info("程序启动")

	util.TaskRunner(
		event.InitUserDB,
		event.InitGinServer,
	)
}

func (a *App) Shutdown(ctx context.Context) {
	log.Info("程序关闭")
}
