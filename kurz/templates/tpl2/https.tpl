{{ $port := .Port -}}
{{ $root := .Root -}}

{{ range $index, $dom := .Domains }}
# $index: {{ velke $dom  }} -> $root
server {
    listen      {{ $port }} ssl http2;
    server_name {{ $dom }};
    root        {{ $root }};

    ssl_certificate     /etc/ssl/{{ $dom }}.crt;
    ssl_certificate_key /etc/ssl/{{ $dom }}.key;
}
{{ end }}