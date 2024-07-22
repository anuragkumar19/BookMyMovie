package auth

const otpLength = 6

func generateOTP() (rawOTP string, hash string, err error) {
	return "", "", nil
}

func generateRandomToken() (string, error) {
	return "", nil
}

func matchOTP(rawOTP string, hash string) bool {
	return rawOTP == hash
}
