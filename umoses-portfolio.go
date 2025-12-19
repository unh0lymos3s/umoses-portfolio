package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	host = "0.0.0.0"
	port = 23234
)

type model struct {
	items    []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		items: []string{"Stack", "Projects", "Experience", "Contact"},

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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

		}

	}
	return m, nil
}

func (m model) View() string {
	s := `[[
                                                                                                                                          
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
                                                                                                                                          
]]
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

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked,item)
		}
	s+= "\n Press Q to quit."
	return s
}

func main() {
	p:= tea.NewProgram(initialModel())
	if _, err := p.Run(); err !=nil {
		fmt.Printf("There been an error %v", err)
		os.Exit(1)
	}
}
