
# Todo CLI

A Command Line Interface Where users can Create Update and Delete Todos.MongoDB has been used as a a database for Todos, Cobra library is used for creating this amazing Todo CLI, and Lip gloss is used for designing the output(Todos) 💙
.


## Installation

To set up the project locally, follow these steps:

```bash
git clone https://github.com/Yashh56/Todo-CLI.git
```

### Build
```bash
go build
```





## Environment Variables

To run this project, you will need to add the following environment variables.

`MONGODB_URI`


## Demo

![Windows PowerShell 2024-06-13 20-34-42 (online-video-cutter com)](https://github.com/Yashh56/Todo-CLI/assets/141008488/d085f50a-d4cf-4a45-9e64-8986c683c5a0)



## Commands



#### Create Todo

```bash
go run main.go todo --add "Go Code"
```
#### Delete Todo From Database

```bash
go run main.go todo --remove "Go-Code"
```
#### Update Todo (completed)

```bash
go run main.go todo --update "Go Code"
```

#### Get All Todos from DB

```bash
go run main.go todos
```
#### Get Pending (Incompleted) Todos from DB

```bash
go run main.go pending
```
#### Delete all Todos from the DB

```bash
go run main.go deleteAll
```
#### Get all completed Todos from DB

```bash
go run main.go completed
```


## Packages used in this Project

 - [Cobra](https://github.com/spf13/cobra)
 - [Lip Gloss](https://github.com/charmbracelet/lipgloss)
 

