package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ddoemonn/go-dot-dot/internal/config"
	"github.com/ddoemonn/go-dot-dot/internal/db"
	"github.com/ddoemonn/go-dot-dot/internal/model"
	"github.com/ddoemonn/go-dot-dot/internal/ui"
	"github.com/ddoemonn/go-dot-dot/internal/utils"
)

// App represents the application
type App struct {
    model  model.Model
    db     *db.Database
    styles *ui.Styles
    keys   *ui.KeyMap
}

// New creates a new application instance
func New(cfg *config.Config) (*App, error) {
    // Connect to the database
    database, err := db.Connect(&cfg.DB)
    if err != nil {
        return nil, err
    }

    // Fetch all table names
    tables, err := database.FetchTables()
    if err != nil {
        return nil, err
    }

    // Initialize styles and keymap
    styles := ui.NewStyles()
    keys := ui.NewKeyMap()

    // Create table list
    tableList := ui.CreateTableList(tables, styles)

    // Set up search input
    searchInput := ui.CreateSearchInput()
    
    // Initialize the model
    m := model.Model{
        Pool:              database.GetPool(),
        TableList:         tableList,
        Tables:            tables,
        Focused:           0,
        SearchInput:       searchInput,
        Help:              help.New(),
        ShowHelp:          false,
        ConnectionDetails: cfg.DB.ConnectionDetails(),
        HorizontalScrollOffset: 0,
    }

    return &App{
        model:  m,
        db:     database,
        styles: styles,
        keys:   keys,
    }, nil
}

// Run starts the application
func (a *App) Run() error {
    p := tea.NewProgram(a, tea.WithAltScreen())
    _, err := p.Run()
    return err
}

// Init initializes the application
func (a App) Init() tea.Cmd {
    return nil
}

// Update handles messages and user input
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Handle search mode separately
        if a.model.SearchMode {
            switch msg.String() {
            case "esc":
                a.model.SearchMode = false
                a.model.SearchQuery = a.model.SearchInput.Value()
                // Apply filter
                a.applySearchFilter()
                return a, nil
            case "enter":
                a.model.SearchMode = false
                a.model.SearchQuery = a.model.SearchInput.Value()
                // Apply filter
                a.applySearchFilter()
                return a, nil
            default:
                var inputCmd tea.Cmd
                a.model.SearchInput, inputCmd = a.model.SearchInput.Update(msg)
                return a, inputCmd
            }
        }

        // Global keys
        switch {
        case key.Matches(msg, a.keys.Quit):
            return a, tea.Quit
        case key.Matches(msg, a.keys.Help):
            a.model.ShowHelp = !a.model.ShowHelp
            return a, nil
        case key.Matches(msg, a.keys.Search):
            if a.model.Focused == 1 && len(a.model.Data) > 0 { // Only allow search in table view with data
                a.model.SearchMode = true
                a.model.SearchInput.Focus()
                a.model.SearchInput.Placeholder = "Type to search table..."
                return a, nil
            }
        case key.Matches(msg, a.keys.ClearSearch):
            if a.model.Focused == 1 && a.model.SearchQuery != "" {
                a.model.SearchQuery = ""
                a.model.SearchInput.Reset()
                a.model.FilteredData = a.model.Data
                if len(a.model.ColumnNames) > 0 && len(a.model.Data) > 0 {
                    a.model.TableData = ui.CreateTableData(a.model.ColumnNames, a.model.FilteredData, a.model.HorizontalScrollOffset)
                }
            }
            return a, nil
        case key.Matches(msg, a.keys.Back):
            // Back button behavior depends on current view
            if a.model.Focused == 2 { // Detail view -> Table view
                a.model.Focused = 1
                return a, nil
            } else if a.model.Focused == 1 { // Table view -> Table list
                a.model.Focused = 0
                a.model.SelectedTable = ""
                return a, nil
            }
        // Handle horizontal scrolling
        case key.Matches(msg, a.keys.ScrollLeft):
            if a.model.Focused == 1 && a.model.HorizontalScrollOffset > 0 {
                a.model.HorizontalScrollOffset--
                if len(a.model.ColumnNames) > 0 && len(a.model.FilteredData) > 0 {
                    a.model.TableData = ui.CreateTableData(a.model.ColumnNames, a.model.FilteredData, a.model.HorizontalScrollOffset)
                }
                return a, nil
            }
        case key.Matches(msg, a.keys.ScrollRight):
            if a.model.Focused == 1 && a.model.HorizontalScrollOffset < len(a.model.ColumnNames)-1 {
                a.model.HorizontalScrollOffset++
                if len(a.model.ColumnNames) > 0 && len(a.model.FilteredData) > 0 {
                    a.model.TableData = ui.CreateTableData(a.model.ColumnNames, a.model.FilteredData, a.model.HorizontalScrollOffset)
                }
                return a, nil
            }
        }

        // Handle based on focus
        if a.model.Focused == 0 { // Table list
            switch {
            case key.Matches(msg, a.keys.Right), key.Matches(msg, a.keys.Select):
                if len(a.model.Tables) > 0 {
                    i, ok := a.model.TableList.SelectedItem().(list.Item)
                    if ok {
                        a.model.SelectedTable = i.FilterValue()
                        var err error
                        a.model.Data, a.model.ColumnNames, err = a.db.FetchTableData(a.model.SelectedTable)
                        if err != nil {
                            a.model.Err = err
                            return a, nil
                        }
                        a.model.FilteredData = a.model.Data
                        a.model.SearchQuery = ""
                        a.model.SearchInput.Reset()
                        // Reset horizontal scroll when selecting a new table
                        a.model.HorizontalScrollOffset = 0
                        
                        if len(a.model.ColumnNames) > 0 && len(a.model.Data) > 0 {
                            a.model.TableData = ui.CreateTableData(a.model.ColumnNames, a.model.Data, a.model.HorizontalScrollOffset)
                        }
                        a.model.Focused = 1
                    }
                }
            }
            a.model.TableList, cmd = a.model.TableList.Update(msg)
            cmds = append(cmds, cmd)
        } else if a.model.Focused == 1 { // Table data
            switch {
            case key.Matches(msg, a.keys.Left):
                // Go back to table list
                a.model.Focused = 0
                return a, nil
            case key.Matches(msg, a.keys.ViewDetails), key.Matches(msg, a.keys.Select):
                // View details of selected row
                if len(a.model.FilteredData) > 0 {
                    rowIndex := a.model.TableData.Cursor()
                    if rowIndex >= 0 && rowIndex < len(a.model.FilteredData) {
                        a.model.SelectedRow = rowIndex
                        a.model.SelectedRowData = make(map[string]string)
                        for i, col := range a.model.ColumnNames {
                            if i < len(a.model.FilteredData[rowIndex]) {
                                a.model.SelectedRowData[col] = a.model.FilteredData[rowIndex][i]
                            }
                        }
                        a.model.Focused = 2 // Switch to detail view
                    }
                }
            default:
                a.model.TableData, cmd = a.model.TableData.Update(msg)
                cmds = append(cmds, cmd)
            }
        } else if a.model.Focused == 2 { // Detail view
            // No special handling needed for detail view beyond global keys
        }

    case tea.WindowSizeMsg:
        a.model.Width = msg.Width
        a.model.Height = msg.Height
        
        // Adjust table list
        listWidth := utils.Min(30, a.model.Width/4)
        a.model.TableList.SetWidth(listWidth)
        a.model.TableList.SetHeight(a.model.Height - 8) // Leave space for headers and footers
        
        // Adjust table data
        headerHeight := 6
        footerHeight := 3
        availableHeight := a.model.Height - headerHeight - footerHeight
        if len(a.model.ColumnNames) > 0 && len(a.model.FilteredData) > 0 {
            a.model.TableData.SetHeight(availableHeight)
            a.model.TableData.SetWidth(a.model.Width - listWidth - 8)
        }
        
        // Update styles based on width
        a.styles.TableListHeader = a.styles.TableListHeader.Width(listWidth)
    }

    return a, tea.Batch(cmds...)
}

// View renders the UI
func (a App) View() string {
    return ui.RenderView(&a.model, a.styles, a.keys)
}

// Apply search filter to table data
func (a *App) applySearchFilter() {
    if a.model.SearchQuery == "" {
        a.model.FilteredData = a.model.Data
    } else {
        searchLower := strings.ToLower(a.model.SearchQuery)
        a.model.FilteredData = nil
        
        // Search through all rows and all columns
        for _, row := range a.model.Data {
            found := false
            for _, cell := range row {
                if strings.Contains(strings.ToLower(cell), searchLower) {
                    found = true
                    break
                }
            }
            if found {
                a.model.FilteredData = append(a.model.FilteredData, row)
            }
        }
    }
    
    // Recreate the table with filtered data
    if len(a.model.ColumnNames) > 0 && len(a.model.FilteredData) > 0 {
        a.model.TableData = ui.CreateTableData(a.model.ColumnNames, a.model.FilteredData, a.model.HorizontalScrollOffset)
    }
}