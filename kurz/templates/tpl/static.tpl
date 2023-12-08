server {
    listen      {{ .Port }};
    server_name {{ range .Domains }}{{ . }} {{ end }};
    root        {{ .Root }};
}
