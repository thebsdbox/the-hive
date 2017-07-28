package task

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/play-with-docker/play-with-docker/docker"
	"github.com/play-with-docker/play-with-docker/event"
	"github.com/play-with-docker/play-with-docker/pwd/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSessionProvider struct {
	mock.Mock
}

func (m *mockSessionProvider) GetDocker(sessionId string) (docker.DockerApi, error) {
	args := m.Called(sessionId)

	return args.Get(0).(docker.DockerApi), args.Error(1)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func TestCollectStats_Name(t *testing.T) {
	e := &event.Mock{}
	f := &docker.FactoryMock{}

	task := NewCollectStats(e, f)

	assert.Equal(t, "CollectStats", task.Name())
	e.M.AssertExpectations(t)
	f.AssertExpectations(t)
}

func TestCollectStats_Run(t *testing.T) {
	d := &docker.Mock{}
	e := &event.Mock{}
	f := &docker.FactoryMock{}

	stats := dockerTypes.StatsJSON{}
	b, _ := json.Marshal(stats)
	i := &types.Instance{
		IP:        "10.0.0.1",
		Name:      "aaaabbbb_node1",
		SessionId: "aaaabbbbcccc",
	}

	f.On("GetForSession", i.SessionId).Return(d, nil)
	d.On("GetContainerStats", i.Name).Return(nopCloser{bytes.NewReader(b)}, nil)
	e.M.On("Emit", CollectStatsEvent, "aaaabbbbcccc", []interface{}{InstanceStats{Instance: i.Name, Mem: "0.00% (0B / 0B)", Cpu: "0.00%"}}).Return()

	task := NewCollectStats(e, f)
	ctx := context.Background()

	err := task.Run(ctx, i)

	assert.Nil(t, err)
	d.AssertExpectations(t)
	e.M.AssertExpectations(t)
	f.AssertExpectations(t)
}
