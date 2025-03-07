package obsidian

// Error represents an API error response
type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

// NoteJson represents a note in the vault with its metadata
type NoteJson struct {
	Content     string                 `json:"content"`
	Frontmatter map[string]interface{} `json:"frontmatter"`
	Path        string                 `json:"path"`
	Stat        NoteStat               `json:"stat"`
	Tags        []string               `json:"tags"`
}

// NoteStat represents file statistics for a note
type NoteStat struct {
	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
	Size  int64 `json:"size"`
}

// ServerInfo represents the server status information
type ServerInfo struct {
	Authenticated bool     `json:"authenticated"`
	OK            string   `json:"ok"`
	Service       string   `json:"service"`
	Versions      Versions `json:"versions"`
}

// Versions represents version information for Obsidian and the plugin
type Versions struct {
	Obsidian string `json:"obsidian"`
	Self     string `json:"self"`
}

// SearchResult represents a search result from the API
type SearchResult struct {
	Filename string      `json:"filename"`
	Result   interface{} `json:"result"`
}

// SimpleSearchResult represents a result from the simple search endpoint
type SimpleSearchResult struct {
	Filename string        `json:"filename"`
	Matches  []SearchMatch `json:"matches"`
	Score    float64       `json:"score"`
}

// SearchMatch represents a match in a simple search result
type SearchMatch struct {
	Context string          `json:"context"`
	Match   SearchMatchSpan `json:"match"`
}

// SearchMatchSpan represents the location of a match in the text
type SearchMatchSpan struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// DirectoryListing represents a list of files in a directory
type DirectoryListing struct {
	Files []string `json:"files"`
}
