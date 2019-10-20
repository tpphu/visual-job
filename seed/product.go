package seed

import (
	"github.com/tpphu/visual-job/model"
	"syreclabs.com/go/faker"
)

func (s Seeder) ProductSeeed() {
	for i := 0; i < 1000000; i++ {
		p := model.Product{}
		p.Name = faker.Commerce().ProductName()
		p.Category = faker.Commerce().Department()
		p.Price = faker.Commerce().Price()
		s.DB.Create(&p)
	}
}
