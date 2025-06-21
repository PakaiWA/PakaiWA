# pakaiwa
[![Go Reference](https://pkg.go.dev/badge/github.com/pakaiwa/pakaiwa.svg)](https://pkg.go.dev/github.com/pakaiwa/pakaiwa)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FPakaiWA%2FPakaiWA.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FPakaiWA%2FPakaiWA?ref=badge_shield)

pakaiwa is a Go library for the WhatsApp web multidevice API.

For questions about the WhatsApp protocol (like how to send a specific type of message), you can also use the [WhatsApp protocol Q&A] section on GitHub discussions.

```shell
protoc -I proto -I proto/waCommon  --go_out=./proto/ --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*/*.proto
```
[WhatsApp protocol Q&A]: https://github.com/pakaiwa/pakaiwa/discussions/categories/whatsapp-protocol-q-a

## Features
Most core features are already present:

* Sending messages to private chats and groups (both text and media)
* Receiving all messages
* Managing groups and receiving group change events
* Joining via invite messages, using and creating invite links
* Sending and receiving typing notifications
* Sending and receiving delivery and read receipts
* Reading and writing app state (contact list, chat pin/mute status, etc)
* Sending and handling retry receipts if message decryption fails
* Sending status messages (experimental, may not work for large contact lists)

Things that are not yet implemented:

* Sending broadcast list messages (this is not supported on WhatsApp web either)
* Calls

***Original Repo: [https://github.com/tulir/whatsmeow](https://github.com/tulir/whatsmeow)***

## 📋 Credit & Attribution

| Komponen / File                              | Asal / Kontributor Asli                                                    | Status Lisensi | Modifikasi oleh | Catatan            | commit id |
| -------------------------------------------- | -------------------------------------------------------------------------- | -------------- | --------------- | ------------------ | --------- |
| `proto`                                      | [repo asli](https://github.com/tulir/whatsmeow)                            | MPL-2.0        | @KAnggara75     | Tidak dimodifikasi |           |
| `util/dbutil/connlog.go`                     | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/database.go`                    | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/deadlock_test.go`               | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/iter.go`                        | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/json.go`                        | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/litestream/nocgo.go`            | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/litestream/register.go`         | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/litestream/register_notrace.go` | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/litestream/register_trace.go`   | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/log.go`                         | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/massinsert.go`                  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/massinsert_test.go`             | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/queryhelper.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/reflectscan.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/01-sample.sql`          | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/04-notxn.sql`           | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/05-compat.sql`          | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/01-postgres.sql` | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/01-sqlite3.sql`  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/04-postgres.sql` | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/04-sqlite3.sql`  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/05-postgres.sql` | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/samples/output/05-sqlite3.sql`  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/transaction.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/upgrades.go`                    | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/upgrades_test.go`               | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/upgradetable.go`                | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/dbutil/upgradetable_test.g`            | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/dbutil)     | MPL-2.0        | -               | Tidak dimodifikasi | 569609e   |
| `util/random/bytes.go`                       | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/random)     | MPL-2.0        | -               | Tidak dimodifikasi | 199bf3e   |
| `util/random/string.go`                      | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/random)     | MPL-2.0        | -               | Tidak dimodifikasi | 199bf3e   |
| `util/random/string_test.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/random)     | MPL-2.0        | -               | Tidak dimodifikasi | 199bf3e   |
| `util/exzerolog/callermarshal.go`            | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exzerolog)  | MPL-2.0        | -               | Tidak dimodifikasi | 32294da   |
| `util/exzerolog/defaults.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exzerolog)  | MPL-2.0        | -               | Tidak dimodifikasi | 32294da   |
| `util/exzerolog/generics.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exzerolog)  | MPL-2.0        | -               | Tidak dimodifikasi | 32294da   |
| `util/exzerolog/writer.go`                   | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exzerolog)  | MPL-2.0        | -               | Tidak dimodifikasi | 32294da   |
| `util/exerrors/dualerror.go`                 | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exerrors)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exerrors/must.go`                      | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exerrors)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exsync/event.go`                       | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exsync)     | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exsync/returnonce.go`                  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exsync)     | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exsync/ringbuffer.go`                  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exsync)     | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exsync/syncmap.go`                     | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exsync)     | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/exsync/syncset.go`                     | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/exsync)     | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/fallocate/doc.go`                      | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/fallocate)  | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/fallocate/fallocate_darwin.go`         | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/fallocate)  | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/fallocate/fallocate_linux.go`          | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/fallocate)  | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/fallocate/fallocate_unknown.go`        | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/fallocate)  | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/jsontime/duration.go`                  | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/jsontime)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/jsontime/helpers.go`                   | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/jsontime)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/jsontime/integer.go`                   | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/jsontime)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/jsontime/string.go`                    | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/jsontime)   | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/ptr/ptr.go`                            | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/ptr)        | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/retryafter/retryafter.go`              | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/retryafter) | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |
| `util/retryafter/retryafter_test.go`         | [mautrix/go-util](https://github.com/mautrix/go-util/tree/main/retryafter) | MPL-2.0        | -               | Tidak dimodifikasi | 1de030f   |


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FPakaiWA%2FPakaiWA.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FPakaiWA%2FPakaiWA?ref=badge_large)