package main

type StoryLine struct {
	Intro     Intro     `json:"intro"`
	NewYork   NewYork   `json:"new-york"`
	Debate    Debate    `json:"debate"`
	SeanKelly SeanKelly `json:"sean-kelly"`
	MarkBates MarkBates `json:"mark-bates"`
	Denver    Denver    `json:"denver"`
	Home      Home      `json:"home"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Intro struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type NewYork struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Debate struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type SeanKelly struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type MarkBates struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Denver struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Home struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []any    `json:"options"`
}
