import "os"
import "strings"
import "math"
import "gl"

type Mesh struct {
    vertices    array
    normals     array
    texcoords   array
    triangles   array
    mtllib      string
    materials   map
    current_mtl string
}

fn parse_vertex_index(token) {
    parts := strings.split(token, "/")
    vi := -1
    vti := -1
    vni := -1

    if len(parts) >= 1 && parts[0] != "" {
        vi = int(parts[0]) - 1
    }
    if len(parts) >= 2 && parts[1] != "" {
        vti = int(parts[1]) - 1
    }
    if len(parts) >= 3 && parts[2] != "" {
        vni = int(parts[2]) - 1
    }
    return [vi, vti, vni]
}

fn triangulate_polygon(face_vertices) {
    n := len(face_vertices)
    if n < 3 { return [] }
    triangles := []
    
    for i := 1; i < n-1; i++ {
        triangles = append(triangles, face_vertices[0][0], face_vertices[0][1], face_vertices[0][2])
        triangles = append(triangles, face_vertices[i][0], face_vertices[i][1], face_vertices[i][2])
        triangles = append(triangles, face_vertices[i+1][0], face_vertices[i+1][1], face_vertices[i+1][2])
    }
    return triangles
}

fn load_obj(path) {
    data := os.read_file(path)
    if is_error(data) {
        println("Error reading file:", path, data)
        return null
    }

    content := string(data)
    lines := strings.split(content, "\n")

    vertices := []
    normals := []
    texcoords := []
    face_list := []
    mtllib := ""
    current_mtl := ""
    materials := {}

    for line in lines {
        line = strings.trim(line, " \t\r")
        if line == "" || strings.has_prefix(line, "#") { continue }

        parts := strings.split(line, " ")
        filtered := []
        for p in parts {
            if p != "" { filtered = append(filtered, p) }
        }
        if len(filtered) == 0 { continue }
        
        cmd := filtered[0]

        if cmd == "v" {
            vertices = append(vertices, float(filtered[1]), float(filtered[2]), float(filtered[3]))
        } else if cmd == "vn" {
            normals = append(normals, float(filtered[1]), float(filtered[2]), float(filtered[3]))
        } else if cmd == "vt" {
            texcoords = append(texcoords, float(filtered[1]), float(filtered[2]))
        } else if cmd == "f" {
            face_verts := []
            for i := 1; i < len(filtered); i++ {
                face_verts = append(face_verts, parse_vertex_index(filtered[i]))
            }
            
            tri_indices := triangulate_polygon(face_verts)
            for j := 0; j < len(tri_indices); j += 3 {
                face_list = append(face_list, tri_indices[j], tri_indices[j+2], tri_indices[j+1])
            }
        } else if cmd == "mtllib" {
            mtllib = filtered[1]
        } else if cmd == "usemtl" {
            current_mtl = filtered[1]
        }
    }
	
    return Mesh{
        vertices: vertices,
        normals: normals,
        texcoords: texcoords,
        triangles: face_list,
        mtllib: mtllib,
        materials: materials,
        current_mtl: current_mtl
    }
}

fn (self Mesh) draw() {
    n_triangles := len(self.triangles) / 3
    if n_triangles == 0 { return }

    gl.begin(gl.TRIANGLES)
    for i := 0; i < len(self.triangles); i += 3 {
        vi := self.triangles[i]
        vni := self.triangles[i+1]
        vti := self.triangles[i+2]

        if vni >= 0 {
            idx := vni * 3
            gl.normal3f(self.normals[idx], self.normals[idx+1], self.normals[idx+2])
        }
        
        if vti >= 0 {
            idx := vti * 2
            gl.tex_coord2f(self.texcoords[idx], self.texcoords[idx+1])
        }

        if vi >= 0 {
            idx := vi * 3
            gl.vertex3f(self.vertices[idx], self.vertices[idx+1], self.vertices[idx+2])
        } else {
            gl.vertex3f(0, 0, 0)
        }
    }
    gl.end()
}

export {
    load_obj: load_obj
}