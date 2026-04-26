package content

func init() {
	Register(&Package{
		Name:       "mime/multipart",
		ImportPath: "mime/multipart",
		Category:   "Networking",
		Summary:    "Read and write multipart messages — file uploads, email parts.",
		Sections: []Section{
			{
				Title: "Reading a multipart body on the server",
				Examples: []Example{
					{
						Title: "File upload handler",
						Notes: "ParseMultipartForm parses then r.FormFile pulls the file out. maxMemory bytes stay in RAM; bigger parts spill to temp files.",
						Code: `func upload(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MiB in memory
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    file, header, _ := r.FormFile("file")
    defer file.Close()
    fmt.Println(header.Filename, header.Size)
    io.Copy(os.Stdout, file)
}`,
					},
				},
			},
			{
				Title: "Writing a multipart body on the client",
				Examples: []Example{
					{
						Title: "Post a file",
						Code: `var body bytes.Buffer
w := multipart.NewWriter(&body)

part, _ := w.CreateFormFile("file", "data.bin")
io.Copy(part, f)
w.WriteField("note", "hello")
w.Close()   // MUST close to write the closing boundary

req, _ := http.NewRequest("POST", url, &body)
req.Header.Set("Content-Type", w.FormDataContentType())
http.DefaultClient.Do(req)`,
					},
				},
			},
			{
				Title: "Low-level — multipart.Reader",
				Examples: []Example{
					{
						Title: "Streaming parts",
						Notes: "Prefer this for huge uploads — never buffers the whole body.",
						Code: `mr := multipart.NewReader(r.Body, boundary)
for {
    part, err := mr.NextPart()
    if err == io.EOF { break }
    fmt.Println(part.FormName(), part.FileName())
    io.Copy(dst, part)
}`,
					},
				},
			},
		},
	})
}
