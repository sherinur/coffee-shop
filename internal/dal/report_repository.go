package dal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"coffee-shop/internal/utils"
	"coffee-shop/models"
)

type ReportRepository interface {
	GetTotalSales() (models.TotalSales, error)
	SetTotalSales(t float64) error
	SaveTotalSales(totalSales models.TotalSales) error
	UpdateTotalSales(income float64) error
	ResetTotalSales(income float64) error
}

type reportRepository struct {
	filePath string
}

func NewReportRepository(filePath string) *reportRepository {
	return &reportRepository{filePath: filePath}
}

func (r *reportRepository) GetTotalSales() (models.TotalSales, error) {
	totalSales := models.TotalSales{}

	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return totalSales, err
	}
	if !exists {
		return totalSales, nil
	}

	file, err := os.Open(r.filePath)
	if err != nil {
		return totalSales, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if utils.FileEmpty(file) {
		return totalSales, nil
	}

	err = decoder.Decode(&totalSales)
	if err != nil {
		return totalSales, err
	}

	return totalSales, nil
}

func (r *reportRepository) SaveTotalSales(totalSales models.TotalSales) error {
	// Checking the existence of a directory for a file
	dir := filepath.Dir(r.filePath)
	err := utils.CreateDir(dir)
	if err != nil {
		return err
	}

	// Checking file write permissions
	exists, err := utils.FileExists(r.filePath)
	if err != nil {
		return err
	}

	if !exists {
		// If the file does not exist, create it
		err := utils.CreateFile(r.filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", r.filePath, err)
		}
	}

	jsonData, err := json.MarshalIndent(totalSales, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) SetTotalSales(t float64) error {
	totalSales, err := r.GetTotalSales()
	if err != nil {
		return err
	}

	totalSales.TotalSales = t

	return r.SaveTotalSales(totalSales)
}

func (r *reportRepository) UpdateTotalSales(income float64) error {
	totalSales, err := r.GetTotalSales()
	if err != nil {
		return err
	}

	totalSales.TotalSales += income
	return r.SaveTotalSales(totalSales)
}

func (r *reportRepository) ResetTotalSales(income float64) error {
	return r.SetTotalSales(0)
}
