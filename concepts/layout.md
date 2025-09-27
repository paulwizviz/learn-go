# Go Project Structure

Organising a new project is a foundational task in software development, yet it can be a source of significant friction, especially when adopting a new language like Go.

My observation is developers coming from environments with established, hierarchical architectures, such as Java or C#, often find Go's idiomatic approach counterintuitive. This guide will explore the challenges, discuss the Go-native principles, and provide a clear path to building well-structured projects.

## Repositories, Modules and Packages

In Go, code is organised into **packages**, which are the fundamental units of compilation and reuse. A collection of related packages is versioned and distributed as a **module**, which is typically hosted in a source code **repository** (like Git). The Go toolchain also recognises special directory names, like `internal`, which enforces a convention that restricts access to its packages, making them private to that specific module. Understanding how these concepts—packages, modules, repositories, and visibility rules—work together is fundamental to structuring a project idiomatically.

* **Repository:** A repository, typically a Git repository, is where the source code (Go and others) is organised. In the case of Go, the repository is intrinsic because its location determines the canonical module path—the unique name Go uses to find, import, and fetch the code. Here is how:

  * **Module Path Identity:** Every Go module's name (defined in go.mod) is conventionally expected to match the URL where it lives, such as github.com/username/repository-name.

  * **The `go get` Command:** When you run `go get github.com/username/repository-name`, the Go tooling knows precisely where to look because it uses the URL structure to query the version control system (Git, in this case).

  * **Decentralised Fetching:** Unlike systems where a central registry (like npm or Maven Central) handles all code location, Go is designed to resolve dependencies directly from the version control repository itself.

* **Module:** A module is a collection of related Go packages that are versioned together as a single unit. It is defined by a `go.mod` file at the root of the project directory. This file declares the module's path (its unique name), the version of Go it is built with, and its dependencies on other modules. A repository typically contains a single module.

* **Package:** A package is a directory containing one or more Go source files.All files within a directory must belong to the same package, declared using the `package` keyword at the top of each file. A package provides a distinct namespace for its exported identifiers (those starting with a capital letter)and represents a single, coherent unit of functionality. It is also a unit of compilation and abstraction:

  * **Single Output:** The compiler takes all .go files that declare the same package name (e.g., package `rest`) and compiles them into a single artifact. For non-main packages, this artifact is a .a file (an archive) containing the package's compiled code.

  * **Inter-File Visibility:** Because all files in a package are compiled together, all identifiers (functions, types, variables) declared in any file within that package are directly visible to all other files in the same package. There are no separate import statements needed within the package.

  * **Private Scope:** The visibility rules (capital letter for exported,
    lowercase for unexported) apply only at the package boundary. Identifiers
    starting with a lowercase letter are private to the entire package,
    regardless of which file they are declared in.
  
  * **No Header Files:** Unlike languages like C/C++, Go eliminates the need for
    header files. The compiler doesn't need to check definitions across
    separate declaration and implementation files; it just processes everything
    in the package together.

  * **Dependencies are Clean:** When you compile a package, the compiler only needs to look at the exported identifiers of the packages it imports. It doesn't need to re-read or re-parse the source code of the imported packages, further speeding up compilation.

  * **Unit of Abstraction:** In object-oriented languages (like C++ or C#),classes often serve as the unit of abstraction, using keywords like `public` and `private` to hide internal methods and fields. In Go, the package takes on this responsibility. It acts as a cohesive unit that encapsulates its internal workings, providing a clean, stable public interface for the rest of the application to rely on.

* **`internal` directory:** This is a special directory used for private packages that cannot be imported by other repositories or `go.mod`. It's the primary mechanism for enforcing encapsulation. Any code you don’t want to be part of your public API or shared with other modules should go here.

## Directories by Conventions

Go's official documentation and the broader community recommend other ways of
organising source code that are purely by convention. In other words, Go tools do not influence its use like the `internal` folder.

One folder I have observed in many projects is `cmd/`. This folder appears to be used to contain the source code for multiple `main` packages. For example, if your project has a web server and a background worker, they would live in `cmd/web` and `cmd/worker`, respectively.

It is also important to avoid:

* `src/` directory. This is a common pattern in other languages but is
  considered an anti-pattern in Go, as the entire Go workspace is already a
  source-code directory.

* `pkg/` folder is often unnecessary and can become a dumping ground for
  packages.

## Repository Layout

There appears to be no fixed way to lay out a Go project beyond the principal
organising components mentioned previously.

My observation is that there are two factors to consider when it comes to laying out the repository:

* Service vs Library Oriented Module
* Mono vs multiple repositories

### Service vs Library Oriented Module

I've found it helpful to categorise Go modules in one of two ways:

* A **service-oriented module** is what I think of as a runnable application. It contains one or more `main` packages, each serving as an entry point for a specific service or executable. It does not typically provide exportable packages for other modules to consume.
* A **library-oriented module** is a collection of reusable code intended to be imported by other modules. It has no `main` packages.

My observation for a **service-oriented module** with a single `main` package is that the `main.go` file is often found at the root, at the same level as the `go.mod`.

```sh
my-project/
  go.mod
  main.go
```

For projects with multiple executables, I've seen that the convention is to organise the `main` packages within a `cmd` folder (e.g., `cmd/appone`, `cmd/apptwo`). The common wisdom seems to be keeping these `main` packages as small as possible, acting only as entry points that delegate to the actual implementation code, which is often organised under an `internal` folder.

Here's a structure I've noted that reflects this:

```sh
my-project/
  go.mod
  cmd/
    gateway/        // Main package for building gateway executable
      main.go
    datasvc/        // Main package for building data service executable
      main.go
  internal/         // "Library" packages to support the cmd executables
    gatewaycli/     // Package containing implementation for the gateway
      gateway.go
    datasvc/        // Package containing implementation for the data service
      datasvc.go
```

Let's consider the layout of a **library oriented module**. It's worth noting that my understanding of a Go "library" has shifted from my experience with languages like Java or C#. The primary reason is that **Go libraries are distributed as source code, not as pre-compiled binaries.**

This source-based approach has several important implications:

* **Distribution:** When I add a Go library as a dependency, I'm pulling its source code directly from its version control repository. In contrast, Java (using `.jar` files) and C# (using `.dll` files) distribute pre-compiled binary artifacts from a central repository like Maven Central or NuGet.

* **Compilation:** Because Go libraries are source code, they are compiled from scratch by the *consumer* as part of their own application's build process. In the Java/C# world, the library author compiles the code, and I simply link against that binary.

* **Transparency:** The Go approach means I always have the source code for my dependencies, which aids in debugging and understanding. There are no "black box" binaries.

This source-based philosophy seems central to Go's goals of simplicity, transparency, and fast build times.

To illustrate, here is a simple structure for a single-package library:

```sh
my-library/
 go.mod
 mylibrary.go
 README.md
```

And for a multi-package library:

```sh
my-library/
 go.mod
 regression/       // Regression package
   regression.go
 neural/           // Neural network package
   neural.go
```

I've also been looking into libraries that involve code generation, like gRPC. These have a more complex structure. A common pattern I've found is to separate the API definitions (`.proto` files) from the auto-generated Go code.

Here is a typical structure I've observed:

```sh
my-grpc-library/
  go.mod
  proto/v1/
    service.proto          # The gRPC service definition file
  gen/go/v1/
    service.pb.go          # Auto-generated Go code for message types
    service_grpc.pb.go     # Auto-generated Go code for client/server stubs
  client/
    client.go              # An optional, hand-written, user-friendly client wrapper
  README.md
```

My notes on this structure:

* **/proto/v1/service.proto:** This is the source of truth for the API, defining the services and messages.
* **/gen/go/v1/:** This directory holds the Go code generated by the `protoc` compiler. The key takeaway for me is that these files shouldn't be edited manually.
* **/client/:** This optional, hand-written package is a great idea for providing a simpler, idiomatic Go client that wraps the low-level generated stubs, making the library easier to use.

### Monorepo vs. Multi-repo

With a **multi-repo** approach, each repository is a self-contained Go module. A typical layout is to have a repository and `go.mod` organised around the top level of the repository.

Here is an example of a service oriented layout

```sh
/             // root level of https://github.com/user/project
 .git
 go.mod  
 cmd/  
  api/        // container for api main package
 internal/  
  handlers/  
  models/ 
```

Here is an example of a library repository.

```sh
/          // root level of https://github.com/user/ai
 .git
 go.mod
 ai.go
```

In this case, the repository's project name (`ai`) also becomes the package name. In other words, if another Go module were to consume this library, it would be imported with the package name `ai`.

A **monorepo** holds the code for Go-related modules and potentially modules in other languages. If we consider only Go-related code, it would typically be composed of a mix of service and library-oriented modules. Here is an example of a Go-based monorepo.

```sh
/  // Root level of https://github.com/user/project 
  .git
  kycapp/             // KYC microservices oriented module
    go.mod  
    cmd/              // Main package
      onboardsvc/ 
      checksvc/ 
    internal/         // Entities
      person/         
      company/  
  authlib/            // Auth library oriented module
    go.mod  
  verifylib/          // Auth library oriented module  
    go.mod
```

In this structure, each service and library has its own independent `go.mod` file, allowing for separate dependency management and versioning.

## Versioning

The ability to align a build artefact with, or consume a library of, a particular version is a crucial aspect of software development. Go uses semantic versioning in the format `vX.X.X`.

For a standard, single-module repository, creating a version is as simple as adding a Git tag.

Consider this structure:

```sh
/          // Root of https://github.com/user/ai
 .git
 go.mod
 ai.go
```

To release version `v0.0.1` of this module, you would create a Git tag named `v0.0.1`.

However, this approach is not suitable for a monorepo containing multiple modules. Consider this layout:

```sh
/  // Root of https://github.com/user/project
  .git
  kycapp/             // KYC microservices module
    go.mod
    cmd/
      onboardsvc/
      checksvc/
    internal/
      person/
      company/
  authlib/            // Auth library module
    go.mod
  verifylib/          // Verification library module
    go.mod
```

To version `kycapp`, `authlib`, and `verifylib` independently, the Go toolchain uses a convention where the module's directory path prefixes the version tag. This allows each module to have a distinct versioning lifecycle within the same repository.

Applying this to the example structure:

* **For `kycapp`:** A release of version `v1.0.0` requires a Git tag named `kycapp/v1.0.0`.
* **For `authlib`:** A release of version `v1.2.3` requires a Git tag named `authlib/v1.2.3`.
* **For `verifylib`:** An initial release of `v0.1.0` requires a Git tag named `verifylib/v0.1.0`.

The `go get` command understands this format natively. For instance, running `go get github.com/user/project/authlib@v1.2.3` correctly resolves the `authlib/v1.2.3` tag and fetches the specified module version.
