{{ define "photos/show.tpl" }}
    {{ template "layouts/header.tpl" .}}
        <img src="{{ .photo.Url }}" alt="Photo" width="500" height="500">
        <p>{{ .photo.Description }}</p>
        <div style="margin-left: 50px;">
            {{ template "comments/list_subview.tpl" .photo.Comments }}
            <a href="/photos/{{.photo.ID}}/comments" role="button">Comments</a>
        </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}