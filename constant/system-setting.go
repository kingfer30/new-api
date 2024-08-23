package constant

var ServerAddress = "http://localhost:3000"
var WorkerUrl = ""
var WorkerValidKey = ""
var HttpProxy = ""

func EnableWorker() bool {
	return WorkerUrl != ""
}
