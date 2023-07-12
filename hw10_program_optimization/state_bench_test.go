package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetDomainStat(b *testing.B) {
	r, err := zip.OpenReader("testdata/users.dat.zip")
	if err != nil {
		b.Fatalf("failed to open zip file: %v", err)
	}
	defer r.Close()

	data, err := r.File[0].Open()
	if err != nil {
		b.Fatalf("failed to open file inside zip: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetDomainStat(data, "biz")
		if err != nil {
			b.Fatalf("failed to get domain stat: %v", err)
		}
	}
}
