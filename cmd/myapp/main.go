package main

func main() {
	sc := InitializeServiceContext("configs/user.yaml")
	sc.Start()
}
