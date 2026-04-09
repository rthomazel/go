// Copyright 2024-Present R. Thomazella. All rights reserved.
// Use of this source code is governed by the BSD-3-Clause
// license that can be found in the LICENSE file and online
// at https://opensource.org/license/BSD-3-clause.

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const usageText = `copyright: check and fix copyright headers in source files.

usage: copyright -c TOKEN [-f] [-s] [-i PATTERN] GLOB[,GLOB...]

flags:
  -c TOKEN     comment token prepended to each header line (required)
  -f           fix: write header to files missing it
  -s           shebang: preserve first line, insert header after it
  -i PATTERN   extended regexp to ignore paths (default: mock_|/.local/)

reads header text from COPYRIGHT_HEADER env var or a file named
'copyright-header' next to the binary or in the working directory.

exits 1 if any files are missing the header (or on error).
`

func main() {
	fComment := flag.String("c", "", "comment token prepended to each header line")
	fFix := flag.Bool("f", false, "write header to files missing it")
	fShebang := flag.Bool("s", false, "preserve first line, insert header after it")
	fIgnore := flag.String("i", `mock_|/.local/`, "extended regexp to ignore paths")
	flag.Usage = func() { fmt.Fprint(os.Stderr, usageText) }
	flag.Parse()

	globs := flag.Arg(0)
	if globs == "" {
		fmt.Fprint(os.Stderr, "error: glob argument required\n")
		os.Exit(1)
	}
	if *fComment == "" {
		fmt.Fprint(os.Stderr, "error: -c TOKEN is required\n")
		os.Exit(1)
	}

	header, err := loadHeader()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	ignore := regexp.MustCompile(*fIgnore)
	commented := commentHeader(header, *fComment)

	missing, err := findMissing(strings.Split(globs, ","), ignore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if len(missing) == 0 {
		return
	}

	for _, f := range missing {
		fmt.Println(f)
	}

	if !*fFix {
		fmt.Fprintf(os.Stderr, "error: %d file(s) missing copyright header\n", len(missing))
		os.Exit(1)
	}

	for _, f := range missing {
		if err := fixFile(f, commented, *fShebang); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", f, err)
			os.Exit(1)
		}
	}
}

func loadHeader() (string, error) {
	if h := os.Getenv("COPYRIGHT_HEADER"); h != "" {
		return h, nil
	}

	candidates := []string{
		filepath.Join(filepath.Dir(os.Args[0]), "copyright-header"),
		"copyright-header",
	}

	for _, p := range candidates {
		data, err := os.ReadFile(p) //nolint:gosec // intentional: path is controlled
		if err == nil {
			return strings.TrimRight(string(data), "\n"), nil
		}
	}

	return "", errors.New("no header found: set COPYRIGHT_HEADER env var or create a 'copyright-header' file")
}

func commentHeader(header, token string) string {
	var builder strings.Builder

	for line := range strings.SplitSeq(header, "\n") {
		builder.WriteString(token)
		builder.WriteString(line)
		builder.WriteByte('\n')
	}

	return builder.String()
}

func findMissing(globs []string, ignore *regexp.Regexp) ([]string, error) {
	var missing []string

	for _, glob := range globs {
		matches, err := filepath.Glob(glob)
		if err != nil {
			return nil, fmt.Errorf("glob %q: %w", g, err)
		}

		// filepath.Glob only matches in cwd; walk for bare name patterns like *.go
		if len(matches) == 0 {
			matches, err = walkGlob(g)
			if err != nil {
				return nil, err
			}
		}

		for _, path := range matches {
			if ignore.MatchString(path) {
				continue
			}

			ok, err := hasHeader(path)
			if err != nil {
				return nil, fmt.Errorf("checking %s: %w", path, err)
			}

			if !ok {
				missing = append(missing, path)
			}
		}
	}

	return missing, nil
}

// walkGlob finds files matching a basename glob under the current directory.
func walkGlob(g string) ([]string, error) {
	base := filepath.Base(g)
	var out []string

	var walkErr error
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() && strings.HasPrefix(entry.Name(), ".") {
			return filepath.SkipDir
		}

		matched, err := filepath.Match(base, entry.Name())
		if err != nil {
			return err
		}

		if matched {
			out = append(out, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walking directory: %w", err)
	}
	return out, nil
}

func hasHeader(path string) (bool, error) {
	f, err := os.Open(path) //nolint:gosec // intentional: path from find
	if err != nil {
		return false, fmt.Errorf("opening %s: %w", path, err)
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for range 10 {
		if !scanner.Scan() {
			break
		}

		if strings.Contains(scanner.Text(), "Copyright") {
			return true, nil
		}
	}

	if scanErr := scanner.Err(); scanErr != nil {
		return false, fmt.Errorf("scanning %s: %w", path, scanErr)
	}
	return false, nil
}

func fixFile(path, header string, shebang bool) error {
	data, err := os.ReadFile(path) //nolint:gosec // intentional: path from find
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	var newContent string

	if shebang {
		firstNL := strings.Index(string(data), "\n")
		if firstNL == -1 {
			return errors.New("no newline found (cannot detect shebang line)")
		}

		newContent = string(data[:firstNL+1]) + header + "\n" + string(data[firstNL+1:])
	} else {
		newContent = header + "\n" + string(data)
	}

	return os.WriteFile(path, []byte(newContent), 0) //nolint:gosec // intentional: path from find
}
