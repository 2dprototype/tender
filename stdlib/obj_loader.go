package stdlib

import (
	"bufio"
	"bytes"
	_ "image"
	// _ "image/jpeg" // Auto-register JPEG decoder
	// _ "image/png"  // Auto-register PNG decoder
	"os"
	"strconv"
	"unsafe"

	"github.com/2dprototype/tender/v/gl"
)

type Mesh struct {
	Vertices []float32
	Normals  []float32
	UVs      []float32
}

// Shared helper to compile unrolled arrays into an immutable GPU Display List
func compileMeshToDisplayList(mesh *Mesh) uint32 {
	listID := gl.GenLists(1)
	gl.NewList(listID, gl.COMPILE)

	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.VertexPointer(3, gl.FLOAT, 0, unsafe.Pointer(&mesh.Vertices[0]))

	if len(mesh.Normals) > 0 {
		gl.EnableClientState(gl.NORMAL_ARRAY)
		gl.NormalPointer(gl.FLOAT, 0, unsafe.Pointer(&mesh.Normals[0]))
	}

	if len(mesh.UVs) > 0 {
		gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
		gl.TexCoordPointer(2, gl.FLOAT, 0, unsafe.Pointer(&mesh.UVs[0]))
	}

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(mesh.Vertices)/3))

	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.DisableClientState(gl.NORMAL_ARRAY)
	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)

	gl.EndList()
	return listID
}

// ParseOBJ parses raw bytes/string data directly from memory
func ParseOBJ(data []byte) (*Mesh, error) {
	var tempVertices [][3]float32
	var tempNormals [][3]float32
	var tempUVs [][2]float32

	tempVertices = make([][3]float32, 0, 1024)
	tempNormals = make([][3]float32, 0, 1024)
	tempUVs = make([][2]float32, 0, 1024)

	mesh := &Mesh{
		Vertices: make([]float32, 0, 3072),
		Normals:  make([]float32, 0, 3072),
		UVs:      make([]float32, 0, 2048),
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
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
			for i := 2; i < len(parts)-1; i++ {
				parseFaceVertex(parts[1], &tempVertices, &tempUVs, &tempNormals, mesh)
				parseFaceVertex(parts[i], &tempVertices, &tempUVs, &tempNormals, mesh)
				parseFaceVertex(parts[i+1], &tempVertices, &tempUVs, &tempNormals, mesh)
			}
		}
	}
	return mesh, scanner.Err()
}

func LoadOBJ(filepath string) (*Mesh, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return ParseOBJ(file)
}

func parseFaceVertex(faceData []byte, tempV *[][3]float32, tempUV *[][2]float32, tempN *[][3]float32, mesh *Mesh) {
	indices := bytes.Split(faceData, []byte("/"))
	vIdx, _ := strconv.Atoi(string(indices[0]))
	if vIdx > 0 && vIdx <= len(*tempV) {
		v := (*tempV)[vIdx-1]
		mesh.Vertices = append(mesh.Vertices, v[0], v[1], v[2])
	}
	if len(indices) > 1 && len(indices[1]) > 0 {
		uvIdx, _ := strconv.Atoi(string(indices[1]))
		if uvIdx > 0 && uvIdx <= len(*tempUV) {
			uv := (*tempUV)[uvIdx-1]
			mesh.UVs = append(mesh.UVs, uv[0], uv[1])
		}
	}
	if len(indices) > 2 && len(indices[2]) > 0 {
		nIdx, _ := strconv.Atoi(string(indices[2]))
		if nIdx > 0 && nIdx <= len(*tempN) {
			n := (*tempN)[nIdx-1]
			mesh.Normals = append(mesh.Normals, n[0], n[1], n[2])
		}
	}
}