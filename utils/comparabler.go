package utils

type Comparabler interface {
	More(interface{}) bool
	Less(interface{}) bool
}
