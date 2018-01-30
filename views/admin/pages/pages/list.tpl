{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Pages</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li class="active">
                <strong>Pages</strong>
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
                <h5>Pages</h5>
                <div class="ibox-tools">
                    <a href="/admin/pages/create" class="btn btn-primary btn-xs">Add New</a>
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
                        <th>Slug</th>
                        <th>Time</th>
                        <th class="text-right">Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{ if .pages}}
                        {{range $page:=.pages}}
                        <tr>
                            <td>{{$page.Id}}</td>
                            <td>{{$page.Title}}</td>
                            <td>{{$page.Slug}}</td>
                            <td>{{date $page.Time "d.m.Y / H:i:s"}}</td>
                            <td class="text-right">
                                <div class="btn-group">
                                    <a href="/admin/pages/show/{{$page.Id}}" class="btn btn-xs btn-primary">View</a>
                                    <a href="/admin/pages/edit/{{$page.Id}}" class="btn btn-xs btn-warning">Edit</a>
                                    <a href="/admin/pages/delete/{{$page.Id}}" class="btn btn-xs btn-danger">Delete</a>
                                </div>
                            </td>
                        </tr>
                        {{end}}
                    {{else}}
                        <tr>
                            <td colspan="5">No pages here</td>
                        </tr>
                    {{end}}
                    </tbody>
                </table>

            </div>
        </div>
    </div>
</div>
{{ end }}