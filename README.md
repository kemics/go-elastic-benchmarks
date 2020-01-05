# ElasticSearch clients benchmarks

## Results

### Perform cat indices requests

`go test -bench=BenchmarkCatIndices -benchtime=10000x -timeout=100m`

#### [Elastic Olivere](https://github.com/olivere/elastic)

```
BenchmarkCatIndicesElasticOlivereJSONIterator-12           10000            810098 ns/op           52544 B/op         77 allocs/op
BenchmarkCatIndicesElasticOlivereNaive-12                  10000            708428 ns/op           52694 B/op         79 allocs/op
BenchmarkCatIndicesElasticOlivereFastHTTP-12               10000            626468 ns/op            5692 B/op         51 allocs/op
```

#### [Official ElasticSearch](https://github.com/elastic/go-elasticsearch)

```
BenchmarkCatIndicesOfficialNaive-12                        10000            706032 ns/op            2749 B/op         34 allocs/op
BenchmarkCatIndicesOfficialFastHTTP-12                     10000            592768 ns/op            1601 B/op         22 allocs/op
```

### Perform match all requests

`go test -bench=BenchmarkMatchAll -benchtime=10000x -timeout=100m`

#### [Elastic Olivere](https://github.com/olivere/elastic)

```
BenchmarkMatchAllElasticOlivereJSONIterator-12             10000            680579 ns/op           54703 B/op        116 allocs/op
BenchmarkMatchAllElasticOlivereNaive-12                    10000            684675 ns/op           55006 B/op        120 allocs/op
BenchmarkMatchAllElasticOlivereFastHTTP-12                 10000            561623 ns/op            8040 B/op         88 allocs/op
```

#### [Official ElasticSearch](https://github.com/elastic/go-elasticsearch)

```
BenchmarkMatchAllOfficialNaive-12                          10000            654073 ns/op           52089 B/op         92 allocs/op
BenchmarkMatchAllOfficialFastHTTP-12                       10000            550271 ns/op            5137 B/op         61 allocs/op
BenchmarkMatchAllOfficialJSONIterator-12                   10000            653290 ns/op           51986 B/op         99 allocs/op
```

## Reproduce

1. To start up a local ElacticSearch use `docker-compose up`
2. To run all benchmarks use `go test -bench=. -benchtime=10000x -timeout=100m`