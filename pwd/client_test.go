package pwd

import (
	"testing"
	"time"

	"github.com/play-with-docker/play-with-docker/config"
	"github.com/play-with-docker/play-with-docker/docker"
	"github.com/play-with-docker/play-with-docker/event"
	"github.com/play-with-docker/play-with-docker/provisioner"
	"github.com/play-with-docker/play-with-docker/pwd/types"
	"github.com/play-with-docker/play-with-docker/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientNew(t *testing.T) {
	_s := &storage.Mock{}
	_f := &docker.FactoryMock{}
	_g := &mockGenerator{}
	_d := &docker.Mock{}
	_e := &event.Mock{}

	ipf := provisioner.NewInstanceProvisionerFactory(provisioner.NewWindowsASG(_f, _s), provisioner.NewDinD(_f))
	sp := provisioner.NewOverlaySessionProvisioner(_f)

	_g.On("NewId").Return("aaaabbbbcccc")
	_f.On("GetForSession", "aaaabbbbcccc").Return(_d, nil)
	_d.On("CreateNetwork", "aaaabbbbcccc").Return(nil)
	_d.On("GetDaemonHost").Return("localhost")
	_d.On("ConnectNetwork", config.L2ContainerName, "aaaabbbbcccc", "").Return("10.0.0.1", nil)
	_s.On("SessionPut", mock.AnythingOfType("*types.Session")).Return(nil)
	_s.On("SessionCount").Return(1, nil)
	_s.On("InstanceCount").Return(0, nil)

	var nilArgs []interface{}
	_e.M.On("Emit", event.SESSION_NEW, "aaaabbbbcccc", nilArgs).Return()

	p := NewPWD(_f, _e, _s, sp, ipf)
	p.generator = _g

	session, err := p.SessionNew(time.Hour, "", "", "")
	assert.Nil(t, err)

	client := p.ClientNew("foobar", session)

	assert.Equal(t, types.Client{Id: "foobar", Session: session, ViewPort: types.ViewPort{Cols: 0, Rows: 0}}, *client)
	assert.Contains(t, session.Clients, client)

	_d.AssertExpectations(t)
	_f.AssertExpectations(t)
	_s.AssertExpectations(t)
	_g.AssertExpectations(t)
	_e.M.AssertExpectations(t)
}

func TestClientCount(t *testing.T) {
	_s := &storage.Mock{}
	_f := &docker.FactoryMock{}
	_g := &mockGenerator{}
	_d := &docker.Mock{}
	_e := &event.Mock{}
	ipf := provisioner.NewInstanceProvisionerFactory(provisioner.NewWindowsASG(_f, _s), provisioner.NewDinD(_f))
	sp := provisioner.NewOverlaySessionProvisioner(_f)

	_g.On("NewId").Return("aaaabbbbcccc")
	_f.On("GetForSession", "aaaabbbbcccc").Return(_d, nil)
	_d.On("CreateNetwork", "aaaabbbbcccc").Return(nil)
	_d.On("GetDaemonHost").Return("localhost")
	_d.On("ConnectNetwork", config.L2ContainerName, "aaaabbbbcccc", "").Return("10.0.0.1", nil)
	_s.On("SessionPut", mock.AnythingOfType("*types.Session")).Return(nil)
	_s.On("SessionCount").Return(1, nil)
	_s.On("InstanceCount").Return(-1, nil)
	var nilArgs []interface{}
	_e.M.On("Emit", event.SESSION_NEW, "aaaabbbbcccc", nilArgs).Return()

	p := NewPWD(_f, _e, _s, sp, ipf)
	p.generator = _g

	session, err := p.SessionNew(time.Hour, "", "", "")
	assert.Nil(t, err)

	p.ClientNew("foobar", session)

	assert.Equal(t, 1, p.ClientCount())

	_d.AssertExpectations(t)
	_f.AssertExpectations(t)
	_s.AssertExpectations(t)
	_g.AssertExpectations(t)
	_e.M.AssertExpectations(t)
}

func TestClientResizeViewPort(t *testing.T) {
	_s := &storage.Mock{}
	_f := &docker.FactoryMock{}
	_g := &mockGenerator{}
	_d := &docker.Mock{}
	_e := &event.Mock{}
	ipf := provisioner.NewInstanceProvisionerFactory(provisioner.NewWindowsASG(_f, _s), provisioner.NewDinD(_f))
	sp := provisioner.NewOverlaySessionProvisioner(_f)

	_g.On("NewId").Return("aaaabbbbcccc")
	_f.On("GetForSession", "aaaabbbbcccc").Return(_d, nil)
	_d.On("CreateNetwork", "aaaabbbbcccc").Return(nil)
	_d.On("GetDaemonHost").Return("localhost")
	_d.On("ConnectNetwork", config.L2ContainerName, "aaaabbbbcccc", "").Return("10.0.0.1", nil)
	_s.On("SessionPut", mock.AnythingOfType("*types.Session")).Return(nil)
	_s.On("SessionCount").Return(1, nil)
	_s.On("InstanceCount").Return(0, nil)
	var nilArgs []interface{}
	_e.M.On("Emit", event.SESSION_NEW, "aaaabbbbcccc", nilArgs).Return()

	_e.M.On("Emit", event.INSTANCE_VIEWPORT_RESIZE, "aaaabbbbcccc", []interface{}{uint(80), uint(24)}).Return()
	p := NewPWD(_f, _e, _s, sp, ipf)
	p.generator = _g

	session, err := p.SessionNew(time.Hour, "", "", "")
	assert.Nil(t, err)
	client := p.ClientNew("foobar", session)

	p.ClientResizeViewPort(client, 80, 24)

	assert.Equal(t, types.ViewPort{Cols: 80, Rows: 24}, client.ViewPort)

	_d.AssertExpectations(t)
	_f.AssertExpectations(t)
	_s.AssertExpectations(t)
	_g.AssertExpectations(t)
	_e.M.AssertExpectations(t)
}
