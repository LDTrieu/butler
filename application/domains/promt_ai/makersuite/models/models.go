package models

type MakersuiteResponseApi struct {
	Candidates []Candidate `json:"candidates"`
	Error      struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Candidate struct {
	Output string `json:"output"`
}

type RequestBody struct {
	Prompt          TextRequest `json:"prompt" structs:"prompt"`
	SafetySettings  []Setting   `json:"safetySettings,omitempty" structs:"safetySettings,omitempty"`
	StopSequences   []string    `json:"stopSequences,omitempty" structs:"stopSequences,omitempty"`
	Temperature     float32     `json:"temperature,omitempty" structs:"temperature,omitempty"`
	CandidateCount  int         `json:"candidate_count,omitempty" structs:"candidate_count,omitempty"`
	MaxOutputTokens int64       `json:"maxOutputTokens,omitempty" structs:"maxOutputTokens,omitempty"`
	TopP            float32     `json:"topP,omitempty" structs:"topP,omitempty"`
	TopK            int         `json:"topK,omitempty" structs:"topK,omitempty"`
}

type Setting struct {
	Category  string `json:"category,omitempty" structs:"text,omitempty"`
	Threshold string `json:"threshold,omitempty" structs:"text,omitempty"`
}

type TextRequest struct {
	Text string `json:"text" structs:"text"`
}
