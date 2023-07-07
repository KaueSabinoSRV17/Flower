package catch

import "log"

func HandleError(message string, err error) {
	if err != nil {
		log.Fatalf("%v\n\t%v", message, err)
	}
}
