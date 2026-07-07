package stdlib

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"unsafe"

	"github.com/2dprototype/tender/v/gl"
)

type MaterialGroup struct {
	MaterialName string
	Vertices     []float32
	Normals      []float32
	UVs          []float32
}

type Mesh struct {
	Groups []*MaterialGroup
}

// compileMeshToDisplayList now bakes material state changes into the GPU list
func compileMeshToDisplayList(mesh *Mesh, materials map[string]*Material) uint32 {
	listID := gl.GenLists(1)
	gl.NewList(listID, gl.COMPILE)

	for _, group := range mesh.Groups {
		if len(group.Vertices) == 0 {
			continue
		}

		// 1. Apply Material State (if materials are provided)
		if materials != nil {
			if mat, ok := materials[group.MaterialName]; ok {
				gl.Materialfv(gl.FRONT_AND_BACK, gl.AMBIENT, &mat.Ambient[0])
				gl.Materialfv(gl.FRONT_AND_BACK, gl.DIFFUSE, &mat.Diffuse[0])
				gl.Materialfv(gl.FRONT_AND_BACK, gl.SPECULAR, &mat.Specular[0])
				gl.Materialf(gl.FRONT_AND_BACK, gl.SHININESS, mat.Shininess)

				if mat.TextureID > 0 {
					gl.Enable(gl.TEXTURE_2D)
					gl.BindTexture(gl.TEXTURE_2D, mat.TextureID)
				} else {
					gl.Disable(gl.TEXTURE_2D)
				}
			}
		}

		// 2. Draw Group via fast array pointers
		gl.EnableClientState(gl.VERTEX_ARRAY)
		gl.VertexPointer(3, gl.FLOAT, 0, unsafe.Pointer(&group.Vertices[0]))

		if len(group.Normals) > 0 {
			gl.EnableClientState(gl.NORMAL_ARRAY)
			gl.NormalPointer(gl.FLOAT, 0, unsafe.Pointer(&group.Normals[0]))
		}

		if len(group.UVs) > 0 {
			gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
			gl.TexCoordPointer(2, gl.FLOAT, 0, unsafe.Pointer(&group.UVs[0]))
		}

		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(group.Vertices)/3))

		gl.DisableClientState(gl.VERTEX_ARRAY)
		gl.DisableClientState(gl.NORMAL_ARRAY)
		gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	}

	gl.EndList()
	return listID
}

func ParseOBJ(data []byte) (*Mesh, error) {
	var tempVertices [][3]float32
	var tempNormals [][3]float32
	var tempUVs [][2]float32

	tempVertices = make([][3]float32, 0, 1024)
	tempNormals = make([][3]float32, 0, 1024)
	tempUVs = make([][2]float32, 0, 1024)

	mesh := &Mesh{Groups: make([]*MaterialGroup, 0)}
	
	// Create a default group for meshes without an MTL file
	currentGroup := &MaterialGroup{MaterialName: "default"}
	mesh.Groups = append(mesh.Groups, currentGroup)

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
		case "usemtl":
			// Switch to a new material group when requested
			matName := string(parts[1])
			currentGroup = &MaterialGroup{MaterialName: matName}
			mesh.Groups = append(mesh.Groups, currentGroup)
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
				parseFaceVertex(parts[1], &tempVertices, &tempUVs, &tempNormals, currentGroup)
				parseFaceVertex(parts[i], &tempVertices, &tempUVs, &tempNormals, currentGroup)
				parseFaceVertex(parts[i+1], &tempVertices, &tempUVs, &tempNormals, currentGroup)
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

func parseFaceVertex(faceData []byte, tempV *[][3]float32, tempUV *[][2]float32, tempN *[][3]float32, group *MaterialGroup) {
	indices := bytes.Split(faceData, []byte("/"))
	vIdx, _ := strconv.Atoi(string(indices[0]))
	if vIdx > 0 && vIdx <= len(*tempV) {
		v := (*tempV)[vIdx-1]
		group.Vertices = append(group.Vertices, v[0], v[1], v[2])
	}
	if len(indices) > 1 && len(indices[1]) > 0 {
		uvIdx, _ := strconv.Atoi(string(indices[1]))
		if uvIdx > 0 && uvIdx <= len(*tempUV) {
			uv := (*tempUV)[uvIdx-1]
			group.UVs = append(group.UVs, uv[0], uv[1])
		}
	}
	if len(indices) > 2 && len(indices[2]) > 0 {
		nIdx, _ := strconv.Atoi(string(indices[2]))
		if nIdx > 0 && nIdx <= len(*tempN) {
			n := (*tempN)[nIdx-1]
			group.Normals = append(group.Normals, n[0], n[1], n[2])
		}
	}
}