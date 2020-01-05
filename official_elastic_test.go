package go_elastic_benchmarks

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkCatIndicesOfficialNaive(b *testing.B) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Cat.Indices()
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkCatIndicesOfficialFastHTTP(b *testing.B) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Transport: &Transport{},
	})
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Cat.Indices()
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		_ = resp
	}
}

func BenchmarkMatchAllOfficialNaive(b *testing.B) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search(client.Search.WithBody(strings.NewReader(`{"query":{"match_all": {}}}`)))
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		var i interface{}
		if err := json.NewDecoder(resp.Body).Decode(&i); err != nil {
			b.Fatalf("Can't decode response from elasticsearch, error: %v", err)
		}
		if err := resp.Body.Close(); err != nil {
			b.Fatalf("Can't clode response body, error: %v", err)
		}
	}
}

func BenchmarkMatchAllOfficialJSONIterator(b *testing.B) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search(client.Search.WithBody(strings.NewReader(`{"query":{"match_all": {}}}`)))
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		var i interface{}

		if err := jsoniter.NewDecoder(resp.Body).Decode(&i); err != nil {
			b.Fatalf("Can't decode response from elasticsearch, error: %v", err)
		}
		if err := resp.Body.Close(); err != nil {
			b.Fatalf("Can't clode response body, error: %v", err)
		}
	}
}

func BenchmarkMatchAllOfficialFastHTTP(b *testing.B) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Transport: &Transport{},
	})
	if err != nil {
		b.Fatalf("Can't connect to elasticsearch, error: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Search(client.Search.WithBody(strings.NewReader(`{"query":{"match_all": {}}}`)))
		if err != nil {
			b.Fatalf("Can't perform request to elasticsearch, error: %v", err)
		}
		var i interface{}
		if err := json.NewDecoder(resp.Body).Decode(&i); err != nil {
			b.Fatalf("Can't decode response from elasticsearch, error: %v", err)
		}
		if err := resp.Body.Close(); err != nil {
			b.Fatalf("Can't clode response body, error: %v", err)
		}
	}
}
