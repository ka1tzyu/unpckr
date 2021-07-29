package storage

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/logger"
	tools2 "github.com/nekovalue/unpckr/pkg/tools"
	"os"
	"strings"
)

type Storage struct {
	Sources      []string
	Destinations []string
}

// AppendSource adds file to sources list
func (store *Storage) AppendSource(source string) {
	store.Sources = append(store.Sources, source)
	logger.Log.Debugf("New source {%s} was appended", source)
}

// AppendDestination adds file to destinations list
func (store *Storage) AppendDestination(destination string) {
	store.Destinations = append(store.Destinations, destination)
	logger.Log.Debugf("New destination {%s} was appended", destination)
}

// RemoveStoragePairByIndex removes pair of Destination and Source
func (store *Storage) RemoveStoragePairByIndex(i int) {
	store.Sources = tools2.RemoveSliceElementByIndex(store.Sources, i)
	store.Destinations = tools2.RemoveSliceElementByIndex(store.Destinations, i)
}

// ClearStorageDestinations removes all elements from destinations list
func (store *Storage) ClearStorageDestinations() {
	store.Destinations = []string{}
	logger.Log.Debug("Storage Destinations were cleared")
}

// GenerateDestinations generating destinations from destination folder
// and source filename corresponding to this destination
func (store *Storage) GenerateDestinations(path string) error {
	for _, src := range store.Sources {
		filename := strings.Split(src, string(os.PathSeparator))
		dest := path + string(os.PathSeparator) + filename[len(filename)-1]
		store.AppendDestination(dest)
	}

	logger.Log.Debug("Destinations were generated")

	return nil
}

// Get destination of source by index
func (store Storage) getDestinationOfSource(source string) string {
	for i := range store.Sources {
		if store.Sources[i] == source {
			return store.Destinations[i]
		}
	}
	return ""
}

// GetIndexOfSource returns index of given source
func (store Storage) GetIndexOfSource(source string) int {
	for i, value := range store.Sources {
		if source == value {
			return i
		}
	}
	return -1
}

func (store Storage) String() string {
	ret := fmt.Sprintf("StorageType\n->Sources\n")
	for i, src := range store.Sources {
		ret += fmt.Sprintf("--> %d: [%v]\n", i, src)
	}
	ret += "\n->Destinations\n"
	for i, dest := range store.Destinations {
		ret += fmt.Sprintf("--> %d: [%v]\n", i, dest)
	}
	return ret
}
