# Jalankan semua unit test di folder src/test dan generate coverage
	go test ./... -coverprofile ./cover.out -covermode atomic -coverpkg ./...

	head -n 1 cover.out > coverage_filter.out

	grep 'src/' cover.out | \
  grep -vE 'database/|dto/|utils/|services/|middlewares/|mocks/|authController.go|bankAccountController.go|pocketController.go|userController.go|models/|repositories/|routes/|container/' >> coverage_filter.out

# Tampilkan hasil coverage per fungsi
	go tool cover -func=coverage_filter.out

# Buka tampilan HTML coverage
	go tool cover -html=coverage_filter.out

# Hapus file coverage (opsional buat bersih-bersih)
	rm -f coverage_filter.out
