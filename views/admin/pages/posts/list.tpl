{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Posts</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li class="active">
                <strong>Posts</strong>
            </li>
        </ol>
    </div>
</div>
{{ end }}

{{ define "content" }}
<div class="row">
    <div class="col-lg-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>Posts</h5>
                <div class="ibox-tools">
                    <a href="/admin/posts/create" class="btn btn-primary btn-xs">Add New</a>
                </div>
            </div>
            <div class="ibox-content">
                {{if .flash.notice}}
                <div class="alert alert-success alert-dismissable">
                    <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
                    {{.flash.notice}}.
                </div>
                {{ end }}
                <table class="table">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Title</th>
                        <th>Category</th>
                        <th>Time</th>
                        <th class="text-right">Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{ if .posts}}
                        {{range $post:=.posts}}
                        <tr>
                            <td>{{$post.Id}}</td>
                            <td>{{$post.Title}}</td>
                            <td>{{$post.Category.Title}}</td>
                            <td>{{date $post.Time "d.m.Y / H:i:s"}}</td>
                            <td class="text-right">
                                <div class="btn-group">
                                    <a href="/admin/posts/show/{{$post.Id}}" class="btn btn-xs btn-primary">View</a>
                                    <a href="/admin/posts/edit/{{$post.Id}}" class="btn btn-xs btn-warning">Edit</a>
                                    <a href="/admin/posts/delete/{{$post.Id}}" class="btn btn-xs btn-danger">Delete</a>
                                </div>
                            </td>
                        </tr>
                        {{end}}
                    {{else}}
                        <tr>
                            <td colspan="6">No posts here</td>
                        </tr>
                    {{end}}
                    </tbody>
                </table>

            </div>
        </div>
    </div>
</div>
{{ end }}