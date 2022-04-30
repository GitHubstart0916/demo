package utils

// TODO:密码加密,解密,比较

func ComparePassword(pwd1 string, pwd2 string) bool {
	if pwd1 == pwd2 {
		return true
	} else {
		return false
	}
}
