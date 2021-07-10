package storage

import "fmt"

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

func (store Storage) getDestinationOfSource(source string) string {
	for i := range store.Sources {
		if store.Sources[i] == source {
			return store.Destinations[i]
		}
	}
	return ""
}

func (store Storage) String() string {
	ret := fmt.Sprintf("StorageType\n->Sources")
	for i, src := range store.Sources {
		ret += fmt.Sprintf("--> %d: [%v]\n", i, src)
	}
	ret += "\n->Destinations\n"
	for i, dest := range store.Destinations {
		ret += fmt.Sprintf("--> %d: [%v]\n", i, dest)
	}
	return ret
}
