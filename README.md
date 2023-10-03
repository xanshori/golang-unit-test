# Panduan Pembuatan Unit Test dalam Go (Golang)

## Aturan File Test

1. Go-Lang memiliki aturan cara membuat file khusus untuk unit test.
2. Untuk membuat file unit test, kita harus menggunakan akhiran `_test`.
3. Sebagai contoh, jika kita memiliki file `hello_world.go`, maka untuk membuat unit testnya, kita harus membuat file dengan nama `hello_world_test.go`.

## Aturan Nama Function Unit Test

1. Selain aturan untuk nama file, di Go-Lang juga telah diatur aturan untuk nama function unit test.
2. Nama function untuk unit test harus diawali dengan kata `Test`.
3. Misalnya, jika kita ingin menguji function `HelloWorld`, maka kita akan membuat function unit test dengan nama `TestHelloWorld`.
4. Tak ada aturan yang mengharuskan nama belakang function unit test harus sama dengan nama function yang akan diuji, yang penting adalah harus diawali dengan kata `Test`.
5. Function unit test harus memiliki parameter `t *testing.T` dan tidak mengembalikan return value.

Dengan mengikuti aturan-aturan di atas, Anda dapat dengan mudah membuat dan menjalankan unit test dalam proyek Go Anda untuk memastikan kualitas dan keandalan kode Anda.
