package main

import "fmt"

/*
依賴反轉原則
說明:
	- 高層模組不應依賴低層模組，兩者之間皆應依賴抽象介面
*/

// 高層模組依自己的需求，定義需要的抽象介面
type Cache interface {
	Set(key, value string) error
}

// 2. 高層模組
type UserRepository struct {
	cache Cache
}

func NewUserRepository(cache Cache) *UserRepository {
	return &UserRepository{
		cache: cache,
	}
}

func (u *UserRepository) Register(name string) error {
	return u.cache.Set("user", name)

}

// 3. 底層模組
type FileCache struct{}

func (f *FileCache) Set(key, value string) error {
	fmt.Printf("Assign %s to the %s key in file.\n", value, key)
	return nil
}

type RedisCache struct{}

func (r *RedisCache) Set(key, value string) error {
	fmt.Printf("Assign %s to the %s key in redis.\n", value, key)
	return nil
}

func main() {
	fileCacheUser := NewUserRepository(&FileCache{})
	fileCacheUser.Register("hao.chen")

	redisCacheUser := NewUserRepository(&RedisCache{})
	redisCacheUser.Register("hao")
}
