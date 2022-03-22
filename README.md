# golang
github belajar golang

================================ CATATAN =========================
#Bagian 2 (Golang Modules)

63.Membuat Module
	go mod init github.com/ismail118/golang
	untuk versi di letakan di tag pada git, format contoh versinya v1.0.0

64.Menambah Dependency
	go get github.com/ismail118/golang

65.Upgrade Module
	setiap kali ada perubahan di code module sekecil apapun itu di haruskan mengganti version yg ada di tag git

66.Upgrade Dependency
	cukup rubah versi dependency yg ada di file go.mod
	setelah itu tinggal jalankan perintah di cmd "go get"

67.Major Upgrade
	setiap kali ada perubahan yg besar pada module yang kita buat, ada tehnik untuk menghindari error di aplikasi yang menggunakan module kita
	yaitu dengan cara mengganti nama module yg ada di file go.mod contohnya dari github.com/ismail118/golang --> github.com/ismail118/golang/v2
	dan pada tag versinya pun di update contoh v1.0.1 --> v2.0.0

#Bagian 3 (Golang Unit Test)

71.Pengenalan Testing Package
	ee documentation about package golang unit testing
	oc https://pkg.go.dev/testing
	x https://pkg.go.dev/testing#pkg-examples

	*esting.T
	olang menyediakan sebuah struct yang bernama testing.T yang di gunakan untuk unit test di golang
	*sting.M
	dalah struct yang disediakan golang untuk mengatur life cycle testing
	*sting.B
	dalah struct yang disediakan golang untuk melakukan benchmarking,
	igolang untuk melakukan benchmark ( mengukur kecepatan kode program ) pun sudah disediakan

72.Membuat Unit Test
	*Aturan File Test
	-untuk membuat file unit test kita harus menggunakan akhiran _test
	-jadi kita misal kita membuat file hello_world.go artinya untuk membuat unit testnya
	 kita membuat file hello_world_test.go
	
	*Aturan Function Unit Test
	-nama function unit test harus diawali Test
	 ex function HeloWorld() maka unit function unit testnya TestHelloWorld()
	-harus memiliki parameter (t *testing.T) dan tidak mengembalikan return

	*Menjalankan Unit Test
	-untuk menjalankan unit test kita bisa menggunakan perintah
	 go test
	-jika kita ingin lihat lebih detail function test apa saja yang sudah running, kita bisa gunakan perintah
	 go test -v
	- dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah
	 go test -v -run TestNamaFunction

	*Menjalankan Semua Unit Test
	-jika kita ingin menjalankan semua unit dari root folder module nya, kita bisa gunakan peirntah
	 go test -v ./...

73.Menggagalkan Unit Test
	-menggagalkan unit test menggunakan panic() bukanlah hal yg bagus
	-golang sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
	-terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test
	
	Penjelasan functin:
	*t.Fail()
	 akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test, Namun diakhir ketika selesai maka unit test tersebut dianggap gagal
	*t.FailNow()
	 akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
	*t.Error(args...)
	 function ini lebih seperti melakukan log (print) error, namun setelah melakukan log error,
	 dia akan otomatis memanggil function t.Fail()
	*t.Fatal()
	 mirip dengan t.Error() hanya saja setelah melakukan log error, dia secara otomatis akan memanggil t.FailNow()

	Note: recomended mengguanakn function t.Error() dan t.Fatal() karena ada log nya

74.Assertions
	-melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
	-apalagi jika result data yang harus di cek itu banyak
	-oleh karena itu, sangat disarankan untuk menggunakan Asserion untuk melakukan pengecekan
	-sayangnya golang tidak menyediakan package untuk assertion, sehingga kita BUTUH MENAMBAHKAN LIBRARY untuk melakukan Assertion ini

	*Testify
	-adalah salah satu library asserion yang paling populer di golang
	-link https://github.com/stretchr/testify
	-untuk menbahkan library gunakan perintah
	 go get github.com/stretchr/testify

	*testify.Assert()
	-saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(),
	 artinya eksekusi unit test akan tetap dilanjutkan
	*testify.Require()
	-saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil FailNow(),
	 artinya eksekusi unit test tidak akan dilanjutkan
	
75.Skip Test
	untuk membatalkan test ( bukan untuk menggagalkan )
	
	*Why We Need This?
	-kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
	-di golang kita bisa membatalkan eksekusi unit test jika kita mau
	-untuk membatalkan unit test kita bisa menggunakan function Skip("message")

76.Before dan After Test
	-biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
	-jika kode yang kita lakukan sebelum dan setelah selalu sama antara unit test function,
	 maka membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
	-untungnya golang terdapat fitur yang bernama testing.M
	-fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk
	 melakukan Before dan After di unit test
	
	*testing.M
	-untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
	-jika terdapat function TestMain tersebut, maka secara otomatis golang akan mengeksekusi function ini setiap kali akan menjalankan unit test di sebuah package
	-dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
	-ingat function TestMain itu dieksekusi hanya sekali per goalng package, bukan pertiap function unit test

77.Sub Test
	-golang mendukung fitur pembuatan function unit test di dalam function unit test
	-fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lain
	-untuk membuat sub tset, kita bisa menggunakan function Run()

	*Menjalankan hanya sub test
	-jika hanya ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah
	 go test -run TestNamaFunction/NamaSubTest