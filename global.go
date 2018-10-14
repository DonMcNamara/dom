package dom

import (
	"github.com/dennwc/dom/js"
	"image"
)

var (
	Doc  = GetDocument()
	Body = Doc.GetElementsByTagName("body")[0]
)

type Value = js.JSRef

func ConsoleLog(args ...interface{}) {
	js.Get("console").Call("log", args...)
}

func Loop() {
	select {}
}

type Point = image.Point
type Rect = image.Rectangle
