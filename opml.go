package opml

import (
	//"strings",
	"encoding/xml"
)

const (
	RFC822 = "Mon, 02 Jan 2006 15:04:05 MST"
)

type Opml1 struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head
	Body    Body
}

type Opml2 struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head2
	Body    Body2
}

type Head struct {
	XMLName         xml.Name `xml:"head"`
	Title           string   `xml:"title"`
	DateCreated     string   `xml:"dateCreated"`
	DateModified    string   `xml:"dateModified"`
	OwnerName       string   `xml:"ownerName"`
	OwnerEmail      string   `xml:"ownerEmail"`
	ExpansionState  string   `xml:"expansionState"`
	VertScrollState string   `xml:"vertScrollState"`
	WindowTop       string   `xml:"windowTop"`
	WindowLeft      string   `xml:"windowLeft"`
	WindowBottom    string   `xml:"windowBottom"`
	WindowRight     string   `xml:"windowRight"`
}

type Head2 struct {
	XMLName         xml.Name `xml:"head"`
	Title           string   `xml:"title,omitempty"`
	DateCreated     string   `xml:"dateCreated,omitempty"`
	DateModified    string   `xml:"dateModified,omitempty"`
	OwnerName       string   `xml:"ownerName,omitempty"`
	OwnerEmail      string   `xml:"ownerEmail,omitempty"`
	OwnerId         string   `xml:"ownerId,omitempty"`
	Docs            string   `xml:"docs,omitempty"`
	ExpansionState  string   `xml:"expansionState,omitempty"`
	VertScrollState string   `xml:"vertScrollState,omitempty"`
	WindowTop       string   `xml:"windowTop,omitempty"`
	WindowLeft      string   `xml:"windowLeft,omitempty"`
	WindowBottom    string   `xml:"windowBottom,omitempty"`
	WindowRight     string   `xml:"windowRight,omitempty"`
}

type Body struct {
	XMLName  xml.Name  `xml:"body"`
	Outlines []Outline `xml:"outline"`
}

type Outline struct {
	XMLName      xml.Name  `xml:"outline"`
	Text         string    `xml:"text,attr,omitempty"`
	Title        string    `xml:"title,attr,omitempty"`
	Type         string    `xml:"type,attr,omitempty"`
	XmlUrl       string    `xml:"xmlUrl,attr,omitempty"`
	HtmlUrl      string    `xml:"htmlUrl,attr,omitempty"`
	IsComment    string    `xml:"isComment,attr,omitempty"`
	IsBreakpoint string    `xml:"isBreakpoint,attr,omitempty"`
	Outlines     []Outline `xml:"outline"`
}

type Body2 struct {
	XMLName  xml.Name   `xml:"body"`
	Outlines []Outline2 `xml:"outline"`
}

type Outline2 struct {
	XMLName      xml.Name   `xml:"outline"`
	Text         string     `xml:"text,attr,omitempty"`
	Title        string     `xml:"title,attr,omitempty"`
	Type         string     `xml:"type,attr,omitempty"`
	XmlUrl       string     `xml:"xmlUrl,attr,omitempty"`
	HtmlUrl      string     `xml:"htmlUrl,attr,omitempty"`
	IsComment    string     `xml:"isComment,attr,omitempty"`
	IsBreakpoint string     `xml:"isBreakpoint,attr,omitempty"`
	Created      string     `xml:"created,attr,omitempty"`
	Category     string     `xml:"category,attr,omitempty"`
	Description  string     `xml:"description,attr,omitempty"`
	Language     string     `xml:"language,attr,omitempty"`
	Version      string     `xml:"version,attr,omitempty"`
	Url          string     `xml:"url,attr,omitempty"`
	Outlines     []Outline2 `xml:"outline"`
}
