package models

import (
	"log"
	"testing"
)

func TestIsNullRetweeted(t *testing.T) {
	if ImageInfoIsRetweeted("aiueo") {
		log.Println("IM ON NOT OK")
		t.Fatal("failed test")
	} else {
		log.Println("I'm OK")
	}
}

func TestIsRetweeted(t *testing.T) {
	if ImageInfoIsRetweeted("http://pbs.twimg.com/media/DqcaS0JUwAAAqFJ.jpg") {
		log.Println("IM ON NOT OK")
	} else {
		log.Println("I'm OK")
		t.Fatal("failed test")
	}
}
