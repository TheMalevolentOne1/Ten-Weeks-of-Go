# Ten Weeks of Go Learning
## Installation (When Week 1 Complete (Not Current))

Clone the repo and build the binary from the `Week` folder. E.g:

```console
$ git clone https://github.com/TheMalevolentOne1/Ten-Weeks-of-Go.git
$ cd "Ten-Weeks-of-Go/Week 1 - CLITaskList"
$ go build -o tasklist .
```

This creates a `tasklist` executable you can run with `./tasklist`.

## [Week 1 — Command-Line Interface Tasklist](https://themalevolentone1.github.io/My-Notes-Collection/Notes/Programming/Go/Ten%20Weeks%20of%20Go/Go%20-%20Ten%20Weeks%20-%20Week%201%20-%20CLI%20Tasklist/)
(11-09-2026) - (18-09-2026)
### Objective

The objective for **Week 1** is to create a command-line interface (CLI) task list in **Go**.

The application will allow the user to:

* View all tasks
* Add a new task
* Remove an existing task
* Persist tasks between program executions

Tasks are stored locally in a **JSON file** and are assigned a unique, incremental ID.

---

## Features

* [ ] List all tasks
* [ ] Add a new task
* [ ] Remove an existing task
* [ ] Persist tasks using a JSON file
* [ ] Sort tasks alphabetically
* [ ] Assign each task a unique incremental ID
* [ ] Implement minimal error handling
* [ ] Unit test primary functionality

---

## Usage

### List All Tasks

```console
$ tasklist all
```

```text
+--------------------+
| TASKLIST MANAGER   |
+--------------------+
| 1. A Task One!     |
| 2. B Task Two!     |
| 3. C Task Three!   |
+--------------------+
```

If no tasks have been added:

```console
$ tasklist all
```

```text
+--------------------+
| TASKLIST MANAGER   |
+--------------------+
|    No Tasks Added  |
+--------------------+
```

---

### Add a Task

```text
tasklist add <TASK>
```

Adds a new task and assigns it a unique ID.

Example:

```console
$ tasklist add "Learn Go"
```

```text
+------------------+
| TASKLIST MANAGER |
+------------------+
|  Task Added - 1  |
+------------------+
```

If the task already exists:

```console
$ tasklist add "Learn Go"
```

```text
+------------------+
| TASKLIST MANAGER |
+------------------+
|  Task Exists - 1 |
+------------------+
```

---

### Remove a Task

```text
tasklist remove <TASK_ID>
```

Removes a task using its unique ID.

Example:

```console
$ tasklist remove 1
```

```text
+------------------+
| TASKLIST MANAGER |
+------------------+
|   Task Deleted   |
+------------------+
```

If the specified task does not exist:

```console
$ tasklist remove 99
```

```text
+------------------+
| TASKLIST MANAGER |
+------------------+
|  No Task Found!  |
+------------------+
```

> **Future consideration:** `delete` may be introduced as an alias for `remove` in a future iteration.

---

## Project Goals
The primary goal of Week 1 is to develop familiarity with:

* Go project structure
* Command-line arguments
* File I/O
* JSON encoding and decoding
* Structs and methods
* Error handling
* Unit testing
