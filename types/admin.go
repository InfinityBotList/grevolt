package types

import "github.com/infinitybotlist/grevolt/types/timestamp"

// Server Stats
//
// <these are technical admin stats>
type Stats struct {
	// Index usage information
	Indices map[string][]*Index `json:"indices"`
	// Collection stats
	CollStats map[string]*CollectionStats `json:"coll_stats"`
}

// Collection index
type Index struct {
	// Index name
	Name string `json:"name"`
	// Access information
	Accesses *IndexAccesses `json:"accesses"`
}

// Access information
type IndexAccesses struct {
	// Operations since timestamp
	Ops uint64 `json:"ops"`
	// Timestamp at which data keeping began
	Since timestamp.Timestamp `json:"since"`
}

// Collection stats
type CollectionStats struct {
	// Namespace
	Ns string `json:"ns"`
	// Local time
	LocalTime timestamp.Timestamp `json:"localTime"`
	// Latency stats
	LatencyStats map[string]*LatencyStats `json:"latencyStats"`
	// Query exec stats
	QueryExecStats *CollectionStatsQueryExecStats `json:"queryExecStats"`
	// Number of documents in collection
	Count uint64 `json:"count"`
}

// Collection latency stats
type LatencyStats struct {
	// Total operations
	Ops uint64 `json:"ops"`
	// Timestamp at which data keeping begun
	Latency uint64 `json:"latency"`
	// Histogram representation of latency data
	Histogram []*LatencyHistogramEntry `json:"histogram"`
}

// Histogram entry
type LatencyHistogramEntry struct {
	// Time
	Micros uint64 `json:"micros"`
	// Count
	Count uint64 `json:"count"`
}

// Query exec stats
type CollectionStatsQueryExecStats struct {
	// Stats regarding collection scans
	CollectionScans *CollectionScans `json:"collectionScans"`
}

// Query collection scan stats
type CollectionScans struct {
	// Number of total collection scans
	Total uint64 `json:"total"`
	// Number of total collection scans not using a tailable cursor
	NonTailable uint64 `json:"nonTailable"`
}
