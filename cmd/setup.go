package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/justwinstuff/hakarutrade/internal"
)

const mt5URL = "https://download.terminal.free/cdn/web/metaquotes.ltd/mt5/mt5setup.exe"
var ids [2]string

func createConfig() {	
	var config internal.Config

	config.Accounts.Source = ids[0]
	config.Accounts.Destinations = ids[1:]
	
	data, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	
	_, err = out.Write(data)
	if err != nil {
		panic(err)
	}
}

func generateId() string {
	return uuid.NewString()
}

func downloadMT5() {
	fmt.Println("Downloading MetaTrader 5...")

	res, err := http.Get(mt5URL)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(fmt.Errorf("MetaTrader 5 download failed: %v", res.Status))
	}
	defer res.Body.Close()

	insPath := filepath.Join(os.TempDir(), "hakarutrade-mt5installer.exe")
	insOut, err := os.Create(insPath)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(insOut, res.Body)
	if err != nil {
		panic(err)
	}
	insOut.Close()

	for i := range ids {
		id := generateId()
		dir := filepath.Join("terminals", id)

		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Installing MetaTrader 5 for account: %v\n", id)
		cmd := exec.Command(insPath, "/auto", "/path:"+dir)
		err = cmd.Run()
		if err != nil {
			panic(err)
		}

		ids[i] = id
	}
}

func Setup() {
	downloadMT5()
	createConfig()
}