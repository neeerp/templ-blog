package handlers

// TODO: Naming doesn't quite look right when importing
func Get() []Todo {
	return []Todo{
		{"Go outside"},
		{"Touch Grass"},
		{"Eat"},
		{"Check Sanity"},
	}
}

type Todo struct {
	Title string
}
