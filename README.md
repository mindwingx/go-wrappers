## Golang Wrappers Collection

This guide introduces a curated set of Golang wrapped drivers, each packaged as third-party modules to help manage them
efficiently using an abstraction layer.
This modular approach facilitates handling features, issues, bugs, and more concurrently while seamlessly integrating
them into multiple microservices.

All the packages in this collection conform to the interfaces provided by the ABSTRACTION module.
It serves as both an entry point and a core service functionality hub.

- **Note**: For implementation best practices, refer to the `example.go` file.

### Packages

##### Abstraction

- This package acts as the foundation for the other packages. By utilizing this module, you can employ interfaces to
  initialize microservices, thereby enhancing flexibility and enabling effortless modifications to related modules.

```go
go get github.com/mindwingx/abstraction
```

##### Registry

- This module wraps the functionalities of the `viper` package, allowing you to effortlessly parse configuration files
  such as `YML`, `YAML`, `.env`, and more.

```go
go get github.com/mindwingx/go-registry-wrapper
```

##### Locale

- This module encapsulates the capabilities of the `go-i18n/v2` package, facilitating the creation of service message
  translations based on the `JSON` format. Other compatible formats can also be utilized.

```go
go get github.com/mindwingx/go-locale-wrapper
```

##### Cache

- This module wrapped the `redis/v8` package functionalities to cache data.

```go
go get github.com/mindwingx/go-cache-wrapper
```

##### SQL

- This module enhances the usage of the `gorm` package, enabling seamless transitions between different database drivers
  like `mongo-db`, `elastic-search`, and more, in a clear and easy manner.

```go
go get github.com/mindwingx/go-sql-wrapper
```

##### RPC

- This module simplifies the use of the Golang built-in RPC functionalities, making service calls more efficient. To
  access the best practices, refer to the <a href="https://github.com/mindwingx/go-rpc" traget="_blank">go-rpc</a>
  repository.

```go
go get github.com/mindwingx/go-rpc-wrapper
```

##### HTTP

- This module provides wrapped functionalities for popular frameworks like `Gin`, `Fiber`, `Echo`, etc. The current
  version of the module supports the `Gin` framework with the `g` tag/version prefix.

```go
go get github.com/mindwingx/go-http-wrapper@g-0.1.0
```

##### Helper

- This module brings together shared functionalities such as `response` handling, `time` manipulation, and more.

```go
go get github.com/mindwingx/go-helper
```