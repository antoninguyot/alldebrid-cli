# alldebrid-cli

A simple command line interface for [Alldebrid](https://alldebrid.com/).

## Installation

Fetch the binary from the [releases page](releases) and put it somewhere in your `$PATH`.

## Usage

The first thing you'll want to do is to login to your Alldebrid account:

```bash
alldebrid auth
```

This will open a web browser and store the resulting token in `$XDG_CONFIG/alldebrid-cli/config.yaml`.

You can then use the `magnets upload` command to upload a magnet link to Alldebrid:

```bash
# Upload a magnet link
alldebrid magnets upload 'magnet:?xt=urn:btih:...'

# Upload a .torrent file
alldebrid magnets upload 'path/to/file.torrent'
```

You can also use the `links unlock` command to unlock a link:

```bash
alldebrid links unlock 'https://uptobox.com/...'
```

You can view the full documentation of the CLI in the [docs/](docs/) directory.
