package model

type UserType int

const (
	StudentUser UserType = iota
	AssistantUser
	TeacherUser
	AdminUser
)

func (ut UserType) String() string {
	return [...]string{"student", "assistant", "teacher", "admin"}[ut]
}
