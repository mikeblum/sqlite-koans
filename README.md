# SQLite Koans

Companion code for [SQLite Koans](https://mblum.me/2023/01/sqlite-koans/)

This project configures `sqlite3` for real world apps:

✅ Foreign Keys

✅ Read as you Write

✅ Write-Ahead Logging

✅ `UTF-8` encoding

✅ graceful table lock error handling via `PRAGMA busy_timeout;`

The repo has a commit covering each koan that we tie to a particular `PRAGMA` config flag:

[SQLite PRAGMA](https://www.sqlite.org/pragma.html)

## Depends On

[SQLite Version 3.40.1](https://www.sqlite.org/releaselog/3_40_1.html)

> Ubuntu apt version is outdated

## Build SQLite from source

1\. Install Build Tools

> `gcc` is required for `CGO` bindings used by `mattn/go-sqlite3`

`sudo apt-get install build-essential gcc`

1a\. confirm `gcc` installation

`gcc --version`

2\. Install latest SQLite snapshot (~v3.40.1 minimum)

```
export SQLITE_SNAPSHOT=202301131932
wget "https://www.sqlite.org/snapshot/sqlite-snapshot-${SQLITE_SNAPSHOT}.tar.gz"
tar -xzvf sqlite-snapshot-$SQLITE_SNAPSHOT.tar.gz
cd sqlite-snapshot-$SQLITE_SNAPSHOT
./configure
make
sudo make install
```

3\. Verify `sqlite3` install

```
❯ sqlite3
SQLite version 3.40.2 2023-01-13 19:32:19
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite>
```
