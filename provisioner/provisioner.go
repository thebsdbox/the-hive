package provisioner

import (
	"errors"
	"io"
	"net"

	"github.com/play-with-docker/play-with-docker/pwd/types"
)

var OutOfCapacityError = errors.New("OutOfCapacity")

func OutOfCapacity(e error) bool {
	return e == OutOfCapacityError
}

type InstanceProvisionerApi interface {
	InstanceNew(session *types.Session, conf types.InstanceConfig) (*types.Instance, error)
	InstanceDelete(session *types.Session, instance *types.Instance) error

	InstanceResizeTerminal(instance *types.Instance, cols, rows uint) error
	InstanceGetTerminal(instance *types.Instance) (net.Conn, error)

	InstanceUploadFromUrl(instance *types.Instance, fileName, dest, url string) error
	InstanceUploadFromReader(instance *types.Instance, fileName, dest string, reader io.Reader) error
}

type SessionProvisionerApi interface {
	SessionNew(session *types.Session) error
	SessionClose(session *types.Session) error
}

type InstanceProvisionerFactoryApi interface {
	GetProvisioner(instanceType string) (InstanceProvisionerApi, error)
}
