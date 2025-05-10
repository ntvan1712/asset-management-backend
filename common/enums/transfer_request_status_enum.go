package enums

type transferRequestEnum struct {
	Pending  string
	Approved string
	Rejected string
}

var TransferRequestEnum = transferRequestEnum{
	Pending:  "pending",
	Approved: "approved",
	Rejected: "rejected",
}
