{{ define "photos/index.tpl"}}
    {{ template "layouts/header.tpl" .}}

        {{if .page}}
        <script>
            function generatePaginationLinks(link, currentPage, totalPages, pageSize) {
                let paginationHTML = '';

                if (totalPages > 1) {
                    paginationHTML += '<div class="pagination">';

                    if (currentPage > 1) {
                        paginationHTML += `<a href="${link}?page=${currentPage - 1}&pageSize=${pageSize}">Previous</a>`;
                    }

                    for (let page = 1; page <= totalPages; page++) {
                        if (page === currentPage) {
                            paginationHTML += `<span class="current-page">${page}</span>`;
                        } else {
                            paginationHTML += `<a href="${link}?page=${page}&pageSize=${pageSize}">${page}</a>`;
                        }
                    }

                    if (currentPage < totalPages) {
                        paginationHTML += `<a href="${link}?page=${currentPage + 1}&pageSize=${pageSize}">Next</a>`;
                    }

                    paginationHTML += '</div>';
                }

                return paginationHTML;
            }

            window.onload = function() {
                const paginationLinks = generatePaginationLinks("/photos", {{.page}}, {{.totalPages}}, {{.pageSize}});

                const contentDiv = document.getElementById('pagination');
                contentDiv.innerHTML = paginationLinks;
            };
        </script>
        {{end}}

        <div>
            <h1>Photos</h1>

            <div id="pagination" class="pagination">
            </div>

            <br/>
            <br/>

            <ul>
                {{ with .photos }}
                    {{ range . }}
                        <li>
                            <div>
                                <a href="/photos/{{.ID}}">
                                    <img src="{{ .Url }}" alt="Photo" width="500" height="500">
                                </a>
                                <p>{{ .Description }}</p>
                            </div>
                            <br/>
                        </li>
                        <div style="margin-left: 50px;">
                            {{ template "comments/list_subview.tpl" .Comments }}
                            <a href="/photos/{{.ID}}/comments" role="button">Comments</a>
                        </div>
                    {{ end }}
                {{ end }}
            </ul>
        </div>
    {{ template "layouts/footer.tpl" .}}
{{end}}