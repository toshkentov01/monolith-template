package models

// Response ...
type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Error ...
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//SuccessMessage return boolean about success or not
type SuccessMessage struct {
	Success bool `json:"success"`
}
