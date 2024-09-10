# Split

Track spending/splitting when cost-sharing between 1 or more people.
Acts as a replacement for apps like Splitwise.

## Stack

### Backend

- Golang 1.23 (no routing framework)
- gorm - Go orm with auto-migrations
- Templ - Templating engine which allows embedding Go into HTML

### Frontend

- Htmx - Gives all html tags ability to make HTTP requests and swap partials of the DOM
- Alpine.js - Small JS library for common DOM manipulation tactics written directly into HTML attributes
- Tailwind CSS - Styling by structured classes
- daisyUI - Tailwind component library

## Local Development

### Local Stack

- `air` is responsible for hot reloading the app on change
  - `.air.toml` configures what files to listen to, what command(s) to run, what the build target is, etc
  - For example, running tailwind collectors
- `dlv` is a golang debugger. It's configured to run on 2345, which is what neovim's DAP is configured to listen for.

### Run

- Install Go dependencies: `go mod tidy`
- Install Node dependencies: `npm install`
- Run dev server with `air`
- Open <http://localhost:8090>
