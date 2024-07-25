package main

import (
	"fmt"
	"strconv"
)

type task struct {
	id    int
	nama  string
	selesai bool
}

var tasks [] task
var nextId int

func main() {
	var command string
	for {
		fmt.Println("masukan perintah(add, list, done, remove, exit,) :")
		fmt.Scan(&command)
		switch command {
		case "1.add":
			//kode untuk menambahkan tugas
			var nama string
			fmt.Println("daftar tugas: ")
			fmt.Scan(&nama)
			tambahTugas(nama)
		case "2.list":
			lihattugas()
		case "3.done":
			var id_tugas string
			fmt.Println("nasukan id tugas :")
			fmt.Scan(&id_tugas)
			id ,_:=strconv.Atoi(id_tugas)
			tugasSelesai(id)
		case "4.remove":
			var id_tugas string
			fmt.Println("masukan id tugas :")
			fmt.Scan(&id_tugas)
			id ,_:=strconv.Atoi(id_tugas)
			hapusTugas(id)
		case "5.exit":
			return
		default:
			fmt.Println("perintah tidak dikenal")


		}
	}
}
func tambahTugas(nama string) {
	tasks=append(tasks, task{
		id: nextId,
		nama: nama,

	})
	nextId++
}

func lihattugas() {
	if len(tasks)==0{
		fmt.Println("tidak ada tugas")
	}
	for _, task:=range tasks{
		status:="belum selesai"
		if task.selesai{
			status="selesai"
		}
		fmt.Printf("%d. %s. [%s.]\n",task.id,task.nama,status)
	}
}
func tugasSelesai (id int) {
	for i:= range tasks {
		if tasks[i].id==id {
			tasks[i].selesai=true
			return
		}
	}
	fmt.Println("tugas dengan id tersebut tidak ditemukan")
	
}
func hapusTugas(id int) {
	for i:= range tasks {
		if tasks[i].id==id {
			tasks=append(tasks[:i],tasks[i+1:]... )
			return
		}
	}
	fmt.Println("tugas dengan id tersebut tidak ditemukan")
}