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

