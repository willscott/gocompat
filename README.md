# gocompat

gocompat is a basic check for whether two go projects use compatible versions
of dependent libraries.

The specific use case is for development of plugins, where the versions
of libraries need to match for the plugin to be loaded at runtime.
gocompat can be used in the primary code base to get alerted on PRs if
they will cause incompatibilities between the project and external
dependencies used by plugins.
