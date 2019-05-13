package domain

type Conf struct {
	Username string
	Buttons  []Button
}

type Button struct {
	Text  string
	Value string
}
