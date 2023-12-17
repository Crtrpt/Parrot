package parrot

func Execute() (err error) {
	return main()
}
func main() (err error) {
	ConfigInit()

	StartWebServer()
	return nil
}
