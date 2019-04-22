type error interface {
  Error() string
}

func (b *Reader) ReadByte() (c byte, err error)
