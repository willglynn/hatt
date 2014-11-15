package interop

import "path/filepath"

var Formats = []Format{
	&md5sum{},
	&sfv{},
}

func FindByName(name string) Format {
	for _, format := range Formats {
		if format.Name() == name {
			return format
		}
	}

	return nil
}

func FindByExtension(path string) Format {
	// extract the file extension
	ext := filepath.Ext(path)

	if len(ext) > 0 {
		// ext contains a dot, so strip it off
		// this could be smarter...
		return FindByName(ext[1:])
	} else {
		return nil
	}
}
