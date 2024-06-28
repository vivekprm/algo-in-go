package module1

func BaseToBase(value string, base, newBase int) string {
	des := BaseToDec(value, base)
	return DecToBase(des, newBase)
}