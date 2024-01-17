package main

type StoryLine map[string]Section

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Section struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
