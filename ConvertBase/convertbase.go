package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	intnbr := AtoiBase(nbr, baseFrom)
	s := PrintNbrBase(intnbr, baseTo)
	return s
}
