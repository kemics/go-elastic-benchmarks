package go_elastic_benchmarks

import (
	"context"
	"net/http"
	"testing"

	"github.com/olivere/elastic/v7"
)

func BenchmarkCatIndicesElasticOlivereNaive(b *testing.B) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.CatIndices().Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkCatIndicesElasticOlivereJSONIterator(b *testing.B) {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetDecoder(&JSONIteratorDecoder{}),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.CatIndices().Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkCatIndicesElasticOlivereFastHTTP(b *testing.B) {
	client, err := elastic.NewClient(
		elastic.SetHttpClient(&http.Client{Transport: &Transport{}}),
		elastic.SetSniff(false),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.CatIndices().Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkMatchAllElasticOlivereNaive(b *testing.B) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search().Query(elastic.NewRawStringQuery(`{"match_all": {}}`)).Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkMatchAllElasticOlivereJSONIterator(b *testing.B) {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetDecoder(&JSONIteratorDecoder{}),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search().Query(elastic.NewRawStringQuery(`{"match_all": {}}`)).Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkMatchAllElasticOlivereFastHTTP(b *testing.B) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHttpClient(&http.Client{Transport: &Transport{}}),
	)
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search().Query(elastic.NewRawStringQuery(`{"match_all": {}}`)).Do(context.Background())
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}
