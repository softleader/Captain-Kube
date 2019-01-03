package captain

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/dur"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/utils"
	"github.com/softleader/captain-kube/pkg/utils/cmd"
	"github.com/softleader/captain-kube/pkg/utils/tcp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGenerateScript(t *testing.T) {
	endpoint := "192.168.1.93"
	port := DefaultPort
	if !tcp.Reachable(endpoint, port, 3) {
		t.Skipf("endpoint %s:%v is not reachable", endpoint, port)
	}

	helm := "helm"
	if !cmd.IsAvailable(helm) {
		t.Skipf("%q command does not exist", helm)
	}

	path, err := ioutil.TempDir(os.TempDir(), "test-generate-script-")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(path)

	cmd := exec.Command(helm, "create", "foo")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	cmd = exec.Command(helm, "package", "foo")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	chart, err := ioutil.ReadFile(filepath.Join(path, "foo-0.1.0.tgz"))
	if err != nil {
		t.Error(err)
	}

	log := logrus.New()
	log.SetFormatter(&utils.PlainFormatter{})
	err = GenerateScript(log, "192.168.1.93:30051", &proto.GenerateScriptRequest{
		Chart: &proto.Chart{
			Content:  chart,
			FileName: "foo-0.1.0.tgz",
		},
		Pull: true,
		Retag: &proto.ReTag{
		},
		Save:    true,
		Load:    true,
		Verbose: true,
	}, dur.DefaultDeadlineSecond)
	if err != nil {
		t.Error(err)
	}
}
