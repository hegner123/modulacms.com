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
	Description     string            `json:"description"`
	Slug            string            `json:"slug"`
	MetaTitle       string            `json:"meta_title"`
	MetaDescription string            `json:"meta_description"`
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
	Meta          Meta              `json:"_meta"`
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Title         string            `json:"title"`
	Slug          string            `json:"slug"`
	ClientName    string            `json:"clientName"`
	Description   string            `json:"description"`
	Challenge     string            `json:"challenge"`
	Solution      string            `json:"solution"`
	Results       string            `json:"results"`
	FeaturedImage string            `json:"featuredImage"`
	Published     bool              `json:"published"`
	RawChildren   []json.RawMessage `json:"children"`
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
	Row     *Row
	Columns *Columns
	Grid    *Grid
	Area    *Area

	// Marketing page sections
	HeroSection      *HeroSection
	CTASection       *CTASection
	BentoGridSection *BentoGridSection
	ContentSection   *ContentSection
	ContentReference *ContentReference

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
	DocSection *Section
	CodeBlock  *CodeBlock
	Reference  *Reference
	StepHeader *StepHeader

	// Marketing page sections (additional)
	StatsSection       *StatsSection
	StatItem           *StatItem
	FeatureSection     *FeatureSection
	FeatureItem        *FeatureItem
	PricingSection     *PricingSection
	PricingTier        *PricingTier
	PricingFeature     *PricingFeature
	TestimonialSection *TestimonialSection
	Testimonial        *Testimonial
	BlogSection        *BlogSection
	BlogPostCard       *BlogPostCard
	FAQSection         *FAQSection
	FAQItem            *FAQItem
	LogoCloudSection   *LogoCloudSection
	LogoItem           *LogoItem
	Banner             *Banner
	ContactSection     *ContactSection
	ContactDetail      *ContactDetail
	NewsletterSection  *NewsletterSection
	NewsletterDetail   *NewsletterDetail
	ErrorPage          *ErrorPage
	PopularPage        *PopularPage
	TeamSection        *TeamSection
	TeamMember         *TeamMember
	TeamSocialLink     *TeamSocialLink
	HeaderSection      *HeaderSection
	HeaderActionLink   *HeaderActionLink
	InlineTestimonial  *InlineTestimonial
	CTABenefit         *CTABenefit
	DocImage           *DocImage

	// Navigation types
	Footer       *Footer
	FooterColumn *FooterColumn
	FooterLink   *FooterLink
	FooterSocial *FooterSocial
	FooterText   *FooterText
	SocialLink   *SocialLink
	Header       *Header
	NavLink      *NavLink
	FlyoutMenu   *FlyoutMenu
	FlyoutLink   *FlyoutLink

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
// Marketing page sections
// ──────────────────────────────────────

type HeroSection struct {
	Meta             Meta   `json:"_meta"`
	ID               string `json:"id"`
	Type             string `json:"type"`
	Heading          string `json:"heading"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	ImageDark        string `json:"image_dark"`
	PrimaryCTAText   string `json:"primary_cta_text"`
	PrimaryCTAURL    string `json:"primary_cta_url"`
	SecondaryCTAText string `json:"secondary_cta_text"`
	SecondaryCTAURL  string `json:"secondary_cta_url"`
	AnnouncementText string `json:"announcement_text"`
	AnnouncementURL  string `json:"announcement_url"`
	ImageURL         string `json:"-"`
	ImageDarkURL     string `json:"-"`
}

type CTASection struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Heading     string `json:"heading"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CTAText     string `json:"cta_text"`
	CTAURL      string `json:"cta_url"`
	ImageURL    string `json:"-"`
}

type BentoGridSection struct {
	Meta     Meta              `json:"_meta"`
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Eyebrow  string            `json:"eyebrow"`
	Heading  string            `json:"heading"`
	RawCells []json.RawMessage `json:"bento_cells"`
	Cells    []BentoCell       `json:"-"`
}

type BentoCell struct {
	Meta         Meta   `json:"_meta"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	ImageDark    string `json:"image_dark"`
	Span         string `json:"span"`
	ImageURL     string `json:"-"`
	ImageDarkURL string `json:"-"`
}

type ContentSection struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Eyebrow  string `json:"eyebrow"`
	Heading  string `json:"heading"`
	Body     string `json:"body"`
	Image    string `json:"image"`
	CTAText  string `json:"cta_text"`
	CTAURL   string `json:"cta_url"`
	ImageURL string `json:"-"`
}

type ContentReference struct {
	Meta     Meta              `json:"_meta"`
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	RawMenus []json.RawMessage `json:"menus"`
	Menus    []Menu            `json:"-"`
}

type StatsSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type StatItem struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type FeatureSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	ImageDark   string            `json:"image_dark"`
	ImageURL    string            `json:"-"`
	ImageDkURL  string            `json:"-"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type FeatureItem struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type PricingSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type PricingTier struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Price       string            `json:"price"`
	Description string            `json:"description"`
	CTAText     string            `json:"cta_text"`
	CTAURL      string            `json:"cta_url"`
	Highlighted bool              `json:"highlighted"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type PricingFeature struct {
	Meta Meta   `json:"_meta"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type TestimonialSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type Testimonial struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Image    string `json:"image"`
	ImageURL string `json:"-"`
}

type BlogSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type BlogPostCard struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      string `json:"author"`
	Image       string `json:"image"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	ImageURL    string `json:"-"`
}

type FAQSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type FAQItem struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type LogoCloudSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type LogoItem struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	URL      string `json:"url"`
	ImageURL string `json:"-"`
}

type Banner struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	URL         string `json:"url"`
	CTAText     string `json:"cta_text"`
	Dismissable bool   `json:"dismissable"`
}

type ContactSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type ContactDetail struct {
	Meta  Meta   `json:"_meta"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
	Value string `json:"value"`
	Icon  string `json:"icon"`
}

type NewsletterSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type NewsletterDetail struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ErrorPage struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	StatusCode  int               `json:"status_code"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type PopularPage struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type TeamSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type TeamMember struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Role        string            `json:"role"`
	Image       string            `json:"image"`
	Bio         string            `json:"bio"`
	ImageURL    string            `json:"-"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type TeamSocialLink struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

type HeaderSection struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Eyebrow     string            `json:"eyebrow"`
	Heading     string            `json:"heading"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type HeaderActionLink struct {
	Meta  Meta   `json:"_meta"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
	URL   string `json:"url"`
	Style string `json:"style"`
}

type InlineTestimonial struct {
	Meta    Meta   `json:"_meta"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Quote   string `json:"quote"`
	Author  string `json:"author"`
	Role    string `json:"role"`
	Company string `json:"company"`
}

type CTABenefit struct {
	Meta Meta   `json:"_meta"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
	Icon string `json:"icon"`
}

type DocImage struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Image    string `json:"image"`
	Alt      string `json:"alt"`
	Caption  string `json:"caption"`
	ImageURL string `json:"-"`
}

// ──────────────────────────────────────
// Navigation types
// ──────────────────────────────────────

type Footer struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Copyright   string            `json:"copyright"`
	Description string            `json:"description"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type FooterColumn struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Title       string            `json:"title"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type FooterLink struct {
	Meta  Meta   `json:"_meta"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type FooterSocial struct {
	Meta  Meta   `json:"_meta"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

type FooterText struct {
	Meta Meta   `json:"_meta"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type SocialLink struct {
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Icon     string `json:"icon"`
}

type Header struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Logo        string            `json:"logo"`
	LogoDark    string            `json:"logo_dark"`
	LogoURL     string            `json:"-"`
	LogoDkURL   string            `json:"-"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type NavLink struct {
	Meta  Meta   `json:"_meta"`
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type FlyoutMenu struct {
	Meta        Meta              `json:"_meta"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Label       string            `json:"label"`
	RawChildren []json.RawMessage `json:"children"`
	Resolved    []Child           `json:"-"`
}

type FlyoutLink struct {
	Meta        Meta   `json:"_meta"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Icon        string `json:"icon"`
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
	Meta     Meta   `json:"_meta"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	ImageID  string `json:"image"`
	AltText  string `json:"altText"`
	Caption  string `json:"caption"`
	MediaURL string `json:"-"`
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

// BuildLog collects non-fatal issues during content resolution and media
// resolution so they can be surfaced in the browser console.
type BuildLog struct {
	entries []string
}

func (b *BuildLog) Add(msg string) {
	b.entries = append(b.entries, msg)
}

func (b *BuildLog) Entries() []string {
	return b.entries
}

// ScriptTag returns a <script> element that logs all entries to the
// browser console. Returns empty string when there are no entries.
func (b *BuildLog) ScriptTag() string {
	if len(b.entries) == 0 {
		return ""
	}
	encoded, err := json.Marshal(b.entries)
	if err != nil {
		return ""
	}
	return fmt.Sprintf(`<script>JSON.parse('%s').forEach(function(e){console.error("[build]",e)})</script>`, string(encoded))
}

type PageData struct {
	Page     Page
	Children []Child
	Log      *BuildLog
	DocsNav  []DocsNavSection
}

// DocsNavItem represents a single link in the docs sidebar navigation.
type DocsNavItem struct {
	Title     string
	Slug      string
	SortOrder int
}

// DocsNavSection groups navigation items under a heading derived from
// the first URL segment after /docs/.
type DocsNavSection struct {
	Title string
	Items []DocsNavItem
}

// GroupDocsNav groups a flat list of nav items into sections by their
// first path segment after /docs/. Items directly at /docs are placed
// in an untitled section at the top.
func GroupDocsNav(items []DocsNavItem) []DocsNavSection {
	type entry struct {
		key   string
		title string
	}
	var order []entry
	groups := map[string][]DocsNavItem{}

	for _, item := range items {
		trimmed := strings.TrimPrefix(item.Slug, "/docs")
		trimmed = strings.TrimPrefix(trimmed, "/")
		parts := strings.SplitN(trimmed, "/", 2)
		key := ""
		if len(parts) > 1 {
			key = parts[0]
		}
		if _, exists := groups[key]; !exists {
			title := segmentToTitle(key)
			order = append(order, entry{key: key, title: title})
		}
		groups[key] = append(groups[key], item)
	}

	sections := make([]DocsNavSection, 0, len(order))
	for _, e := range order {
		sections = append(sections, DocsNavSection{
			Title: e.title,
			Items: groups[e.key],
		})
	}
	return sections
}

// segmentToTitle converts a URL segment like "getting-started" to "Getting Started".
// Common acronyms (api, sdks, cli, css, etc.) are fully uppercased.
func segmentToTitle(seg string) string {
	if seg == "" {
		return ""
	}
	acronyms := map[string]string{
		"api": "API", "sdks": "SDKs", "sdk": "SDK", "cli": "CLI",
		"css": "CSS", "html": "HTML", "http": "HTTP", "url": "URL",
		"rbac": "RBAC", "oauth": "OAuth", "mcp": "MCP", "faq": "FAQ",
	}
	words := strings.Split(seg, "-")
	for i, w := range words {
		if replacement, ok := acronyms[w]; ok {
			words[i] = replacement
		} else if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// HeaderMenus returns all menus with position "header" from content
// references in the page tree.
func (pd PageData) HeaderMenus() []Menu {
	var menus []Menu
	for _, child := range pd.Children {
		if child.ContentReference == nil {
			continue
		}
		for _, m := range child.ContentReference.Menus {
			if m.Position == "header" {
				menus = append(menus, m)
			}
		}
	}
	return menus
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

	// Marketing page sections
	case "hero_section":
		var v HeroSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal HeroSection: %w", idx, err)
		}
		return &Child{HeroSection: &v}, nil
	case "cta_section":
		var v CTASection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal CTASection: %w", idx, err)
		}
		return &Child{CTASection: &v}, nil
	case "bento_grid_section":
		var v BentoGridSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal BentoGridSection: %w", idx, err)
		}
		if err := resolveBentoGrid(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve BentoGridSection: %w", idx, err)
		}
		return &Child{BentoGridSection: &v}, nil
	case "content_section":
		var v ContentSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal ContentSection: %w", idx, err)
		}
		return &Child{ContentSection: &v}, nil

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
	case "Column", "Columns":
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
	case "doc_section":
		var v Section
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal DocSection: %w", idx, err)
		}
		return &Child{DocSection: &v}, nil
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

	// Content references
	case "menu_reference", "doc-menu-reference":
		var v ContentReference
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal ContentReference: %w", idx, err)
		}
		if err := resolveContentReference(&v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve ContentReference: %w", idx, err)
		}
		return &Child{ContentReference: &v}, nil

	// Menu components
	case "menu_link", "menu_icon_link":
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

	// Marketing page sections (additional)
	case "stats_section":
		var v StatsSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal StatsSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "StatsSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve StatsSection: %w", idx, err)
		}
		return &Child{StatsSection: &v}, nil
	case "stat_item":
		var v StatItem
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal StatItem: %w", idx, err)
		}
		return &Child{StatItem: &v}, nil
	case "feature_section":
		var v FeatureSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FeatureSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "FeatureSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve FeatureSection: %w", idx, err)
		}
		return &Child{FeatureSection: &v}, nil
	case "feature_item":
		var v FeatureItem
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FeatureItem: %w", idx, err)
		}
		return &Child{FeatureItem: &v}, nil
	case "pricing_section":
		var v PricingSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal PricingSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "PricingSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve PricingSection: %w", idx, err)
		}
		return &Child{PricingSection: &v}, nil
	case "pricing_tier":
		var v PricingTier
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal PricingTier: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "PricingTier", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve PricingTier: %w", idx, err)
		}
		return &Child{PricingTier: &v}, nil
	case "pricing_feature":
		var v PricingFeature
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal PricingFeature: %w", idx, err)
		}
		return &Child{PricingFeature: &v}, nil
	case "testimonial_section":
		var v TestimonialSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal TestimonialSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "TestimonialSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve TestimonialSection: %w", idx, err)
		}
		return &Child{TestimonialSection: &v}, nil
	case "testimonial":
		var v Testimonial
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Testimonial: %w", idx, err)
		}
		return &Child{Testimonial: &v}, nil
	case "blog_section":
		var v BlogSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal BlogSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "BlogSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve BlogSection: %w", idx, err)
		}
		return &Child{BlogSection: &v}, nil
	case "blog_post_card":
		var v BlogPostCard
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal BlogPostCard: %w", idx, err)
		}
		return &Child{BlogPostCard: &v}, nil
	case "faq_section":
		var v FAQSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FAQSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "FAQSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve FAQSection: %w", idx, err)
		}
		return &Child{FAQSection: &v}, nil
	case "faq_item":
		var v FAQItem
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FAQItem: %w", idx, err)
		}
		return &Child{FAQItem: &v}, nil
	case "logo_cloud_section":
		var v LogoCloudSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal LogoCloudSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "LogoCloudSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve LogoCloudSection: %w", idx, err)
		}
		return &Child{LogoCloudSection: &v}, nil
	case "logo_item":
		var v LogoItem
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal LogoItem: %w", idx, err)
		}
		return &Child{LogoItem: &v}, nil
	case "banner":
		var v Banner
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Banner: %w", idx, err)
		}
		return &Child{Banner: &v}, nil
	case "contact_section":
		var v ContactSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal ContactSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "ContactSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve ContactSection: %w", idx, err)
		}
		return &Child{ContactSection: &v}, nil
	case "contact_detail":
		var v ContactDetail
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal ContactDetail: %w", idx, err)
		}
		return &Child{ContactDetail: &v}, nil
	case "newsletter_section":
		var v NewsletterSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal NewsletterSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "NewsletterSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve NewsletterSection: %w", idx, err)
		}
		return &Child{NewsletterSection: &v}, nil
	case "newsletter_detail":
		var v NewsletterDetail
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal NewsletterDetail: %w", idx, err)
		}
		return &Child{NewsletterDetail: &v}, nil
	case "error_page":
		var v ErrorPage
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal ErrorPage: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "ErrorPage", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve ErrorPage: %w", idx, err)
		}
		return &Child{ErrorPage: &v}, nil
	case "popular_page":
		var v PopularPage
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal PopularPage: %w", idx, err)
		}
		return &Child{PopularPage: &v}, nil
	case "team_section":
		var v TeamSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal TeamSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "TeamSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve TeamSection: %w", idx, err)
		}
		return &Child{TeamSection: &v}, nil
	case "team_member":
		var v TeamMember
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal TeamMember: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "TeamMember", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve TeamMember: %w", idx, err)
		}
		return &Child{TeamMember: &v}, nil
	case "team_social_link":
		var v TeamSocialLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal TeamSocialLink: %w", idx, err)
		}
		return &Child{TeamSocialLink: &v}, nil
	case "header_section":
		var v HeaderSection
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal HeaderSection: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "HeaderSection", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve HeaderSection: %w", idx, err)
		}
		return &Child{HeaderSection: &v}, nil
	case "header_action_link":
		var v HeaderActionLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal HeaderActionLink: %w", idx, err)
		}
		return &Child{HeaderActionLink: &v}, nil
	case "inline_testimonial":
		var v InlineTestimonial
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal InlineTestimonial: %w", idx, err)
		}
		return &Child{InlineTestimonial: &v}, nil
	case "cta_benefit":
		var v CTABenefit
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal CTABenefit: %w", idx, err)
		}
		return &Child{CTABenefit: &v}, nil
	case "doc_image":
		var v DocImage
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal DocImage: %w", idx, err)
		}
		return &Child{DocImage: &v}, nil

	// Navigation types
	case "footer":
		var v Footer
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Footer: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "Footer", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Footer: %w", idx, err)
		}
		return &Child{Footer: &v}, nil
	case "footer_column":
		var v FooterColumn
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FooterColumn: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "FooterColumn", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve FooterColumn: %w", idx, err)
		}
		return &Child{FooterColumn: &v}, nil
	case "footer_link":
		var v FooterLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FooterLink: %w", idx, err)
		}
		return &Child{FooterLink: &v}, nil
	case "footer_social":
		var v FooterSocial
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FooterSocial: %w", idx, err)
		}
		return &Child{FooterSocial: &v}, nil
	case "footer_text":
		var v FooterText
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FooterText: %w", idx, err)
		}
		return &Child{FooterText: &v}, nil
	case "social_link":
		var v SocialLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal SocialLink: %w", idx, err)
		}
		return &Child{SocialLink: &v}, nil
	case "header":
		var v Header
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal Header: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "Header", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve Header: %w", idx, err)
		}
		return &Child{Header: &v}, nil
	case "nav_link":
		var v NavLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal NavLink: %w", idx, err)
		}
		return &Child{NavLink: &v}, nil
	case "flyout_menu":
		var v FlyoutMenu
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FlyoutMenu: %w", idx, err)
		}
		if err := resolveGenericChildren(v.RawChildren, &v.Resolved, "FlyoutMenu", v.ID); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: resolve FlyoutMenu: %w", idx, err)
		}
		return &Child{FlyoutMenu: &v}, nil
	case "flyout_link":
		var v FlyoutLink
		if err := json.Unmarshal(msg, &v); err != nil {
			return nil, fmt.Errorf("parseChildren[%d]: unmarshal FlyoutLink: %w", idx, err)
		}
		return &Child{FlyoutLink: &v}, nil

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

// resolveGenericChildren is a reusable helper for any type with
// RawChildren []json.RawMessage + Resolved []Child.
func resolveGenericChildren(raw []json.RawMessage, resolved *[]Child, typeName, id string) error {
	if len(raw) == 0 {
		return nil
	}
	children, err := ParseChildren(raw)
	if err != nil {
		return fmt.Errorf("resolve %s %s children: %w", typeName, id, err)
	}
	*resolved = children
	return nil
}

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

func resolveContentReference(ref *ContentReference) error {
	for _, raw := range ref.RawMenus {
		var menu Menu
		if err := json.Unmarshal(raw, &menu); err != nil {
			return fmt.Errorf("resolve content reference %s menu: %w", ref.ID, err)
		}
		if err := ResolveMenu(&menu); err != nil {
			return err
		}
		ref.Menus = append(ref.Menus, menu)
	}
	return nil
}

func resolveBentoGrid(bg *BentoGridSection) error {
	for _, raw := range bg.RawCells {
		var cell BentoCell
		if err := json.Unmarshal(raw, &cell); err != nil {
			return fmt.Errorf("resolve bento grid %s cell: %w", bg.ID, err)
		}
		bg.Cells = append(bg.Cells, cell)
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

// MediaRef pairs a media ID with a pointer to the URL field that should
// be populated after resolution.
type MediaRef struct {
	MediaID string
	URL     *string
}

// CollectMediaRefs walks the content tree and returns all media fields
// that need URL resolution.
func CollectMediaRefs(children []Child) []MediaRef {
	var refs []MediaRef
	for i := range children {
		c := &children[i]
		if c.HeroSection != nil {
			if c.HeroSection.Image != "" {
				refs = append(refs, MediaRef{c.HeroSection.Image, &c.HeroSection.ImageURL})
			}
			if c.HeroSection.ImageDark != "" {
				refs = append(refs, MediaRef{c.HeroSection.ImageDark, &c.HeroSection.ImageDarkURL})
			}
		}
		if c.CTASection != nil && c.CTASection.Image != "" {
			refs = append(refs, MediaRef{c.CTASection.Image, &c.CTASection.ImageURL})
		}
		if c.ContentSection != nil && c.ContentSection.Image != "" {
			refs = append(refs, MediaRef{c.ContentSection.Image, &c.ContentSection.ImageURL})
		}
		if c.BentoGridSection != nil {
			for j := range c.BentoGridSection.Cells {
				cell := &c.BentoGridSection.Cells[j]
				if cell.Image != "" {
					refs = append(refs, MediaRef{cell.Image, &cell.ImageURL})
				}
				if cell.ImageDark != "" {
					refs = append(refs, MediaRef{cell.ImageDark, &cell.ImageDarkURL})
				}
			}
		}
		if c.Image != nil && c.Image.ImageID != "" {
			refs = append(refs, MediaRef{c.Image.ImageID, &c.Image.MediaURL})
		}
		// New image-bearing types
		if c.FeatureSection != nil {
			if c.FeatureSection.Image != "" {
				refs = append(refs, MediaRef{c.FeatureSection.Image, &c.FeatureSection.ImageURL})
			}
			if c.FeatureSection.ImageDark != "" {
				refs = append(refs, MediaRef{c.FeatureSection.ImageDark, &c.FeatureSection.ImageDkURL})
			}
		}
		if c.Testimonial != nil && c.Testimonial.Image != "" {
			refs = append(refs, MediaRef{c.Testimonial.Image, &c.Testimonial.ImageURL})
		}
		if c.BlogPostCard != nil && c.BlogPostCard.Image != "" {
			refs = append(refs, MediaRef{c.BlogPostCard.Image, &c.BlogPostCard.ImageURL})
		}
		if c.LogoItem != nil && c.LogoItem.Image != "" {
			refs = append(refs, MediaRef{c.LogoItem.Image, &c.LogoItem.ImageURL})
		}
		if c.TeamMember != nil && c.TeamMember.Image != "" {
			refs = append(refs, MediaRef{c.TeamMember.Image, &c.TeamMember.ImageURL})
		}
		if c.DocImage != nil && c.DocImage.Image != "" {
			refs = append(refs, MediaRef{c.DocImage.Image, &c.DocImage.ImageURL})
		}
		if c.Header != nil {
			if c.Header.Logo != "" {
				refs = append(refs, MediaRef{c.Header.Logo, &c.Header.LogoURL})
			}
			if c.Header.LogoDark != "" {
				refs = append(refs, MediaRef{c.Header.LogoDark, &c.Header.LogoDkURL})
			}
		}
		// Recurse into layout types
		if c.Row != nil {
			for j := range c.Row.Columns {
				refs = append(refs, CollectMediaRefs(c.Row.Columns[j].Resolved)...)
			}
			refs = append(refs, CollectMediaRefs(c.Row.Resolved)...)
		}
		if c.Columns != nil {
			refs = append(refs, CollectMediaRefs(c.Columns.Resolved)...)
		}
		if c.Grid != nil {
			for j := range c.Grid.Areas {
				refs = append(refs, CollectMediaRefs(c.Grid.Areas[j].Resolved)...)
			}
			refs = append(refs, CollectMediaRefs(c.Grid.Resolved)...)
		}
		// Recurse into new container types
		if c.StatsSection != nil {
			refs = append(refs, CollectMediaRefs(c.StatsSection.Resolved)...)
		}
		if c.FeatureSection != nil {
			refs = append(refs, CollectMediaRefs(c.FeatureSection.Resolved)...)
		}
		if c.PricingSection != nil {
			refs = append(refs, CollectMediaRefs(c.PricingSection.Resolved)...)
		}
		if c.PricingTier != nil {
			refs = append(refs, CollectMediaRefs(c.PricingTier.Resolved)...)
		}
		if c.TestimonialSection != nil {
			refs = append(refs, CollectMediaRefs(c.TestimonialSection.Resolved)...)
		}
		if c.BlogSection != nil {
			refs = append(refs, CollectMediaRefs(c.BlogSection.Resolved)...)
		}
		if c.FAQSection != nil {
			refs = append(refs, CollectMediaRefs(c.FAQSection.Resolved)...)
		}
		if c.LogoCloudSection != nil {
			refs = append(refs, CollectMediaRefs(c.LogoCloudSection.Resolved)...)
		}
		if c.ContactSection != nil {
			refs = append(refs, CollectMediaRefs(c.ContactSection.Resolved)...)
		}
		if c.NewsletterSection != nil {
			refs = append(refs, CollectMediaRefs(c.NewsletterSection.Resolved)...)
		}
		if c.ErrorPage != nil {
			refs = append(refs, CollectMediaRefs(c.ErrorPage.Resolved)...)
		}
		if c.TeamSection != nil {
			refs = append(refs, CollectMediaRefs(c.TeamSection.Resolved)...)
		}
		if c.TeamMember != nil {
			refs = append(refs, CollectMediaRefs(c.TeamMember.Resolved)...)
		}
		if c.HeaderSection != nil {
			refs = append(refs, CollectMediaRefs(c.HeaderSection.Resolved)...)
		}
		if c.Footer != nil {
			refs = append(refs, CollectMediaRefs(c.Footer.Resolved)...)
		}
		if c.FooterColumn != nil {
			refs = append(refs, CollectMediaRefs(c.FooterColumn.Resolved)...)
		}
		if c.Header != nil {
			refs = append(refs, CollectMediaRefs(c.Header.Resolved)...)
		}
		if c.FlyoutMenu != nil {
			refs = append(refs, CollectMediaRefs(c.FlyoutMenu.Resolved)...)
		}
	}
	return refs
}

// CollectImages returns pointers to all Image blocks in the tree
// so callers can resolve their MediaURL fields.
// Deprecated: Use CollectMediaRefs instead.
func CollectImages(children []Child) []*Image {
	var images []*Image
	for i := range children {
		c := &children[i]
		if c.Image != nil {
			images = append(images, c.Image)
		}
		if c.Row != nil {
			for j := range c.Row.Columns {
				images = append(images, CollectImages(c.Row.Columns[j].Resolved)...)
			}
			images = append(images, CollectImages(c.Row.Resolved)...)
		}
		if c.Columns != nil {
			images = append(images, CollectImages(c.Columns.Resolved)...)
		}
		if c.Grid != nil {
			for j := range c.Grid.Areas {
				images = append(images, CollectImages(c.Grid.Areas[j].Resolved)...)
			}
			images = append(images, CollectImages(c.Grid.Resolved)...)
		}
	}
	return images
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

// Slugify converts a heading string into a URL-safe anchor ID.
func Slugify(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	prev := false
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			if prev {
				b.WriteByte('-')
				prev = false
			}
			b.WriteRune(r)
		} else {
			if b.Len() > 0 {
				prev = true
			}
		}
	}
	return b.String()
}
