package footer

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Kind uint

const (
	Mini Kind = iota
	Mega
)

func (k Kind) String() string {
	switch k {
	case Mini:
		return "mini"
	case Mega:
		return "mega"
	}
	return ""
}

type Footer struct {
	vecty.Core
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (f *Footer) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+f.Kind.String()+"-footer"] = true
	return elem.Footer(
		c, f.Children,
	)
}

type Section struct {
	Kind     Kind
	Left     bool
	Right    bool
	Children vecty.MarkupOrComponentOrHTML
}

func (s *Section) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	pos := ""
	if s.Left {
		pos = "left"
	}
	if s.Right {
		pos = "right"
	}
	c["mdl-"+s.Kind.String()+"-footer__"+pos+"-section"] = true
	return elem.Div(
		c, s.Children,
	)
}

type Logo struct {
	Children vecty.MarkupOrComponentOrHTML
}

func (l *Logo) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-logo"] = true
	return elem.Div(
		c, l.Children,
	)
}

type SociatButton struct {
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (s *SociatButton) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+s.Kind.String()+"-footer__social-btn"] = true
	return elem.Div(
		c, s.Children,
	)
}

type LinkList struct {
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (l *LinkList) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+l.Kind.String()+"-footer__link-list"] = true
	return elem.Div(
		c, l.Children,
	)
}
