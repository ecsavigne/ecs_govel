package logg

// func Log(Function string, CompanyPhone string, ContactPhone string, Log string, Error bool) {
// 	// 0. print to stdout
// 	if Error {
// 		err := errors.New(Log)
// 		ErrorLogger.Printf("\033[31m %v \033[0m\n", err)
// 	} else {
// 		GeneralLogger.Printf("\033[36m %v \033[0m\n", Log)
// 	}

// 	// 1. send to datadog

// 	// 2. print to ErrorLogger
// 	if Error {
// 		go ErrorLogger.Println("\033[31m"+Function+" ", Log, "\033[0m")
// 	} else {
// 		go GeneralLogger.Println("\033[36m" + Function + " " + Log + " \033[0m")
// 	}

// }
