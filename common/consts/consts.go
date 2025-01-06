package consts

const (
	// CurrentBlockNumKey The cache key for the block number being polled
	CurrentBlockNumKey = "current_block_num"
)

const (
	Not_audit = iota
	Audit_has_been_passed
	Audit_failed
)
