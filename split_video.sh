#!/bin/zsh

# Check if the correct number of arguments are provided
if [ "$#" -ne 4 ]; then
    echo "Usage: $0 <input_video> <initial_duration> <min_increment> <max_increment>"
    exit 1
fi

# Input video file from the first argument
input_video=$1

# Initial duration in seconds from the second argument
duration=$2

# Range for random increment
min_increment=$3
max_increment=$4

# Get the total duration of the video in seconds
total_duration=$(ffmpeg -i "$input_video" 2>&1 | grep "Duration" | awk '{print $2}' | tr -d , | awk -F: '{ print ($1 * 3600) + ($2 * 60) + $3 }' | cut -d. -f1)

if [ -z "$total_duration" ]; then
    echo "Could not determine video duration."
    exit 1
fi

echo "Total video duration: $total_duration seconds"

# Initial start time
start_time=0
clip_number=0

while [ "$start_time" -lt "$total_duration" ]; do
    output_video="clip_${clip_number}.mp4"

    ffmpeg -i "$input_video" -ss "$start_time" -t "$duration" -c copy "$output_video"

    if [ $? -ne 0 ]; then
        echo "An error occurred or end of video reached, stopping the loop."
        break
    fi

    # Update start time and duration for the next clip
    start_time=$((start_time + duration))
    duration=$((duration + RANDOM % (max_increment - min_increment + 1) + min_increment))
    clip_number=$((clip_number + 1))
done

echo "Script completed."

