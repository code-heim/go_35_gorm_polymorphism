{{ define "comments/list_subview.tpl"}}
    {{ if . }}
        <h2>Comments</h2>
        <div>
            {{ with . }}
                {{ range . }}
                    <div style="border: 2px solid #000000; padding: 20px; margin: 20px;">
                        <h5>{{ .Content }}</h5>
                        <small>Posted on {{ .UpdatedAt }}</small>
                    </div>
                    <hr class="solid"></hr>
                {{ end }}
            {{ end }}
        </div>
    {{ end }}
{{ end }}