{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Categories</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li class="active">
                <strong>Categories</strong>
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
                <h5>Categories</h5>
                <div class="ibox-tools">
                    <a href="/admin/categories/create" class="btn btn-primary btn-xs">Add New</a>
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
                            <th>Description</th>
                            <th class="text-right">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ if .categories}}
                            {{range $category:=.categories}}
                            <tr>
                                <td>{{$category.Id}}</td>
                                <td>{{$category.Title}}</td>
                                <td>{{$category.Slug}}</td>
                                <td>{{html2str $category.Description}}</td>
                                <td class="text-right">
                                    <div class="btn-group">
                                        <a href="/admin/categories/show/{{$category.Id}}" class="btn btn-xs btn-primary">View</a>
                                        <a href="/admin/categories/edit/{{$category.Id}}" class="btn btn-xs btn-warning">Edit</a>
                                        <a href="/admin/categories/delete/{{$category.Id}}" class="btn btn-xs btn-danger">Delete</a>
                                    </div>
                                </td>
                            </tr>
                            {{end}}
                        {{else}}
                            <tr>
                                <td colspan="5">No categories here</td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>

            </div>
        </div>
    </div>
</div>
{{ end }}