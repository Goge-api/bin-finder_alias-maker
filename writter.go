package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func PathMaker(ssp, s string) (pathSub string) {
	pathSub = ssp + s
	return
}

func FileWritter(ff string, path string) {

	f, err := os.OpenFile("binsFound.log",
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	//make a rune of the path to make a substring of the basefile of file found
	runes := []rune(path)
	subString := string(runes[0 : len(path)-len(filepath.Base(path))])

	if _, err := f.WriteString("bitcoind.exe" + ", " + PathMaker(subString, "bitcoind.exe") + ", btcd, " + AliasMaker(subString+"bitcoind.exe", "btcd") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("bitcoin-cli.exe" + ", " + PathMaker(subString, "bitcoin-cli.exe") + ", btcli, " + AliasMaker(subString+"bitcoin-cli.exe", "btcli") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("bitcoin-tx.exe" + ", " + PathMaker(subString, "bitcoin-tx.exe") + ", btctx, " + AliasMaker(subString+"bitcoin-tx.exe", "btctx") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("bitcoin-wallet.exe" + ", " + PathMaker(subString, "bitcoin-wallet.exe") + ", btcwall, " + AliasMaker(subString+"bitcoin-wallet.exe", "btcwall") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("bitcoin-qt.exe" + ", " + PathMaker(subString, "bitcoin-qt.exe") + ", btcqt, " + AliasMaker(subString+"bitcoin-qt.exe", "btcqt") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("test_bitcoin.exe" + ", " + PathMaker(subString, "test_bitcoin.exe") + ", tbtc, " + AliasMaker(subString+"test_bitcoin.exe", "tbtc")); err != nil {
		log.Println(err)
	}

	// fmt.Println("File Writter done")
}

func FileWritterAppend(ff string, path string) {

	f, err := os.OpenFile("binsFound.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	runes := []rune(path)
	// ... Convert back into a string from rune slice.
	subString := string(runes[0 : len(path)-len(filepath.Base(path))])

	print(subString)

	if _, err := f.WriteString("bitcoind.exe" + ", " + AliasMaker(subString+"bitcoind.exe", "btcd") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("bitcoin-cli.exe" + ", " + AliasMaker(subString+"bitcoin-cli.exe", "btcli") + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("test_bitcoin.exe" + ", " + AliasMaker(subString+"test_bitcoin.exe", "test_btc")); err != nil {
		log.Println(err)
	}

	fmt.Println("File Writter done")
}
