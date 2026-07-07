package stdlib

import (
	"bytes"
	"os"
	"strconv"
)

type Material struct {
	Name       string
	Ambient    [4]float32
	Diffuse    [4]float32
	Specular   [4]float32
	Shininess  float32
	DiffuseMap string // Path to the texture image
	TextureID  uint32 // Resolved OpenGL texture ID
}

func ParseMTL(data []byte) (map[string]*Material, error) {
	materials := make(map[string]*Material)
	var currentMat *Material

	lines := bytes.Split(data, []byte{'\n'})
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := bytes.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch string(parts[0]) {
		case "newmtl":
			name := string(parts[1])
			currentMat = &Material{
				Name:    name,
				Ambient: [4]float32{0.2, 0.2, 0.2, 1.0},
				Diffuse: [4]float32{0.8, 0.8, 0.8, 1.0},
			}
			materials[name] = currentMat
		case "Ka":
			if currentMat != nil && len(parts) > 3 {
				currentMat.Ambient[0] = parseFloat(parts[1])
				currentMat.Ambient[1] = parseFloat(parts[2])
				currentMat.Ambient[2] = parseFloat(parts[3])
			}
		case "Kd":
			if currentMat != nil && len(parts) > 3 {
				currentMat.Diffuse[0] = parseFloat(parts[1])
				currentMat.Diffuse[1] = parseFloat(parts[2])
				currentMat.Diffuse[2] = parseFloat(parts[3])
			}
		case "Ks":
			if currentMat != nil && len(parts) > 3 {
				currentMat.Specular[0] = parseFloat(parts[1])
				currentMat.Specular[1] = parseFloat(parts[2])
				currentMat.Specular[2] = parseFloat(parts[3])
			}
		case "Ns":
			if currentMat != nil && len(parts) > 1 {
				currentMat.Shininess = parseFloat(parts[1])
			}
		case "map_Kd":
			if currentMat != nil && len(parts) > 1 {
				// Safely extract the full path even if it contains spaces
				prefix := []byte("map_Kd")
				idx := bytes.Index(line, prefix)
				if idx != -1 {
					currentMat.DiffuseMap = string(bytes.TrimSpace(line[idx+len(prefix):]))
				}
			}
		}
	}
	return materials, nil
}

func LoadMTL(filepath string) (map[string]*Material, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return ParseMTL(data)
}

func parseFloat(b []byte) float32 {
	f, _ := strconv.ParseFloat(string(b), 32)
	return float32(f)
}