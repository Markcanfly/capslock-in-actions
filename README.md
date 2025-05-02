# Capslock in Action(s)

[Capslock](https://github.com/google/capslock) is a tool for static analysis of golang source files to determine what capabilities they use; capabilities like **network access**, **filesystem access** and others that may be sensitive for your usecase.

The main goal of using Capslock in Continous Integration (CI) processes is twofold:

- Raise questions and warrant explanation when a capability gets added/removed from the source code
- Raise an alert when a dependency adds or removes capabilities.

The goal of this repository is to experiment with Capslock in the GitHub Actions automation tool and demonstrate what Capslock can do.
