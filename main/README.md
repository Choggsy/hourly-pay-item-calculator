# hourly-pay-item-calculator
As an experienced Java and mobile developer, I wanted to explore Go for my new job role by building something simple but useful: 
a CLI tool (command line interface) that calculates how many hours of work an item or service costs. 
You enter your hourly wage and the item's price, and it tells you how long you'd need to work to afford it.

---

##  Iteration 1: CLI Foundations

### ✅ Learning Goals
- 1.1 Get comfortable with Go syntax and structure
- 1.2 Build a simple CLI tool from scratch
- 1.3 Work with user input and basic arithmetic
- 1.4 Organize code using packages and modules
- 1.5 Handle basic error checking and formatting

### ✅ End Product Requirements
- 2.1 Accept hourly wage and item price as input
- 2.2 Calculate and display hours required to afford item
- 2.3 Run as a standalone CLI app
- 2.4 Maintain code readability following SOLID principles

---

##  🎯 Iteration 2: Wage Management & Front-End Integration

###  Learning Goals
- 2.1 Store and manage user wage data (hourly and annual)
- 2.2 Accept and calculate monthly work schedule (days/week, hours/day)
- 2.3 Calculate monthly income and item affordability
- 2.4 Build a basic front-end interface to interact with Go logic
- 2.5 Serve web pages and handle form submissions using Go’s `net/http`

### End Product Requirements
- 3.1 Accept and store user’s annual and hourly wage
- 3.2 Accept and store user’s monthly work schedule
- 3.3 Calculate monthly hours worked and monthly income
- 3.4 Determine item affordability based on monthly income
- 3.5 Build a basic front-end interface (HTML/CSS/JS)
- 3.6 Connect front-end to Go backend using HTTP handlers

---

## 📦 Tech Stack

- Go (CLI + backend logic)
- HTML/CSS/JavaScript (basic front-end)
- `net/http` for serving and handling web requests

---

## 🛠️ Next Steps

- [ ] Create `UserProfile` structure to store wage and schedule
- [ ] Build HTML form for user input
- [ ] Set up Go web server and handlers
- [ ] Connect front-end to backend logic
- [ ] Add monthly income and affordability calculations  

hourly-pay-item-calculator/
├── main/
│   ├── main.model/
│   │   └── work_calculator.go
│   ├── main.model.utils/
│   │   ├── calculator.go
│   │   ├── calculator_test.go
│   │   ├── input_validator.go
│   │   └── input_validator_test.go
│   ├── web/
│   │   ├── handlers.go         # HTTP handlers for form processing
│   │   ├── templates/
│   │   │   └── form.html       # HTML form for user input
│   │   └── web.go              # Web server entry point
│   ├── main.go                 # CLI entry point
│   ├── go.mod
│   ├── README.md
│   ├── debug-log.md
│   └── live-template-log.md
├── assets/
│   └── img.png



## Development Log

1. download GO language (https://go.dev/dl/), I have the Windows version
2. For IDE's, my favourite is always Jetbrains IntelliJ IDEA. There is a GO plugin which is what il be using for this project.
3. Run go mod init [repository name], this set up declares the project as GO but also acts as the dependency manager like Maven or Gradle in other stacks.
4. Start writing some code! .... and now refer to my debug log 😂
5. Now for tests starting on extracted logic... ideally I'd write tests first but since im learning the language I won't be attempting TDD today
6. Unit testing is similar enough to Junit. IntelliJ is very helpful to creating test functioning using alt + insert. or ctrl + shift + t to make and locate test files. Not loving how Go structures unit tests in the same package as the unit being tested instead of a separate test package like in java. I have found how to make 'nested' classes in the test file like in java, i like how the test cases for t.run () have more readable names then the test function naming convention
7. 7. Extracted out common testing logic like `t.Errorf(message,expected,result)` into a live template. 
8. finished basic CLI tool. now making new requirements for additional features and an improved web front end
<br> 


https://github.com/user-attachments/assets/3bca3a8a-5f34-4e05-8b94-05b26eccd36a


<br>

         
