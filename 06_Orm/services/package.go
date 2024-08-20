package services

import (
	"orm/database"
	"orm/models"
)

type PackageServices struct{}

func (ps *PackageServices) GetAll() ([]models.Package, error) {
	var packages []models.Package
	if err := database.DB.Find(&packages).Error; err != nil {
		return []models.Package{}, err
	}
	return packages, nil
}

func (ps *PackageServices) GetByID(id string) (models.Package, error) {
	var pack models.Package
	if err := database.DB.First(&pack, "id = ?", id).Error; err != nil {
		return models.Package{}, err
	}
	return pack, nil
}

func (ps *PackageServices) Create(packReq models.PackageRequest) (models.Package, error) {
	var pack models.Package = models.Package{
		Name:             packReq.Name,
		Sender:           packReq.Sender,
		Receiver:         packReq.Receiver,
		Receiverlocation: packReq.Receiverlocation,
		Senderlocation:   packReq.Senderlocation,
		Fee:              packReq.Fee,
		Weight:           packReq.Weight,
	}
	result := database.DB.Create(&pack)
	if err := result.Error; err != nil {
		return models.Package{}, err
	}
	if err := result.Last(&pack).Error; err != nil {
		return models.Package{}, err
	}
	return pack, nil
}

func (ps *PackageServices) Update(id string, packReq models.PackageRequest) (models.Package, error) {
	pack, err := ps.GetByID(id)

	if err != nil {
		return models.Package{}, err
	}

	pack.Name = packReq.Name
	pack.Sender = packReq.Sender
	pack.Receiver = packReq.Receiver
	pack.Senderlocation = packReq.Senderlocation
	pack.Receiverlocation = packReq.Receiverlocation
	pack.Fee = packReq.Fee
	pack.Weight = packReq.Weight
	if err := database.DB.Save(&pack).Error; err != nil {
		return models.Package{}, err
	}

	return pack, nil
}

func (ps *PackageServices) Delete(id string) error {
	pack, err := ps.GetByID(id)

	if err != nil {
		return err
	}
	if err := database.DB.Delete(&pack).Error; err != nil {
		return err
	}
	return nil
}
