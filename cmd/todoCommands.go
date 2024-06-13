package cmd

import (
	"fmt"
	"log"

	"github.com/Yashh56/todo/controllers"
	"github.com/Yashh56/todo/model"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

var AddFlag bool
var DeleteFlag bool
var UpdateFlag bool
var RemoveAllFlag bool
var TodoCmd = &cobra.Command{
	Use:   "todo [flags]",
	Short: "Manage your todo items (Create,Delete,Update)",
	Long:  `A tool to manage your todo items with add, delete, and get all functionalities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		if AddFlag {
			todo := model.Model{
				Title: title,
			}

			addTodo, err := controllers.AddTodo(todo)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(addTodo)
		} else if DeleteFlag {
			result, err := controllers.DeleteTodo(title)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		} else if UpdateFlag {
			result, err := controllers.UpdateTodo(title)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		} else {
			log.Fatal("You must specify an action with either --add or --delete")
		}
	},
}
var GetAllCmd = &cobra.Command{
	Use:   "todos",
	Short: "Get all todo items",
	Long:  `Retrieve and display all done todo items from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := controllers.GetAllTodos()
		if err != nil {
			fmt.Println("Failed to retrieve todos:", err)
			return
		}

		HeaderStyle := lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center).Width(12)
		EvenRowStyle := lipgloss.NewStyle().Foreground(lightGray).Padding(1, 2)
		OddRowStyle := lipgloss.NewStyle().Foreground(gray).Padding(1, 2)

		var rows [][]string

		for _, todo := range todos {
			createdAt := todo.CreatedAt.Format("2006-01-02 15:04:05")
			doneStr := fmt.Sprintf("%v", todo.Done)
			row := []string{todo.Title, createdAt, doneStr}
			rows = append(rows, row)
		}
		if len(rows) == 0 {
			fmt.Println("No data found")
			return
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == 0:
					return HeaderStyle
				case row%2 == 0:
					return EvenRowStyle
				default:
					return OddRowStyle
				}
			}).
			Headers("Todos", "Time", "Completed").
			Rows(rows...)
		fmt.Println(t.Render())
	},
}
var GetAllCompletedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Get the list of all completed todos",
	Long:  `Retrieve and display all done todo items from the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		todos, err := controllers.GetAllDoneTodos()
		if err != nil {
			fmt.Println("Failed to retrieve todos:", err)
			return
		}

		HeaderStyle := lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center).Width(12)
		EvenRowStyle := lipgloss.NewStyle().Foreground(lightGray).Padding(1, 2)
		OddRowStyle := lipgloss.NewStyle().Foreground(gray).Padding(1, 2)

		// Prepare table headers and rows
		var rows [][]string

		for _, todo := range todos {
			createdAt := todo.CreatedAt.Format("2006-01-02 15:04:05")
			doneStr := fmt.Sprintf("%v", todo.Done)
			row := []string{todo.Title, createdAt, doneStr}
			rows = append(rows, row)
		}

		if len(rows) == 0 {
			fmt.Println("No data found")
			return
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == 0:
					return HeaderStyle
				case row%2 == 0:
					return EvenRowStyle
				default:
					return OddRowStyle
				}
			}).
			Headers("Todos", "Time", "Completed").
			Rows(rows...)

		fmt.Println(t.Render())
	},
}
var GetPendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "Get the list of all completed todos",
	Long:  `Retrieve and display all done todo items from the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		todos, err := controllers.GetPendingTodo()
		if err != nil {
			fmt.Println("Failed to retrieve todos:", err)
			return
		}

		HeaderStyle := lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center).Width(12)
		EvenRowStyle := lipgloss.NewStyle().Foreground(lightGray).Padding(1, 2)
		OddRowStyle := lipgloss.NewStyle().Foreground(gray).Padding(1, 2)

		// Prepare table headers and rows
		var rows [][]string

		for _, todo := range todos {
			createdAt := todo.CreatedAt.Format("2006-01-02 15:04:05")
			doneStr := fmt.Sprintf("%v", todo.Done)
			row := []string{todo.Title, createdAt, doneStr}
			rows = append(rows, row)
		}
		if len(rows) == 0 {
			fmt.Println("No Pending Todo found")
			return
		}
		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == 0:
					return HeaderStyle
				case row%2 == 0:
					return EvenRowStyle
				default:
					return OddRowStyle
				}
			}).
			Headers("Todos", "CreatedAt", "Completed").
			Rows(rows...)

		fmt.Println(t.Render())
	},
}
var DeleteAllCmd = &cobra.Command{
	Use:   "deleteAll",
	Short: "Delete all todo items",
	Long:  `Delete all todo items from the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		result, err := controllers.DeleteAllTodos()

		if err != nil {
			fmt.Println("Failed to delete todos:", err)
			return
		}
		fmt.Println(result)
	},
}
