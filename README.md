# Split

Track spending/splitting when cost-sharing between 1 or more people.
Acts as a replacement for apps like Splitwise.

<img src="https://private-user-images.githubusercontent.com/23506992/419607669-ed2c1c81-053d-46b3-9321-b7f5b21f96c1.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NDEyMDI0NDMsIm5iZiI6MTc0MTIwMjE0MywicGF0aCI6Ii8yMzUwNjk5Mi80MTk2MDc2NjktZWQyYzFjODEtMDUzZC00NmIzLTkzMjEtYjdmNWIyMWY5NmMxLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTAzMDUlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUwMzA1VDE5MTU0M1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTQ3YTE1M2QyYzQwMDZkMDY0NjAxOWViYWQ1MGEwYTA4Yjg3ZDU5ZDlkMGQ1YmMzZGQ4YTEzMzgwMDRmNzdjNjAmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.8NqbH8ndR7dCaN5TWZqZC3DfQFIM2To7XWzT8nw_lPA" alt="image" height="600">

## Features

- FX Rate support
  - Pulls FX rates on a schedule from [fxratesapi](https://fxratesapi.com/)
  - Configure what FX rates you want to track
  - Simply select what currency you're paying in and it will convert to your configured base currency
- Category management
  - Change icons (fontawesome)
- Scheduled expenses
  - Set up recurring expenses such as rent, internet, etc which will be automatically added on the configured schedule

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
