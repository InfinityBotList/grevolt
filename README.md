# Grevolt

A low-level library for Revolt focused on being up-to-date and feature-complete with Revolts current API while also being well tested with unit tests and providing a high level of control over all parts of the library to allow both small and large bots to thrive.

Grevolt seperates low-level ``Client`` operations from more higher level abstractions by providing ``extra``, a collection of additional utilities for common tasks such as more complex event handling, command handling (TODO) and more!

## Why grevolt?

Other libraries had many flaws when inspected. Not to blame these projects, but here are some of the issues I found.

- [Panicking on unsuccessful json marshalling](https://github.com/ben-forster/revolt/blob/main/websocket.go#L38)

Revolt goes down a *lot* and in the future may even spew out invalid JSON as it grows, and panicking here is simply infeasible for larger bots. Grevolt instead returns an error and allows the user to handle it for example by restarting the bot or logging the error in a database for example. In addition to this, these libraries often don't `Ping` the 'official client' way (`data` is supposed to be a timestamp and [not a constant](https://github.com/sentinelb51/revoltgo/blob/main/websocket.go#L115))

- Poor concurrency, such as [blocking calls](https://github.com/ben-forster/revolt/blob/main/websocket.go#L46) and generally lower quality code (including [hardcoded JSONs](https://github.com/ben-forster/revolt/blob/main/websocket.go#L59C3-L59C3) in some libraries such as [this library](https://github.com/ben-forster/revolt).)

## Testing

Run ``go test -v ./...`` to test stuff

## TODO

- Better usage examples
- Command handling framework

## Credits

- https://github.com/sentinelb51/revoltgo for some models (embeds), utilities (merge), and inspirations (permissions; they're beautiful, aren't they?)

**Heavy work in progress**
