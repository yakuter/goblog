{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Categories</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li>
                <a href="/admin/categories">Categories</a>
            </li>
            <li class="active">
                <strong>{{.category.Title}}</strong>
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
                <h5>View Category</h5>
                <div class="ibox-tools">
                    <a href="/admin/categories/create" class="btn btn-primary btn-xs">Add New</a>
                    <a href="/admin/categories/edit/{{.category.Id}}" class="btn btn-primary btn-xs">Edit</a>
                </div>
            </div>
            <div class="ibox-content">
                <h3>{{.category.Title}}</h3>
                <p>{{str2html .category.Description}}</p>
                <span class="btn btn-primary btn-xs"><strong>ID:</strong> #{{.category.Id}}</span>
                <span class="btn btn-primary btn-xs"><strong>Slug:</strong> {{.category.Slug}}</span>
            </div>
        </div>
    </div>
</div>
{{ end }}