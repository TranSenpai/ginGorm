package repo

import (
	"main/internal/entity"

	"gorm.io/gorm"
)

func migrateTable(connection *gorm.DB) error {
	return connection.AutoMigrate(&entity.Contract{})
}

func createTableContract(dbConnection *gorm.DB) {
	err := migrateTable(dbConnection)
	if err != nil {
		panic(err)
	}
}

func createPartitionContract(dbConnection *gorm.DB) {
	dbConnection.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE contracts
  	ADD COLUMN registry_month integer unsigned GENERATED ALWAYS AS (MONTH(registry_at)) STORED`)

	dbConnection.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE contracts
	PARTITION BY LIST COLUMNS (id, registry_month) (
		PARTITION p01 VALUES IN (1),
		PARTITION p02 VALUES IN (2),
		PARTITION p03 VALUES IN (3),
		PARTITION p04 VALUES IN (4),
		PARTITION p05 VALUES IN (5),
		PARTITION p06 VALUES IN (6),
		PARTITION p07 VALUES IN (7),
		PARTITION p08 VALUES IN (8),
		PARTITION p09 VALUES IN (9),
		PARTITION p10 VALUES IN (10),
		PARTITION p11 VALUES IN (11),
		PARTITION p12 VALUES IN (12)
	)`)
}
