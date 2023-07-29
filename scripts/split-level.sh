#!/bin/bash

usage() {
    echo "Usage: $0 -f <input_file> [-r <start_row>] [-c <start_col>] [-w <sub_rectangle_width>] [-h <sub_rectangle_height>] [-p <output_prefix>]"
}

script_dir=$(realpath "${BASH_SOURCE[0]}")
project_root=$(dirname "$(dirname "${script_dir}")")
input_file=$project_root/internal/assets/levels/level_1.png
f=$(basename "$input_file")
level_name="${f%.*}"
output_dir=$(dirname "$input_file")/${level_name}_tiles
mkdir -p "$output_dir"

split_png() {
    local input_file=$1
    local start_row=${2:-0}
    local start_col=${3:-0}
    local sub_rect_width=${4:-0}
    local sub_rect_height=${5:-0}
    local output_prefix=${6:-"split"}

    # Check if the input file exists
    if [[ ! -f "$input_file" ]]; then
        echo "Error: Input file '$input_file' not found."
        exit 1
    fi

    # If sub_rectangle_width and sub_rectangle_height are not provided, use the entire PNG dimensions
    if [[ $sub_rect_width -eq 0 ]] || [[ $sub_rect_height -eq 0 ]]; then
        read image_width image_height < <(identify -format "%w %h" "$input_file")
        sub_rect_width=$((image_width / chunk_size))
        sub_rect_height=$((image_height / chunk_size))
    fi

    # Calculate the total number of rows and columns for the sub-rectangle
    total_rows=$((sub_rect_height))
    total_columns=$((sub_rect_width))

    # Use ImageMagick's 'convert' command to split the PNG
    for ((row = 0; row < total_rows; row++)); do
        for ((col = 0; col < total_columns; col++)); do
            crop_row=$((start_row + row * chunk_size))
            crop_col=$((start_col + col * chunk_size))
            convert "$input_file" -crop "${chunk_size}x${chunk_size}+${crop_col}+${crop_row}" \
                "${output_prefix}_${row}_${col}.png" 2>/dev/null
        done
    done
}

# Default chunk size
chunk_size=16

# Parse command line options using getopts
while getopts ":f:r:c:w:h:p:" opt; do
    case $opt in
        f) input_file=$OPTARG ;;
        r) start_row=$OPTARG ;;
        c) start_col=$OPTARG ;;
        w) sub_rect_width=$OPTARG ;;
        h) sub_rect_height=$OPTARG ;;
        p) output_prefix=$OPTARG ;;
        \?) echo "Invalid option: -$OPTARG" >&2
            usage
            exit 1 ;;
        :) echo "Option -$OPTARG requires an argument." >&2
           usage
           exit 1 ;;
    esac
done

# Check if the input file is provided
if [ -z "$input_file" ]; then
    echo "Error: Input file not specified."
    usage
    exit 1
fi

pushd "$output_dir" &>/dev/null || exit
    # Call the split_png function with the provided arguments
    split_png "$input_file" "$start_row" "$start_col" "$sub_rect_width" "$sub_rect_height" "$level_name"

    # Detect and remove duplicate PNGs using aHash
    shopt -s nullglob
    duplicates=()
    for file in *; do
        ihash=$(identify -verbose "$file" | grep signature)
        if [[ " ${duplicates[*]} " == *" $ihash "* ]]; then
            rm "$file"
        else
            duplicates+=("$ihash")
        fi
    done
popd "$output_dir" &>/dev/null || exit
