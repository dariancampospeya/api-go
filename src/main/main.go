package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
)

func main() {
	// api.Run()
	shell := ishell.New()

	// display info.
	shell.Println("Darian Campos Test")

	//Consider the unicode characters supported by the users font
	shell.SetMultiChoicePrompt(" >>", " - ")
	shell.SetChecklistOptions("[ ] ", "[X] ")

	// handle login.
	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Println("Let's simulate login")

			// prompt for input
			c.Print("Username: ")
			username := c.ReadLine()
			c.Print("Password: ")
			password := c.ReadPassword()

			// do something with username and password
			c.Println("Your inputs were", username, "and", password+".")

		},
		Help: "simulate a login",
	})

	// handle "greet".
	shell.AddCmd(&ishell.Cmd{
		Name:    "Saludo",
		Aliases: []string{"Hola", "Bienvenido!!"},
		Help:    "Saludo Usuario",
		Func: func(c *ishell.Context) {
			name := "Stranger"
			if len(c.Args) > 0 {
				name = strings.Join(c.Args, " ")
			}
			c.Println("Hola", name)
		},
	})

	// handle "default".
	shell.AddCmd(&ishell.Cmd{
		Name: "default",
		Help: "readline with default input",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			defaultInput := "Texto por default."
			if len(c.Args) > 0 {
				defaultInput = strings.Join(c.Args, " ")
			}

			c.Print("input: ")
			read := c.ReadLineWithDefault(defaultInput)

			if read == defaultInput {
				c.Println("No has realizado ningun cambio.")
			} else {
				c.Printf("Modificaste el texto por defecto: '%s'", read)
				c.Println()
			}
		},
	})
	// read multiple lines with "multi" command
	shell.AddCmd(&ishell.Cmd{
		Name: "multi",
		Help: "input in multiple lines",
		Func: func(c *ishell.Context) {
			c.Println("Input multiple lines and end with semicolon ';'.")
			lines := c.ReadMultiLines(";")
			c.Println("Done reading. You wrote:")
			c.Println(lines)
		},
	})

	// Agregar nuevo módulo choice
	red := color.New(color.FgRed).SprintFunc()
	shell.AddCmd(&ishell.Cmd{
		Name: "Add",
		Help: "Agregar modulos a su proyecto",
		Func: func(c *ishell.Context) {
			choice := c.MultiChoice([]string{
				"Android",
				"iOS",
				"Backend",
				"Otro",
			}, red("¿Qué scaffolding quieres crear?"))
			if choice == 0 {
				c.Println("Creando scaffolding de Android")
			} else if choice == 1 {
				c.Println("Creando scaffolding de iOS")
			} else if choice == 2 {
				choice2 := c.MultiChoice([]string{
					"Go",
					"NodeJs",
					"Python",
				}, "¿Qué lenguaje desea utilizar?")
				languages := []string{"DB", "Docker", "CI-CD", "Athena", "New Relic", "Data Dog"}
				if choice2 == 0 {
					c.Println("Creando scaffolding de Go")
					choices := c.Checklist(languages,
						"Agregar pre-configuraciones",
						nil)
					out := func() (c []string) {
						for _, v := range choices {
							c = append(c, languages[v])
						}
						return
					}
					c.Println("Configuraciones agregadas", strings.Join(out(), ", "))
				} else if choice2 == 1 {
					c.Println("Creando scaffolding de NodeJS")
					choices := c.Checklist(languages,
						"Agregar pre-configuraciones",
						nil)
					out := func() (c []string) {
						for _, v := range choices {
							c = append(c, languages[v])
						}
						return
					}
					c.Println("Configuraciones agregadas", strings.Join(out(), ", "))
				} else if choice2 == 2 {
					c.Println("Creando scaffolding de Python")
					choices := c.Checklist(languages,
						"Agregar pre-configuraciones",
						nil)
					out := func() (c []string) {
						for _, v := range choices {
							c = append(c, languages[v])
						}
						return
					}
					c.Println("Configuraciones agregadas", strings.Join(out(), ", "))
				}
			} else {
				c.Println("En contrucción..")
			}
		},
	})

	// multiple choice
	shell.AddCmd(&ishell.Cmd{
		Name: "module",
		Help: "Agranda tu combo",
		Func: func(c *ishell.Context) {
			languages := []string{"DB", "Docker", "CI-CD", "Athena", "New Relic", "Data Dog"}
			choices := c.Checklist(languages,
				"Agregar extras",
				nil)
			out := func() (c []string) {
				for _, v := range choices {
					c = append(c, languages[v])
				}
				return
			}
			c.Println("Your choices are", strings.Join(out(), ", "))
		},
	})

	// progress bars
	{
		// determinate
		shell.AddCmd(&ishell.Cmd{
			Name: "det",
			Help: "determinate progress bar",
			Func: func(c *ishell.Context) {
				c.ProgressBar().Start()
				for i := 0; i < 101; i++ {
					c.ProgressBar().Suffix(fmt.Sprint(" ", i, "%"))
					c.ProgressBar().Progress(i)
					time.Sleep(time.Millisecond * 100)
				}
				c.ProgressBar().Stop()
			},
		})

		// indeterminate
		shell.AddCmd(&ishell.Cmd{
			Name: "ind",
			Help: "indeterminate progress bar",
			Func: func(c *ishell.Context) {
				c.ProgressBar().Indeterminate(true)
				c.ProgressBar().Start()
				time.Sleep(time.Second * 10)
				c.ProgressBar().Stop()
			},
		})
	}

	// subcommands and custom autocomplete.
	{
		var words []string
		autoCmd := &ishell.Cmd{
			Name: "suggest",
			Help: "try auto complete",
			LongHelp: `Try dynamic autocomplete by adding and removing words.
Then view the autocomplete by tabbing after "words" subcommand.
This is an example of a long help.`,
		}
		autoCmd.AddCmd(&ishell.Cmd{
			Name: "add",
			Help: "add words to autocomplete",
			Func: func(c *ishell.Context) {
				if len(c.Args) == 0 {
					c.Err(errors.New("missing word(s)"))
					return
				}
				words = append(words, c.Args...)
			},
		})

		autoCmd.AddCmd(&ishell.Cmd{
			Name: "clear",
			Help: "clear words in autocomplete",
			Func: func(c *ishell.Context) {
				words = nil
			},
		})

		autoCmd.AddCmd(&ishell.Cmd{
			Name: "words",
			Help: "add words with 'suggest add', then tab after typing 'suggest words '",
			Completer: func([]string) []string {
				return words
			},
		})

		shell.AddCmd(autoCmd)
	}

	shell.AddCmd(&ishell.Cmd{
		Name: "paged",
		Help: "show paged text",
		Func: func(c *ishell.Context) {
			lines := ""
			line := `%d. This is a paged text input.
This is another line of it.
`
			for i := 0; i < 100; i++ {
				lines += fmt.Sprintf(line, i+1)
			}
			c.ShowPaged(lines)
		},
	})

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	boldRed := color.New(color.FgRed, color.Bold).SprintFunc()
	shell.AddCmd(&ishell.Cmd{
		Name: "color",
		Help: "color print",
		Func: func(c *ishell.Context) {
			c.Print(cyan("cyan\n"))
			c.Println(yellow("yellow"))
			c.Printf("%s\n", boldRed("bold red"))
		},
	})

	// when started with "exit" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}
}
