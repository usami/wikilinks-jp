# Wikilinks JP

A Wikipedia entity linker for [SHINRA2021-LinkJP Task](http://shinra-project.info/shinra2021linkjp/).

This uses a simple, existing link based approach for linking.

## Usage

Clone this repository first.
```
$ git clone https://github.com/usami/wikilinks-jp.git
```

Then run `./do link-sample` downloads the sample data, builds the linker and run the linker for the downloaded sample data.

```
$ ./do link-sample
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

The linker can be used as a CLI tool.

```
Usage: ./bin/linker [category] [annotation-file] [html-dir] [title-pageid-file] [output-file]
```
