package main

import (
	"fmt"
	"os"

	"abcd-optics/internal/element"
	"abcd-optics/internal/system"
)

const usage = `abcd-optics: paraxial ABCD ray-transfer accounting.

Reads an optical element sequence from a JSON spec, composes the 2x2 ray
transfer matrix M = M_n ... M_1, and reports the system matrix, the effective
focal length, and for a given object distance the image distance and lateral
magnification.

usage:
  abcd-optics trace <spec.json>
  abcd-optics telescope <f1> <f2> <spacing>
  abcd-optics help

commands:
  trace       compose the system from spec.json and report imaging
  telescope   analyze a two thin-lens afocal combination
  help        show this message

spec.json fields:
  name             optional label
  object_distance  object plane distance in front of the first element
  elements         ordered element list; first element is met first:
                     {"kind":"space",      "length":L}
                     {"kind":"thin_lens",  "focal":f}
                     {"kind":"refraction", "radius":R, "n1":n1, "n2":n2}

Validation:
  negative propagation length, non-positive refractive index, zero focal
  length and unknown element kinds are rejected with a non-zero exit code.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "trace":
		if err := runTrace(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "abcd-optics: %v\n", err)
			os.Exit(1)
		}
	case "telescope":
		if err := runTelescope(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "abcd-optics: %v\n", err)
			os.Exit(1)
		}
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "abcd-optics: unknown command %q\n\n%s", os.Args[1], usage)
		os.Exit(2)
	}
}

func runTrace(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("trace needs exactly one spec file")
	}
	spec, err := element.LoadSpecFile(args[0])
	if err != nil {
		return err
	}
	report, err := system.BuildReport(spec)
	if err != nil {
		return err
	}
	fmt.Print(report.String())
	return nil
}

func runTelescope(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("telescope needs f1, f2 and spacing")
	}
	var f1, f2, spacing float64
	if _, err := fmt.Sscanf(args[0], "%g", &f1); err != nil {
		return fmt.Errorf("f1: %w", err)
	}
	if _, err := fmt.Sscanf(args[1], "%g", &f2); err != nil {
		return fmt.Errorf("f2: %w", err)
	}
	if _, err := fmt.Sscanf(args[2], "%g", &spacing); err != nil {
		return fmt.Errorf("spacing: %w", err)
	}
	t, err := system.AnalyzeTelescope(f1, f2, spacing)
	if err != nil {
		return err
	}
	fmt.Print(t.String())
	return nil
}
