package main

import (
	"context"
	"fmt"
	crt "project101/api/createAcc"
	"sync"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var (
	wg = &sync.WaitGroup{}
)

type city struct{
	name string
	lat float64
	lon float64
}
//points on the globe
var cities = []city{
	{"Nairobi",-1.2921,36.8219},
	{"New York", 40.7128,-74.0060},
	{"Tokyo", 35.6895,139.692},
	{"London", 51.5074, -0.1278},
	{"Sao Paulo", -23.5505,-46.6333},
	{"Sydney", -33.8688, 151.2093},
	{"Cairo", 30.0444, 31.2357},
}
//messages
type tickMsg time.Time
type finishedMsg string
//model
type viewState int

const(
	menuView viewState = iota
	runningView 
)

type model struct{
	choices []string
	cursor int
	state viewState
	status string
	angle float64
	actions map[int]func() tea.Cmd
}

func initialModel()model{
	m := model{
		choices: []string{"file input", "type input"},
		state: menuView,
		status: "Ready- select an option",
	}
	m.actions = map[int]func() tea.Cmd{
		0:login,
		1:createAcc,
	}

	return m
	
}

func (m model) Init() tea.Cmd{
	return tea.Tick(80*time.Millisecond,func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func login()tea.Cmd{
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	return func() tea.Msg {
		lgn.GetUser().LoginApi(ctx,wg)	
		wg.Wait()

		return finishedMsg("Logged in successfully")	    	
	}
}

func createAcc()tea.Cmd{
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	return func() tea.Msg {
		crt.GetUser().CreateApi(ctx,wg)
		wg.Wait()
		return finishedMsg("created account successfully")
	}	
}

func beep(){fmt.Print("\a")}

func(m model)Update(msg tea.Msg)(tea.Model,tea.Cmd){
	switch msg := msg.(type){
		case tickMsg:
		m.angle += 0.08
		return m , tea.Tick(80*time.Millisecond,func(t time.Time)tea.Msg{
			return tickMsg(t)
		})
		//process completed ...restore initial state
		case finishedMsg:
		m.state = menuView
		m.status = string(msg)
		beep()
		return m, nil

		case tea.KeyMsg:
		switch msg.String(){
		case "q","ctrl+c","esc":
		return m, tea.Quit

		case "enter":
		if m.state == menuView{
			m.state = runningView
			m.status = fmt.Sprintf("Running %s...", m.choices[m.cursor])
			if fn, ok := m.actions[m.cursor];ok{
				return m, fn()
			}
		}

		case "up","k":
		if m.state == menuView && m.cursor > 0 {
			m.cursor --
			beep()
		}

		case "down","j":
		if m.state == menuView && m.cursor < len(m.choices)-1{
			m.cursor++
			beep()
	  }
	
	}
  }
return m, nil
}


func (m model) View()string{

	menuStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#555")).Padding(1,2).Width(100)
	g
    for i, ch := range m.choices{
     cursor := ""
     if i== m.cursor && m.state == menuView{
     cursor = ">"
     line := lipgloss.NewStyle().Background(lipgloss.Color("#2A4A3A")).Foreground(lipgloss.Color("#4EDE7F")).Bold(true).Render(fmt.Sprintf("%s %s",cursor, ch))

     }else{
     line = fmt.Sprintf("%s %s", cursor,ch)
     }
  
}
return menuStyle.String()
}

func main(){
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil{
		fmt.Println(err)
	}
}