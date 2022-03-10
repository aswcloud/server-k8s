package main

import "time"

func main() {
	CreateNamespace()
	CreatePersistent()
	CreateDeployment()
	CreateService()
	time.Sleep(time.Second * 5)
	DeleteAll()

	// // 님은 방구 뿡뿡이~ 님은 방구 뿡뿡잉~~
}
