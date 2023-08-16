package lol

func Lol(nick string) (profile interface{}) {
	profile = SendRequest(nick)
	return profile
}
