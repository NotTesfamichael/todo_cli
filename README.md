
## Run Locally

Clone the project

```bash
  git clone https://github.com/tesfamichael28/todo_cli
```

Go to the project directory

```bash
  cd todo_cli
```

Build the application

```bash
  go build
```

Run this command to list

```bash
  ./main -list
```

Run this command to edit the current

```bash
  ./main -edit "1:new_title"
```
Run this command to add a new todo list

```bash
  ./main -add "Your todo"
```
Run this command to toggle from No to YES or vice versa

```bash
  ./main -toggle #
  
  example
  /main -toggle 2
```
