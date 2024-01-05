package storage

type IStorage interface {
	Get(key string) Result
	Set(key string, value string) Result
	Delete(key string) Result
}

type Result struct {
	Out string
}
