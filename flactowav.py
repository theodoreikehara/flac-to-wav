import os
import glob
import argparse
from pydub import AudioSegment
from tqdm import tqdm

def convert(input_dir, output_dir):
    # loop through all subdirectories and files in input directory
    for root, dirs, files in os.walk(input_dir):
        for file in files:
            if file.endswith('.flac'):
                # construct full input and output paths
                input_path = os.path.join(root, file)
                output_path = os.path.join(output_dir, os.path.relpath(input_path, input_dir))
                output_path = os.path.splitext(output_path)[0] + '.wav'
    
                # create output directory if it doesn't exist
                os.makedirs(os.path.dirname(output_path), exist_ok=True)
    
                # load FLAC file and convert to WAV
                flac_audio = AudioSegment.from_file(input_path, format='flac')
                flac_audio.export(output_path, format='wav')

def convert_with(input_dir, output_dir):
    # get list of all FLAC files in input directory
    flac_files = [os.path.join(root, file) for root, dirs, files in os.walk(input_dir) for file in files if file.endswith('.flac')]
    
    # loop through all FLAC files and convert to WAV
    for input_path in tqdm(flac_files, desc='Converting'):
        # construct output path
        output_path = os.path.join(output_dir, os.path.relpath(input_path, input_dir))
        output_path = os.path.splitext(output_path)[0] + '.wav'

        # create output directory if it doesn't exist
        os.makedirs(os.path.dirname(output_path), exist_ok=True)

        # load FLAC file and convert to WAV
        flac_audio = AudioSegment.from_file(input_path, format='flac')
        flac_audio.export(output_path, format='wav')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Convert FLAC files to WAV.')
    parser.add_argument('input', help='Input directory or file path')
    parser.add_argument('output', help='Output directory or file path')
    args = parser.parse_args()

    # convert = tqdm(convert)
    convert_with(args.input, args.output)
