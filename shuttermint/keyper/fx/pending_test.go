package fx

import (
	"encoding/gob"
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"
	"os"

	"github.com/shutter-network/shutter/shuttermint/keyper/observe"
)

var myactions []IAction

type MyAction struct {
	ID ActionID
}

func (a *MyAction) IsExpired(_ observe.World) bool {
	return false
}

func init() {
	gob.Register(&MyAction{})
	for i := 0; i < 20; i++ {
		myactions = append(myactions, &MyAction{ID: ActionID(i)})
	}
}

func TestAddActions(t *testing.T) {
	tmp, err := os.MkdirTemp("", "*")
	assert.NilError(t, err)
	pending := NewPendingActions(filepath.Join(tmp, "actions.gob"))
	pending.AddActions(ActionID(0), myactions[0:10])

	for i := 3; i < 5; i++ {
		pending.RemoveAction(ActionID(i))
	}

	assert.Equal(t, 8, len(pending.SortedIDs()))
	pending.AddActions(ActionID(3), myactions[3:5])
	assert.Equal(t, 8, len(pending.SortedIDs()))
}
