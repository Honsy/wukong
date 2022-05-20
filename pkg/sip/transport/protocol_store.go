package transport

import "sync"

type protocolKey string

// Thread-safe protocols pool.
type protocolStore struct {
	protocols map[protocolKey]Protocol
	mu        sync.RWMutex
}

func newProtocolStore() *protocolStore {
	return &protocolStore{
		protocols: make(map[protocolKey]Protocol),
	}
}

func (store *protocolStore) put(key protocolKey, protocol Protocol) {
	store.mu.Lock()
	store.protocols[key] = protocol
	store.mu.Unlock()
}

func (store *protocolStore) get(key protocolKey) (Protocol, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	protocol, ok := store.protocols[key]
	return protocol, ok
}

func (store *protocolStore) drop(key protocolKey) bool {
	if _, ok := store.get(key); !ok {
		return false
	}
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.protocols, key)
	return true
}

func (store *protocolStore) all() []Protocol {
	all := make([]Protocol, 0)
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, protocol := range store.protocols {
		all = append(all, protocol)
	}

	return all
}
