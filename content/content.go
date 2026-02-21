package content

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type Meta struct {
	AuthorID     string `json:"authorId"`
	RouteID      string `json:"routeId"`
	DateCreated  string `json:"dateCreated"`
	DateModified string `json:"dateModified"`
}

// ──────────────────────────────────────
// Root types
// ──────────────────────────────────────

type Page struct {
	Meta            Meta              `json:"_meta"`
	ID              string            `json:"id"`
	Type            string            `json:"type"`
	Title           string            `json:"title"`
	Slug            string            `json:"slug"`
	MetaTitle       string            `json:"metaTitle"`
	MetaDescription string            `json:"metaDescription"`
	Published       bool              `json:"published"`
	RawChildren     []json.RawMessage `json:"children"`
}

type Post struct {
	Meta            Meta              `json:"_meta"`
	ID              string            `json:"id"`
	Type            string            `json:"type"`
	Title           string            `json:"title"`
	Slug            string            `json:"slug"`
	MetaTitle       string            `json:"metaTitle"`
	MetaDescription string            `json:"metaDescription"`
	Published       bool              `json:"published"`
	RawChildren     []json.RawMessage `json:"children"`
}

type CaseStudy struct {
	Meta            Meta              `json:"_meta"`
	ID              string            `json:"id"`
	Type            string            `json:"type"`
	Title           string            `json:"title"`
	Slug            string            `json:"slug"`
	ClientName      string            `json:"clientName"`
	Description     string            `json:"description"`
	Challenge       string            `json:"challenge"`
	Solution        string            `json:"solution"`
	Results         string            `json:"results"`
	FeaturedImage   string            `json:"featuredImage"`
	Published       bool              `json:"published"`
	RawChildren     []json.RawMessage `json:"children"`
}

type Documentation struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Published   bool              `json:"published"`
	RawChildren []json.RawMessage `json:"children"`
}

type Menu struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Position    string            `json:"position"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

// ──────────────────────────────────────
// Child union
// ──────────────────────────────────────

// Child is a discriminated union: exactly one field is non-nil.
type Child struct {
	// Layout
	Row    *Row
	Columns *Columns
	Grid   *Grid
	Area   *Area

	// Content blocks
	CTA      *CTA
	Card     *Card
	RichText *RichText
	Text     *Text
	Image    *Image
	Button   *Button

	// Post
	PostContent *PostContent

	// Doc components
	Section    *Section
	CodeBlock  *CodeBlock
	Reference  *Reference
	StepHeader *StepHeader

	// Menu components
	MenuLink       *MenuLink
	MenuList       *MenuList
	MenuListLink   *MenuListLink
	MenuNestedList *MenuNestedList
	MenuNestedLink *MenuNestedLink
}

// ──────────────────────────────────────
// Layout types
// ──────────────────────────────────────

type Row struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	FullWidth   bool              `json:"fullWidth"`
	Columns     []Columns         `json:"columns"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type Columns struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Count       string            `json:"count"`
	CTAs        []CTA             `json:"ctas"`
	Cards       []Card            `json:"cards"`
	RichTexts   []RichText        `json:"rich texts"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

// GridStyle returns a CSS style string for grid-template-columns.
// A plain number (e.g. "3") produces "grid-template-columns: 1fr 1fr 1fr".
// Anything else (e.g. "1fr 2fr 1fr") is used verbatim.
// Returns empty string when Count is blank.
func (c Columns) GridStyle() string {
	count := strings.TrimSpace(c.Count)
	if count == "" {
		return ""
	}
	n, err := strconv.Atoi(count)
	if err != nil {
		return "grid-template-columns: " + count
	}
	parts := make([]string, n)
	for i := range n {
		parts[i] = "1fr"
	}
	return "grid-template-columns: " + strings.Join(parts, " ")
}

type Grid struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Columns     string            `json:"columns"`
	Rows        string            `json:"rows"`
	Gap         string            `json:"gap"`
	Areas       []Area            `json:"areas"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type Area struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	ColumnStart int               `json:"columnStart"`
	ColumnEnd   int               `json:"columnEnd"`
	RowStart    int               `json:"rowStart"`
	RowEnd      int               `json:"rowEnd"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

// ──────────────────────────────────────
// Content blocks
// ──────────────────────────────────────

type CTA struct {
	Meta       Meta   `json:"_meta"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	Heading    string `json:"heading"`
	Subheading string `json:"subheading"`
	ButtonText string `json:"buttonText"`
	ButtonURL  string `json:"buttonUrl"`
}

type Card struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	LinkURL     string `json:"linkUrl"`
}

type RichText struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Text struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Image struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	ImageID string `json:"image"`
	AltText string `json:"altText"`
	Caption string `json:"caption"`
}

type Button struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Label   string `json:"label"`
	URL     string `json:"url"`
	Variant string `json:"variant"`
}

type Animation struct {
	Meta       Meta   `json:"_meta"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	AnimType   string `json:"animationType"`
	Duration   string `json:"duration"`
	Delay      string `json:"delay"`
	Easing     string `json:"easing"`
	Direction  string `json:"direction"`
	Iterations string `json:"iterations"`
}

type Settings struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Margin  string `json:"margin"`
	Padding string `json:"padding"`
}

// ──────────────────────────────────────
// Post children
// ──────────────────────────────────────

type PostContent struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// ──────────────────────────────────────
// Doc components
// ──────────────────────────────────────

type Section struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Heading string `json:"heading"`
	Content string `json:"content"`
}

type CodeBlock struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Language string `json:"language"`
	Code     string `json:"code"`
	Caption  string `json:"caption"`
}

type Reference struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type StepHeader struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	StepNumber  int    `json:"stepNumber"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ──────────────────────────────────────
// Menu components
// ──────────────────────────────────────

type MenuLink struct {
	Meta   Meta   `json:"_meta"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	Label  string `json:"label"`
	URL    string `json:"url"`
	Target string `json:"target"`
	Icon   string `json:"icon"`
}

type MenuList struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Label       string            `json:"label"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type MenuListLink struct {
	Meta   Meta   `json:"_meta"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	Label  string `json:"label"`
	URL    string `json:"url"`
	Target string `json:"target"`
}

type MenuNestedList struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Label       string            `json:"label"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type MenuNestedLink struct {
	Meta   Meta   `json:"_meta"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	Label  string `json:"label"`
	URL    string `json:"url"`
	Target string `json:"target"`
}

// ──────────────────────────────────────
// Resolved data for rendering
// ──────────────────────────────────────

type PageData struct {
	Page     Page
	Children []Child
}

// ──────────────────────────────────────
// Parser
// ──────────────────────────────────────

type typeProbe struct {
	Type string `json:"type"`
}

// ParseChildren resolves mixed JSON nodes into typed Child values.
// Unknown types are logged and skipped.
func ParseChildren(raw []json.RawMessage) ([]Child, error) {
	children := make([]Child, 0, len(raw))
	for i, msg := range raw {
		var probe typeProbe
		if err := json.Unmarshal(msg, &probe); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: probe type: %w", i, err)
		}
		child, err := unmarshalChild(i, probe.Type, msg)
		if err != nil {
			return nil, err
		}
		if child != nil {
			children = append(children, *child)
		}
	}
	return children, nil
}

func unmarshalChild(idx int, typeName string, msg json.RawMessage) (*Child, error) {
	switch typeName {

	// Layout
	case "Row":
		var v Row
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Row: %w", idx, err)
		}
		if err := resolveRow(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Row: %w", idx, err)
		}
		return &Child{Row: &v}, nil
	case "Columns":
		var v Columns
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Columns: %w", idx, err)
		}
		if err := resolveColumns(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Columns: %w", idx, err)
		}
		return &Child{Columns: &v}, nil
	case "Grid":
		var v Grid
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Grid: %w", idx, err)
		}
		if err := resolveGrid(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Grid: %w", idx, err)
		}
		return &Child{Grid: &v}, nil
	case "Area":
		var v Area
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Area: %w", idx, err)
		}
		if err := resolveArea(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Area: %w", idx, err)
		}
		return &Child{Area: &v}, nil

	// Content blocks
	case "CTA":
		var v CTA
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal CTA: %w", idx, err)
		}
		return &Child{CTA: &v}, nil
	case "Card":
		var v Card
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Card: %w", idx, err)
		}
		return &Child{Card: &v}, nil
	case "Rich Text":
		var v RichText
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal RichText: %w", idx, err)
		}
		return &Child{RichText: &v}, nil
	case "Text":
		var v Text
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Text: %w", idx, err)
		}
		return &Child{Text: &v}, nil
	case "Image":
		var v Image
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Image: %w", idx, err)
		}
		return &Child{Image: &v}, nil
	case "Button":
		var v Button
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Button: %w", idx, err)
		}
		return &Child{Button: &v}, nil

	// Post children
	case "Content":
		var v PostContent
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal PostContent: %w", idx, err)
		}
		return &Child{PostContent: &v}, nil

	// Doc components
	case "Section":
		var v Section
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Section: %w", idx, err)
		}
		return &Child{Section: &v}, nil
	case "Code Block":
		var v CodeBlock
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal CodeBlock: %w", idx, err)
		}
		return &Child{CodeBlock: &v}, nil
	case "Reference":
		var v Reference
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Reference: %w", idx, err)
		}
		return &Child{Reference: &v}, nil
	case "Step Header":
		var v StepHeader
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal StepHeader: %w", idx, err)
		}
		return &Child{StepHeader: &v}, nil

	// Menu components
	case "Menu Link":
		var v MenuLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal MenuLink: %w", idx, err)
		}
		return &Child{MenuLink: &v}, nil
	case "Menu List":
		var v MenuList
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal MenuList: %w", idx, err)
		}
		if err := resolveMenuList(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve MenuList: %w", idx, err)
		}
		return &Child{MenuList: &v}, nil
	case "Menu List Link":
		var v MenuListLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal MenuListLink: %w", idx, err)
		}
		return &Child{MenuListLink: &v}, nil
	case "Menu Nested List":
		var v MenuNestedList
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal MenuNestedList: %w", idx, err)
		}
		if err := resolveMenuNestedList(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve MenuNestedList: %w", idx, err)
		}
		return &Child{MenuNestedList: &v}, nil
	case "Menu Nested Link":
		var v MenuNestedLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal MenuNestedLink: %w", idx, err)
		}
		return &Child{MenuNestedLink: &v}, nil

	// Settings and Animation are metadata, not rendered — skip silently
	case "Settings", "Animation":
		return nil, nil

	default:
		slog.Warn("parseChildren: unknown node type, skipping", "index", idx, "type", typeName)
		return nil, nil
	}
}

// ──────────────────────────────────────
// Tree resolution
// ──────────────────────────────────────

func resolveRow(row *Row) error {
	if len(row.RawChildren) > 0 {
		resolved, err := ParseChildren(row.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve row %s children: %w", row.ID, err)
		}
		row.Resolved = resolved
	}
	for i := range row.Columns {
		if err := resolveColumns(&row.Columns[i]); err != nil {
			return err
		}
	}
	return nil
}

func resolveColumns(cols *Columns) error {
	if len(cols.RawChildren) > 0 {
		resolved, err := ParseChildren(cols.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve columns %s children: %w", cols.ID, err)
		}
		cols.Resolved = resolved
	}
	return nil
}

func resolveGrid(grid *Grid) error {
	if len(grid.RawChildren) > 0 {
		resolved, err := ParseChildren(grid.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve grid %s children: %w", grid.ID, err)
		}
		grid.Resolved = resolved
	}
	for i := range grid.Areas {
		if err := resolveArea(&grid.Areas[i]); err != nil {
			return err
		}
	}
	return nil
}

func resolveArea(area *Area) error {
	if len(area.RawChildren) > 0 {
		resolved, err := ParseChildren(area.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve area %s children: %w", area.ID, err)
		}
		area.Resolved = resolved
	}
	return nil
}

// ResolveMenu parses a Menu's RawChildren into its Resolved field.
func ResolveMenu(menu *Menu) error {
	if len(menu.RawChildren) > 0 {
		resolved, err := ParseChildren(menu.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve menu %s children: %w", menu.ID, err)
		}
		menu.Resolved = resolved
	}
	return nil
}

func resolveMenuList(ml *MenuList) error {
	if len(ml.RawChildren) > 0 {
		resolved, err := ParseChildren(ml.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve menu list %s children: %w", ml.ID, err)
		}
		ml.Resolved = resolved
	}
	return nil
}

func resolveMenuNestedList(mnl *MenuNestedList) error {
	if len(mnl.RawChildren) > 0 {
		resolved, err := ParseChildren(mnl.RawChildren)
		if err != nil {
			return fmt.Errorf("resolve menu nested list %s children: %w", mnl.ID, err)
		}
		mnl.Resolved = resolved
	}
	return nil
}
