package stdlib

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// Mesh holds the unrolled, flat arrays optimized for glDrawArrays
type Mesh struct {
	Vertices []float32
	Normals  []float32
	UVs      []float32
	VBO      uint32 // If using modern GL
}

// LoadOBJ loads an OBJ file and unrolls faces into flat float32 slices for fast rendering
func LoadOBJ(filepath string) (*Mesh, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tempVertices [][3]float32
	var tempNormals [][3]float32
	var tempUVs [][2]float32

	// Pre-allocate to reduce reallocation overhead during parsing
	tempVertices = make([][3]float32, 0, 1024)
	tempNormals = make([][3]float32, 0, 1024)
	tempUVs = make([][2]float32, 0, 1024)

	mesh := &Mesh{
		Vertices: make([]float32, 0, 3072),
		Normals:  make([]float32, 0, 3072),
		UVs:      make([]float32, 0, 2048),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := bytes.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch string(parts[0]) {
		case "v":
			x, _ := strconv.ParseFloat(string(parts[1]), 32)
			y, _ := strconv.ParseFloat(string(parts[2]), 32)
			z, _ := strconv.ParseFloat(string(parts[3]), 32)
			tempVertices = append(tempVertices, [3]float32{float32(x), float32(y), float32(z)})
		case "vt":
			u, _ := strconv.ParseFloat(string(parts[1]), 32)
			v, _ := strconv.ParseFloat(string(parts[2]), 32)
			tempUVs = append(tempUVs, [2]float32{float32(u), float32(v)})
		case "vn":
			nx, _ := strconv.ParseFloat(string(parts[1]), 32)
			ny, _ := strconv.ParseFloat(string(parts[2]), 32)
			nz, _ := strconv.ParseFloat(string(parts[3]), 32)
			tempNormals = append(tempNormals, [3]float32{float32(nx), float32(ny), float32(nz)})
		case "f":
			// Triangulate on the fly (basic fan triangulation for quads/ngons)
			for i := 2; i < len(parts)-1; i++ {
				parseFaceVertex(parts[1], &tempVertices, &tempUVs, &tempNormals, mesh)
				parseFaceVertex(parts[i], &tempVertices, &tempUVs, &tempNormals, mesh)
				parseFaceVertex(parts[i+1], &tempVertices, &tempUVs, &tempNormals, mesh)
			}
		}
	}

	return mesh, scanner.Err()
}

func parseFaceVertex(faceData []byte, tempV *[][3]float32, tempUV *[][2]float32, tempN *[][3]float32, mesh *Mesh) {
	indices := bytes.Split(faceData, []byte("/"))
	
	// Vertex Index (1-based in OBJ)
	vIdx, _ := strconv.Atoi(string(indices[0]))
	if vIdx > 0 && vIdx <= len(*tempV) {
		v := (*tempV)[vIdx-1]
		mesh.Vertices = append(mesh.Vertices, v[0], v[1], v[2])
	}

	// UV Index
	if len(indices) > 1 && len(indices[1]) > 0 {
		uvIdx, _ := strconv.Atoi(string(indices[1]))
		if uvIdx > 0 && uvIdx <= len(*tempUV) {
			uv := (*tempUV)[uvIdx-1]
			mesh.UVs = append(mesh.UVs, uv[0], uv[1])
		}
	}

	// Normal Index
	if len(indices) > 2 && len(indices[2]) > 0 {
		nIdx, _ := strconv.Atoi(string(indices[2]))
		if nIdx > 0 && nIdx <= len(*tempN) {
			n := (*tempN)[nIdx-1]
			mesh.Normals = append(mesh.Normals, n[0], n[1], n[2])
		}
	}
}