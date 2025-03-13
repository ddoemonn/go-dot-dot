package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/ddoemonn/go-dot-dot/internal/model"
)

// RenderView renders the main UI view
func RenderView(m *model.Model, styles *Styles, keys *KeyMap) string {
    if m.Err != nil {
        return styles.Base.Render(fmt.Sprintf("Error: %v\n\nPress q to quit.", m.Err))
    }
    
    // App title bar
    appTitle := styles.AppTitle.Width(m.Width).Render("PostgreSQL Database Explorer")
    
    // Connection info badge
    connectionInfo := styles.InfoBox.Render(fmt.Sprintf("🔌 %s", m.ConnectionDetails))
    
    // Context-sensitive help based on current view
    contextHelp := ""
    switch m.Focused {
    case 0:
        contextHelp = styles.StatusMessage.Render("Select a table with Enter or → | ? for help")
    case 1:
        if len(m.FilteredData) > 0 {
            contextHelp = styles.StatusMessage.Render("Press v or Enter to view row details | / to search | ? for help")
        } else {
            contextHelp = styles.StatusMessage.Render("No data to display | Esc to go back | ? for help")
        }
    case 2:
        contextHelp = styles.StatusMessage.Render("Viewing row details | Esc to go back | ? for help")
    }
    
    // Help view
    helpView := ""
    if m.ShowHelp {
        helpView = styles.Help.Render(m.Help.View(*keys))
    }
    
    // Main content based on focused view
    var content string
    
    if m.Focused == 2 {
        // Detail view
        detailContent := RenderDetailView(m.SelectedRowData, m.Width-10, m.SelectedRow, styles)
        content = styles.DetailCard.Width(m.Width - 10).Render(detailContent)
    } else {
        // Table list view with title
        tableListHeader := styles.TableListHeader.Render("DATABASE TABLES")
        tableListView := m.TableList.View()
        
        if m.Focused == 0 {
            tableListView = styles.Focused.Render(tableListView)
        } else {
            tableListView = styles.Unfocused.Render(tableListView)
        }
        
        // Table data view with title
        var tableDataView string
        if m.SelectedTable != "" {
            tableCount := fmt.Sprintf(" (%d rows)", len(m.Data))
            tableDataHeader := styles.TableDataHeader.Render(fmt.Sprintf(" TABLE: %s%s ", 
                strings.ToUpper(m.SelectedTable), 
                styles.StatusMessage.Render(tableCount)))
            
            // Search UI
            searchUI := ""
            if m.SearchMode {
                searchUI = styles.SearchPrompt.Render("🔍 ") + m.SearchInput.View()
            } else if m.SearchQuery != "" {
                resultsCount := fmt.Sprintf(" (%d/%d rows)", len(m.FilteredData), len(m.Data))
                searchUI = styles.SearchPrompt.Render("🔍 ") + 
                    styles.FilterIndicator.Render(m.SearchQuery) + 
                    styles.StatusMessage.Render(resultsCount)
            }
            
            // Status message
            statusMsg := ""
            if len(m.FilteredData) == 0 && m.SearchQuery != "" {
                statusMsg = styles.StatusMessage.Render("No matching results. Press Ctrl+X to clear filter.")
            } else if len(m.Data) == 0 {
                statusMsg = styles.StatusMessage.Render("Empty table")
            }
            
            dataView := m.TableData.View()
            if m.Focused == 1 {
                dataView = styles.Focused.Render(dataView)
            } else {
                dataView = styles.Unfocused.Render(dataView)
            }
            
            tableDataView = lipgloss.JoinVertical(lipgloss.Left,
                tableDataHeader,
                searchUI,
                statusMsg,
                dataView,
            )
        } else {
            tableDataView = styles.Unfocused.Render(
                lipgloss.Place(
                    m.Width-styles.TableListHeader.GetWidth()-10, 
                    m.Height-10, 
                    lipgloss.Center, 
                    lipgloss.Center, 
                    styles.StatusMessage.Render("← Select a table to view data"), 
                    lipgloss.WithWhitespaceChars(""),
                ),
            )
        }
        
        // Arrange horizontally
        content = lipgloss.JoinHorizontal(lipgloss.Top,
            lipgloss.JoinVertical(lipgloss.Left, tableListHeader, tableListView),
            tableDataView,
        )
    }
    
    // Final layout
    return styles.App.Render(lipgloss.JoinVertical(lipgloss.Left, 
        appTitle,
        connectionInfo,
        styles.Divider.Render(strings.Repeat("─", m.Width)),
        contextHelp,
        content, 
        helpView,
    ))
}

// RenderDetailView renders a detailed view of a row
func RenderDetailView(data map[string]string, width int, rowIndex int, styles *Styles) string {
    if len(data) == 0 {
        return "No data available"
    }
    
    // Title with row number
    title := styles.Title.Copy().Width(width - 6).Align(lipgloss.Center).
        Foreground(lipgloss.Color(ColorBackground)).
        Background(lipgloss.Color(ColorPrimary)).
        Padding(0, 2).
        Render(fmt.Sprintf("ROW DETAILS (Row #%d)", rowIndex+1))
    
    // Find the longest key for alignment
    maxKeyLen := 0
    for k := range data {
        if len(k) > maxKeyLen {
            maxKeyLen = len(k)
        }
    }
    
    // Sort keys alphabetically
    var keys []string
    for k := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    // Build rows
    var rows []string
    rows = append(rows, title)
    rows = append(rows, "")
    
    // Add all fields
    for _, k := range keys {
        v := data[k]
        
        // Format the value nicely
        formattedValue := v
        if v == "NULL" {
            formattedValue = styles.DetailNull.Render("NULL")
        }
        
        label := styles.DetailLabel.Copy().Width(maxKeyLen + 2).Render(k + ":")
        value := styles.DetailValue.Render(formattedValue)
        
        row := fmt.Sprintf("%s %s", label, value)
        rows = append(rows, row)
    }
    
    return lipgloss.JoinVertical(lipgloss.Left, rows...)
}