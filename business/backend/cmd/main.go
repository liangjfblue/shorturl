package main

func main() {
	server, cleanup, err := wireServer()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = server.Run()
	if err != nil {
		panic(err)
	}

	if err = server.Stop(); err != nil {
		panic(err)
	}
}
