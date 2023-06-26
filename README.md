# Grevolt

A low-level library for Revolt focused on being up-to-date and feature-complete with Revolts current API while also being well tested with unit tests and providing a high level of control over all parts of the library to allow both small and large bots to thrive.

Grevolt seperates low-level ``Client`` operations from more higher level abstractions by providing ``extra``, a collection of additional utilities for common tasks such as more complex event handling, command handling (TODO) and more!

## Why grevolt?

Other libraries had many flaws when inspected. Not to blame these projects, but here are some of the issues I found.

- [Panic's on connection failures](https://github.com/sentinelb51/revoltgo/blob/6690504750626ba063fb10a6ca86ceb5e4e57111/websocket.go#L325) and when [failing to read a message](https://github.com/sentinelb51/revoltgo/blob/6690504750626ba063fb10a6ca86ceb5e4e57111/websocket.go#L95C4-L95C9) or when (different library) [json marshalling](https://github.com/ben-forster/revolt/blob/main/websocket.go#L38)

Revolt goes down a *lot* and in the future may even spew out invalid JSON as it grows, and panicking here is simply infeasible for larger bots. Grevolt instead returns an error and allows the user to handle it for example by restarting the bot or logging the error in a database for example. In addition to this, these libraries often don't `Ping` the 'official client' way (`data` is supposed to be a timestamp and not a constant)

Also, panicking on failing to read a message isn't ideal and libraries should aim to try and restart the websocket instead of outright panicking.

- Poor concurrency, such as [using maps without a mutex](https://github.com/sentinelb51/revoltgo/blob/main/state.go) and generally lower quality code (including hardcoded JSONs in some libraries such as [this library](https://github.com/ben-forster/revolt).)

Concurrent map writes in Go are not allowed making such libraries highly infeasible for large bots.

## Testing

Run ``go test -v ./...`` to test stuff

## TODO

- Better usage examples
- Command handling framework

## Credits

- https://github.com/sentinelb51/revoltgo for some models (embeds)

**Heavy work in progress**