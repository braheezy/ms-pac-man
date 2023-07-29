#!/bin/bash

script_dir=$(realpath "${BASH_SOURCE[0]}")
project_root=$(dirname "$(dirname "${script_dir}")")
input_file=$project_root/internal/assets/sprites/all-characters.png
output_dir=$(dirname "$input_file")
mkdir -p "$output_dir"

# Default value for N (chunk size)
N=${2-32}

# File names to assign to chunks as they are generated sequentially.
# "" means skip.
chunk_files=(
  "mspac_Rwide"
  "mspac_Rnorm"
  "mspac_Rclose"
  "cherry"
  "strawberry"
  "orange"
  "pretzel"
  "apple"
  "pear1"
  "banana1"
  ""
  ""
  "mspac_Lwide"
  "mspac_Lnorm"
  "mspac_Lclose"
  "100pts"
  "200pts"
  "500pts"
  "700pts"
  "1000pts"
  "2000pts"
  "5000pts"
  ""
  ""
  "mspac_Uwide"
  "mspac_Unorm"
  "mspac_Uclose"
  ""
  ""
  ""
  ""
  ""
  "pear2"
  "banana2"
  ""
  ""
  "mspac_Dwide"
  "mspac_Dnorm"
  "mspac_Dclose"
  ""
  ""
  ""
  ""
  ""
  ""
  ""
  ""
  ""
  "blinky_R1"
  "blinky_R2"
  "blinky_L1"
  "blinky_L2"
  "blinky_U1"
  "blinky_U2"
  "blinky_D1"
  "blinky_D2"
  "dead_blue1"
  "dead_blue2"
  "dead_white1"
  "dead_white2"
  "pinky_R1"
  "pinky_R2"
  "pinky_L1"
  "pinky_L2"
  "pinky_U1"
  "pinky_U2"
  "pinky_D1"
  "pinky_D2"
  "eyes_R"
  "eyes_L"
  "eyes_U"
  "eyes_D"
  "inky_R1"
  "inky_R2"
  "inky_L1"
  "inky_L2"
  "inky_U1"
  "inky_U2"
  "inky_D1"
  "inky_D2"
  ""
  ""
  ""
  ""
  "sue_R1"
  "sue_R2"
  "sue_L1"
  "sue_L2"
  "sue_U1"
  "sue_U2"
  "sue_D1"
  "sue_D2"
  ""
  ""
  ""
  ""
  "200"
  "400"
  "800"
  "1600"
  ""
  ""
  ""
  ""
  ""
  ""
  ""
  ""
)

# Function to split the PNG sheet into chunks and save each chunk to a new file
split_and_save_chunks() {
  local input="$1"
  local size="$2"
  local filenames=("${@:3}")

  local width
  width=$(identify -format "%w" "$input")

  local height
  height=$(identify -format "%h" "$input")

  local row
  local col
  local chunk_num=0
  local total_chunks=${#filenames[@]}

  for ((row = 0; row < height / size; row++)); do
    for ((col = 0; col < width / size; col++)); do
      local x=$((col * size))
      local y=$((row * size))

      if ((chunk_num >= total_chunks)); then
        return
      fi

      local output_chunk="${filenames[$chunk_num]}"
      chunk_num=$((chunk_num + 1))
      if [[ "$output_chunk" == "" ]]; then
        continue
      fi
      convert "$input" -crop "${size}x${size}+${x}+${y}" -define colorspace:auto-grayscale=false \
                      "$output_chunk.png"
    done
  done
}


pushd "$output_dir" &>/dev/null || exit
    split_and_save_chunks "$input_file" "$N" "${chunk_files[@]}"
popd "$output_dir" &>/dev/null || exit
