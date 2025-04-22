# Jalankan semua unit test di folder src/test dan generate coverage
test:
	go test -coverprofile=coverage.out ./src/...

# Tampilkan hasil coverage per fungsi
cover:
	go tool cover -func=coverage.out

# Buka tampilan HTML coverage
cover-html:
	go tool cover -html=coverage.out

# Hapus file coverage (opsional buat bersih-bersih)
clean:
	rm -f coverage.out
