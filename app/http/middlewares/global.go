package middlewares

import (
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func GlobalMiddlewares(engine *gin.Engine) {
	engine.Use(
		cors.Default(),
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
