# golang-snippets
A collection of Golang snippets I've gathered ranging from problem-solving 
challenges to more practical applications

## Setup
1. Install Go. See [here](https://golang.org/doc/install) for instructions.
2. Clone this repository.
3. Install dependencies.
    ```bash
    go mod download
    ```
4. Run the program.
    ```bash
    go run main.go
    ```
   Note: Most snippets in this repository can be run by importing the package and
   calling the necessary functions. For example, to run the `Quiz` snippet, you
   can import the `quiz` package and call the `Quiz` function.
   ```go
   package main
   import "GolangSnippets/quiz"
   
   ...

    func main() {
         quiz.Quiz()
    }
   ```
   Another thing to note is that some snippets aren't callable for the time being.
   These snippets are yet to be refactored or still in progress.
   
