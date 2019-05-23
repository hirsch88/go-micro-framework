package middlewares

import (
	"github.com/aviddiviner/gin-limit"
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/thinkerou/favicon"
)

func GlobalMiddlewares(server *gin.Engine, config *config.AppConfig, logMiddleware *LogMiddleware) {
	server.Use(
		logMiddleware.Handler(),
		cors.Default(),
		limit.MaxAllowed(config.Connection),
		gzip.Gzip(gzip.DefaultCompression),
		helmet.NoSniff(),
		helmet.DNSPrefetchControl(),
		helmet.FrameGuard(),
		helmet.SetHSTS(true),
		helmet.IENoOpen(),
		helmet.XSSFilter(),
		helmet.NoCache(),
		favicon.New("./public/favicon.ico"),
		static.Serve("/", static.LocalFile("./public", false)),
	)
}
