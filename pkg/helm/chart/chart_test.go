package chart

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/utils/command"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLoadArchive(t *testing.T) {
	helm := "helm"
	if !command.IsAvailable(helm) {
		t.Skipf("%q command does not exist", helm)
	}

	log := logrus.New()

	path, err := ioutil.TempDir(os.TempDir(), "test-load-archive-")
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

	tpls, err := LoadArchive(log, filepath.Join(path, "foo-0.1.0.tgz"))
	if err != nil {
		t.Error(err)
	}
	if len(tpls) == 0 {
		t.Error("no template found")
	}
}
