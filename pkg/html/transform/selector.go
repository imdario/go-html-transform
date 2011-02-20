/*
 Copyright 2010 Jeremy Wall (jeremy@marzhillstudios.com)
 Use of this source code is governed by the Artistic License 2.0.
 That License is included in the LICENSE file.
*/
package transform

import (
	. "html"
	v "container/vector"
	s "strings"
)

type SelectorQuery struct {
	*v.Vector
}

type Selector struct {
	Type    byte
	TagType string
	Key     string
	Val     string
}

const (
	TAGNAME byte = iota // zero value so the default
	CLASS   byte = '.'
	ID      byte = '#'
	PSEUDO  byte = ':'
	ANY     byte = '*'
	ATTR    byte = '['
)

func newAnyTagClassOrIdSelector(str string) *Selector {
	return &Selector{
		Type:    str[0],
		TagType: "*",
		Val:     str[1:],
	}
}

func newAnyTagSelector(str string) *Selector {
	return &Selector{
		Type:    str[0],
		TagType: "*",
	}
}

func splitAttrs(str string) []string {
	attrs := s.FieldsFunc(str[1:len(str)-1], func(c int) bool {
		if c == '=' {
			return true
		}
		return false
	})
	return attrs
}

func newAnyTagAttrSelector(str string) *Selector {
	attrs := splitAttrs(str)
	return &Selector{
		TagType: "*",
		Type:    str[0],
		Key:     attrs[0],
		Val:     attrs[1],
	}
}

func newTagNameSelector(str string) *Selector {
	return &Selector{
		Type:    TAGNAME,
		TagType: str,
	}
}

func newTagNameWithConstraints(str string, i int) *Selector {
	// TODO(jwall): indexAny use [CLASS,...]
	var selector = new(Selector)
	switch str[i] {
	case CLASS, ID, PSEUDO: // with class or id
		selector = newAnyTagClassOrIdSelector(str[i:])
	case ATTR: // with attribute
		selector = newAnyTagAttrSelector(str[i:])
	default:
		panic("Invalid constraint type for the tagname selector")
	}
	selector.TagType = str[0:i]
	//selector.Type = TAGNAME
	return selector
}

func NewSelector(str string) *Selector {
	str = s.TrimSpace(str) // trim whitespace
	var selector *Selector
	switch str[0] {
	case CLASS, ID: // Any tagname with class or id
		selector = newAnyTagClassOrIdSelector(str)
	case ANY: // Any tagname
		selector = newAnyTagSelector(str)
	case ATTR: // any tagname with attribute
		selector = newAnyTagAttrSelector(str)
	default: // TAGNAME
		// TODO(jwall): indexAny use [CLASS,...]
		if i := s.IndexAny(str, ".:#["); i != -1 {
			selector = newTagNameWithConstraints(str, i)
		} else { // just a tagname
			selector = newTagNameSelector(str)
		}
	}
	return selector
}

func NewSelectorQuery(sel ...string) *SelectorQuery {
	q := SelectorQuery{}
	for _, str := range sel {
		q.Insert(0, *NewSelector(str))
	}
	return &q
}

func testNode(node *Node, sel Selector) bool {
	/*
	if sel.TagType == "*" {
		attrs := node.Attr
		// TODO(jwall): abstract this out
		switch sel.Type {
		case ID:
			if attrs["id"] == sel.Val {
				return true
			}
		case CLASS:
			if attrs["class"] == sel.Val {
				return true
			}
		case ATTR:
			if attrs[sel.Key] == sel.Val {
				return true
			}
			//case PSEUDO:
			//TODO(jwall): implement these
		}
	} else {
		if node.nodeValue == sel.TagType {
			attrs := node.nodeAttributes
			switch sel.Type {
			case ID:
				if attrs["id"] == sel.Val {
					return true
				}
			case CLASS:
				if attrs["class"] == sel.Val {
					return true
				}
			case ATTR:
				if attrs[sel.Key] == sel.Val {
					return true
				}
			//case PSEUDO:
			//TODO(jwall): implement these
			default:
				return true
			}
		}
	}
	*/
	return false
}

/*
 Apply the css selector to a document.

 Returns a Vector of nodes that the selector matched.
*/
func (sel *SelectorQuery) Apply(doc *Document) *v.Vector {
	interesting := new(v.Vector)
	return interesting
}

/*
 Replace each node the selector matches with the passed in node.

 Applies the selector against the doc and replaces the returned
 Nodes with the passed in n HtmlNode.
*/
func (sel *SelectorQuery) Replace(doc *Document, n *Node) {
	nv := sel.Apply(doc)
	for i := 0; i <= nv.Len(); i++ {
		// Change to take into account new usage of *Node
		//nv.At(i).(*Node).Copy(n)
	}
	return
}
