package server

import (
	"time"
)

type trend struct {
	Name       string
	ID         string
	Heat       int64
	Length     int64
	CreateTime time.Time
}

type recommend struct {
	ID   int
	Name string
}

type list struct {
	Torrent []listdata
	Count   int64
}

// use by es
type listdata struct {
	Name       string
	Infohash   string
	Length     int64
	FileCount  int64
	Heat       int64
	CreateTime time.Time
}

// use by es
type estorrent struct {
	Name       string
	Infohash   string
	Length     int64
	FileCount  int64
	Heat       int64
	Files      []file
	CreateTime time.Time
}

type file struct {
	Name   string
	Length int64
}
