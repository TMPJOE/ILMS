# ILMS - Task Manager

A desktop Task Manager application built with **Wails** (Go backend + Vue.js frontend). ILMS helps you organize and track tasks with a clean, modern interface and SQLite persistence.

---

## Features

- **Full CRUD Operations** — Create, read, update, and delete tasks
- **Task Status Tracking** — Four workflow statuses: *To Do*, *In Progress*, *Pending Review*, *Done*
- **Color-Coded Status Borders** — Visual feedback on task cards based on status
- **Cursor-Based Pagination** — Efficient loading of 10 tasks per page with Previous/Next navigation
- **Client-Side Form Validation** — Required fields with character limits (name: 4–50, description: 4–200)
- **Modal Dialogs** — Add/Edit forms in a modal with blurred backdrop
- **Dot-Menu Actions** — Edit and delete tasks from a per-card dropdown menu
- **Dark Theme** — Dark UI with teal accent (`#008e7f`)
- **Structured Logging** — JSON-formatted logs written to `logs/tasks.log`

---

## Tech Stack

### Backend (Go)

| Layer | Technology |
|-------|-----------|
| Framework | [Wails v2](https://wails.io/) |
| Database | SQLite (`go-sqlite3`) |
| Logging | `log/slog` (structured JSON) |
| Architecture | Repository + Service pattern |

### Frontend (Vue.js)

| Library | Purpose |
|---------|---------|
| [Vue 3](https://vuejs.org/) | UI framework (Composition API, `<script setup>`) |
| [TypeScript](https://www.typescriptlang.org/) | Type safety |
| [Vite](https://vitejs.dev/) | Build tool & dev server |
| [vee-validate](https://vee-validate.logaretm.com/) | Form validation |
| [yup](https://github.com/jquense/yup) | Schema validation |
| [@vueuse/core](https://vueuse.org/) | Composition utilities (`onClickOutside`) |

---

## Project Structure

```
ILMS/
├── main.go                     # Wails app entry point, dependency injection
├── App.go                      # Main app struct (startup, greeting)
├── wails.json                  # Wails configuration
├── go.mod / go.sum             # Go dependencies
├── package.json                # Frontend dependencies
│
├── backend/
│   ├── database/
│   │   └── db.go               # SQLite connection pool
│   ├── log/
│   │   └── logger.go           # Structured logger factory (JSON handler)
│   ├── models/
│   │   └── task.go             # Task data models (Task, TaskInput, TaskOutput, etc.)
│   ├── repositories/
│   │   └── task.go             # TaskRepo — CRUD SQL operations
│   ├── services/
│   │   └── task.go             # TaskService — business logic layer
│   └── sql/
│       ├── dbMigration.sql     # Table schema (task table)
│       └── cleanup.sql         # Cleanup query (DELETE FROM task)
│
├── frontend/
│   ├── src/
│   │   ├── main.ts             # Vue entry point
│   │   ├── App.vue             # Root component (title + task list)
│   │   ├── style.css           # Global styles
│   │   └── components/
│   │       ├── ListOfTasks.vue # Task list with pagination & modal
│   │       ├── Task.vue        # Individual task card with status dropdown & menu
│   │       └── validatedForm.vue  # Add/Edit form with vee-validate + yup
│   └── wailsjs/                # Wails-generated Go→JS bindings
│
└── logs/
    └── tasks.log               # Application log file (auto-created)
```

---

## Architecture

### Backend Layers

```
┌─────────────────────────────────────────────┐
│  main.go                                    │
│  - Initializes logger, database, services   │
│  - Binds TaskService to Wails runtime       │
├─────────────────────────────────────────────┤
│  TaskService (services/task.go)             │
│  - Business logic, input validation         │
│  - Logging of all operations                │
├─────────────────────────────────────────────┤
│  TaskRepo (repositories/task.go)            │
│  - Direct SQL queries (Create, Select,      │
│    Update, Delete)                          │
├─────────────────────────────────────────────┤
│  SQLite (backend/database/tasks.db)         │
│  - Persistent storage                       │
└─────────────────────────────────────────────┘
```

### Frontend Components

```
┌─────────────────────────────────────────────┐
│  App.vue                                    │
│  - Title: "ILMS - Task Manager"             │
│  - Renders ListOfTasks                      │
├─────────────────────────────────────────────┤
│  ListOfTasks.vue                            │
│  - Fetches tasks via Wails TaskService      │
│  - Cursor-based pagination (10 per page)    │
│  - Opens ValidatedForm modal for add/edit   │
├─────────────────────────────────────────────┤
│  Task.vue                                   │
│  - Displays task name, description, date    │
│  - Status dropdown (4 statuses)             │
│  - Dot-menu: Edit, Delete                   │
│  - Color-coded borders by status            │
├─────────────────────────────────────────────┤
│  ValidatedForm.vue                          │
│  - vee-validate + yup validation            │
│  - Add mode / Edit mode                     │
│  - Emits task-added / task-updated events   │
└─────────────────────────────────────────────┘
```

---

## Data Models

| Model | Fields |
|-------|--------|
| `Task` | `id`, `status`, `name`, `desc`, `date` |
| `TaskInput` | `name`, `desc` |
| `TaskOutput` | `id`, `status`, `name`, `desc`, `date` (string) |
| `TaskResponse` | `Tasks[]`, `LastId`, `FirstPage` |
| `TaskUpdate` | `id`, `status`, `name`, `desc` |
| `TaskSwapBack` | `name`, `desc` |

### Task Statuses

| Value | Status |
|-------|--------|
| 0 | To Do |
| 1 | In Progress |
| 2 | Pending Review |
| 3 | Done |

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.18+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
- [Node.js](https://nodejs.org/) 16+
- [GCC](https://gcc.gnu.org/) (required for `go-sqlite3` CGO compilation)

### Installation

```bash
# Clone the repository
git clone <repo-url>
cd ILMS

# Install frontend dependencies
npm install

# (Optional) Install Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

---

## Live Development

To run in live development mode, run:

```bash
wails dev
```

This starts a Vite development server with hot reload. A dev server is also available at `http://localhost:34115` for browser-based development with access to Go methods via devtools.

---

## Building

To build a redistributable, production-mode binary:

```bash
wails build
```

Output will be placed in the `build/` directory.

### Platform-Specific Builds

- **Windows:** `wails build -platform windows/amd64`
- **macOS:** `wails build -platform darwin/amd64`
- **Linux:** `wails build -platform linux/amd64`

---

## Logging

Application logs are written to `logs/tasks.log` in JSON format with source location tracking. The log file is created automatically on app startup.

Example log entry:
```json
{
  "time": "2026-04-02T10:00:00Z",
  "level": "INFO",
  "source": {"file": "services/task.go", "line": 42},
  "msg": "Task Added",
  "Name": "My Task"
}
```

---

## Database

The application uses **SQLite** for persistent storage. The database file is located at `backend/database/tasks.db`.

### Schema

```sql
CREATE TABLE IF NOT EXISTS "task" (
    id INTEGER PRIMARY KEY,
    status INT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Author

**tmpjoe** — [josedegracia1212@gmail.com](mailto:josedegracia1212@gmail.com)

---

## License

See `LICENSE` file if applicable.
