package job

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/tpphu/visual-job/model"
	"github.com/tpphu/visual-job/repo"
)

// ExportProduct use to export products to csv
type ExportProduct struct {
	chanCancel  chan bool
	hasCancel   bool
	ProductRepo repo.ProductRepo
	Data        model.Job
}

// Process is a main func to process data
func (j *ExportProduct) Process() {
	allProducts := [][]string{}
	for {
		if j.hasCancel == true {
			break
		}
		products, err := j.ProductRepo.List(j.Data.Offset, j.Data.Limit)
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		for _, p := range products {
			allProducts = append(allProducts, []string{
				fmt.Sprintf("%d", p.ID),
				p.Name,
				p.Category,
				fmt.Sprintf("%.3f", p.Price),
			})
		}
		if len(products) < j.Data.Limit {
			break
		}
		if len(allProducts) >= 100 {
			err = j.writeDatatFile(allProducts)
			if err != nil {
				time.Sleep(time.Second * 3)
				continue
			}
			allProducts = [][]string{}
		}
	}
	if len(allProducts) > 0 {
		j.writeDatatFile(allProducts)
	}
}

func (j *ExportProduct) WaitCancel() {
	j.hasCancel = <-j.chanCancel
}

// HasCancel returns true or false
func (j *ExportProduct) HasCancel() bool {
	return j.hasCancel
}

func (j *ExportProduct) writeDatatFile(data [][]string) error {
	filename := fmt.Sprintf("file_%s_%s.csv", data[0][0], data[len(data)-1][0])
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	return err
}
