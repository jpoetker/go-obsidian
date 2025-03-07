# go-obsidian

A Go client library for the Obsidian Local REST API. This library provides a convenient way to interact with Obsidian's Local REST API plugin.

## Installation

```bash
go get github.com/jpoetker/go-obsidian
```

## Usage

First, make sure you have the [Local REST API](https://github.com/coddingtonbear/obsidian-local-rest-api) plugin installed in Obsidian and have obtained an API key.

### Creating a Client

```go
import "github.com/jpoetker/go-obsidian"

client := obsidian.NewClient("your-api-key")

// Optionally, configure the client with options
client = obsidian.NewClient(
    "your-api-key",
    obsidian.WithBaseURL("https://custom-host:27124"),
    obsidian.WithHTTPClient(customHTTPClient),
)
```

### Checking Server Status

```go
ctx := context.Background()

// Get server status and authentication status
status, err := client.Status.GetStatus(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Server: %s\n", status.Service)
fmt.Printf("Authenticated: %v\n", status.Authenticated)
fmt.Printf("Obsidian Version: %s\n", status.Versions.Obsidian)
```

### Working with Vault Files

```go
ctx := context.Background()

// List files in a directory
files, err := client.Vault.ListDirectory(ctx, "path/to/directory")
if err != nil {
    log.Fatal(err)
}

// Get file content
note, err := client.Vault.GetFile(ctx, "path/to/note.md")
if err != nil {
    log.Fatal(err)
}

// Create or update a file
err = client.Vault.CreateOrUpdateFile(ctx, "new-note.md", "# My New Note\n\nContent here")
if err != nil {
    log.Fatal(err)
}

// Append to a file
err = client.Vault.AppendToFile(ctx, "existing-note.md", "\n\nAppended content")
if err != nil {
    log.Fatal(err)
}

// Delete a file
err = client.Vault.DeleteFile(ctx, "note-to-delete.md")
if err != nil {
    log.Fatal(err)
}
```

### Searching

```go
// Simple text search
results, err := client.Search.SimpleSearch(ctx, "search query", &obsidian.SimpleSearchOptions{
    ContextLength: 100,
})
if err != nil {
    log.Fatal(err)
}

// Advanced search using DQL
dqlResults, err := client.Search.Search(ctx, obsidian.SearchQuery{
    Query:     "TABLE time-played FROM #game SORT rating DESC",
    QueryType: "dql",
})
if err != nil {
    log.Fatal(err)
}

// Advanced search using JsonLogic
jsonLogicQuery := map[string]interface{}{
    "in": []interface{}{
        "myTag",
        map[string]interface{}{
            "var": "tags",
        },
    },
}
jsonLogicResults, err := client.Search.Search(ctx, obsidian.SearchQuery{
    Query:     jsonLogicQuery,
    QueryType: "jsonlogic",
})
if err != nil {
    log.Fatal(err)
}
```

### Working with Commands

```go
// List available commands
commands, err := client.Commands.List(ctx)
if err != nil {
    log.Fatal(err)
}

// Execute a command
err = client.Commands.Execute(ctx, "graph:open")
if err != nil {
    log.Fatal(err)
}
```

## Features

- Full support for Obsidian's Local REST API
- Vault file operations (create, read, update, delete, append)
- Simple and advanced search capabilities
- Command execution
- Context support for cancellation and timeouts
- Configurable client with custom HTTP client and base URL
- Type-safe API with proper error handling

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 