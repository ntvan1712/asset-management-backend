package enums

type sortOrderEnum struct {
	Asc  string
	Desc string
}

var SortOrderEnum = sortOrderEnum{
	Asc:  "asc",
	Desc: "desc",
}
