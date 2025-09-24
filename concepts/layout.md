# Go Project Structure: A Guide for Developers

## The Go Way: Principles of Organisation

Go's built-in tools and community conventions provide a solid foundation for organising projects.

### Go Modules and Packages

`go.mod` is the heart of modern Go development. It defines a module, which is a collection of related packages that are versioned together. This file is critical for managing dependencies and allowing other projects to import your code.

In Go, a directory name typically becomes the package name. This simple rule encourages developers to think about logical grouping rather than architectural layering. A directory named `user` should contain all the code related to the `user` entity, regardless of whether it’s a model, a repository, or a service.

This is a stark contrast to how logical grouping is handled in languages like C#, which uses namespaces. A C# namespace can be spread across multiple files and even multiple assemblies (DLLs), and it does not need to have a one-to-one relationship with the physical directory structure. A developer can put files for the `MyApp.BusinessLogic` namespace in a variety of folders. Go, on the other hand, forces a strict one-to-one mapping: all files in a single directory must belong to the same package. This constraint is what drives the Go community's preference for a flatter project structure.

Beyond just a logical grouping, Go packages are also a unit of abstraction. Unlike C#, it is idiomatic to use a package name to represent a business entity or a specific concern. For example, a business entity like a `Person` would be naturally represented by a package named `person` that contains all the related types and functions, rather than a `Person` struct within a larger `entities` package. While Go's tooling does not enforce this convention, idiomatic Go strongly recommends this approach to keep packages small, cohesive, and focused on a single responsibility.

### The `cmd` and `internal` Directories

Go's official documentation and the broader community recommend a specific directory layout to solve common problems:

* `cmd/`: This directory should contain the source code for each of your main applications or executables. For example, if your project has a web server and a background worker, they would live in cmd/web and cmd/worker, respectively. This clearly separates your application entry points (which contain main packages) from the reusable library code.

* `internal/`: This is a special directory used for private packages that cannot be imported by other repositories. It's the primary mechanism for enforcing encapsulation. Any code you don’t want to be part of your public API or shared with other modules should go here.

It is also important to note what to avoid: a top-level `src/` directory. This is a common pattern in other languages but is considered an anti-pattern in Go, as the entire Go workspace is already a source-code directory. Likewise, a top-level `pkg/` folder is often unnecessary and can become a dumping ground for packages.

## Monorepo vs. Multi-repo

The choice between a monorepo and a multi-repo depends on your organisation's needs. To make an informed decision, it helps to first understand the two primary types of Go projects: **service-oriented** and **library-oriented**.

A **service-oriented** Go project is a complete, runnable application. It typically includes one or more `main` packages, each serving as an entry point for a specific service or executable. These `main` packages are always organised within a `cmd` folder. This type of project also uses the `internal` folder to house shared packages that are not intended to be imported by modules outside of the repository.

A **library-oriented** Go project, on the other hand, is a collection of reusable code. It contains no `main` packages. Instead, its artefacts (functions, types, etc.) are exported for use by other Go modules.

With these definitions in mind, here is how the two repository patterns are used.

### Multi-repo Organisation

With a multi-repo approach, each repository is a self-contained Go module. This is the simplest and most common setup for small to medium-sized projects or for a microservices architecture. Each repository will typically follow the standard project layout, with a go.mod file at the root.

* my-service-a/  
  * go.mod  
  * cmd/  
    * api/ (main package)  
  * internal/  
    * handlers/  
    * models/  
  * ...other library packages

### Monorepo Organisation

A monorepo holds the code for all your services and applications in a single repository. It's an ideal structure when you need to manage a mix of service-oriented and library-oriented projects in one place.

A key consideration for a monorepo is whether to have one top-level module or multiple modules. A powerful approach is to use multiple go.mod files, with each sub-directory representing a self-contained Go module.

* my-mono-repo/  
  * kycapp/ (Service-oriented module)  
    * go.mod  
    * cmd/  
      * onboardsvc/ (main package)  
      * crecksvc/ (main package)  
    * internal/  
      * person/ (internal packages for kycapp)  
      * company/  
  * authlib/ (Library-oriented module)  
    * go.mod  
  * verifylib/ (Library-oriented module)  
    * go.mod

In this structure, each service and library has its own independent go.mod file, allowing for separate dependency management and versioning. This provides greater flexibility and autonomy for each component within the monorepo, while still centralising the code in a single repository.