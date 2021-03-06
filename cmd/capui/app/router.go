package app

import (
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/softleader/captain-kube/pkg/utils/strutil"
	"html/template"
	"net/http"
	"strings"
)

// NewCapUIServer 建立 capui server
func NewCapUIServer(cmd *capUICmd) (r *gin.Engine) {
	r = gin.Default()

	r.Use(healthcheck.Default()) // for health probe

	r.SetFuncMap(template.FuncMap{
		"Contains": strutil.Contains,
		"NotContains": func(vs []string, s string) bool {
			return !strutil.Contains(vs, s)
		},
		"Join": strings.Join,
		"HasPrefix": strings.HasPrefix,
	})

	// static and template
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*.html")

	// index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"requestURI": c.Request.RequestURI,
			"config":     &cmd,
		})
	})

	chartsRoute := r.Group("/charts")
	{
		charts := &Charts{
			cmd,
		}
		chartsRoute.GET("/", charts.View)
		chartsRoute.POST("/", charts.Install)
	}

	cleanUpRoute := r.Group("/cleanup")
	{
		cleanup := &CleanUp{
			cmd,
		}
		cleanUpRoute.GET("/", cleanup.View)
		cleanUpRoute.POST("/", cleanup.Clean)
	}

	contextsRoute := r.Group("/contexts")
	{
		ctxs := &Contexts{
			cmd,
		}
		contextsRoute.GET("/", ctxs.ListContext)
		contextsRoute.GET("/versions", ctxs.ListContextVersions)
	}

	return
}
