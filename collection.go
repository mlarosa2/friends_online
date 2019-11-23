package main

type Collection struct {
	Object
	summary    string
	totalItems int64
	items      []Object
}
