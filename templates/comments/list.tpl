{{ define "comments/list.tpl"}}
    {{ template "layouts/header.tpl" .}}
    <div>
        {{ template "comments/list_subview.tpl" .comments }}
        
        <h4>Add a Comment</h4>
        <form action="" method="POST">
            <div>
                <label for="comment" class="form-label">Comment</label>
                <textarea class="form-control" id="comment" rows="3" name="comment" placeholder="Type your comment here"></textarea>
            </div>
            <button type="submit" class="btn btn-primary">post</button>
        </form>
    </div>

{{ end }}