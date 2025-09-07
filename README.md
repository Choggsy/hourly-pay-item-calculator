# hourly-pay-item-calculator
As an experienced Java and mobile developer, I wanted to explore Go for my new job role by building something simple but useful: 
a CLI tool (command line interface) that calculates how many hours of work an item or service costs. 
You enter your hourly wage and the item's price, and it tells you how long you'd need to work to afford it.

## 🧠 Learning Goals

- 1.1  Get comfortable with Go syntax and structure 
- 1.2  Build a simple CLI tool from scratch
- 1.3  Work with user input and basic arithmetic
- 1.4  Organize code using packages and modules
- 1.5  Handle basic error checking and formatting

## ✅ End Product Requirements

- 2.1 Accept hourly wage and item price as input
- 2.2 Calculate and display hours required to afford item
- 2.3 Run as a standalone CLI app
- 2.4 Maintaine code readable following SOLID

work-cost-calculator/

├── main.go # Entry point of the CLI app 

├── **go.mod** # Go module definition

├── **go.sum** # Go dependencies checksum

├── /models # Structs and data models 

│ └── **item.go** # Item struct (price, name, etc.) 

├── /utils # Helper functions 

│ └── **calculator.go** # Core logic for calculating work hours 

├── _/data # Optional:_ store saved items like user hourly rate or configs 

│ └── **items.json** # sample data file

└── README.md # Project overview and setup


## Development Log

1. download GO language (https://go.dev/dl/), I have the Windows version
2. For IDE's, my favourite is always Jetbrains IntelliJ IDEA. There is a GO plugin which is what il be using for this project.
3. Run go mod init [repository name], this set up declares the project as GO but also acts as the dependency manager like Maven or Gradle in other stacks.
4. Start writing some code and now refer to my debug log!!! 😂
5. Now for tests starting on extracted logic... ideally I'd write tests first but since im learning the language I won't be attempting TDD today
6. Unit testing is similar enough to Junit. IntelliJ is very helpful to creating test functioning using alt + insert. or ctrl + shift + t to make and locate test files. Not loving how Go structures unit tests in the same package as the unit being tested instead of a separate test package like in java
7. Extracted out common testing logic like `t.Errorf(message,expected,result)` into a live template. 
8. 
