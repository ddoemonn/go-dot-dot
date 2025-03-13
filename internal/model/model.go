package model

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Model represents the application state
type Model struct {
	Pool                   *pgxpool.Pool
	TableList              list.Model
	TableData              table.Model
	SelectedTable          string
	Tables                 []string
	ColumnNames            []string
	Data                   [][]string
	FilteredData           [][]string
	Width                  int
	Height                 int
	Focused                int // 0: table list, 1: table data, 2: detail view
	SearchMode             bool
	SearchInput            textinput.Model
	SearchQuery            string
	ShowHelp               bool
	Help                   help.Model
	Err                    error
	SelectedRow            int
	SelectedRowData        map[string]string // Column name -> value
	ConnectionDetails      string
	HorizontalScrollOffset int // Track horizontal scroll position
}

// TableItem represents a database table in the list
type TableItem struct {
	Name string
}

// FilterValue returns the value to filter on
func (i TableItem) FilterValue() string {
	return i.Name
}

// Title returns the title of the item
func (i TableItem) Title() string {
	return i.Name
}

// Description returns the description of the item
func (i TableItem) Description() string {
	return ""
}
