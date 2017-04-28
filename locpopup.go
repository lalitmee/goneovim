package gonvim

import (
	"github.com/dzhou121/ui"
)

// Locpopup is the location popup
type Locpopup struct {
	box     *ui.Box
	locType *SpanHandler
	text    *SpanHandler
}

func initLocpopup() *Locpopup {
	box := ui.NewHorizontalBox()

	typeHandler := &SpanHandler{}
	typeSpan := ui.NewArea(typeHandler)
	typeHandler.span = typeSpan

	textHandler := &SpanHandler{}
	textSpan := ui.NewArea(textHandler)
	textHandler.span = textSpan

	box.Append(textSpan, false)
	box.Append(typeSpan, false)

	return &Locpopup{
		box:     box,
		locType: typeHandler,
		text:    textHandler,
	}
}

func (l *Locpopup) show(loc map[string]interface{}) {
	locType := loc["type"].(string)
	switch locType {
	case "E":
		l.locType.SetText("Error")
		l.locType.SetFont(editor.font.font)
		l.locType.SetColor(newRGBA(255, 255, 255, 1))
		l.locType.SetBackground(newRGBA(204, 62, 68, 1))
		l.locType.span.SetPosition(5, editor.font.shift)
		l.locType.setSize(editor.font.width*5, editor.font.height)
	case "W":
		l.locType.SetText("Warning")
		l.locType.SetFont(editor.font.font)
		l.locType.SetColor(newRGBA(255, 255, 255, 1))
		l.locType.SetBackground(newRGBA(203, 203, 65, 1))
		l.locType.span.SetPosition(5, editor.font.shift)
		l.locType.setSize(editor.font.width*7, editor.font.height)
	}
	text := loc["text"].(string)
	l.text.SetText(text)
	l.text.SetFont(editor.font.font)
	l.text.SetColor(newRGBA(14, 17, 18, 1))
	l.text.SetBackground(newRGBA(212, 215, 214, 1))
	l.text.paddingLeft = editor.font.width*len(l.locType.text) + 10
	l.text.paddingRight = 5
	w, _ := l.text.getSize()
	l.text.setSize(w, editor.font.lineHeight)
	l.text.paddingTop = editor.font.shift
	ui.QueueMain(func() {
		l.box.SetPosition(0, editor.font.lineHeight)
		l.locType.span.Show()
		l.text.span.Show()
		l.locType.span.QueueRedrawAll()
		l.text.span.QueueRedrawAll()
		l.box.SetSize(w, editor.font.lineHeight)
		l.box.Show()
	})
}

func (l *Locpopup) hide() {
	ui.QueueMain(func() {
		l.box.Hide()
	})
}
