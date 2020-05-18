package auth

// UserAuth .
type UserAuth struct {
	ID  int64 `xorm:"pk  autoincr 'id'"`
	UID int64 `xorm:"'user_id'"`
}
