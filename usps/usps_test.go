package usps

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFormat(t *testing.T) {
	Convey("Format tracking number", t, func() {
		So(Format("9400109699936261459205"), ShouldEqual, "9400 1096 9993 6261 4592 05")
	},
	)
}
