{{ template "../../layouts/master.tpl" . }}

{{ define "pagehead" }}
<link href="/static/admin/css/plugins/summernote/summernote.css" rel="stylesheet">
<link href="/static/admin/css/plugins/summernote/summernote-bs3.css" rel="stylesheet">
{{ end }}

{{ define "pagefoot" }}
<!-- SUMMERNOTE -->
<script src="/static/admin/js/plugins/summernote/summernote.min.js"></script>

<script type="text/javascript">
    jQuery(function($){
        $('#description').summernote({height: 200});
    });
</script>

{{ end }}

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
                <strong>New Category</strong>
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
                <h5>New Category</h5>
            </div>
            <div class="ibox-content">

                {{if .flash.error}}
                <div class="alert alert-danger alert-dismissable">
                    <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
                    {{.flash.error}}.
                </div>
                {{end}}
                <form method="post" action="/admin/categories/create" class="form-horizontal">
                    {{ .xsrfdata }}
                    <div class="form-group"><label class="col-sm-2 control-label">Title</label>
                        <div class="col-sm-10"><input type="text" name="title" id="title" placeholder="Sample Title" class="form-control slug-source"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Slug</label>
                        <div class="col-sm-10"><input type="text" name="slug" id="slug" placeholder="sample-title" class="form-control slug-target"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Description</label>
                        <div class="col-sm-10"><textarea name="description" id="description" class="form-control"></textarea></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <div class="col-sm-6 col-sm-offset-2">
                            <!--<button class="btn btn-white" type="submit">Cancel</button>-->
                            <button class="btn btn-primary" type="submit">Save</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
{{ end }}