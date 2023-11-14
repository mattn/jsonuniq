# jsonuniq

Uniq for ndjson.

## Usage

Uniq each lines composed of JSON by the value pointed by jsonpath.

```
cat output.json | jsonuniq -p $.id
```

## Installation

```
go install github.com/mattn/jsonuniq
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
