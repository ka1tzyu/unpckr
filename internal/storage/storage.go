package storage

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/tools"
	"os"
	"strings"
)

type Storage struct {
	Sources      []string
	Destinations []string
}

func (store *Storage) AppendSource(source string) {
	store.Sources = append(store.Sources, source)
}

func (store *Storage) AppendDestination(destination string) {
	store.Destinations = append(store.Destinations, destination)
}

// RemoveStoragePairByIndex Removes pair of Destination and Source
func (store *Storage) RemoveStoragePairByIndex(i int) {
	store.Sources = tools.RemoveSliceElementByIndex(store.Sources, i)
	store.Destinations = tools.RemoveSliceElementByIndex(store.Destinations, i)
}

func (store *Storage) ClearStorageDestinations() {
	store.Destinations = []string{}
}

func (store *Storage) GenerateDestinations(path string) error {
	for _, src := range store.Sources {
		filename := strings.Split(src, string(os.PathSeparator))
		dest := path + string(os.PathSeparator) + filename[len(filename)-1]
		store.AppendDestination(dest)
	}
	return nil
}

func (store Storage) getDestinationOfSource(source string) string {
	for i := range store.Sources {
		if store.Sources[i] == source {
			return store.Destinations[i]
		}
	}
	return ""
}

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
