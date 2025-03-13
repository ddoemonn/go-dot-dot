package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

// UI styles for the setup wizard
var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#BB9AF7")).
		MarginBottom(1).
		Padding(1, 2)

	inputLabelStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7AA2F7")).
		Width(15).
		PaddingRight(1)

	inputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A9B1D6"))

	infoStyle = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#565F89"))

	buttonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1A1B26")).
		Background(lipgloss.Color("#9ECE6A")).
		Padding(0, 3).
		Bold(true)

	focusedButtonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1A1B26")).
		Background(lipgloss.Color("#F7768E")).
		Padding(0, 3).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F7768E")).
		Italic(true)
)

// SetupModel represents the UI state for the setup wizard
type SetupModel struct {
	inputs       []textinput.Model
	focusIndex   int
	err          error
	width        int
	height       int
	buttonFocus  bool
	confirmSave  bool
	config       *Config
	loading      bool
	loadingMsg   string
}

// NewSetupModel creates a new setup wizard model
func NewSetupModel() SetupModel {
	// Create text inputs for each field
	inputs := make([]textinput.Model, 5)
	
	// DB User
	inputs[0] = textinput.New()
	inputs[0].Placeholder = "postgres"
	inputs[0].Focus()
	inputs[0].Width = 30
	inputs[0].Prompt = "› "
	inputs[0].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7"))
	
	// DB Password
	inputs[1] = textinput.New()
	inputs[1].Placeholder = "password"
	inputs[1].Width = 30
	inputs[1].Prompt = "› "
	inputs[1].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7"))
	inputs[1].EchoMode = textinput.EchoPassword
	inputs[1].EchoCharacter = '•'
	
	// DB Name
	inputs[2] = textinput.New()
	inputs[2].Placeholder = "postgres"
	inputs[2].Width = 30
	inputs[2].Prompt = "› "
	inputs[2].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7"))
	
	// DB Host
	inputs[3] = textinput.New()
	inputs[3].Placeholder = "localhost"
	inputs[3].Width = 30
	inputs[3].Prompt = "› "
	inputs[3].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7"))
	
	// DB Port
	inputs[4] = textinput.New()
	inputs[4].Placeholder = "5432"
	inputs[4].Width = 30
	inputs[4].Prompt = "› "
	inputs[4].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7"))
	
	return SetupModel{
		inputs:      inputs,
		focusIndex:  0,
		buttonFocus: false,
		confirmSave: false,
		config:      &Config{},
		loading:     false,
	}
}

// Init initializes the model
func (m SetupModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles user input
func (m SetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// If loading, ignore key presses
		if m.loading {
			return m, nil
		}
		
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
			
		case "tab", "shift+tab", "up", "down":
			if !m.buttonFocus {
				// Cycle through inputs
				if msg.String() == "up" || msg.String() == "shift+tab" {
					m.focusIndex--
					if m.focusIndex < 0 {
						m.focusIndex = len(m.inputs) - 1
					}
				} else {
					m.focusIndex++
					if m.focusIndex >= len(m.inputs) {
						m.focusIndex = 0
					}
				}
				
				// Update focus states
				for i := 0; i < len(m.inputs); i++ {
					if i == m.focusIndex {
						cmds = append(cmds, m.inputs[i].Focus())
					} else {
						m.inputs[i].Blur()
					}
				}
			}
			
		case "enter":
			if m.focusIndex == len(m.inputs)-1 && !m.buttonFocus {
				// Move focus to the button
				m.buttonFocus = true
				m.inputs[m.focusIndex].Blur()
			} else if m.buttonFocus {
				// Save configuration
				if m.validateInputs() {
					// Start loading state
					m.loading = true
					m.loadingMsg = "Saving configuration..."
					return m, m.saveConfigCmd
				}
			}
			
		case "backspace":
			if m.buttonFocus {
				// Move focus back to the last input
				m.buttonFocus = false
				m.focusIndex = len(m.inputs) - 1
				cmds = append(cmds, m.inputs[m.focusIndex].Focus())
			}
		}
		
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	
	// Handle input updates
	if !m.buttonFocus {
		var cmd tea.Cmd
		m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
		cmds = append(cmds, cmd)
	}
	
	return m, tea.Batch(cmds...)
}

// View renders the UI
func (m SetupModel) View() string {
	if m.confirmSave {
		return titleStyle.Render("Configuration saved successfully!") + "\n\n" +
			infoStyle.Render("Starting application...") + "\n\n" +
			"Please wait a moment."
	}
	
	if m.loading {
		return titleStyle.Render("PostgreSQL Database Configuration") + "\n\n" +
			infoStyle.Render(m.loadingMsg) + "\n\n" +
			"Please wait..."
	}
	
	title := titleStyle.Render("PostgreSQL Database Configuration")
	
	// Render inputs with labels
	inputs := []string{
		renderLabeledInput("DB User:", m.inputs[0].View()),
		renderLabeledInput("DB Password:", m.inputs[1].View()),
		renderLabeledInput("DB Name:", m.inputs[2].View()),
		renderLabeledInput("DB Host:", m.inputs[3].View()),
		renderLabeledInput("DB Port:", m.inputs[4].View()),
	}
	
	// Render save button
	var button string
	if m.buttonFocus {
		button = focusedButtonStyle.Render("Save Configuration")
	} else {
		button = buttonStyle.Render("Save Configuration")
	}
	
	// Render error if any
	errorMsg := ""
	if m.err != nil {
		errorMsg = "\n" + errorStyle.Render(m.err.Error())
	}
	
	// Render help text
	help := "\n" + infoStyle.Render("Tab/Shift+Tab: Navigate • Enter: Confirm • Esc: Quit")
	
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s%s%s",
		title,
		strings.Join(inputs, "\n"),
		button,
		errorMsg,
		help,
	)
}

// Helper function to render a labeled input
func renderLabeledInput(label string, input string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		inputLabelStyle.Render(label),
		inputStyle.Render(input),
	)
}

// Validate inputs before saving
func (m *SetupModel) validateInputs() bool {
	// Check if port is a number
	if m.inputs[4].Value() != "" {
		for _, c := range m.inputs[4].Value() {
			if c < '0' || c > '9' {
				m.err = fmt.Errorf("Port must be a number")
				return false
			}
		}
	}
	
	m.err = nil
	return true
}

// Command to save config in a goroutine
func (m *SetupModel) saveConfigCmd() tea.Msg {
	// Create config from inputs
	m.config.DB = DBConfig{
		User:     getValue(m.inputs[0].Value(), "postgres"),
		Password: getValue(m.inputs[1].Value(), ""),
		Name:     getValue(m.inputs[2].Value(), "postgres"),
		Host:     getValue(m.inputs[3].Value(), "localhost"),
		Port:     getValue(m.inputs[4].Value(), "5432"),
	}
	
	// Create .env file content
	content := fmt.Sprintf(
		"DB_USER=%s\nDB_PASSWORD=%s\nDB_NAME=%s\nDB_HOST=%s\nDB_PORT=%s\n",
		m.config.DB.User,
		m.config.DB.Password,
		m.config.DB.Name,
		m.config.DB.Host,
		m.config.DB.Port,
	)
	
	// Write to .env file
	err := os.WriteFile(".env", []byte(content), 0644)
	if err != nil {
		m.err = fmt.Errorf("Failed to save configuration: %v", err)
		m.loading = false
		return nil
	}
	
	// Set confirmSave to true and immediately return to quit the program
	m.confirmSave = true
	
	// Show success message briefly
	fmt.Println("\nConfiguration saved successfully!")
	fmt.Println("Starting application...")
	
	return tea.Quit()
}

// Helper function to get value with fallback
func getValue(input, fallback string) string {
	if input == "" {
		return fallback
	}
	return input
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	// Try to load .env file first
	_ = godotenv.Load()
	
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// RunSetup runs the configuration setup wizard
func RunSetup() (*Config, error) {
	model := NewSetupModel()
	p := tea.NewProgram(model, tea.WithAltScreen())
	
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}
	
	m, ok := finalModel.(SetupModel)
	if !ok {
		return nil, fmt.Errorf("could not convert model")
	}
	
	// Check if .env file exists regardless of confirmSave flag
	if _, err := os.Stat(".env"); err == nil {
		// .env file exists, so configuration was saved
		return &Config{
			DB: DBConfig{
				User:     getEnv("DB_USER", "postgres"),
				Password: getEnv("DB_PASSWORD", ""),
				Name:     getEnv("DB_NAME", "postgres"),
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "5432"),
			},
		}, nil
	}
	
	if m.confirmSave {
		// Return the config immediately without waiting
		return m.config, nil
	}
	
	return nil, fmt.Errorf("setup cancelled")
} 