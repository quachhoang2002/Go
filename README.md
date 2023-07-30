# TANCA EVENT
## 1. Tools
#### SQLBoiler
- **Used:** for generating database models and queries for the database
- **Installation**: 
  - Following the instructions [here](https://github.com/volatiletech/sqlboiler#download)

    ```bash
    go install github.com/volatiletech/sqlboiler/v4@latest
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
    ```

- **Usage:**: To generate the database models
  - Run `make gen-models` in the *root directory* of the project

#### Mockery
- **Used** for generating mocks for testing
- **Installation**: 
  - Following the instructions [here](https://vektra.github.io/mockery/v2.32/installation/)
  
    ```bash
    go install github.com/vektra/mockery/v2@v2.32.0
    ```
- **Usage:**
  - Add `//go:generate mockery --name=<InterfaceName>` above the interface declaration
  - Run `go generate ./...` in the *root directory* of the project