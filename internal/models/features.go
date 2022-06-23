package models

import "time"

type (
	Features struct {
		Id            uint      `db:"id"`
		CreatedAt     time.Time `db:"created_at"`
		FeaturesKey   string    `db:"featureskey"`
		FeaturesValue string    `db:"featuresvalue"`
		ProductId     uint      `db:"product_id"`
	}
	FeaturesKey struct {
		Id        uint      `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		Name      string    `db:"name"`
	}
	FeaturesValue struct {
		Id            uint      `db:"id"`
		CreatedAt     time.Time `db:"created_at"`
		Name          string    `db:"name"`
		FeaturesKeyId uint      `db:"featuresKeyid"`
	}
	// many to many field(FeaturesKey, Category)
	FeaturesKeyCategory struct {
		Id            uint      `db:"id"`
		CreatedAt     time.Time `db:"created_at"`
		CategoryId    uint      `db:"categoryid"`
		FeaturesKeyId uint      `db:"featureskeyid"`
	}
	// many to many field(FeaturesValue, Product)
	FeaturesValueProduct struct {
		Id              uint      `db:"id"`
		CreatedAt       time.Time `db:"created_at"`
		ProductId       uint      `db:"productid"`
		FeaturesValueId uint      `db:"featuresvalueid"`
	}
)
