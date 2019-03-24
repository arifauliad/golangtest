package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Museumdata struct {
	Data []struct {
		MuseumID          string `json:"museum_id"`
		KodePengelolaan   string `json:"kode_pengelolaan"`
		Nama              string `json:"nama"`
		Sdm               string `json:"sdm"`
		AlamatJalan       string `json:"alamat_jalan"`
		DesaKelurahan     string `json:"desa_kelurahan"`
		Kecamatan         string `json:"kecamatan"`
		KabupatenKota     string `json:"kabupaten_kota"`
		Propinsi          string `json:"propinsi"`
		Lintang           string `json:"lintang"`
		Bujur             string `json:"bujur"`
		Koleksi           string `json:"koleksi"`
		SumberDana        string `json:"sumber_dana"`
		Pengelola         string `json:"pengelola"`
		Tipe              string `json:"tipe"`
		Standar           string `json:"standar"`
		TahunBerdiri      string `json:"tahun_berdiri"`
		Bangunan          string `json:"bangunan"`
		LuasTanah         string `json:"luas_tanah"`
		StatusKepemilikan string `json:"status_kepemilikan"`
	} `json:"data"`
}

func getdata(finished chan bool) {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?"

	resp, err := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	trimbody := bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	bodyString := string(trimbody)

	if err != nil {
		fmt.Print(err)
	}

	file, writeerr := os.Create("temp.json")
	if writeerr != nil {
		fmt.Print(writeerr)
	}
	file.WriteString(bodyString)
	file.Close()
	finished <- true
}

func inputdata(data []string, filename string) {
	f, err := os.OpenFile((filename + ".csv"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()
	defer f.Close()
}

func main() {
	finished := make(chan bool)
	fmt.Println("masih bekrja")
	go getdata(finished)
	fmt.Println("Menunggu")
	<-finished

	b, error := ioutil.ReadFile("temp.json")
	if error != nil {
		fmt.Print(error)
	}
	data := string(b)

	dat := Museumdata{}
	json.Unmarshal([]byte(data), &dat)
	for _, obj := range dat.Data {
		var record []string
		record = append(record, obj.MuseumID)
		record = append(record, obj.KodePengelolaan)
		record = append(record, obj.Nama)
		record = append(record, obj.Sdm)
		record = append(record, obj.AlamatJalan)
		record = append(record, obj.DesaKelurahan)
		record = append(record, obj.Kecamatan)
		record = append(record, obj.KabupatenKota)
		record = append(record, obj.Propinsi)
		record = append(record, obj.Lintang)
		record = append(record, obj.Bujur)
		record = append(record, obj.Koleksi)
		record = append(record, obj.SumberDana)
		record = append(record, obj.Pengelola)
		record = append(record, obj.Tipe)
		record = append(record, obj.Standar)
		record = append(record, obj.TahunBerdiri)
		record = append(record, obj.Bangunan)
		record = append(record, obj.LuasTanah)
		record = append(record, obj.StatusKepemilikan)
		inputdata(record, obj.KabupatenKota)
	}

	fmt.Println("selesai")
}
