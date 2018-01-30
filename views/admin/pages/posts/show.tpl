{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Posts</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li>
                <a href="/admin/posts">Posts</a>
            </li>
            <li class="active">
                <strong>{{.post.Title}}</strong>
            </li>
        </ol>
    </div>
</div>
{{ end }}

{{ define "content" }}
<div class="row">
    <div class="col-lg-12">

        <div class="ibox">
            <div class="ibox-title">
                <h5>Post Detail</h5>
                <div class="ibox-tools">
                    <a href="/admin/posts/create" class="btn btn-primary btn-xs">Add New</a>
                    <a href="/admin/posts/edit/{{.post.Id}}" class="btn btn-primary btn-xs">Edit</a>
                </div>
            </div>
            <div class="ibox-content">
                    <h3>{{.post.Title}}</h3>
                <div class="small m-b-xs">
                    <!--<strong>Erhan Yakut</strong>--> <span class="text-muted"><i class="fa fa-clock-o"></i> {{.post.Time}}</span>

                </div>
                <p>{{str2html .post.Content}}</p>
                <span class="btn btn-primary btn-xs"><strong>ID:</strong> #{{.post.Id}}</span>
                <span class="btn btn-primary btn-xs"><strong>Category:</strong> {{.post.Category.Title}}</span>
                <span class="btn btn-primary btn-xs"><strong>Slug:</strong> {{.post.Slug}}</span>
            </div>
        </div>
    </div>
</div>
{{ end }}