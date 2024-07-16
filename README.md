# Go source code obfuscator

The goal with this project is to provide a go tool that can take go source code
and obfuscate "user" defined symbols, e.g., variable names, funcation names,
package names, etc.

The resulting source code must be compilable and work identical to the input.

Eventually the tool should be able to obfuscate a whole project as an input.
