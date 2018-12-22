package dockerctl

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/cli/command"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/softleader/captain-kube/pkg/helm/chart"
	"github.com/softleader/captain-kube/pkg/logger"
	"github.com/softleader/captain-kube/pkg/proto"
	"io"
	"io/ioutil"
	"os"
)

func Pull(log *logger.Logger, image chart.Image, registryAuth *proto.RegistryAuth) (io.ReadCloser, error) {
	ctx := context.Background()

	// Use DOCKER_HOST to set the url to the docker server.
	// Use DOCKER_API_VERSION to set the version of the API to reach, leave empty for latest.
	// Use DOCKER_CERT_PATH to load the TLS certificates from.
	// Use DOCKER_TLS_VERIFY to enable or disable TLS verification, off by default.
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	log.Printf("pulling image: %s\n", image)
	opt := types.ImagePullOptions{}
	if registryAuth != nil {
		if opt.RegistryAuth, err = encode(registryAuth); err != nil {
			return nil, err
		}
	}
	rc, err := cli.ImagePull(ctx, image.String(), opt)
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func ReTag(log *logger.Logger, source chart.Image, target chart.Image, registryAuth *proto.RegistryAuth) (io.ReadCloser, error) {
	ctx := context.Background()

	// Use DOCKER_HOST to set the url to the docker server.
	// Use DOCKER_API_VERSION to set the version of the API to reach, leave empty for latest.
	// Use DOCKER_CERT_PATH to load the TLS certificates from.
	// Use DOCKER_TLS_VERIFY to enable or disable TLS verification, off by default.
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	log.Printf("tagging image from %q to %q \n", source, target)
	if err := cli.ImageTag(ctx, source.String(), target.String()); err != nil {
		return nil, err
	}

	log.Printf("pushing image: %s \n", target)
	opt := types.ImagePushOptions{}
	if registryAuth != nil {
		if opt.RegistryAuth, err = encode(registryAuth); err != nil {
			return nil, err
		}
	}
	rc, err := cli.ImagePush(ctx, target.String(), opt)
	if err := cli.ImageTag(ctx, source.String(), target.String()); err != nil {
		return nil, err
	}

	return rc, nil
}

func PullAndSync(log *logger.Logger, request *proto.InstallChartRequest) error {
	var tpls chart.Templates
	if request.Pull || request.Sync {
		// mk temp file
		tmpFile, err := ioutil.TempFile(os.TempDir(), "capui-*.tgz")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.Write(request.Chart.Content); err != nil {
			return err
		}

		// load chart templates
		tpls, err = chart.LoadArchive(log, tmpFile.Name())
		if err != nil {
			return err
		}

	}

	if request.Pull {
		// pull all image from chart
		for _, tpl := range tpls {
			for _, image := range tpl {
				log.Println("pulling ", image)
				result, err := Pull(log, *image, request.RegistryAuth)
				if err != nil {
					log.Println("pull image failed: ", image, ", error: ", err)
				}
				jsonmessage.DisplayJSONMessagesToStream(result, command.NewOutStream(log.GetOutput()), nil)
			}
		}
	}

	if request.Sync {
		if len(request.Retag.From) > 0 && len(request.Retag.To) > 0 {
			// retag and push all image from chart
			for _, tpl := range tpls {
				for _, image := range tpl {
					if image.Host == request.Retag.From {
						log.Println("syncing ", image)
						result, err := ReTag(log, *image, chart.Image{
							Host: request.Retag.To,
							Repo: image.Repo,
							Tag:  image.Tag,
						}, request.RegistryAuth)
						if err != nil {
							log.Println( "sync image failed: ", image, ", error: ", err)
						}
						jsonmessage.DisplayJSONMessagesToStream(result, command.NewOutStream(log.GetOutput()), nil)
					}
				}
			}
		} else {
			return errors.New("Registry and SourceRegistry is required when retag mode")
		}
	}

	return nil
}

func encode(ra *proto.RegistryAuth) (string, error) {
	b, err := json.Marshal(ra)
	if err != nil {
		return "", nil
	}
	return base64.URLEncoding.EncodeToString(b), nil
}