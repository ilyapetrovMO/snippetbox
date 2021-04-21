package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	err := e[field]
	if len(err) == 0 {
		return ""
	}

	return err[0]
}
