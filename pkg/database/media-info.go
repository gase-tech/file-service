package database

import (
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/helper"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"math"
	"time"
)

type MediaInfo struct {
	ID          uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CustomerID  string    `json:"customerId"`
	TenantID    string    `json:"tenantId"`
	Name        string    `json:"name"`
	Extension   string    `json:"extension"`
	Description string    `json:"description"`
}

type MediaInfoDao struct {
}

func (m MediaInfoDao) Save(info *MediaInfo) error {
	var equivalentInfo MediaInfo
	err := m.FindByName(&equivalentInfo, info.Name)

	if err != nil && err.Error() != "record not found" {
		log.Err(err)
		return err
	}

	err = store.Connection.Create(info).Error

	if err != nil {
		log.Err(err)
		return err
	} else {
		return nil
	}
}

func (m MediaInfoDao) DeleteByID(id uuid.UUID) {
	store.Connection.Delete(&MediaInfo{}, id)
}

func (m MediaInfoDao) FindByName(info *MediaInfo, name string) error {
	err := store.Connection.Where("name = ?", name).Find(info).Error

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (m MediaInfoDao) FindByID(info *MediaInfo, id uuid.UUID) error {
	err := store.Connection.Where("id = ?", id).Find(info).Error

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (m MediaInfoDao) GetPageable(page int, size int) helper.Pagination {
	var infos = []MediaInfo{}

	store.Connection.Scopes(helper.Paginate(page, size)).Find(&infos)

	var totalRecord int64 = 0

	store.Connection.Model(MediaInfo{}).Count(&totalRecord)

	totalPage := int(math.Ceil(float64(totalRecord) / float64(size)))

	var nextPage int

	if page == totalPage {
		nextPage = totalPage
	} else {
		nextPage = page + 1
	}

	var prevPage int

	if page <= 1 {
		prevPage = 1
	} else {
		prevPage = page - 1
	}

	return helper.Pagination{
		Limit:       size,
		Page:        page,
		Data:        infos,
		TotalRecord: totalRecord,
		TotalPage:   totalPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
}
