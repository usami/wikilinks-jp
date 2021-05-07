# Wikilinks JP

Article link based Wikification linker for [SHINRA2021-LinkJP Task](http://shinra-project.info/shinra2021linkjp/).

This uses a simple, existing link based approach for linking.

[![usami](https://circleci.com/gh/usami/wikilinks-jp.svg?style=svg)](https://app.circleci.com/pipelines/github/usami/wikilinks-jp)

## Usage

Clone this repository first.
```
$ git clone https://github.com/usami/wikilinks-jp.git
```

Then run `./do link-sample` downloads the sample data, builds the linker and run the linker against the sample data.

```
$ ./do link-sample
...
2021/05/07 10:47:02 linker[airport]: load annotaions
2021/05/07 10:47:02 linker[airport]: load pages
2021/05/07 10:47:02 linker[airport]: load title to pageid mappings
2021/05/07 10:47:08 linker[airport]: check links
2021/05/07 10:47:08 linker[airport]: output analyzed results
...
2021/05/07 10:47:23 linker[person]: load annotaions
2021/05/07 10:47:23 linker[person]: load pages
2021/05/07 10:47:23 linker[person]: load title to pageid mappings
2021/05/07 10:47:28 linker[person]: check links
2021/05/07 10:47:28 linker[person]: output analyzed results
```

Example outputs can be found under `output/sample`.

```
$ ls output/sample/
airport.json  city.json  company.json  compound.json  person.json
```

The linker can be used as a CLI tool.

```
Usage: ./bin/linker [category] [annotation-file] [html-dir] [title-pageid-file] [output-file]
```

## Requirements

### Softwares

The linker and `do` script assumes the following commands are installed:

- go (1.16)
- curl
- unzip
- gunzip

### Files

The linker requires Wikipedia title to pageid mapping file. A mapping file is bandled with this repo (`data/jawiki-20190120-title2pageid.json.gz`). You can download the latest version [here](https://drive.google.com/drive/folders/1ncZnWgDPFuoKQyqAVIaDnnx85sjsW5cN?usp=sharing).
