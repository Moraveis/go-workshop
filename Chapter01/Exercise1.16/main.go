// Exercise 1.16: Constants

package main

import "fmt"

const GlobalLimit = 100

const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}

	cache[key] = val
}

// GetBook - get value by key stored on cache
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// SetBook - add a value on cache
func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// GetCD - get value by key stored on cache
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// SetCD - add a value on cache
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}
func main() {

	cache = make(map[string]string)

	SetBook("1234-5678", "Get Ready To Go")

	SetCD("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book :", GetBook("1234-5678"))

	fmt.Println("CD :", GetCD("1234-5678"))

}
