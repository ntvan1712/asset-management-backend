package enums

type borrowRequestEnum struct {
	Pending  string
	Approved string
	Rejected string
}

var BorrowRequestEnum = borrowRequestEnum{
	Pending:  "pending",
	Approved: "approved",
	Rejected: "rejected",
}
