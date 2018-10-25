/*
Package commands contains helpers to create stateful tests based on commands.

Testers have to implement the Commands interface providing generators for the
initial state and the commands. For convenience testers may also use the
ProtoCommands as prototype.

The commands themselves have to implement the Command interface, whereas
testers might choose to use ProtoCommand as prototype.
*/
package commands
