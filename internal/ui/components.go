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
    listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.
        Foreground(lipgloss.Color(ColorBackground)).
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
func CreateTableData(columns []string, data [][]string) table.Model {
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
    
    // Create table columns
    t := table.New(
        table.WithColumns(makeColumns(columns)),
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

// Create table columns with appropriate widths
func makeColumns(headers []string) []table.Column {
    columns := make([]table.Column, len(headers))
    for i, header := range headers {
        // Adjust width based on header length
        width := len(header) + 4
        if width < 12 {
            width = 12
        } else if width > 30 {
            width = 30
        }
        
        columns[i] = table.Column{
            Title: header,
            Width: width,
        }
    }
    return columns
}