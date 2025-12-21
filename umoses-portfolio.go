package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
 	"github.com/charmbracelet/lipgloss"
//	wish "github.com/charmbracelet/wish"
//	bubbles "github.com/charmbracelet/bubbles"
//	harmonica "github.com/charmbracelet/harmonica"
)

type sessionState int
const (
	homeView sessionState = iota
	stackView
//	aboutView
//	projsNexpView 
	host = "0.0.0.0"
	port = 23234
)

func main() {
	//initialStack := []string{"Go", "Rust", "Python"}
	p:= tea.NewProgram(initialModel())
	
	if _, err := p.Run(); err !=nil {
		fmt.Printf("There been an error %v", err)
		os.Exit(1)
	}
}

type Styles struct{
	Title	lipgloss.Style
}

func DefaultStyles() *Styles{
	s:= new(Styles)
	s.Title = lipgloss.NewStyle().BorderForeground(lipgloss.Color("9")).BorderStyle(lipgloss.DoubleBorder()).Padding(1).Align(lipgloss.Center, lipgloss.Center).Bold(true).Background(lipgloss.Color("239")).Foreground(lipgloss.Color("195"))
	return s
}

type home struct {
	styles	*Styles
	state	sessionState
	items   []string
	cursor   int
	selected map[int]struct{}
	stack		stack
	width int 
	height int
}



func initialModel() home {
	return home{
		state: homeView,
		items: []string{"Stack", "Projects", "Experience", "Contact"},
		stack: initialStack(),
		selected: make(map[int]struct{}),
		styles: DefaultStyles(),
	}
}

func initialStack() stack {
	return stack{
		languages: []string{"1","2","3"},
	}
}


//initialize home
func (m home) Init() tea.Cmd {
	return nil
}
//update home
func (m home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "s":
			m.state = stackView
		}
		
	}
	switch m.state{
	case stackView:
		updatedStack, cmd := m.stack.Update(msg)
		m.stack = updatedStack.(stack)
		
		switch msg:= msg.(type){
		case tea.KeyMsg:
			switch msg.String(){
			case "h":
				m.state = homeView

		}
		return m,cmd
	}


	}
	return m,nil
}

//view home
func (m home) View() string{
	switch m.state{
	case homeView:
		return m.homeView()
	
	case stackView:
		return m.stack.View()
	}
	return "unknown"
}

//helper function to view homepage
func (m home) homeView() string{
	homepageString:=`███    █▄  ███▄▄▄▄      ▄█    █▄     ▄██████▄   ▄█       ▄██   ▄     ▄▄▄▄███▄▄▄▄    ▄██████▄     ▄████████    ▄████████    ▄████████ 
███    ███ ███▀▀▀██▄   ███    ███   ███    ███ ███       ███   ██▄ ▄██▀▀▀███▀▀▀██▄ ███    ███   ███    ███   ███    ███   ███    ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███▄▄▄███ ███   ███   ███ ███    ███   ███    █▀    ███    █▀    ███    █▀  
███    ███ ███   ███  ▄███▄▄▄▄███▄▄ ███    ███ ███       ▀▀▀▀▀▀███ ███   ███   ███ ███    ███   ███         ▄███▄▄▄       ███        
███    ███ ███   ███ ▀▀███▀▀▀▀███▀  ███    ███ ███       ▄██   ███ ███   ███   ███ ███    ███ ▀███████████ ▀▀███▀▀▀     ▀███████████ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███   ███ ███   ███   ███ ███    ███          ███   ███    █▄           ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███▌    ▄ ███   ███ ███   ███   ███ ███    ███    ▄█    ███   ███    ███    ▄█    ███ 
████████▀   ▀█   █▀    ███    █▀     ▀██████▀  █████▄▄██  ▀█████▀   ▀█   ███   █▀   ▀██████▀   ▄████████▀    ██████████  ▄████████▀  
                                               ▀                                                                                     `
	// add terminal hyperlinks
	homepageString+= "\n\n"
	homepageString+= "\n\n[A] About\n\n[S] Stack\n\n[E] Experience\n\n[P] Projects\n\n[R] Resume\n\n\n"	
	homepageString+= "\n Press Q to quit."
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center, 
		lipgloss.Center,
		m.styles.Title.Width(m.width - 2).Height(m.height-3).Render(homepageString),
	)
}
//struct for Stack info
type stack struct{
	languages []string
}


//Initialize Stack
func (s stack) Init() tea.Cmd{
	return nil
}


//Updaate stack
func (s stack) Update(msg tea.Msg) (tea.Model, tea.Cmd){

	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return s, tea.Quit
		}
	}
	return s,nil
}

// ViewStack
func (s stack) View() string{

	stackString:= "I am really interested in backend development along with TUIs and AI so my tech stack is rather simple"

	return stackString
}
