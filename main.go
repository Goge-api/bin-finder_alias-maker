package main

func main() {

	// fmt.Println("Do you want to autorun the finder for bitcoin? [Y,y]")
	// var auto string
	// fmt.Scanln(&auto)
	// // if auto == "y" || auto == "Y" {
	// WalkAndFind(GetRoot(), "bitcoind.exe")
	for i := 0; i < len(BinLogRead()); i++ {

		AliasWritter(BinLogRead()[i].fullAlias)
	}

	// } else {

	// fmt.Println("Enter the file name you're looking for.")

	// var fileName string
	// fmt.Scanln(&fileName)

	// WalkAndFind(GetRoot(), fileName)
	// alias doged='/mnt/volume_sfo3_02/dogecoin-1.14.2/bin/./dogecoind
	// }

	println("thing done")
}
