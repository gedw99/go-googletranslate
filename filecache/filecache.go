package filecache

import (
	"container/list"
	"encoding/json"
	"os"
	"sync"
)

/*
 * FileCache is a simple LRU cache that stores its data in a file.
 * It is thread-safe.
 */
type FileCache struct {
	storage  map[string]*list.Element
	lru      *list.List
	filename string
	maxSize  int
	mutex    sync.RWMutex
}

type FileCacheItem struct {
	Key   string
	Value string
}

func NewFileCache(filename string, maxSize int) (*FileCache, error) {

	storage := make(map[string]*list.Element)

	if _, err := os.Stat(filename); err == nil {

		file, err := os.Open(filename)

		if err != nil {
			return nil, err
		}

		defer file.Close()

		var jsonData map[string]string
		err = json.NewDecoder(file).Decode(&jsonData)

		if err != nil {
			return nil, err
		}

		lru := list.New()
		for k, v := range jsonData {
			storage[k] = lru.PushBack(FileCacheItem{k, v})
		}

	}

	return &FileCache{storage, list.New(), filename, maxSize, sync.RWMutex{}}, nil

}

func (fc *FileCache) Get(key string) (string, bool) {

	fc.mutex.RLock()
	defer fc.mutex.RUnlock()

	element, found := fc.storage[key]

	if !found {
		return "", false
	}

	fc.lru.MoveToFront(element)

	item := element.Value.(FileCacheItem)
	return item.Value, true

}

func (fc *FileCache) Set(key string, value string) error {

	fc.mutex.Lock()
	defer fc.mutex.Unlock()

	if element, found := fc.storage[key]; found {

		fc.lru.MoveToFront(element)

		item := element.Value.(FileCacheItem)
		item.Value = value

	} else {

		if len(fc.storage) >= fc.maxSize {

			element := fc.lru.Back()
			fc.lru.Remove(element)

			item := element.Value.(FileCacheItem)
			delete(fc.storage, item.Key)

		}

		fc.storage[key] = fc.lru.PushFront(FileCacheItem{key, value})

	}

	jsonData := make(map[string]string)
	for k, v := range fc.storage {
		jsonData[k] = v.Value.(FileCacheItem).Value
	}

	file, err := os.Create(fc.filename)
	if err != nil {
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(jsonData)

	return err

}
