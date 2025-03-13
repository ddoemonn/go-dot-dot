package ui

import (
    "github.com/charmbracelet/bubbles/key"
)

// KeyMap defines the keybindings
type KeyMap struct {
    Up         key.Binding
    Down       key.Binding
    Left       key.Binding
    Right      key.Binding
    Select     key.Binding
    Back       key.Binding
    Quit       key.Binding
    Search     key.Binding
    ClearSearch key.Binding
    Help       key.Binding
    ViewDetails key.Binding
    PageUp     key.Binding
    PageDown   key.Binding
    Home       key.Binding
    End        key.Binding
    ScrollLeft  key.Binding
    ScrollRight key.Binding
}

// NewKeyMap creates a new keymap with default bindings
func NewKeyMap() *KeyMap {
    return &KeyMap{
        Up: key.NewBinding(
            key.WithKeys("up", "k"),
            key.WithHelp("↑/k", "up"),
        ),
        Down: key.NewBinding(
            key.WithKeys("down", "j"),
            key.WithHelp("↓/j", "down"),
        ),
        Left: key.NewBinding(
            key.WithKeys("left", "h"),
            key.WithHelp("←/h", "left"),
        ),
        Right: key.NewBinding(
            key.WithKeys("right", "l"),
            key.WithHelp("→/l", "right"),
        ),
        Select: key.NewBinding(
            key.WithKeys("enter"),
            key.WithHelp("enter", "select"),
        ),
        Back: key.NewBinding(
            key.WithKeys("esc"),
            key.WithHelp("esc", "back"),
        ),
        Quit: key.NewBinding(
            key.WithKeys("ctrl+c", "q"),
            key.WithHelp("ctrl+c/q", "quit"),
        ),
        Search: key.NewBinding(
            key.WithKeys("/"),
            key.WithHelp("/", "search"),
        ),
        ClearSearch: key.NewBinding(
            key.WithKeys("ctrl+x"),
            key.WithHelp("ctrl+x", "clear search"),
        ),
        Help: key.NewBinding(
            key.WithKeys("?"),
            key.WithHelp("?", "toggle help"),
        ),
        ViewDetails: key.NewBinding(
            key.WithKeys("v", "space"),
            key.WithHelp("v/space", "view details"),
        ),
        PageUp: key.NewBinding(
            key.WithKeys("pgup"),
            key.WithHelp("pgup", "page up"),
        ),
        PageDown: key.NewBinding(
            key.WithKeys("pgdown"),
            key.WithHelp("pgdown", "page down"),
        ),
        Home: key.NewBinding(
            key.WithKeys("home"),
            key.WithHelp("home", "first item"),
        ),
        End: key.NewBinding(
            key.WithKeys("end"),
            key.WithHelp("end", "last item"),
        ),
        ScrollLeft: key.NewBinding(
            key.WithKeys("shift+left", "shift+h"),
            key.WithHelp("shift+←/shift+h", "scroll left"),
        ),
        ScrollRight: key.NewBinding(
            key.WithKeys("shift+right", "shift+l"),
            key.WithHelp("shift+→/shift+l", "scroll right"),
        ),
    }
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k KeyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Up, k.Down, k.Select, k.ViewDetails, k.Search, k.ScrollLeft, k.ScrollRight, k.Back, k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k KeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down, k.Left, k.Right, k.PageUp, k.PageDown},
        {k.Home, k.End, k.Select, k.ViewDetails, k.Back},
        {k.ScrollLeft, k.ScrollRight, k.Search, k.ClearSearch, k.Help, k.Quit},
    }
}