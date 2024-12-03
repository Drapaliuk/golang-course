package lru

type LruCache interface {
	// Якщо наш кеш вже повний (ми досягли нашого capacity)
	// то має видалитись той елемент, який ми до якого ми доступались (читали) найдавніше
	Put(key, value string)
	Get(key string) (string, bool)
}

func NewLruCache(capacity int) LruCache {
	// TODO: Implement
	return nil
}
