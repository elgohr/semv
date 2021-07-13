# semv

semv (pronounce [zɛmpf]) is a command line tool for working with semantic versions (semver).

## Installation

### Go
```bash
go install github.com/elgohr/semv@latest
```

### Binaries
Find the binaries in the releases.

## Usage

### Increment
```bash
semv increment (--patch/--minor/--major) ${currentVersion}
```
Returns the incremented version.

### Compare
```bash
semv compare ${firstVersion} ${secondVersion}
```

| Result | Description                                     |
| ------ | ----------------------------------------------- |
| -1     | ${firstVersion} is lower than ${secondVersion}  |
| 0      | ${firstVersion} is equal than ${secondVersion}  |
| 1      | ${firstVersion} is higher than ${secondVersion} |

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)