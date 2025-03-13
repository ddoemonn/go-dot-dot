package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	"github.com/ddoemonn/go-dot-dot/internal/model"
)

// CreateTableItems converts a slice of table names to list items
func CreateTableItems(tables []string) []list.Item {
	items := make([]list.Item, len(tables))
	for i, table := range tables {
		items[i] = model.TableItem{Name: table}
	}
	return items
}

// CreateTableList creates a styled list for table selection
func CreateTableList(tables []string, styles *Styles) list.Model {
	listDelegate := list.NewDefaultDelegate()
	listDelegate.SetSpacing(0) // Reduce the spacing between items to 0
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.
		//Foreground(lipgloss.Color(ColorBackground)).
		Background(lipgloss.Color(ColorSecondary))
	listDelegate.Styles.NormalTitle = listDelegate.Styles.NormalTitle.
		Foreground(lipgloss.Color(ColorText))

	tableList := list.New(CreateTableItems(tables), listDelegate, 0, 0)
	tableList.SetShowStatusBar(false)
	tableList.SetFilteringEnabled(false)
	tableList.SetShowHelp(false)
	tableList.Styles.Title = styles.Title.Copy()

	return tableList
}

// CreateSearchInput creates and configures a text input for search
func CreateSearchInput() textinput.Model {
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Width = 30
	ti.Prompt = "› "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorSecondary))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorText))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorAccent))
	return ti
}

// CreateTableData creates a styled table based on column names and data
func CreateTableData(columns []string, data [][]string, horizontalScrollOffset int) table.Model {
	rows := make([]table.Row, len(data))
	for i, d := range data {
		// Make sure we don't go out of bounds if the data has more columns than headers
		maxCols := len(columns)
		if len(d) < maxCols {
			maxCols = len(d)
		}

		row := make(table.Row, maxCols)
		for j := 0; j < maxCols; j++ {
			if j < len(d) {
				// Truncate long values to prevent UI issues
				if len(d[j]) > 100 {
					row[j] = d[j][:97] + "..."
				} else {
					row[j] = d[j]
				}
			} else {
				row[j] = ""
			}
		}
		rows[i] = row
	}

	// Create table columns with horizontal scrolling
	t := table.New(
		table.WithColumns(makeColumns(columns, horizontalScrollOffset)),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	// Style the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(ColorMuted)).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color(ColorAccent)).
		Align(lipgloss.Center)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color(ColorBackground)).
		Background(lipgloss.Color(ColorSecondary)).
		Bold(true)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color(ColorText))

	t.SetStyles(s)

	return t
}

// Create table columns with appropriate widths and horizontal scrolling
func makeColumns(headers []string, horizontalScrollOffset int) []table.Column {
	columns := make([]table.Column, len(headers))

	// Apply horizontal scrolling offset
	visibleHeaders := headers
	if horizontalScrollOffset > 0 && horizontalScrollOffset < len(headers) {
		visibleHeaders = headers[horizontalScrollOffset:]
	}

	for i, header := range visibleHeaders {
		// Adjust width based on header length
		width := len(header) + 8
		if width < 16 {
			width = 16
		} else if width > 40 {
			width = 40
		}

		// Calculate the actual index in the original headers array
		actualIndex := i + horizontalScrollOffset
		if actualIndex < len(headers) {
			columns[actualIndex] = table.Column{
				Title: header,
				Width: width,
			}
		}
	}

	return columns
}
