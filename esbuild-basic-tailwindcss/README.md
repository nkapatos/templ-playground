## Features

- **Templ Components**: Type-safe HTML templating with Go
- **Live Reload**: Hot reloading for Go code, templates, and assets
- **ESBuild**: Fast JavaScript/TypeScript bundling with PostCSS
- **Tailwind CSS**: Utility-first CSS framework
- **Task Runner**: Taskfile for streamlined development workflow

## Tech Stack

- **Backend**: Go 1.24.2
- **Templating**: [Templ](https://templ.guide/) v0.3.898
- **Build Tool**: ESBuild v0.25.5
- **CSS**: Tailwind CSS v4.1.8
- **Task Runner**: Taskfile
- **Live Reload**: Air

## Project Structure

```
esbuild-basic-tailwindcss/
├── assets/
│   ├── css/
│   │   └── index.css          # Tailwind CSS with Templ source scanning
│   └── ts/
│       └── main.ts            # TypeScript entry point
├── templ/
│   ├── components/            # Reusable Templ components
│   ├── layouts/              # Page layouts
│   └── views/                # Page views
├── build.js                  # ESBuild configuration
├── main.go                   # Go server entry point
├── taskfile.yaml             # Task definitions
└── package.json              # Node.js dependencies
```

## Prerequisites

- Go 1.24.2+
- Node.js (for ESBuild and Tailwind)
- pnpm (package manager)

## Installation

1. Clone the repository
2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Install Node.js dependencies:
   ```bash
   pnpm install
   ```

## Development

### Start Live Development

Run all live reload services simultaneously:

```bash
task live
```

This command starts:
- **Templ watcher**: Generates Go files from `.templ` files
- **ESBuild watcher**: Bundles TypeScript/CSS with Tailwind
- **Go server**: Runs the application with Air for live reload

### Individual Commands

- **Live reload Go app**: `task live:app`
- **Live reload templates**: `task live:templ`
- **Live reload assets**: `task live:esbuild`
- **Format Templ files**: `task format:templ`

### Manual Build

Build assets manually:
```bash
node build.js
```

Run Go server:
```bash
go run main.go
```

## Configuration

### ESBuild (build.js)
- Entry points: `assets/ts/main.ts` and `assets/css/index.css`
- Output directory: `dist/`
- PostCSS plugin for Tailwind CSS processing
- Source maps enabled
- Watch mode for development

### Tailwind CSS
- Configured to scan Templ files for classes
- Uses Tailwind CSS v4 with PostCSS integration
- Source scanning includes `./templ/**/*.{templ}`

### Air (Live Reload)
- Proxy enabled (app: 8080, proxy: 8081)
- Watches Go files, CSS, and JS
- Excludes `node_modules`, `assets`, `tmp`, `vendor`
- Clean exit enabled

## Development Workflow

1. **Templates**: Edit `.templ` files in the `templ/` directory
2. **Styling**: Modify `assets/css/index.css` with Tailwind classes
3. **Logic**: Add TypeScript in `assets/ts/main.ts`
4. **Backend**: Update Go code in `main.go` and other `.go` files

All changes are automatically detected and the application reloads in real-time.

## Ports

- **Application**: 8080
- **Proxy**: 8081 (for live reload)
