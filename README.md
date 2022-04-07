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

78.Table Test
	-sebelumnya kita sudah belajar tentang sub test
	-jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
	-dan fitur sub test ini, bisa di gunakan oleh programmer golang untuk membuat test dengan konsep table test
	-table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan ekspetasi hasil dari unit test
	-lalu slice tersebut kita iterasi menggunakan sub test

79.Mock
	-moch adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal
	-mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di tsting
	-contohnya kita ingin membuat unit test, namun ada kode program yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test,
	 karena unit testing kita harus selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita mau
	-pada kasus seperti ini cocok sekali untuk menggunakan mock object

	*Testify Mock
	-untuk membuat mock object, tidak ada fiture bawaan goalng, namu kita bisa mengguakan library testify yang sebelumnya kita gunakan untuk assertion
	-testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
	-namun perlu diperhatikan, jika desain kod program kita jelek, akan sulit untuk melakukan mocking, 
	 jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik

	*contoh kasusu
	#Aplikasi Query ke Database
	-kita akan coba contoh khusus dengan membuat contoh aplikasi golang yang melakukan query ke database
	-diaman kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
	-agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa interface	

80.Benchmark
	-selain unit test, golang testing package juga mendukung malekukan benchmark
	-benchmark adalah mekanisme meghitung kecepatan performa kode aplikasi kita
	-benchmark di golang dilakukan dengan secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
	-kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

	*testing.B
	-adalah struct yang digunakan untuk melakukan benchmark
	-testing.B mirip seperti testing.T, terdapat function Fail(), FailNow() , Error(), Fatal() dan lain-lain
	-yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
	-salah satunya adlaah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

	*Cara Kerja Benchmark
	-cara kerja benchmark di golang sangat sederhana
	-dimana kita hanya perlu membuat perulangan sejumlah N attribute
	-nanti secara otomatis golang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proces tersebut berjalan,
	 dan disimpulkan performa benchmark nya dalam waktu

81.Membuat Benchmark
	*Benchmark Function
	-mirip seperti unit test, untuk benchmark pun, di golang sudah ditentukan nama functionnya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld
	-selain itu, harus memiliki parameter (b 8testing.B)
	-dan tidak boleh mengembalikan return value
	-untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go

	#Menjalankan Benchmark
	-untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench:
	 go test -v -bench=.
	-jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah:
	 go test -v -run=NotMathUnitTest -bench=.
	-kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah:
	 go test -v -run=NotMathUnitTest -bench=BenchmarkTest
	-jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah:
	 go test -v -bench=. ./...

82.Sub Benchmark
	* konsepnya sama seperti #77.Sub Test

83.Table Benchmark
	-sama seperti di unit test, programmer golang terbiasa membuat table benchmark juga
	-ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function

86.Pengenalan Concurrency dan Parallel Programming
	*Pengenalan Parallel Programming
	-saat ini kita hidup di era multicore, dimana jarang sekali kita menggunakan processor yang single core
	-semakin canggih perangkat kerjas, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel di aplikasi.
	-parallel programming sederhananya adalah memecahkan suatu masalahdengan cara membaginya menjadi yang lebih kecil, 
	 dan dijalankan secara bersamaan pada waktu yang bersamaan pula

	Contoh Parallel:
	-menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser dan lain-lain)
	-beberapa koki menyiapkan makana di restoran, dimana tiap koki membuat makanan masing-masing
	-antrian di bank, dimana tiap teller melayani nasabah nya masing-masing

	Process vs Thread
	#Process
	-process adalah sebuah eksekusi program
	-process mengkonsumsi memory besar
	-process saling terisolasi dengan process lain
	-process lama untuk dijalankan dihentikan
	#Thread
	-Thread adalah segmen dari process
	-Thread menggunakan memory kecil
	-Thread bisa saling berhubungan jika dalam process yang sama
	-Thread cepat untuk dijalankan dan dihentikan

	Parallel vs Concurrency
	-berbeda dengan parallel ( menjalankan beberapa pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian
	-dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrency kita hanya membutuhkan sedikit Thread

	Contoh Concurrency
	-saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi, dan seterusnya.
	 Tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukan satu hal pada satu waktu,
	 namun bisa bergantian kapanpun kita mau

	#CPU-bound vs I/O-bound
	CPU-bound
	-banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya, Algoritma jenis ini biasanya sangat tergantung dengan kecepatan CPU
	-contoh yang paling populer adalah Machine Learning, oleh karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan G-PU
	 karena memiliki core yang lebih banyak dibanding CPU biasanya
	-jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implementasi Parallel Programming
	
	I/O-bound
	-I/O-bound adalah kebalikan dari sebelumnya, dimana biasanya algoritma atau aplikasinya sangat tergantung dengan kecepatan input output device yang digunakan
	-contohnya aplikasi seperti membaca data dari file, database, dan lain-lain
	-kebanyakan saat ini, biasanya kita akan membuat aplikasi jenis seperti ini
	-aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Concurrency Programming
	-bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapat balasan dari database, 
	 padahal waktu 1 detik itu jika menggunakan Concurrency Programming, bisa di gunakan untuk melakukan hal lain lagi

87.Pengenalan Goroutine
	-goroutine adalah sebuah thread ringan yang dikelola oleh go-runtime
	-ukuran go-runtime sangat kecil sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb
	-namun tidak seperti thread yang berjalan parallel, go-runtime berjalan secara concurrent

	#Cara Kerja Goroutine
	-sebenarnya Goroutine dijalankan oleh Go Sheduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS ( biasanya sejumlah core CPU )
	-jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan di atas Thread
	-namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur ole Go Scheduler

88.Membuat Project

89.Membuat Goroutine
	-untuk membuat goroutine di golang sangatlah sederhana
	-kita hanya cukup menambahkan perintah "go" sebelum memanggil function yang akan kita jalankan dalam goroutine
	-saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronouse, artinya tidak akan ditunggu sampai function tersebut selesai
	-aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

90.Goroutine Sangat Ringan

91.Pengenalan Channel
	-channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
	-di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
	-saat melakukan pengiriman data ke channel, goroutine akan terblock sampai ada yang menerima data tersebut
	-maka dari itu, channel disebut sebagai alat komunikasi synchronous (block)
	-channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa bahasa pemrogramman

	*Karakteristik Channel
	-secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel diambil
	-channel hanya bisa menerima satu jenis data
	-channel bisa diambil lebih dari satu goroutine
	-channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak

92.Membuat Channel
	-channel di golang direpresentasikan dengan tipe data chan
	-untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
	-namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan ke dalam channel tersebut
	ex channel := make(chan string)

	#Mengirim dan Menerima Data dari Channel
	-seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
	-untuk mengirim data, kita bisa gunnakan kode: channel <- data
	-sedangkan untuk menerima data, bisa gunakan kode: data <- channel
	-jika selesai, jangan lupa untuk menutup channel menggunakan function close()

93.Channel Sebagai Parameter
	-dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
	-sebelumnya kita tahu bahkan di golang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter,
	 sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference)
	-berbeda dengan channel, kita tidak perlu melakukan hal tersebut

94.Channel In dan Out
	-saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
	-kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk emngirim data,
	 atau hanya dapat digunakan untuk menerima data
	-hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerima data)

95.Buffered Channel
	-seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
	-aritnya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
	-kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel,
	 maka otomatis pengirim akan ikut lambat juga
	-untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di channel

	#Buffer Capacity
	-kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
	-jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
	-jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
	-ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data

	cara membuat Buffer Channel hampir sama seperti membuat channel biasa, hanya ada perbedaan sedikit, berikut contohnya:
	channel := make(chan string, 5)
	
96.Range Channel
	-kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
	-dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
	-salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
	-ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
	-ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual

97.Select Channel
	-kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
	-lalu kita ingin mendapatkan data dari semua channel tersebut
	-untuk melakukan hal tersebut, kita bisa menggunakan select channel di golang
	-dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang
	 secara bersamaan di beberapa channel, maka akan dipilih secara random

98.Default Select
	-apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
	-maka kita akan menunggu sampai data ada
	-kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
	-dalam select, kita bisa menambahkan default, dimana ini akan diekseuksi jika memang di semua channel yang kita select tidak ada datanya

99.Race Condition
	Masalah Dengan Goroutine
	-saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapu bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
	-hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
	-hal ini bisa menyebabkan masalah yang namanya Race Condition

100.Sync.Mutex (Mutual Exclusion)
	-untuk mengatasi masalah #99.Race Condition, di golang terdapat sebuah struct bernama sync.Mutex
	-Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan
	 locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
	-dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine
	 yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi
	-ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi

101.Sync.RWMutex
	-kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
	-kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
	-di golang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write

102.Sync.Deadlock
	-hati-hati saat membuat aplikasi yang parallel atau concurent, masalah yang sering kita hadapi adalah Deadlock
	-deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
	-sekarang kita coba simulasikan proses deadlock

103.Sync.WaitGroup
	-waitgroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses dilakukan
	-hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine,
	 tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
	-kasus seperti ini bisa menggunakan WaitGroup
	-untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int),
	 setelah proses goroutine selesai, kita bisa gunakan method Done()
	-untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()

104.Sync.Once
	-once adalah fitur di golang yang bisa kita gunakan untuk memastikan bahasa sebuah function di eksekusi hanya sekali
	-jadi berapa banyakpun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya
	-goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi

105.Sync.Pool
	-pool adalah implementasi design pattern bernama object pool pattern
	-sederhananya design pattern pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari pool,
	 dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke poolnya
	-implementasi pool digolang ini sudah aman dari problem race condition

106.Sync.Map
	-golang memiliki sebuah struct bernama sync.Map
	-map ini mirip golang map, namun yang membedakan, Map ini aman untuk menggunakan concurrent menggunakan goroutine
	-ada beberapa function yang bisa kita gunakan di Map:
		~Store(key, value) untuk menimpan data ke map
		~Load(key) untuk mengambil data dari Map menggunakan key
		~Delete(key) untuk menghapus data di Map menggunakan key
		~Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map

107.Sync.Cond
	-cond adalah implementasi locking berbasis kondisi
	-cond membutuhkan locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi lockingnya, namun berbeda dengan locker biasanya, di cond terdapat function Wait()
	 untuk menunggu apakah perlu menunggu atau tidak
	-function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() 
	 digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
	-untuk membuat Cond, kita bias menggunakan function sync.NewCond(Locker)

108.Atomic
	-golang memiliki package yang bernama sync/atomic
	-atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
	-contohnya, sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter,
	 hal ini sebenarnya bisa digunakan menggunakan Atomic package
	-ada banyak sekali function di atomic package, kita bisa ekplore sendiri di halaman dokumentasinya

109.time.Timer
	-timer adalah representasi satu kejadian
	-ketika wkatu timer sudah expire, maka event akan dikirim ke dalam channel
	-untuk membuat timer kita bisa menggunakan time.NewTimer(duration

	#time.After()
	-kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
	-untuk melakukan hal itu kita bisa menggunakan function time.After(duration)

	#time.AfterFunc()
	-kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
	-kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunct()
	-kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadian

110.time.Ticker
	-ticker adalah representasi kejadian yang berulang
	-ketika wkatu ticker sudah expire, maka event akan dikirim ke dalam channel
	-untuk membuat ticker, kita bisa menggunakan time,NewTicker(duration)
	-untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()

	#time.Tick()
	-kadang kita tidak butuh data Ticker nya, kita hnay butuh channel nya saja
	-jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
	 mengembalikan Ticker, hanya mengembalikan channel timer nya saja

111.GOMAXPROCS
	-sebelumnya diawal kita sudah bahas bahwa gorutine itu sebenarnya dijalankan di dalam Thread
	-pertanyaannya seberapa banyak Thread yang ada di golang ketika aplikasi kita berjalan?
	-untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah
	 sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
	-secara default, jumlah thread di golang itu sebanyak jumlah CPU di komputer kita
	-kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime,NumCpu()

#Bagian 5 (Golang Context)

114.Pengenalan Context
	-context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
	-context biasanya dibuat per request (misal setiap ada requset masuk ke server web melalui http request)\
	-context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses

	#Kenapa Context Perlu Dipelajari?
	-context di golang biasanya digunakan untuk mengirim data request atau sinyal ke proses lain
	-dengan menggunakan context, ketika kita ingin membatalkan semua proses, kita cukup mengirim sinyal ke context
	 maka secara otomatis semua proses akan dibatalkan
	-hampir semua bagian di golang memanfaatkan context, seperti database, http sever, http client, dan lain-lain
	-bahkan di google sendiri, ketika menggunakan golang, context wajib digunakan dan selalu dikirm ke setiap function yang dikirim

	#Package Context
	-context direpresentasikan di dalam sebuah interface Context
	-interface Context terdapat dalam package context
	-https://pkg.go.dev/context

115.Membuat Context
	-karena context adalah sebuah inteface, untuk membuat context kita butuh sebuah struct yang sesuai dengan kontreak interface context
	-namun kita tidak perlu membuatnya secara manual
	-di golang package context terdapat function yang bisa kita gunakan untuk membuat context

	#Function Membuat Context
	-context.Background()
	 membuat context kosong, tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun, Biasanya digunakan di main function atau dalam test,
	 atau dalam awal proses request terjadi
	-context.TODO()
	 membuat context kosong seperti Background(), namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan

116.Parent dan Child Context
	-context menganut konsep parent dan child
	-artinya, saat kita membuat context, kita bisa membuat child context dari context yang sudah ada
	-parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context
	-konsep ini mirip dengan pewarisan di pemrograman berorientasi object

	#Hubungan Antara Parent dan Child Context
	-parent dan child context akan selalu terhubung
	-saat nanti kita melakukan misal pembatalan context A, maka semua child dan sub child dari context A akan ikut dibatalkan
	-namun jika misal kita membatalkan context B, hanya context B dan semua child dan sub child nya yang dibatalkan,
	 parent context B tidak akan ikut dibatalkan
	-begitu juga nanti saat kita menyisipkan data ke dalam context A, semua child dan sub child nya bisa mendapatkan data tersebut
	-namun jika kita menyisipkan data di context B, hanya context B dan semua child dan sub child nya yang mendapat data, parent context B tidak akan mendapat data

	#Immutable
	-context merupakan object yang immutabel, artinya setelah context dibuat, dia tidak bisa diubah lagi
	-ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dan yang lainnya, secara otomatis akan membentuk child context baru, 
	 bukan merubah context tersebut

117.Context With Value
	-pada saat awal membuat context, context tidak memiliki value
	-kita bisa menambah sebuah value dengan data Pair (key - value) ke dalam context
	-saat kita menambah value ke context, secara otomatis akan tercipta child context baru, 
	 artinya original context nya tidak akan berubah sama sekali
	-untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)

118.Context With Cancel
	-selain menambahkan value ke context, kita juga bisa menambahkan sinyal cencel ke context
	-kapan sinyal cencel diperlukan dalam context?
	-biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
	-biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, 
	 kita bisa mengirim sinyal cancel ke context nya
	-namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap context nya, jika tidak, tidak ada gunannya
	-untuk membuat context dengan cancel signal, kita bisa menggunakan function context.WithCancel(parent)

119.Conext With Timeout
	-selain menambahkan value ke context, dan juga sinyal cencel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis
	 dengan menggunakan pengaturan timeout
	-dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual,
	 cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
	-penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api,
	 namun ingin menetukan batas maksimal timeout nya
	-untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan
	 function context.WithTimeout(parent, duration)

120.Context With Deadline
	-selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline
	-pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang,
	 kalo deadline ditentukan kapan waktu timeoutnya, misal jam 12 siang hari ini
	-untuk membuat context dengan cancel secara otomatis menggunakan deadline, kita bisa menggunakan function context.WithDeadline(parent, time)

123.Pengenalan Package Database
	-bahasa pemrograman golang secara default memiliki sebuah package bernama database
	-package database adalah package yang berisikan kumpulan standard interface yang menjadi standard untuk berkomunikasi ke database
	-hal ini menjadikan kode program yang kita buat untuk mengakses jenis database apapun bisa menggunakan kode yang sama
	-yang berbeda hanya kode SQL yang perlu kita gunakan sesuai dengan database yang kita gunakan

124.Menambah Database Driver
	-sebelum kita membuat kode program menggunakan database di golang, terlebih dahulu kita wajib menambahkan driver database nya
	-tanpa driver database, maka package database di golang tidak mengerti apapun, karena hanya berisi kontrak interface saja
	-doc https://golang.org/s/sqldrivers

	#Menambahkan Module Database MySQL
	go get -u github.com/go-sql-driver/mysql
	
125.Membuka Koneksi ke Database
	-hal yang pertama kali kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah melakukan koneksi ke database nya
	-untuk melakukan koneksi ke database di golang, kita bisa membuat project sql.DB menggunakan function sql.Open(driver, dataSourceName)
	-untuk menggunakan database MySQL, kita bisa menggunakan driver "mysql"
	-sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing,
	 misal di MySQL, kita bisa menggunakan format dataSourceName seperti dibawah ini:
	 format : username:password@tcp(host:port)/database_name
	 contoh : root:Colonelgila123@tcp(localhost:3306)/test
	-jika object sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan function Close()

126.Database Pooling
	-sql.DB di golang sebenarnya bukanlah sebuah koneksi ke database
	-melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
	-di dalam sql.DB, golang melakukan management koneksi ke database secara otomatis,
	 hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
	-dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat oleh golang,
	 sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa di tangani oleh database yang kita gunakan

	#Pengaturan Database Pooling
	-(DB) SetMaxIdleConns(number)
	 pengaturan berapa jumlah koneksi minimal yang di buat
	-(DB) SetMaxOpenConns(number)
	 pengaturan berapa jumlah koneksi maksimal yang di buat
	-(DB) SetConnMaxIdleTime(duration)
	 pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	-(DB) SecConnMaxLifetime(duration)
	 pengaturan berapa lama koneksi boleh digunakan

127.Eksekusi Perintah SQL
	-saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
	-digolang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function (DB)ExecContext(context, sql, params)
	-ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah perinah kita pelajari di course golang context, dengan context, kita bisa
	 mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya
	
	#Kode Membuat Table Customer
	CREATE TABLE customer ( id VARCHAR(100) NOT NULL, name VARCHAR(100) NOT NULL, PRIMARY KEY (id) ) ENGINE = InnoDB;

128.Query SQL
	-untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika kita membutuhkan result,
	 seperti SELECT SQL, kita bisa menggunakan function yang berbeda
	-function untuk melakukan query ke database, bisa menggunakan function (DB) QueryContext(context, sql, params)

	#Rows
	-hasil query function adalah sebuah data struct sql.Rows
	-rows digunakan untuk melakukan iterasi terhadap hasil dari query
	-kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi terhadap data hasil query,
	 jika return data false, artinya sudah tidak ada data lagi didalam result
	-untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
	-dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan (Rows) Close()

129.Tipe Data Column
	-sebelumnya kita hanya membuat table dengan tipe data di kolomnya berupa VARCHAR
	-untuk varchar di database, biasanya kita gunakan string di golang
	-bagaimana dengan tipe data yang lain?
	-apa representasinya di golang, misal tipe data timestamp, date dan lain-lain

	#Buat Table Baru
	CREATE TABLE buyer ( id VARCHAR(100) NOT NULL, name VARCHAR(100) NOT NULL, email VARCHAR(100), balance INTEGER DEFAULT 0, 
rating DOUBLE DEFAULT 0.0, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, birth_date DATE, married BOOLEAN DEFAULT false, PRIMARY KEY (id) ) ENGINE = InnoDB;

	#Mapping Tipe Data di Database --> Tipe Data di Golang
	VARCHAR, CHAR --> string
	INT, BIGINT --> int32, int64
	FLOAT, DOUBLE --> float32, float64
	BOOLEAN --> bool
	DATE, DATETIME, TIME, TIMESTAMP --> time.Time

	#Kode Insert Data buyer
	INSERT INTO buyer (id, name, email, balance, rating, birth_date, married) 
VALUES ('budi', 'BUDI', 'budi@gmail.com', 100000, 5.0, '1999-9-9', true),('eko', 'Eko', 'eko@gmail.com', 100000, 5.0, '1999-9-9', true);
	
	#Error Tipe Data Date
	-secara default, Driver MySql untuk golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte / []uint8,
	 dimana ini bisa dikonversi menjadi String, lalu di parsing menjadi time.Time
	-namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL 
	 untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true ( pada saat membuat koneksi ke db)0
	ex: db, err := sql.Open("mysql", "root:Colonelgila123@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	 
	#Nullable Type
	-golang database tidak mengerti dengan tipe data NULL di database
	-oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita
	 melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di golang

	#Error Data Null
	-konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
	-oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam package.sql
	
	Tipe Data Nullable
	#Tipe Data Golang --> Tipe Data Nullable
	string --> database/sql.NullString
	bool --> database/sql.NullBool
	float64 --> database/sql.NullFloat64
	int32 --> database/sql.NullInt32
	int64 --> database/sql.NullInt64
	time.Time --> database/sql.NullTime

130.SQL Injection
	-saat membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode golang kita
	-biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user, dan mengirimnya menggunakan perintah SQL

	#Kode Membuat Table User
	CREATE TABLE user ( username VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL, PRIMARY KEY (username) ) ENGINE = InnoDB;
	INSERT INTO user (username, password) VALUES ('admin', 'admin');

	#SQL Injection
	-sql injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan basis data sebuah aplikasi
	-biasa, sql injection dilakukan dengan mengirim input dari user dengan perintah yang salah,
	 sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
	-sql injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman

	#Solusinya?
	-jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
	-jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute atau Query dengan parameter yang akan kita bahas di chapter selanjutnya

131.SQL Dengan Parameter
	-sekarang kita sudah tahu bahaya nya SQL injection jika menggabungkan string ketika membuat query
	-jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan
	 yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL Query yang kita buat
	-untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)

	#Contoh SQL
	SELECT username FROM user WHERE username= ? AND password= ? LIMIT 1
	INSERT INTO user(username, password) VALUES (?, ?)
	Dan lain-lain

132.Auto Increment
	-kadang kita membuat sebuah table dengan id auto increment
	-dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
	-sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
	-tapi untungnya di golang ada cara yang lebih mudah
	-kita bisa menggunakan function (Result) LastInsertId() untuk mendapatkan Id terakhir yang dibuat secara auto increment
	-result adalah object yang dikembalikan ketika kita menggunakan function Exec

	#Kode Membuat Table
	CREATE TABLE comments ( id INT NOT NULL AUTO_INCREMENT, email VARCHAR(100) NOT NULL, comment TEXT, PRIMARY KEY (id)) ENGINE InnoDB;

133.Prepare Statement
	#Query atau Exec dengan parameter
	-saat kita menggunakan function Query() atau Exec() yang menggunakan parameter, sebenarnya implementasi dibawahnya menggunakan Prepare Statement
	-jadi tahapan pertama statement nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
	-kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya, misal insert data langsung banyak
	-pembuatan prepare statement bisa dilakukan dengan manual, tanpa harus menggunakan Query() atau Exec() dengan parameter

	#Prepare Statement
	-saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
	-sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, 
	 maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statement nya hanya sekali diawal saja
	-jika menggunakan Query() atau Exec() dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama,
	 oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
	-untuk membuat Prepare Statement, kita bisa menggunakan function (DB) Prepare(context, sql)
	-prepare statement direpresentasikan dalam struct database/sql.Stmt
	-sama seperti resource sql lainnya, Stmt harus di Close() jika sudah tidak digunakan lagi
	
134.Databse Transaction`
	-salah satu fitur andalan di database adalah transaction
	-materi database transaction sudah saya bahas dengan tuntas di curse MySql database, jadi silahakan pelajari di course tersebut
	-di course ini kita akan fokus bagaimana menggunakan database transaction di golang

	#Transaction di Golang
	-secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya auto commit
	-namun kita bisa menggunakan transaksi sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
	-untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
	-struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada Tx, seperti Exec, Query atau Prepare
	-setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit atau Rollback()

135.Repository Pattern
	-dalam buku Domain-Driven Design, Eric Evans menjelaskan bahwa 
	 "repository is a mechanism for encapsulation storage, retrieval, and search behaviro, which emulates a collection of object"
	-Pattern Repository ini biasanya digunakan sebagai jembatan antara business logic aplikasi kita dengan semua perintah SQL ke database
	-jadi semua perintah SQL akan ditulis di Repository, sedangkan business logic kode program kita hanya cukup menggunakan Repository tersebut

	#Entity / Model
	-dalam pemrograman berorientasi object, biasanya sebuah tabel di database akan selalu dibuat representasinya sebagai class Entity atau Model,
	 namun di golang, karena tidak mengenal Class jadi kita akan representasikan data dalam bentuk Struct
	-ini bisa mempermudah ketika membuat kode program
	-misal ketika kita query ke Repository, dibandingkan mengembalikan array, alangkah baiknya Repository melakukan konversi terlebih dahulu ke struct Entity / Model,
	 sehingga kita tinggal gunakan objectnya saja

================= comand go get driver MySQL ================
	go get -u github.com/go-sql-driver/mysql

#Bagian 7 (Golang Embed)

138.Pengenalan Embed Package
	-sejak golang versi 1.16 terdapat package baru dengan nama embed
	-package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan isi file nya dalam variable
	doc https://pkg.go.dev/embed

	#Cara Embed File
	-untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
	-selanjutnya kita bisa tambahkan komentar //go:embed diikuti dengan nama file nya, diatas variable yang kita tuju
	-variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara otomatis ketika kode golang di compile
	-variable yang dituju tidak bisa disimpan di dalam function
	
139.Embed File ke String
	-embed file bisa kita lakukan ke variable dengan tipe data string
	-secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut

140.Embed File ke Byte
	-selain ke tipe data string, embed file juga bisa dilakukan ke variable tipe data []byte
	-ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain

141.Embed Multiple Files
	-kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
	-hal ini juga bisa dilakukan menggunakan embed package
	-kita bisa menambahkan komentar //go:embed lebih dari satu baris
	-selain itu variable nya bisa kita gunakan tipe data embed.FS

142.Patch Matcher
	-selain manual satu per satu, kita bisa menggunakan patch matcher untuk membaca multiple file yang kita inginkan
	-ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
	-caranya, kita perlu menggunakan path matcher pada package function path.Match

143.Hasil Emebed di Compile
	-perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golangnya
	-artinya bukan dilakukan secara realtime membaca file yang ada diluar
	-hal ini menjadikan jika binary file golang sudah di complie, kita tidak butuh lagi file externalnya,
	 dan bahkan jika diubah file externalnya, isi variable nya tidak akan berubah lagi

	NOTE: KODE NYA TIDAK SAYA TULIS, LIAT VIDEO TUTORIAL UNTUK LEBIH JELAS

146.Pengenalan Web
	#Kenapa Web?
	-saat ini web digunakan oleh jutaan, bahkan mungkin milyaran orang setiap hari
	-dengan web, kita bisa melakukan belajar online, mendengarkan musik online, nonton video online, belanja online sampai memesan makanan secara online
	-namun perlu diperhatikan, web bukanlah internet

	#Web
	-web merupakan kumpulan informasi yang tersedia dalam sebuah komputer yang terkoneksi secara terus menerus melalui internet
	-web bisa berisi informasi dalam bentuk apapun, seperti teks, gambar, audio, video dan lain-lain
	-web berjalan di aplikasi yang bernama web server, yaitu aplikasi yang digunakan untuk menyimpan dan menyampaikan isi informasi web

	#Web Host
	-pemilik web, biasanya tidak menjalankan aplikasi Web Server di komputer pribadi nya
	-biasanya mereka akan menyewa komputer di tempat penyedia data center (kumpulan komputer) yang terjamin keandalan dan kecepatan koneksi internetnya
	-pihak penyedia komputer untuk web server biasanya disebut web host

	#Domain
	-saat komputer web terhubung ke internet, biasanya dia memiliki alamat
	-alamat ini bernama ip address, formatnya misal nya 172.217.194.94
	-karena alamat ip address sangat menyulitkan untuk diingat
	-untung saja ada yang namanya nama domain
	-nama domain adalah alamat yang bisa digunakan sebagai alias ke ip address
	-misal seperti google.co.id, blibli.com, dan lain-lain
	-dengan nama domain, sebagai manusia kita akan mudah mengingat dibandingkan ip address
	-namun, saat kita menggunakan nama domain, sebenarnya komputer tetap akan mengakses web menggunakan alamat ip address

147.Client dan Server
	-web adalah aplikasi berbasis Client dan Server, sekarang pertanyaannya, aapa itu Client dan Server?
	-sederhananya client server merupakan konsep arsitektur aplikasi yang menghubungkan dua pihak, sistem client dan sistem server
	-sistem client dan sistem server yang saling berkomunikasi melalui jaringan komputer, internet atau juga bisa di komputer yang sama

	#Tugas Client dan Server
	-aplikasi Client bertugas mengirim request ke Server, dan menerima response dari server
	-sedangkan, aplikasi server bertugas menerima request dari Client, memproses data, dan mengembalikan hasil proses data ke Client

	#Keuntungan Client dan Server
	-perubahan aplikasi bisa dilakukan dengan mudah di server, tanpa harus membuat perubahan di client, aplikasi jika client nya di lokasi yang sulit dijangkau
	-bisa digunakan oleh banyak client pada saat yang bersamaan, walaupun server tidak banyak
	-bisa diakses dari mana saja, asal terhubung satu jaringan dengan server
	
	#Contoh Client dan Server
	-web adalah salah satu contoh arsitektur client server
	-aplikasi yang bertugas sebagai Client adalah Web Browser
	-aplikasi yang bertugas sebagai Server adalah Web Server, dimana di dalam web server terdapat kode program Web kita

148.Golang Web
	-golang saat ini populer dijadikan salah satu pilihan bahasa pemrograman untuk membuat Web terutama Web API (backend)
	-selain itu, di golang juga sudah disediakan package untuk membuat Web, bahkan di sertakan pula packagae untuk implementasi unit testing untuk Web
	-hal ini menjadikan pembuatan Web menggunakan golang lebih mudah, karena tidak butuh menggunakan library atau framework

	#Cara Kerja Golang Web
	-Web Browser akan melakukan HTTP Request ke Web Server
	-golang menerima HTTP Request, lalu mengeksekusi request tersebut sesuai dengan yang diminta
	-setelah melakukan eksekusi request, golang akan mengembalikan data dan di render sesuai
	 dengan kebutuhannya, misal HTML, CSS, JavaScript dan lain-lain
	-golang akan mengembalikan content hasil render tersebut sebagai HTTP Response ke Web Browser
	-web browser menerima content dari web server, lalu me-render content tersebut sesuai dengan tipe content nya

	#Package net/http
	-pada beberapa bahasa pemrograman lain, seperti java misalnya, untuk membuat web biasanya dibutuhkan tambahan library atau framework
	-sedangkan di golang sudah disediakan package untuk membuat web bernama package net/http
	-sehingga untuk membuat web menggunakan golang, kita tidak butuh lagi library tambahan, kita bisa menggunakan package yang sudah tersedia
	-walaupun memang saat kita membuat web dalam skala besar, direkomendasikan menggunakan framework karena beberapa hal sudah dipermudah oleh web framework
	-namun pada course ini, kita akan fokus menggunakan package net/http untuk membuat web nya. 
	 karena semua framework web golang akan menggunakan net/http sebagai basis dasar framework nya
	
149.Sever
	-server adalah sturct yang terdapat di package net/http yang digunakan sebagai representasi Web Server di golang
	-untuk membuat web, kita wajib membuat server
	-untuk membuat data server, ada beberapa hal yang perlu kita tentukan, seperti host dan juga port tempat dimana web kita berjalan
	-setelah membuat server, kita bisa menjalankan server tersebut menggunakan function ListenAndServe()

150.Handler
	-server hanya bertugas sebagai web server, sedangkan untuk menerima http request yang masuk ke server, kita butugh yang namanya handler
	-handler golang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah function bernama ServeHTTP() 
	 yang digunakan sebagai function yang akan di eksekusi ketika menerima HTTP Request

151.ServeMux
	-saat membuat web, kita biasanya ingin membuat banyak sekali endpoint URL
	-HandlerFunc sayangnya tidak mendukung itu
	-alternative implementasi dari Handler adalah ServeMux
	-ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint

	#URL Pattern
	-URL Pattern dalam ServeMux sederhana, kita tinggal menambahkan string yang ingin kita gunakan
	 sebagai endpoint, tanpa perlu memasukan domain web kita
	-jika URL Pattern dalam ServeMux kita tambahkan di akhirnya dengan garis miring, artinya semua url
	 tersebut akan menerima path dengan awalan tersebut, misal /images/ artinya dieksekusi jika endpoint nya /images/,
	 /images/contoh, /images/contoh/lagi
	-namun jika terdapat URL Pattern yang lebih panjang, maka akan diprioritaskan yang lebih panjang,
	 misal jika terdapat URL /images/ dan /images/thumbnails/, maka jika mengakses /images/thumbnails/ akan mengakses /images/thumbnails/, bukan /images

152.Request
	-request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh Web Browser
	-semua informasi request yang dikirim bisa kita dapatkan di Request
	-seperti URL, http method, http header, http body dan lain-lain

153.HTTP Test
	-golang sudah menyediakan package khusus untuk membuat unit test terhadap fitur web yang kita buat
	-semuanya ada di dalam package net/http/httptest
	 doc https://pkg.go.dev/net/http/httptest
	-dengan menggunakan package ini, kita bisa melakukan testing handler web di golang tanpa harus menjalankan aplikasi web nya
	-kita bisa langsung fokus terhadap handler function yang ingin kita test

	#httptest.NewRequest()
	-NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.Request
	-kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
	-selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, dan lain-lain

	#httptest.NewRecorder()
	-httptest.NewRecorder() merupakan function yang digunakan untuk membuat ResponseRecorder
	-ResponseRecorder merupakan struct bantuan untuk merekam HTTP response dari hasil testing yang kita lakukan

154.Query Parameter
	-query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
	-query parameter biasanya digunakan untuk mengirim data dari client ke server
	-query parameter di tempatkan pada URL
	-untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya

	#url.URL
	-dalam parameter Request, terdapat attribute URL yang berisi data dari URL
	-dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari clinet dengan menggunakan
	 method Query() yang mengembalikan map

	#Multiple Query Parameter
	-dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
	-ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
	-untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya

	#Multiple Value Query Parameter
	-sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
	-artinya dalam satu key query parameter, kita bisa memasukan beberapa value
	-caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda,
	 misal:
	 name=eko&name=kurniawan

155.Header
	-selain query parameter, dalam HTTP, ada juga yang bernama header
	-Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
	-jadi dalam header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan informasi header
	-saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser,
	 seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain

	#Request Header
	-untuk menangkap request header yang dikirim oleh client, kita bisa mengambilnya di Request.Header
	-Header mirip seperti Query Parameter, isinya adalah map[string][]stirng
	-berbeda dengan Query Parameter yang case sensitive, secara spesfifikasi, Header key tidaklah case sensitive

	#Response Header
	-sedangkan jika kita ingin menambahkan header pada response, kita bisa menggunakan function ResponseWriter.Header()

156.Form Post
	-saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method GET atau POST
	-jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
	-sedangkan jika menggunakan POST, maka semua data di form akan dikirim via body HTTP Request
	-di golang untuk mengambil data Form Post sangatlah mudah

	#Request.PostForm
	-semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute Request.PostForm
	-namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method Request.ParseForm() terlebih dahulu,
	 method ini di gunakan untuk melakukan parsing data body apakah bisa di parsing menjadi form data atau tidak,
	 jika tidak bisa di parsing, maka akan menyebabkan error