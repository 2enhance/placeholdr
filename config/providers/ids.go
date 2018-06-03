package providers

import (
	"crypto/md5"

	"github.com/icrowley/fake"
)

func BrandName(brand string) string {
	if brand == "" {
		return fake.Brand()
	}

	return brand
}

func Seed(id string) int {
	if id == "" {
		id = fake.Sentence()
	}

	hash := md5.Sum([]byte(id))
	seedValue := 0

	for _, r := range hash {
		seedValue += int(r)
	}

	return seedValue
}
