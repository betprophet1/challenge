package database

type Table interface {
	Fields() []string
}

func Fields(table Table) []string {
	return table.Fields()
}
