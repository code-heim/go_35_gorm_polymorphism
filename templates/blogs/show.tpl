{{ define "blogs/show.tpl" }}
    {{ template "layouts/header.tpl" .}}
        <h1 class="card-title">{{ .blog.Title }}</h1>
        <p class="card-text">{{ .blog.Content }}</p>

        <div style="margin-left: 50px;">
            {{ template "comments/list_subview.tpl" .blog.Comments }}
            <a href="/blogs/{{.blog.ID}}/comments" role="button">Comments</a>
        </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}