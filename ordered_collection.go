package main

type OrderedCollection struct {
	Collection
	summary      string
	totalItems   int64
	orderedItems []Object
}
