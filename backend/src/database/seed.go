package database

import (
	"fmt"
	"math/rand"
	"time"

	"myapp-go/src/models"
	"myapp-go/src/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RandomBalance(min, max int) float64 {
	return float64(rand.Intn(max-min+1) + min)
}

func SeedDB() {
	// 1. Seed Users + Customers
	users := []models.User{
		{ID: uuid.New(), Username: "admin", Role: "admin", Email: "admin@example.com", Password: "today1234"},
		{ID: uuid.New(), Username: "asepwil", Role: "user", Email: "mike@geller.com", Password: "today1234"},
		{ID: uuid.New(), Username: "luna", Role: "user", Email: "luna@email.com", Password: "luna1today1234234"},
		{ID: uuid.New(), Username: "krystal", Role: "user", Email: "krystal@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "febri", Role: "user", Email: "febri@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "hendrick", Role: "user", Email: "hendrick@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "tatang", Role: "user", Email: "tatang@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "both", Role: "user", Email: "both@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "kutu", Role: "user", Email: "kutu@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "dadi", Role: "user", Email: "dadi@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "zulfan", Role: "user", Email: "zulfan@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "ipul", Role: "user", Email: "ipul@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "yaman", Role: "user", Email: "yaman@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "vikri", Role: "user", Email: "vikri@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "fitra", Role: "user", Email: "fitra@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "odie", Role: "user", Email: "odie@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "gan", Role: "user", Email: "gan@email.com", Password: "today1234"},
		{ID: uuid.New(), Username: "car", Role: "user", Email: "car@email.com", Password: "today1234"},
	}

	for _, user := range users {
		var existing models.User
		if err := DB.Where("email = ?", user.Email).First(&existing).Error; err == gorm.ErrRecordNotFound {
			
			hashedPassword, _ := utils.HashPassword(user.Password)
			user.Password = hashedPassword

			DB.Create(&user)
			fmt.Println("User created:", user.Email)
			
			if user.Role == "user" {
				customer := models.Customer{
					ID:          uuid.New(),
					UserID:      user.ID,
					FullName:    user.Username,
					PhoneNumber: fmt.Sprintf("08%09d", rand.Intn(1000000000)),
					DateOfBirth: time.Date(1990+rand.Intn(10), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC),
					Address:     "Jl. " + user.Username + " Street",
				}
				DB.Create(&customer)
				fmt.Println("Customer created:", customer.FullName)
	
				accountNumber := fmt.Sprintf("92%010d", rand.Intn(1000000000))
	
				// 2. Seed BankAccount
				bankAccount := models.BankAccount{
					ID:         uuid.New(),
					CustomerID: customer.ID,
					AccountType: "Savings",
					Balance:    RandomBalance(100000, 5000000),
					AccountNumber:   accountNumber,
					Currency: "IDR",
				}
				DB.Create(&bankAccount)
	
				// 3. Seed Pocket
				pockets := []models.Pocket{
					{
						ID:            uuid.New(),
						BankAccountID: bankAccount.ID,
						PocketName:    "Liburan "+ user.Username,
						Balance:       max(0, 1500000),
						GoalAmount:    max(0, 10000000),
					},
					{
						ID:            uuid.New(),
						BankAccountID: bankAccount.ID,
						PocketName:    "test"+ user.Username ,
						Balance:       max(0, 15000000),
						GoalAmount:    max(0, 100000000),
					},
				}
				for _, pocket := range pockets {
					DB.Create(&pocket)
				}
	
				// 4. Seed Deposit
				termDeposits := []models.TermDeposit{
					{
						ID:            uuid.New(),
						BankAccountID: bankAccount.ID,
						DepositAmount: 1000000, 
						InterestRate:  5.00, 
						Tenor: 12,
						BalancePlusInterest: models.CalculateBalancePlusInterest(1000000, 5.00),
						StartDate:     time.Now(),
						Status:        "Active",
					},
					{
						ID:            uuid.New(),
						BankAccountID: bankAccount.ID,
						DepositAmount: 5000000,
						InterestRate:  6.50,
						Tenor: 12,
						BalancePlusInterest: models.CalculateBalancePlusInterest(5000000, 6.50),
						StartDate:     time.Now(),
						Status:        "Active",
					},
				}
				for _, termDeposits := range termDeposits {
					DB.Create(&termDeposits)
				}
			}
		}
	}
	fmt.Println("Seeding completed!")
}
