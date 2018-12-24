package service

import (
	"bytes"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/softleader/captain-kube/cmd/cap-ui/app/server/comm"
	"github.com/softleader/captain-kube/pkg/captain"
	"github.com/softleader/captain-kube/pkg/dockerctl"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/sse"
	"github.com/softleader/captain-kube/pkg/utils"
	"github.com/softleader/captain-kube/pkg/utils/strutil"
	"io"
	"net/http"
	"strconv"
)

type InstallRequest struct {
	Platform       string   `form:"platform"`
	Namespace      string   `form:"namespace"`
	Tags           []string `form:"tags"`
	SourceRegistry string   `form:"sourceRegistry"`
	Registry       string   `form:"registry"`
	Verbose        bool     `form:"verbose"`
}

type Install struct {
	Log *logrus.Logger // 這個是 server 自己的 log
	Cfg *comm.Config
}

func (s *Install) View(c *gin.Context) {
	c.HTML(http.StatusOK, "install.html", gin.H{
		"config": &s.Cfg,
	})
}

func (s *Install) Chart(c *gin.Context) { //(path string, r *gin.Engine, cfg *comm.Config) {
	var form InstallRequest

	log := logrus.New()
	log.SetFormatter(&utils.PlainFormatter{})
	log.SetOutput(sse.NewWriter(c))
	if v, _ := strconv.ParseBool(c.Request.FormValue("verbose")); v {
		log.SetLevel(logrus.DebugLevel)
	}

	if err := c.Bind(&form); err != nil {
		//sw.WriteStr(fmt.Sprint("binding form data error:", err))
		log.Println("binding form data error:", err)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		//sw.WriteStr(fmt.Sprint("loading form file error:", err))
		log.Println("loading form file error:", err)
		return
	}

	// ps. 在讀完request body後才可以開始response, 否則body會close

	log.Println("call: POST /install")

	log.Println("form:", form)
	log.Println("file:", file)

	buf := bytes.NewBuffer(nil)
	if readed, err := io.Copy(buf, file); err != nil {
		log.Println("reading file failed:", err)
		return
	} else {
		log.Println("readed ", readed, " bytes")
	}

	// prepare rquest
	request := proto.InstallChartRequest{
		Chart: &proto.Chart{
			FileName: header.Filename,
			Content:  buf.Bytes(),
			FileSize: header.Size,
		},
		Pull: strutil.Contains(form.Tags, "p"),
		Sync: strutil.Contains(form.Tags, "r"),
		Retag: &proto.ReTag{
			From: form.SourceRegistry,
			To:   form.Registry,
		},
		RegistryAuth: &proto.RegistryAuth{
			Username: s.Cfg.RegistryAuth.Username,
			Password: s.Cfg.RegistryAuth.Password,
		},
		Tiller: &proto.Tiller{
			Endpoint:          s.Cfg.Tiller.Endpoint,
			Username:          s.Cfg.Tiller.Username,
			Password:          s.Cfg.Tiller.Password,
			Account:           s.Cfg.Tiller.Account,
			SkipSslValidation: s.Cfg.Tiller.SkipSslValidation,
		},
	}

	if err := dockerctl.PullAndSync(log, &request); err != nil {
		log.Println("Pull/Sync failed:", err)
	}

	if err := captain.InstallChart(log, s.Cfg.DefaultValue.CaptainUrl, &request, 300); err != nil {
		fmt.Fprintln(c.Writer, "call captain InstallChart failed:", err)
	}
	fmt.Fprintln(c.Writer, "InstallChart finish")
}
