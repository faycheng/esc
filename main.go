package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"esc/embed"
)

func main() {
	conf := &embed.Config{
		Invocation: strings.Join(os.Args[1:], " "),
	}

	flag.StringVar(&conf.OutputFile, "o", "", "Output file, else stdout.")
	flag.StringVar(&conf.Package, "pkg", "main", "Package.")
	flag.StringVar(&conf.Prefix, "prefix", "", "Prefix to strip from filesnames.")
	flag.StringVar(&conf.Ignore, "ignore", "", "Regexp for files we should ignore (for example \\\\.DS_Store).")
	flag.StringVar(&conf.Include, "include", "", "Regexp for files to include. Only files that match will be included.")
	flag.StringVar(&conf.ModTime, "modtime", "", "Unix timestamp to override as modification time for all files.")
	flag.BoolVar(&conf.Private, "private", false, "If true, do not export autogenerated functions.")
	flag.BoolVar(&conf.NoCompression, "no-compress", false, "If true, do not compress files.")
	flag.Parse()
	conf.Files = flag.Args()

	var err error
	out := os.Stdout
	if conf.OutputFile != "" {
		if out, err = os.Create(conf.OutputFile); err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}
	if err = embed.Run(conf, out); err != nil {
		log.Fatal(err)
	}
}
