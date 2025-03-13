package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Color constants
const (
	// Base colors
	ColorBackground = "#1A1B26" // Dark blue-black
	ColorText       = "#A9B1D6" // Soft lavender
	ColorMuted      = "#565F89" // Muted indigo

	// Accent colors
	ColorPrimary   = "#BB9AF7" // Bright purple
	ColorSecondary = "#7AA2F7" // Vibrant blue
	ColorAccent    = "#F7768E" // Coral pink
	ColorSuccess   = "#9ECE6A" // Lime green
	ColorWarning   = "#E0AF68" // Amber
	ColorInfo      = "#2AC3DE" // Cyan
)

// Styles holds all application styles
type Styles struct {
	// Base styles
	Base          lipgloss.Style
	Focused       lipgloss.Style
	Unfocused     lipgloss.Style
	App           lipgloss.Style
	Title         lipgloss.Style
	StatusMessage lipgloss.Style
	SearchPrompt  lipgloss.Style
	ColumnHeader  lipgloss.Style
	Help          lipgloss.Style

	// Detail view styles
	DetailLabel lipgloss.Style
	DetailValue lipgloss.Style
	DetailNull  lipgloss.Style
	DetailCard  lipgloss.Style

	// Header styles
	AppTitle lipgloss.Style
	InfoBox  lipgloss.Style

	// Table styles
	TableListHeader lipgloss.Style
	FilterIndicator lipgloss.Style
	TableDataHeader lipgloss.Style
	Divider         lipgloss.Style
	ScrollIndicator lipgloss.Style
}

// NewStyles creates and initializes all application styles
func NewStyles() *Styles {
	s := &Styles{}

	// Base styles
	s.Base = lipgloss.NewStyle().
		//Background(lipgloss.Color(ColorBackground)).
		Foreground(lipgloss.Color(ColorText))

	s.Focused = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(ColorSecondary)).
		Padding(1).
		MarginRight(2)

	s.Unfocused = lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Padding(1).
		MarginRight(2)

	s.App = lipgloss.NewStyle().
		Padding(1)
		//Background(lipgloss.Color(ColorBackground))

	s.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(ColorPrimary)).
		MarginBottom(1).
		MarginLeft(1)

	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMuted)).
		Italic(true)

	s.SearchPrompt = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorSecondary))

	s.ColumnHeader = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(ColorAccent))

	s.Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorText)).
		//Background(lipgloss.Color(ColorBackground)).
		Padding(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(ColorMuted))

	// Detail view styles
	s.DetailLabel = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(ColorSecondary)).
		PaddingRight(1)

	s.DetailValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorText))

	s.DetailNull = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMuted)).
		Italic(true)

	s.DetailCard = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(ColorPrimary)).
		Padding(2)

	s.AppTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(ColorBackground)).
		Background(lipgloss.Color(ColorPrimary)).
		Padding(1, 2).
		Align(lipgloss.Center)

	s.InfoBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(ColorInfo)).
		Padding(1).
		MarginTop(1).
		MarginBottom(1)

	s.TableListHeader = lipgloss.NewStyle().
		Background(lipgloss.Color(ColorMuted)).
		Foreground(lipgloss.Color(ColorBackground)).
		Bold(true).
		Width(25).
		Align(lipgloss.Center)

	s.FilterIndicator = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWarning)).
		Bold(true)

	s.TableDataHeader = lipgloss.NewStyle().
		Background(lipgloss.Color(ColorSecondary)).
		Foreground(lipgloss.Color(ColorBackground)).
		Bold(true)

	s.Divider = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMuted))

	// Scroll indicator style
	s.ScrollIndicator = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorInfo)).
		Bold(true)

	return s
}
