package models

import (
	"encoding/json"
	"fmt"
)

type Root struct {
	Node *Node `json:"root"`
}

type Node struct {
	Datatype Datatype `json:"datatype"`
	Fields   []Field  `json:"fields"`
	Nodes    []*Node  `json:"nodes"`
}

type Datatype struct {
	Info    Datatypes   `json:"info"`
	Content ContentData `json:"content"`
}

type Field struct {
	Info    Fields        `json:"info"`
	Content ContentFields `json:"content"`
}

func (n Node) MarshalJSON() ([]byte, error) {
	// Create a custom structure that mirrors Node but without circular references
	type CustomNode struct {
		Datatype Datatype `json:"datatype"`
		Fields   []Field  `json:"fields"`
		Nodes    []*Node  `json:"nodes,omitempty"`
	}

	// Create a copy to avoid circular reference
	var nodes []*Node
	if n.Nodes != nil {
		// Avoid self-references
		for _, node := range n.Nodes {
			if node != &n { // Avoid self-reference
				nodes = append(nodes, node)
			}
		}
	}

	custom := CustomNode{
		Datatype: n.Datatype,
		Fields:   n.Fields,
		Nodes:    nodes,
	}

	return json.Marshal(custom)
}

func (n *Node) UnmarshalJSON(data []byte) error {
	type NodeAlias Node // Create an alias to avoid infinite recursion

	aux := &struct {
		*NodeAlias
	}{
		NodeAlias: (*NodeAlias)(n),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	return nil
}

// Render renders this node and all its children recursively
func (r Root) Render() string {
	j, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	return string(j)
}

type ContentData struct {
	ContentDataID int64      `json:"content_data_id"`
	ParentID      NullInt64  `json:"parent_id"`
	RouteID       int64      `json:"route_id"`
	DatatypeID    int64      `json:"datatype_id"`
	AuthorID      int64      `json:"author_id"`
	DateCreated   NullString `json:"date_created"`
	DateModified  NullString `json:"date_modified"`
	History       NullString `json:"history"`
}

type ContentFields struct {
	ContentFieldID int64      `json:"content_field_id"`
	RouteID        NullInt64  `json:"route_id"`
	ContentDataID  int64      `json:"content_data_id"`
	FieldID        int64      `json:"field_id"`
	FieldValue     string     `json:"field_value"`
	AuthorID       int64      `json:"author_id"`
	DateCreated    NullString `json:"date_created"`
	DateModified   NullString `json:"date_modified"`
	History        NullString `json:"history"`
}

type Datatypes struct {
	DatatypeID   int64      `json:"datatype_id"`
	ParentID     NullInt64  `json:"parent_id"`
	Label        string     `json:"label"`
	Type         string     `json:"type"`
	AuthorID     int64      `json:"author_id"`
	DateCreated  NullString `json:"date_created"`
	DateModified NullString `json:"date_modified"`
	History      NullString `json:"history"`
}

type Fields struct {
	FieldID      int64      `json:"field_id"`
	ParentID     NullInt64  `json:"parent_id"`
	Label        string     `json:"label"`
	Data         string     `json:"data"`
	Type         string     `json:"type"`
	AuthorID     int64      `json:"author_id"`
	DateCreated  NullString `json:"date_created"`
	DateModified NullString `json:"date_modified"`
	History      NullString `json:"history"`
}

type Media struct {
	MediaID      int64      `json:"media_id"`
	Name         NullString `json:"name"`
	DisplayName  NullString `json:"display_name"`
	Alt          NullString `json:"alt"`
	Caption      NullString `json:"caption"`
	Description  NullString `json:"description"`
	Class        NullString `json:"class"`
	Mimetype     NullString `json:"mimetype"`
	Dimensions   NullString `json:"dimensions"`
	Url          NullString `json:"url"`
	Srcset       NullString `json:"srcset"`
	AuthorID     int64      `json:"author_id"`
	DateCreated  NullString `json:"date_created"`
	DateModified NullString `json:"date_modified"`
}

type MediaDimensions struct {
	MdID        int64      `json:"md_id"`
	Label       NullString `json:"label"`
	Width       NullInt64  `json:"width"`
	Height      NullInt64  `json:"height"`
	AspectRatio NullString `json:"aspect_ratio"`
}
