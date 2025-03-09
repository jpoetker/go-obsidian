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

// Create a client with responses (recommended)
client, err := obsidian.NewAuthenticedClientWithResponses("http://localhost:27124", "your-api-key")
if err != nil {
    log.Fatal(err)
}
```

### Checking Server Status

```go
ctx := context.Background()

// Get server status and authentication status
response, err := client.GetWithResponse(ctx)
if err != nil {
    log.Fatal(err)
}

if response.JSON200 != nil {
    fmt.Printf("Service: %s\n", *response.JSON200.Service)
    fmt.Printf("Authenticated: %v\n", *response.JSON200.Authenticated)
    if response.JSON200.Versions != nil {
        fmt.Printf("Obsidian Version: %s\n", *response.JSON200.Versions.Obsidian)
    }
}
```

### Working with Vault Files

```go
ctx := context.Background()

// List files in a directory
response, err := client.GetVaultPathToDirectoryWithResponse(ctx, "path/to/directory")
if err != nil {
    log.Fatal(err)
}
if response.JSON200 != nil {
    for _, file := range *response.JSON200.Files {
        fmt.Println(file)
    }
}

// Get file content
fileResponse, err := client.GetVaultFilenameWithResponse(ctx, "path/to/note.md")
if err != nil {
    log.Fatal(err)
}

// Create or update a file
content := strings.NewReader("# My New Note\n\nContent here")
_, err = client.PutVaultFilenameWithBodyWithResponse(ctx, "new-note.md", "text/markdown", content)
if err != nil {
    log.Fatal(err)
}

// Delete a file
_, err = client.DeleteVaultFilenameWithResponse(ctx, "note-to-delete.md")
if err != nil {
    log.Fatal(err)
}
```

### Working with Active Note

```go
ctx := context.Background()

// Get active note
activeResponse, err := client.GetActiveWithResponse(ctx)
if err != nil {
    log.Fatal(err)
}

// Update active note
content := strings.NewReader("# Updated Content")
_, err = client.PutActiveWithBodyWithResponse(ctx, "text/markdown", content)
if err != nil {
    log.Fatal(err)
}
```

### Working with Commands

```go
ctx := context.Background()

// List available commands
commandsResponse, err := client.GetCommandsWithResponse(ctx)
if err != nil {
    log.Fatal(err)
}
if commandsResponse.JSON200 != nil && commandsResponse.JSON200.Commands != nil {
    for _, cmd := range *commandsResponse.JSON200.Commands {
        fmt.Printf("Command: %s (ID: %s)\n", *cmd.Name, *cmd.Id)
    }
}

// Execute a command
_, err = client.PostCommandsCommandIdWithResponse(ctx, "graph:open")
if err != nil {
    log.Fatal(err)
}
```

## Features

- Full support for Obsidian's Local REST API
- Type-safe API with proper error handling
- Context support for cancellation and timeouts
- Vault file operations (create, read, update, delete)
- Active note management
- Command execution
- Configurable client with custom HTTP client and base URL options

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 