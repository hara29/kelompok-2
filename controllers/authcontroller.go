package controllers

import (
	"Kelompok2/config"
	"Kelompok2/entities"
	"Kelompok2/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

//membuat struct untuk menampung data yang diterima dari form input login
type UserInput struct{
	Username string
	Password string
}
//Membuat struct untuk menampung data yang diterima dari form pilih pertemuan
type MingguPertemuan struct{
	Pert string
}

var userModel = models.NewUserModel()
var mhsModel = models.NewMhsModel()
//handler, fungsi untuk penanganan request ke root yang ditentukan
func Index(w http.ResponseWriter, r *http.Request){
	session, _:= config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}else{
		if session.Values["loggedIn"] != true{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}else{
			data:= map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}
			t, _ := template.ParseFiles("views/index.html")
			t.Execute(w, data)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		t, err := template.ParseFiles("views/login.html")
		if err != nil{
			panic(err)
		}
		t.Execute(w, nil)
	} else if r.Method == http.MethodPost{
		//proses login
		r.ParseForm()
		//membuat struct berdasarkan inputan di html
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		//check username ke database
		var user entities.User
		userModel.Where(&user, "id_user", UserInput.Username)

		var message error
		if user.Username == ""{
			//username tidak ada di database
			message = errors.New("username atau password salah")
		}else{
			//Pengecekan kecocokan password ke database, password dienskripsi
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil{
				//password salah!
				message = errors.New("username atau password salah")
				fmt.Println("Password salah")
			}
			fmt.Println("Password benar")
		}
		if message != nil{
			//gagal login
			data := map[string]interface{}{
				"error": message,
			}

			t, _ := template.ParseFiles("views/login.html")
			t.Execute(w, data)
		}else{
			//set session
			session, _:= config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["username"] = user.Username
			session.Values["nama_lengkap"] = user.NamaLengkap

			session.Save(r, w)

			http.Redirect(w, r, "/nilai-mahasiswa", http.StatusSeeOther)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request){
	session, _:= config.Store.Get(r, config.SESSION_ID)

	//delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Input(w http.ResponseWriter, r *http.Request){
	session, _:= config.Store.Get(r, config.SESSION_ID)
	//mengambil data pertemuan ke berapa dari form
	r.ParseForm()
	MingguPertemuan := &MingguPertemuan{
		Pert :r.Form.Get("pert"),
	}
	if r.Method == http.MethodGet{
		fmt.Println(MingguPertemuan.Pert)
		//menampilkan tabel untuk input nilai mahasiswa
		var mhs entities.Mahasiswa
		mahasiswa, _ := mhsModel.FindAll(&mhs, "pert", MingguPertemuan.Pert)
		data:= map[string]interface{}{
			"nama_lengkap": session.Values["nama_lengkap"],
			"Pert" : MingguPertemuan.Pert,
			"mahasiswa": mahasiswa, 
		}
		t, _ := template.ParseFiles("views/index.html")
		t.Execute(w, data)

	}else if r.Method == http.MethodPost{
		r.ParseForm()
		//mengambil inputan uts dan uas dari form
		uts := make([]string, 6,7)
		uas := make([]string, 6,7)
		uts[1]= r.Form.Get("uts[1]")
		uas[1]= r.Form.Get("uas[1]")
		uts[2]= r.Form.Get("uts[2]")
		uas[2]= r.Form.Get("uas[2]")
		uts[3]= r.Form.Get("uts[3]")
		uas[3]= r.Form.Get("uas[3]")
		uts[4]= r.Form.Get("uts[4]")
		uas[4]= r.Form.Get("uas[4]")
		uts[5]= r.Form.Get("uts[5]")
		uas[5]= r.Form.Get("uas[5]")
		utsf := make([]float64, 6,7)
		uasf := make([]float64, 6,7)
		for i:=0; i<5; i++{
			utsf[i+1],_ = strconv.ParseFloat(uts[i+1], 32)
			uasf[i+1],_ = strconv.ParseFloat(uas[i+1], 32)
		}
		total := make([]float64, 6,7)
		for i:=0; i<5; i++{
			total[i+1] = utsf[i+1]*0.7 + uasf[i+1]*0.3
		}
		fmt.Println(uts)
		fmt.Println(uas)
		fmt.Println(total)
		//memasukkan data ke database
		fmt.Println("pertemuan ke", MingguPertemuan)
		var mhs entities.Mahasiswa
		for i:=1; i<6; i++{
			a := mhsModel.Create(mhs,utsf[i], uasf[i], total[i], MingguPertemuan.Pert, i)
			if a == true{
				fmt.Println("data berhasil di input")
			}else{
				fmt.Println("data tidak berhasil diinput")
			}
		}
		mahasiswa, _ := mhsModel.FindAll(&mhs, "pert", MingguPertemuan.Pert)
		data:= map[string]interface{}{
			"nama_lengkap": session.Values["nama_lengkap"],
			"Pert" : MingguPertemuan.Pert,
			"mahasiswa": mahasiswa,
		}
		t, _ := template.ParseFiles("views/index.html")
		t.Execute(w, data)
	}	
}
// func Edit(w http.ResponseWriter, r *http.Request){
// 	session, _:= config.Store.Get(r, config.SESSION_ID)
// 	if r.Method == http.MethodGet{
// 		pert := r.URL.Query().Get("pert")
// 		no, _ := strconv.ParseInt(r.URL.Query().Get("no"), 10, 64)
// 		fmt.Println("pertemuan ke ", pert)
// 		fmt.Println("no ke ", no)

// 		var mahasiswa entities.Mahasiswa
// 		mhsModel.Find(no, pert, &mahasiswa)

// 		fmt.Println(mahasiswa)
// 		data := map[string]interface{}{
// 			"mahasiswa": mahasiswa,
// 		}
// 		t, err := template.ParseFiles("views/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}
// 		t.Execute(w, data)
// 	}else if r.Method == http.MethodPost{
// 		r.ParseForm()
// 		var mhs entities.Mahasiswa
// 		mhs.UTS,_ = strconv.ParseFloat(r.Form.Get("uts"), 32)
// 		mhs.UAS,_ = strconv.ParseFloat(r.Form.Get("uas"), 32)

// 		mahasiswa, _ := mhsModel.FindAll(&mhs, "pert", mhs.Pert)
// 		data:= map[string]interface{}{
// 			"nama_lengkap": session.Values["nama_lengkap"],
// 			"Pert" : mhs.Pert,
// 			"mahasiswa": mahasiswa,
// 		}
// 		t, _ := template.ParseFiles("views/index.html")
// 		t.Execute(w, data)
// 	}
// }