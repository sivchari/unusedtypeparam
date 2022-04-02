# unusedtypeparam
unusedtypeparam is an analyzer that detects unused type parameter.

## Instruction

```sh
go install github.com/sivchari/unusedtypeparam/cmd/unusedtypeparam@latest
```

## Usage

```go
package a

type C interface {
	string | ~int
}

func ok1[E C](arg E) {
	arg2 := arg
	_ = arg2
	var arg3 E
	_ = arg3
}

func ok2[E C](arg any) {
	var arg3 E
	_ = arg3
}

func ng[E C](arg any) { // want "This func unused type parameter."
	arg2 := arg
	_ = arg2
}
```

```console
go vet -vettool=(which unusedtypeparam) ./...

# command-line-arguments
testdata/src/a/a.go:19:1: This func unused type parameter.
```

## CI

### CircleCI

```yaml
- run:
    name: install unusedtypeparam
    command: go install github.com/sivchari/unusedtypeparam/cmd/unusedtypeparam@latest

- run:
    name: run unusedtypeparam
    command: go vet -vettool=`which unusedtypeparam` ./...
```

### GitHub Actions

```yaml
- name: install unusedtypeparam
  run: go install github.com/sivchari/unusedtypeparam/cmd/unusedtypeparam@latest

- name: run unusedtypeparam
  run: go vet -vettool=`which unusedtypeparam` ./...
```
