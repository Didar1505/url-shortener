package store

type URLStore interface {
	Save(code string, originalURL string) error
	Get(code string) (string, bool)
}