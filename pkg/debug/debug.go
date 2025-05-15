package debug

import (
	"log"
)

func Info(msg string) {
	//if os.Getenv("APP_DEBUG") == "true" {
	log.Println(msg)
	//}
}

func Log(title string, msg ...interface{}) {
	//if os.Getenv("APP_DEBUG") == "true" {
	if len(msg) == 0 {
		Info(title)
	}

	for i := 0; i < len(msg); i++ {
		log.Printf("%s:%s\n", title, msg[0])
	}
	//}
}

func Error(title string, err error) {
	//if os.Getenv("APP_DEBUG") == "true" {
	log.Printf("%s:%s", title, err.Error())
	//}
}

func HasError(title string, err error) bool {
	if err != nil {
		//if os.Getenv("APP_DEBUG") == "true" {
		log.Printf("%s:%s", title, err.Error())
		//}
		return true
	}
	return false
}

func Panic(name string, err error) {
	panic(name + ":" + err.Error())
}
