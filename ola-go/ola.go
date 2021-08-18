package main

const prefixHey = "Hey "
const prefixEspanol = "Holla "
const prefixPortugues = "Olá "

func Ola(name string, idiom string) string {
	if name == ""{
		name = "you"
	}

	prefix := prefixHey

	switch idiom {
	case "espanol":
		prefix = prefixEspanol
	case "portugues":
		prefix = prefixPortugues
	}
	
	return prefix + name
}
