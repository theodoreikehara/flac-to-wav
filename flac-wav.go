// converts a path of flac to wav files

package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "github.com/mewkiz/flac"
    // "github.com/mewkiz/flac/meta"
    "github.com/mewkiz/pkg/errutil"
    "github.com/mewkiz/pkg/osutil"
)

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Fprintln(os.Stderr, "Usage: flac-to-wav input.flac")
        os.Exit(1)
    }
    inputPath := flag.Arg(0)
    if !osutil.Exists(inputPath) {
        fmt.Fprintf(os.Stderr, "No such file: %q\n", inputPath)
        os.Exit(1)
    }

    // Open input FLAC file.
    fr, err := os.Open(inputPath)
    if err != nil {
        errutil.Fatal(err)
    }
    defer fr.Close()

    // Create output WAV file.
    outputDir, outputPath := filepath.Split(inputPath)
    outputName := fmt.Sprintf("%s.wav", outputPath[:len(outputPath)-len(filepath.Ext(outputPath))])
    output := filepath.Join(outputDir, outputName)
    fw, err := os.Create(output)
    if err != nil {
        errutil.Fatal(err)
    }
    defer fw.Close()

    // Decode FLAC to WAV.
    dec, err := flac.NewDecoder(fr)
    if err != nil {
        errutil.Fatal(err)
    }
    defer dec.Close()
    info := dec.Info
    enc, err := info.NewEncoder(fw)
    if err != nil {
        errutil.Fatal(err)
    }
    defer enc.Close()
    fmt.Printf("Decoding %q to %q\n", inputPath, output)
    for {
        sample, err := dec.ParseNext()
        if err != nil {
            if err == flac.ErrEndOfStream {
                break
            }
            errutil.Fatal(err)
        }
        if err := enc.Write(sample); err != nil {
            errutil.Fatal(err)
        }
    }
    fmt.Println("Done.")
}

