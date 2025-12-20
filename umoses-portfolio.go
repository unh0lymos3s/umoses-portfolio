package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
//	gloss "github.com/charmbracelet/lipgloss"
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
	p:= tea.NewProgram(initialModel())
	if _, err := p.Run(); err !=nil {
		fmt.Printf("There been an error %v", err)
		os.Exit(1)
	}
}

type home struct {
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
		items: []string{"Stack", "Projects", "Experience", "Contact"},

		selected: make(map[int]struct{}),
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
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor <= len(m.items)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "s":
			m.state = stackView
		}


	var cmd tea.Cmd
	switch m.state {
	case stackView:
			m.stack, cmd = m.stack.Update(msg)
	}
	return m, nil
}

//view home
func (m home) View() string {
	
	homepageString := `
                                                                                                                                          
                             *                  ***                         *****   **    **                                              
                           **                    ***                     ******  ***** *****                                              
                           **                     **                    **   *  *  ***** *****                                            
                           **                     **                   *    *  *   * **  * **                                             
**   ****                  **           ****      **    **   ****          *  *    *     *        ****       ****                 ****    
 **    ***  * ***  ****    **  ***     * ***  *   **     **    ***  *     ** **    *     *       * ***  *   * **** *    ***      * **** * 
 **     ****   **** **** * ** * ***   *   ****    **     **     ****      ** **    *     *      *   ****   **  ****    * ***    **  ****  
 **      **     **   ****  ***   *** **    **     **     **      **       ** **    *     *     **    **   ****        *   ***  ****       
 **      **     **    **   **     ** **    **     **     **      **       ** **    *     *     **    **     ***      **    ***   ***      
 **      **     **    **   **     ** **    **     **     **      **       ** **    *     **    **    **       ***    ********      ***    
 **      **     **    **   **     ** **    **     **     **      **       *  **    *     **    **    **         ***  *******         ***  
 **      **     **    **   **     ** **    **     **     **      **          *     *      **   **    **    ****  **  **         ****  **  
  ******* **    **    **   **     **  ******      **      *********      ****      *      **    ******    * **** *   ****    * * **** *   
   *****   **   ***   ***  **     **   ****       *** *     **** ***    *  *****           **    ****        ****     *******     ****    
                 ***   ***  **    **               ***            ***  *     **                                        *****              
                                  *                        *****   *** *                                                                  
                                 *                       ********  **   **                                                                
                                *                       *      ****                                                                       
                               *                                                                                                          
                                                                                                                                          

`	
	for i, item := range m.items{
		cursor := " "
			if m.cursor == i{
			cursor = ">"
			}

		checked := " "
			if _, ok := m.selected[i]; ok {
			checked = "X"
		  }

		homepageString += fmt.Sprintf("%s [%s] %s\n", cursor, checked,item)
		}
	homepageString+= "\n Press Q to quit."
	return homepageString
}
//struct for Stack info
type stack struct{
	languages []string
}


//Initialize Stack
func (s stack) Init (tea.Cmd) {
	return nil
}


//Updaate stack
func (s stack) Update(msg tea.Msg) (tea.Msg, tea.Cmd){

	switch msg := msg.(type){
	case msg.KeyMsg:
		switch msg.String(){
		case "q":
			return s, tea.Quit
		}
	}
	return s,nil
}

// ViewStack
func (s stack) View() string{
	stackString:= "My stack basically comprises of Python, Rust, Go and C/C++, However, I would classify my strongest suite as python"

	return stackString
}
