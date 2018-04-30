package bootstrap

import (
	"time"

	"github.com/cubeee/steamtracker/web/tags"
	"github.com/cubeee/steamtracker/web/util/engine"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

const (
	StaticAssetsPath  = "/static"
	StaticAssetsDir   = "./resources/static"
	TemplateDir       = "./resources/templates"
	TemplateExtension = ".tpl"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName               string
	AppSpawnDate          time.Time
	DisableVersionChecker bool
}

type GlobalCtx struct {
	GoogleAnalyticsId string
}

func New(appName string, configurators ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:               appName,
		AppSpawnDate:          time.Now(),
		Application:           iris.New(),
		DisableVersionChecker: true,
	}
	b.Configure(configurators...)
	return b
}

func (b *Bootstrapper) SetupViews(reload bool) {
	templateEngine := engine.Pongo2(TemplateDir, TemplateExtension)
	templateEngine.Reload(reload)
	templateEngine.RegisterTag("loop", tags.LoopTagParser)
	b.RegisterView(templateEngine)
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("err", err)
		ctx.ViewData("title", "Error")
		ctx.View("_error.tpl")
	})
}

func (b *Bootstrapper) Configure(configurators ...Configurator) {
	for _, configurator := range configurators {
		configurator(b)
	}
}

func (b *Bootstrapper) Bootstrap(env *string, globalContext *GlobalCtx) *Bootstrapper {
	devMode := *env == "dev"
	b.SetupViews(devMode)
	b.SetupErrorHandlers()

	b.StaticWeb(StaticAssetsPath, StaticAssetsDir)

	if globalContext != nil {
		b.Use(func(ctx iris.Context) {
			ctx.ViewData("googleAnalyticsId", globalContext.GoogleAnalyticsId)
			ctx.Next()
		})
	}
	b.Use(recover.New())
	if devMode {
		b.Use(logger.New())
	}
	return b
}

func (b *Bootstrapper) Listen(addr string, configurators ...iris.Configurator) {
	b.Run(iris.Addr(addr), configurators...)
}
