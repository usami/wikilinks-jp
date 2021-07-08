import argparse
from src.linker.linker import Linker


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "category",
        help="Specify category [Airport City Company Compound Person]"
    )
    parser.add_argument(
        "annotation_file",
        help="Specify annotation json filepath"
    )
    parser.add_argument(
        "html_dir",
        help="Specify base directory path that cotains html files"
    )
    parser.add_argument(
        "title_pageid_file",
        help="Specify title2pageid.json filepath"
    )
    parser.add_argument(
        "output_file",
        help="Specify output filepath"
    )
    args = parser.parse_args()

    l = Linker(args.category)
    l.load(args.annotation_file, args.html_dir, args.title_pageid_file)

    l.run()

    l.output(args.output_file)


if __name__ == "__main__":
    main()
