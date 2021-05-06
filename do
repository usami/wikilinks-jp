#!/usr/bin/env bash
set -eu -o pipefail

SAMPLE_DATE=210408

help_download_sample="Download SHINRA2021-LinkJP task sample data."
download-sample() {
  mkdir -p tmp
  mkdir -p data

  rm -rf "data/linkjp-sample-${SAMPLE_DATE}"

  local fileid=1rH-0L2E7Cxd8JIhss6AL1RZdPkSJLyJ-
  local out=tmp/linkjp-sample.zip
  curl -L -o ${out} "https://drive.google.com/uc?export=download&id=${fileid}"
  unzip ${out} -d data/
}

help_install_gotestsum="Install gotestsum."
install-gotestsum() {
	go get gotest.tools/gotestsum
}

help_test="Run all tests."
test() {
	gotestsum -- ${GOTESTFLAGS:-} "${@:-./...}"
}

help_build="Build Wikilinks-jp linker."
build() {
  root=$(pwd)

  mkdir -p bin

  cd cmd/linker
  go build

  mv linker $root/bin
  cd $root
}

help_link_sample="Run Wikilinks-jp linker for sample data."
link-sample() {
  local sample_dir="data/linkjp-sample-${SAMPLE_DATE}"
  local output_dir=output/sample
  local title_pageid_json=data/jawiki-20190120-title2pageid.json

  [ ! -d "${sample_dir}" ] && download-sample
  [ ! -f "${title_pageid_json}" ] && gunzip ${title_pageid_json}.gz

  build

  rm -rf $output_dir
  mkdir -p $output_dir

  for cat in airport city company compound person
	do
		local cat_title=$(title-case $cat)
		local annotation_json=${sample_dir}/ene_annotation/${cat_title}.json
		local html_dir=${sample_dir}/html/${cat_title}
		bin/linker $cat $annotation_json $html_dir $title_pageid_json $output_dir/$cat.json
	done
}

title-case() {
	echo $1 | sed 's/.*/\L&/; s/[a-z]*/\u&/g'
}

list() {
  declare -F | awk '{print $3}'
}

help_help="Print help text, or detailed help for a task."
help() {
	local item="${1-}"
	if [ -n "${item}" ]; then
		local help_name="help_${item//-/_}"
		echo "${!help_name-}"
		return
	fi

	type -t help-text-intro > /dev/null && help-text-intro
	for item in $(list); do
		local help_name="help_${item//-/_}"
		local text="${!help_name-}"
		[ -n "$text" ] && printf "%-20s\t%s\n" $item "$(echo "$text" | head -1)"
	done
}

case "${1-}" in
  list) list;;
  ""|"help") help "${2-}";;
  *) "$@";;
esac
