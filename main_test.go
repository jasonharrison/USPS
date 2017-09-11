package main

import (
	"github.com/atotto/clipboard"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtil(t *testing.T) {
	Convey("Command: format", t, func() {
		tracking := "9400109699936261459205"
		clipboard.WriteAll(tracking) // write tracking to clipboard
		formatcmd(nil)
		text, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}
		So(text, ShouldEqual, "9400 1096 9993 6261 4592 05")
	})
}
