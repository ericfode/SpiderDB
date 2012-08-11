package spiderDB

type KeyNotFoundError string

type dbError string

func (e KeyNotFoundError) Error() string {}
