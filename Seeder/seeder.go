package seeder

import "github.com/fathimasithara01/ecommerce/utils/seeder"

func GroupSeeder() {
	seeder.SeedCategories()
	seeder.SeedBrands()
	seeder.SeedProducts()
}
