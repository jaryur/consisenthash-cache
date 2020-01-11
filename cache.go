package consisenthash_cache

type Cache interface{

	Add(key,value interface{}) bool

	Remove(key interface{}) bool

	Get(key interface{}) (value interface{},ok bool)

	Len() int

}