package common

// this pkg define the output DTO types

// server status
type ServerStatus struct {
	Version    string `json:"version"`
	Goroutines int    `json:"goroutines"`
	Workers    int    `json:"workers"`
	Queue      int64  `json:"queue"`
}
