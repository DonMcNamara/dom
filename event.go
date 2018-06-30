package dom

import (
	"fmt"
	"syscall/js"
)

type EventTarget interface {
	Value
	AddEventListener(typ string, fnc func(e Event))
	// TODO: removeEventListener
	// TODO: dispatchEvent
}

type Event interface {
	Value
	Bubbles() bool
	Cancelable() bool
	Composed() bool
	CurrentTarget() *Element
	DefaultPrevented() bool
	Target() *Element
	Type() string
	IsTrusted() bool
	Path() NodeList
}

type EventConstructor func(e BaseEvent) Event

func RegisterEventType(typ string, fnc EventConstructor) {
	cl := global.Get(typ)
	if cl == js.Null() || cl == js.Undefined() {
		panic(fmt.Errorf("class undefined: %q", typ))
	}
	eventClasses = append(eventClasses, eventClass{
		Class: cl, New: fnc,
	})
}

func init() {
	RegisterEventType("MouseEvent", func(e BaseEvent) Event {
		return &MouseEvent{e}
	})
}

type eventClass struct {
	Class js.Value
	New   EventConstructor
}

var (
	eventClasses []eventClass
)

func convertEvent(v js.Value) Event {
	e := BaseEvent{v: v}
	// TODO: get class name directly
	for _, cl := range eventClasses {
		if v.InstanceOf(cl.Class) {
			return cl.New(e)
		}
	}
	return &e
}

type BaseEvent struct {
	v js.Value
}

func (e *BaseEvent) getBool(name string) bool {
	return e.v.Get(name).Bool()
}
func (e *BaseEvent) Bubbles() bool {
	return e.getBool("bubbles")
}

func (e *BaseEvent) Cancelable() bool {
	return e.getBool("cancelable")
}

func (e *BaseEvent) Composed() bool {
	return e.getBool("composed")
}

func (e *BaseEvent) CurrentTarget() *Element {
	return AsElement(e.v.Get("currentTarget"))
}

func (e *BaseEvent) DefaultPrevented() bool {
	return e.getBool("defaultPrevented")
}

func (e *BaseEvent) IsTrusted() bool {
	return e.getBool("isTrusted")
}

func (e *BaseEvent) JSValue() js.Value {
	return e.v
}

func (e *BaseEvent) Type() string {
	return e.v.Get("type").String()
}

func (e *BaseEvent) Target() *Element {
	return AsElement(e.v.Get("target"))
}

func (e *BaseEvent) Path() NodeList {
	return AsNodeList(e.v.Get("path"))
}

type MouseEvent struct {
	BaseEvent
}

func (e *MouseEvent) getPos(nameX, nameY string) (x, y float64) {
	x = e.v.Get(nameX).Float()
	y = e.v.Get(nameY).Float()
	return
}

func (e *MouseEvent) getPosPref(pref string) (x, y float64) {
	return e.getPos(pref+"X", pref+"Y")
}

type MouseButton int

func (e *MouseEvent) Button() MouseButton {
	return MouseButton(e.v.Get("button").Int())
}

func (e *MouseEvent) ClientPos() (x, y float64) {
	return e.getPosPref("client")
}

func (e *MouseEvent) OffsetPos() (x, y float64) {
	return e.getPosPref("offset")
}

func (e *MouseEvent) PagePos() (x, y float64) {
	return e.getPosPref("page")
}

func (e *MouseEvent) ScreenPos() (x, y float64) {
	return e.getPosPref("screen")
}

func (e *MouseEvent) AltKey() bool {
	return e.v.Get("altKey").Bool()
}

func (e *MouseEvent) CtrlKey() bool {
	return e.v.Get("ctrlKey").Bool()
}

func (e *MouseEvent) ShiftKey() bool {
	return e.v.Get("shiftKey").Bool()
}

func (e *MouseEvent) MetaKey() bool {
	return e.v.Get("metaKey").Bool()
}
