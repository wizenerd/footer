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
	vecty.Core
	Kind     Kind
	Top      bool
	Left     bool
	Right    bool
	Bottom   bool
	Middle   bool
	Children vecty.MarkupOrComponentOrHTML
}

func (s *Section) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	pos := ""
	if s.Top {
		pos = "top"
	}
	if s.Left {
		pos = "left"
	}
	if s.Right {
		pos = "right"
	}
	if s.Bottom {
		pos = "bottom"
	}
	if s.Middle {
		pos = "middle"
	}
	c["mdl-"+s.Kind.String()+"-footer__"+pos+"-section"] = true
	return elem.Div(
		c, s.Children,
	)
}

type Logo struct {
	vecty.Core
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
	vecty.Core
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (s *SociatButton) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+s.Kind.String()+"-footer__social-btn"] = true
	return elem.Button(
		c, s.Children,
	)
}

type DropDown struct {
	vecty.Core
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (s *DropDown) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+s.Kind.String()+"-footer__drop-down-section"] = true
	return elem.Div(
		c, s.Children,
	)
}

type LinkList struct {
	vecty.Core
	Kind     Kind
	Children vecty.MarkupOrComponentOrHTML
}

func (l *LinkList) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+l.Kind.String()+"-footer__link-list"] = true
	return elem.UnorderedList(
		c, l.Children,
	)
}

func H1(k Kind, m ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	c := make(vecty.ClassMap)
	c["mdl-"+k.String()+"-footer__heading"] = true
	var v vecty.List
	v = append(v, c)
	v = append(v, m...)
	return elem.Heading1(v)
}
