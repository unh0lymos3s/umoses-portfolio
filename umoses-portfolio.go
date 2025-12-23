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
	StackStyle lipgloss.Style
	hyperLinks lipgloss.Style
	homeBodyText lipgloss.Style
	stackLogo	lipgloss.Style
}

func DefaultStyles() *Styles{
	s:= new(Styles)
	s.Title = lipgloss.NewStyle().BorderForeground(lipgloss.Color("9")).Padding(1).MarginBottom(-2).Align(lipgloss.Center, lipgloss.Center).Bold(true).Foreground(lipgloss.Color("195"))
	s.StackStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("9")).Padding(1).Align(lipgloss.Center, lipgloss.Top).Bold(true).Foreground(lipgloss.Color("195"))
	s.hyperLinks = lipgloss.NewStyle().Underline(true)
	s.homeBodyText = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center, lipgloss.Center).Padding(1)
	s.stackLogo = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Padding(1)	
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
		styles: DefaultStyles(),		
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
		m.stack.width = msg.Width
		m.stack.height = msg.Height
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
	homepageStringTitle:=`███    █▄  ███▄▄▄▄      ▄█    █▄     ▄██████▄   ▄█       ▄██   ▄     ▄▄▄▄███▄▄▄▄    ▄██████▄     ▄████████    ▄████████    ▄████████ 
███    ███ ███▀▀▀██▄   ███    ███   ███    ███ ███       ███   ██▄ ▄██▀▀▀███▀▀▀██▄ ███    ███   ███    ███   ███    ███   ███    ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███▄▄▄███ ███   ███   ███ ███    ███   ███    █▀    ███    █▀    ███    █▀  
███    ███ ███   ███  ▄███▄▄▄▄███▄▄ ███    ███ ███       ▀▀▀▀▀▀███ ███   ███   ███ ███    ███   ███         ▄███▄▄▄       ███        
███    ███ ███   ███ ▀▀███▀▀▀▀███▀  ███    ███ ███       ▄██   ███ ███   ███   ███ ███    ███ ▀███████████ ▀▀███▀▀▀     ▀███████████ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███   ███ ███   ███   ███ ███    ███          ███   ███    █▄           ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███▌    ▄ ███   ███ ███   ███   ███ ███    ███    ▄█    ███   ███    ███    ▄█    ███ 
████████▀   ▀█   █▀    ███    █▀     ▀██████▀  █████▄▄██  ▀█████▀   ▀█   ███   █▀   ▀██████▀   ▄████████▀    ██████████  ▄████████▀  
                                               ▀                                                                                     `
	// add terminal hyperlinks
	homepageString:= ""
	GitUrl:= "https://github.com/unh0lymos3s"
	GitText:= "Github"
	homepageString+= fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\ ", GitUrl, GitText)
	lnkdinUrl := "https://linkedin.com/in/moosabinnaseem"
	lnkdinText := "LinkedIn"
	homepageString+= fmt.Sprintf("\t | \t \x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", lnkdinUrl, lnkdinText)
	//homepageString+=fmt.Sprintf("%s", gitLink)
	homepageString+= "\n\n\n[A] About\t|\t[S] Stack\t|\t[R] Resume\t|\t[P] Projects\t|\t[E] Experience\n\n\n"	
	homepageString+= "\n Press Q to quit."
	homePageStringTitleRendered := m.styles.Title.Align(lipgloss.Center, lipgloss.Center).Render(homepageStringTitle)
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center, 
		lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, homePageStringTitleRendered,
		m.styles.homeBodyText.Render(homepageString)),
	)
}
//struct for Stack info
type stack struct{
	width int
	height int
	languages []string
	styles *Styles
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
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
	}

	return s,nil
}

// ViewStack
func (s stack) View() string{
	stackStringTitle:= `    ███        ▄████████  ▄████████    ▄█    █▄            ▄████████     ███        ▄████████  ▄████████    ▄█   ▄█▄ 
▀█████████▄   ███    ███ ███    ███   ███    ███          ███    ███ ▀█████████▄   ███    ███ ███    ███   ███ ▄███▀ 
   ▀███▀▀██   ███    █▀  ███    █▀    ███    ███          ███    █▀     ▀███▀▀██   ███    ███ ███    █▀    ███▐██▀   
    ███   ▀  ▄███▄▄▄     ███         ▄███▄▄▄▄███▄▄        ███            ███   ▀   ███    ███ ███         ▄█████▀    
    ███     ▀▀███▀▀▀     ███        ▀▀███▀▀▀▀███▀       ▀███████████     ███     ▀███████████ ███        ▀▀█████▄    
    ███       ███    █▄  ███    █▄    ███    ███                 ███     ███       ███    ███ ███    █▄    ███▐██▄   
    ███       ███    ███ ███    ███   ███    ███           ▄█    ███     ███       ███    ███ ███    ███   ███ ▀███▄ 
   ▄████▀     ██████████ ████████▀    ███    █▀          ▄████████▀     ▄████▀     ███    █▀  ████████▀    ███   ▀█▀ 
                                                                                                           ▀         `
	stackStringRust :=`██████╗ ██╗   ██╗███████╗████████╗
██╔══██╗██║   ██║██╔════╝╚══██╔══╝
██████╔╝██║   ██║███████╗   ██║   
██╔══██╗██║   ██║╚════██║   ██║   
██║  ██║╚██████╔╝███████║   ██║   
╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝   
                                  `
	stackStringPython:= `██████╗ ██╗   ██╗████████╗██╗  ██╗ ██████╗ ███╗   ██╗
██╔══██╗╚██╗ ██╔╝╚══██╔══╝██║  ██║██╔═══██╗████╗  ██║
██████╔╝ ╚████╔╝    ██║   ███████║██║   ██║██╔██╗ ██║
██╔═══╝   ╚██╔╝     ██║   ██╔══██║██║   ██║██║╚██╗██║
██║        ██║      ██║   ██║  ██║╚██████╔╝██║ ╚████║
╚═╝        ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
                                                     `
	stackStringGo:=` ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗
██║  ███╗██║   ██║
██║   ██║██║   ██║
╚██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝ 
                  `
	pipe:= `  ██  `
	padding:="\n\n\n\n"
	stackStringBody:= "I've been working with Machine Learning and Data Science for almost 2 years now, so Python is my daily driver. Having some backend experience with Node.JS and Flask, I have started to hack away at Rust and GO as I transition more into systems programming and HPC; currently building Pherrous (A distributed network computing system written purely in rust) and this portfolio using Go and Bubbletea framework."
	stackStringBody+= "\n\n\n [Q] Quit  |  [H] Home"
	rustRendered:= s.styles.stackLogo.Foreground(lipgloss.Color("#E97451")).Render(stackStringRust)
	pyRendered:= s.styles.stackLogo.Render(stackStringPython)
	goRendered:= s.styles.stackLogo.Foreground(lipgloss.Color("#2596be")).Render(stackStringGo)
	titleRendered:= s.styles.StackStyle.Width(s.width  - 2).Render(stackStringTitle)
	stackBody:= s.styles.homeBodyText.Width(s.width-4).Render(stackStringBody)
	return lipgloss.Place(
		s.width,
		s.height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical( lipgloss.Center, titleRendered,padding, lipgloss.JoinHorizontal(lipgloss.Center, rustRendered,pipe,pyRendered,pipe,goRendered), stackBody),

	)
}
