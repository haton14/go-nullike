package nullike

type NullString struct {
	s *string
}

func (s *NullString) Valid() bool {
	return s.s != nil

}
