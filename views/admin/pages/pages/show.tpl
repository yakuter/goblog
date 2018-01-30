{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Pages</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li>
                <a href="/admin/pages">Pages</a>
            </li>
            <li class="active">
                <strong>{{.page.Title}}</strong>
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
                <h5>Page Detail</h5>
                <div class="ibox-tools">
                    <a href="/admin/pages/create" class="btn btn-primary btn-xs">Add New</a>
                    <a href="/admin/pages/edit/{{.page.Id}}" class="btn btn-primary btn-xs">Edit</a>
                </div>
            </div>
            <div class="ibox-content">
                    <h3>{{.page.Title}}</h3>
                <div class="small m-b-xs">
                    <!--<strong>Erhan Yakut</strong>--> <span class="text-muted"><i class="fa fa-clock-o"></i> {{.page.Time}}</span>

                </div>
                <p>{{str2html .page.Content}}</p>
                <span class="btn btn-primary btn-xs"><strong>ID:</strong> #{{.page.Id}}</span>
                <span class="btn btn-primary btn-xs"><strong>Slug:</strong> {{.page.Slug}}</span>
            </div>
        </div>
    </div>
</div>
{{ end }}