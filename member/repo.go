package member

type MemberRepo interface {
	SetStatus(int) error
}

type MySQL struct {

}
