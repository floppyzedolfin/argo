package load

// Request is the structure of a load request, specifying which local file to ingest
type Request struct {
	Input string `json:"input"`
}

// Response is the structure of a load response - it's empty, as there's nothing much to say
type Response struct{}
