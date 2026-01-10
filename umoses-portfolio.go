package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
 	"github.com/charmbracelet/lipgloss"
//	"github.com/charmbracelet/glamour"
//	bubbles "github.com/charmbracelet/bubbles"
//	harmonica "github.com/charmbracelet/harmonica"
)

type sessionState int
const (
	homeView sessionState = iota
	stackView
	aboutView
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
	pyt lipgloss.Style
	hon lipgloss.Style
	stackBodyText lipgloss.Style
	aboutBody lipgloss.Style
}

func DefaultStyles() *Styles{
	s:= new(Styles)
	s.Title = lipgloss.NewStyle().Padding(1).MarginBottom(-2).Align(lipgloss.Center, lipgloss.Center).Bold(true).Foreground(lipgloss.AdaptiveColor{Light:"#1c9305", Dark:"#09FF00"})
	s.StackStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("9")).Padding(1).Align(lipgloss.Center, lipgloss.Top).Bold(true).Foreground(lipgloss.AdaptiveColor{Light:"#a0a2a0", Dark:"#555954"})
	s.hyperLinks = lipgloss.NewStyle().Underline(true)
	s.homeBodyText = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center, lipgloss.Center).Padding(1)
	s.stackBodyText = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center, lipgloss.Center).Padding(1).MarginLeft(3).MarginRight(3)
	s.stackLogo = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Padding(1)	
	s.aboutBody = lipgloss.NewStyle()
	return s
}

type home struct {
	styles	*Styles
	state	sessionState
	items   []string
	cursor   int
	selected map[int]struct{}
	stack		stack
	about	about
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
		//languages: []string{"1","2","3"},
		styles: DefaultStyles(),		
	}
}

func initialAbout() about{
	return about{
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
		case "a":
			m.state = aboutView
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
	case aboutView:
		updatedAbout,cmd := m.about.Update(msg)
		m.about =  updatedAbout.(about)

		switch msg := msg.(type){
		case tea.KeyMsg:
			switch msg.String(){
			case "h":
				m.state = homeView
			}
		}
		return m, cmd
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
/*	homepageStringTitle:=`███    █▄  ███▄▄▄▄      ▄█    █▄     ▄██████▄   ▄█       ▄██   ▄     ▄▄▄▄███▄▄▄▄    ▄██████▄     ▄████████    ▄████████    ▄████████ 
███    ███ ███▀▀▀██▄   ███    ███   ███    ███ ███       ███   ██▄ ▄██▀▀▀███▀▀▀██▄ ███    ███   ███    ███   ███    ███   ███    ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███▄▄▄███ ███   ███   ███ ███    ███   ███    █▀    ███    █▀    ███    █▀  
███    ███ ███   ███  ▄███▄▄▄▄███▄▄ ███    ███ ███       ▀▀▀▀▀▀███ ███   ███   ███ ███    ███   ███         ▄███▄▄▄       ███        
███    ███ ███   ███ ▀▀███▀▀▀▀███▀  ███    ███ ███       ▄██   ███ ███   ███   ███ ███    ███ ▀███████████ ▀▀███▀▀▀     ▀███████████ 
███    ███ ███   ███   ███    ███   ███    ███ ███       ███   ███ ███   ███   ███ ███    ███          ███   ███    █▄           ███ 
███    ███ ███   ███   ███    ███   ███    ███ ███▌    ▄ ███   ███ ███   ███   ███ ███    ███    ▄█    ███   ███    ███    ▄█    ███ 
████████▀   ▀█   █▀    ███    █▀     ▀██████▀  █████▄▄██  ▀█████▀   ▀█   ███   █▀   ▀██████▀   ▄████████▀    ██████████  ▄████████▀  
                                               ▀                                                                                     `*/
	homepageStringTitle:=`         _               _______ _               _______ _______ _______ ______  _______ 
|\     /( (    /|\     /(  __   | \  |\     /|  (       |  ___  |  ____ Y ___  \(  ____ \
| )   ( |  \  ( | )   ( | (  )  | (  ( \   / )  | () () | (   ) | (    \|/   \  \ (    \/
| |   | |   \ | | (___) | | /   | |   \ (_) /   | || || | |   | | (_____   ___) / (_____ 
| |   | | (\ \) |  ___  | (/ /) | |    \   /    | |(_)| | |   | (_____  ) (___ ((_____  )
| |   | | | \   | (   ) |   / | | |     ) (     | |   | | |   | |     ) |     ) \     ) |
| (___) | )  \  | )   ( |  (__) | (____/\ |     | )   ( | (___) /\____) /\___/  /\____) |
(_______)/    )_)/     \(_______|_______|_/     |/     \(_______)_______)______/\_______)
                                                                                         `

	// add terminal hyperlinks
	homepageString:= ""
	GitUrl:= "https://github.com/unh0lymos3s"
	GitText:= " Github"
	homepageString+= fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\ ", GitUrl, GitText)
	lnkdinUrl := "https://linkedin.com/in/moosabinnaseem"
	lnkdinText := " LinkedIn"
	homepageString+= fmt.Sprintf("\t | \t \x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", lnkdinUrl, lnkdinText)
	hfUrl := "https://huggingface.co/Moosa01"
	hfText := "\U0001F917 HuggingFace"
	homepageString+= fmt.Sprintf("\t | \t \x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", hfUrl, hfText)
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
	stackStringTitle:= `_______________  _______            _______________________ _______ _       
\__   __/ ___  \(  ____ \\     /|  (  ____ \__   __(  ___  |  ____ \ \    /\
   ) (  \/   \  \ (    \/ )   ( |  | (    \/  ) (  | (   ) | (    \/  \  / /
   | |     ___) / |     | (___) |  | (_____   | |  | (___) | |     |  (_/ / 
   | |    (___ (| |     |  ___  |  (_____  )  | |  |  ___  | |     |   _ (  
   | |        ) \ |     | (   ) |        ) |  | |  | (   ) | |     |  ( \ \ 
   | |  /\___/  / (____/\ )   ( |  /\____) |  | |  | )   ( | (____/\  /  \ \
   )_(  \______/(_______//     \|  \_______)  )_(  |/     \(_______/_/    \/
                                                                            `
	stackStringRust :=`██████╗ ██╗   ██╗███████╗████████╗
██╔══██╗██║   ██║██╔════╝╚══██╔══╝
██████╔╝██║   ██║███████╗   ██║   
██╔══██╗██║   ██║╚════██║   ██║   
██║  ██║╚██████╔╝███████║   ██║   
╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝   
                                  `
	stackStringPyt:= `██████╗ ██╗   ██╗████████╗
██╔══██╗╚██╗ ██╔╝╚══██╔══╝
██████╔╝ ╚████╔╝    ██║   
██╔═══╝   ╚██╔╝     ██║   
██║        ██║      ██║   
╚═╝        ╚═╝      ╚═╝   
                          `
	stackStringHon:= `██╗  ██╗ ██████╗ ███╗   ██╗
██║  ██║██╔═══██╗████╗  ██║
███████║██║   ██║██╔██╗ ██║
██╔══██║██║   ██║██║╚██╗██║
██║  ██║╚██████╔╝██║ ╚████║
╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
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
	pytRendered:= s.styles.stackLogo.Foreground(lipgloss.Color("#2596be")).Render(stackStringPyt)
	honRendered:= s.styles.stackLogo.Foreground(lipgloss.Color("#fec008")).PaddingLeft(-1).Render(stackStringHon)
	goRendered:= s.styles.stackLogo.Foreground(lipgloss.Color("#2596be")).Render(stackStringGo)
	titleRendered:= s.styles.StackStyle.Width(s.width  - 2).Render(stackStringTitle)
	stackBody:= s.styles.stackBodyText.Width(s.width-5).Render(stackStringBody)
	return lipgloss.Place(
		s.width,
		s.height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical( lipgloss.Center, titleRendered,padding, lipgloss.JoinHorizontal(lipgloss.Center, rustRendered,pipe,pytRendered,honRendered,pipe,goRendered), stackBody),

	)
}


type about struct{
	styles *Styles
	state sessionState
	about []string
	width int
	height int
}

func (a about) Init() tea.Cmd{
	return nil
}

func (a about) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return a, tea.Quit
		}
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
	}

	return a,nil
}

func (a about) View() string{
	aboutString:= `skjdfsdljkfsn`
	abtRendered:= a.styles.aboutBody.Render(aboutString)
	return lipgloss.Place(
		a.height,
		a.width,
		lipgloss.Center,
		lipgloss.Top,
		abtRendered,
	)
}
