package cluster

import (
	"fmt"

	"github.com/docker/engine-api/types"
)

// Volume is exported
type Volume struct {
	types.Volume

	Engine *Engine
}

// Volumes represents an array of volumes
type Volumes []*Volume

// Get returns a volume using its ID or Name
func (volumes Volumes) Get(IDOrName string) (*Volume, error) {
	// Abort immediately if the IDOrName is empty.
	if len(IDOrName) == 0 {
		return nil, fmt.Errorf("Length of volume name cannot be 0")
	}

	candidates := []*Volume{}

	// Match name or engine/name.
	for _, volume := range volumes {
		if volume.Name == IDOrName ||
			volume.Engine.ID+"/"+volume.Name == IDOrName ||
			volume.Engine.Name+"/"+volume.Name == IDOrName {
			candidates = append(candidates, volume)
		}
	}

	// Return if we found a unique match.
	if size := len(candidates); size == 1 {
		return candidates[0], nil
	} else if size > 1 {
		return nil, fmt.Errorf("More than one volume named by (%s) in cluster."+
			" Please use Engine_ID/Volume_Name to specify one.", IDOrName)
	}

	// There is no candidate.
	// Match /name and return as soon as we find one.
	for _, volume := range volumes {
		if volume.Name == "/"+IDOrName {
			candidates = append(candidates, volume)
		}
	}

	return nil
}
