package models

type Query struct {
	Title    string `query:"title"`
	Category string `query:"category"`
}
