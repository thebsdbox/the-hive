package pwd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientNew(t *testing.T) {
	docker := &mockDocker{}
	tasks := &mockTasks{}
	broadcast := &mockBroadcast{}
	storage := &mockStorage{}

	p := NewPWD(docker, tasks, broadcast, storage)

	session, err := p.SessionNew(time.Hour, "", "")
	assert.Nil(t, err)

	client := p.ClientNew("foobar", session)

	assert.Equal(t, Client{Id: "foobar", session: session, viewPort: ViewPort{Cols: 0, Rows: 0}}, *client)
	assert.Contains(t, session.clients, client)
}

func TestClientResizeViewPort(t *testing.T) {
	docker := &mockDocker{}
	tasks := &mockTasks{}
	broadcast := &mockBroadcast{}

	broadcastedSessionId := ""
	broadcastedEventName := ""
	broadcastedArgs := []interface{}{}

	broadcast.broadcastTo = func(sessionId, eventName string, args ...interface{}) {
		broadcastedSessionId = sessionId
		broadcastedEventName = eventName
		broadcastedArgs = args
	}

	storage := &mockStorage{}

	p := NewPWD(docker, tasks, broadcast, storage)

	session, err := p.SessionNew(time.Hour, "", "")
	assert.Nil(t, err)
	client := p.ClientNew("foobar", session)

	p.ClientResizeViewPort(client, 80, 24)

	assert.Equal(t, ViewPort{Cols: 80, Rows: 24}, client.viewPort)
	assert.Equal(t, session.Id, broadcastedSessionId)
	assert.Equal(t, "viewport resize", broadcastedEventName)
	assert.Equal(t, uint(80), broadcastedArgs[0])
	assert.Equal(t, uint(24), broadcastedArgs[1])
}
