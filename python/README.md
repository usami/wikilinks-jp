# Wikilinks JP (Python version)

This is the Python version of the Go language implementation published by the official.

## Usage

Clone this repository first.
```
$ git clone https://github.com/usami/wikilinks-jp.git
```

Create a virtual environment with pipenv and install the necessary libraries.

```
$ cd wikilinks-jp/
$ pipenv install --dev (or pipenv sync --dev)
```

Then running `./do_python.sh link-sample` downloads the sample data, builds the linker and runs the linker against the sample data.

```
$ cd python/
$ ./do_python.sh link-sample
```

Example outputs can be found under `output/sample`.

```
$ ls output/sample/
airport.json  city.json  company.json  compound.json  person.json
```

The linker can be used as a CLI tool.

```
Usage: python main.py [category] [annotation-file] [html-dir] [title-pageid-file] [output-file]
```

## Test

It is possible to test with pytest.

```
cd python
pytest -v tests/
```

# Requirements

### Softwares

The linker and `do_python.sh` script assumes the following commands are installed:

- python 3.8+
- python libraries (in Pipfile)
- curl
- unzip
- gunzip

### Files

The linker requires Wikipedia title to pageid mapping file. A mapping file is bandled with this repo (`wikilinks-jp/data/jawiki-20190120-title2pageid.json.gz`). You can download the latest version [here](https://drive.google.com/drive/folders/1ncZnWgDPFuoKQyqAVIaDnnx85sjsW5cN?usp=sharing).