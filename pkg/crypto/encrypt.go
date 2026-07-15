package crypto

// PswEnc 密码加密，采用bcrypt (pass-through: bcrypt is done on the frontend)
func PswEnc(psw string) string {
	return psw
}

// PswDec 密码解密 (pass-through: bcrypt is done on the frontend)
func PswDec(psw string) string {
	return psw
}
