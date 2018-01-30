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
        $('#info').summernote({height: 200});
    });
</script>
{{ end }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Users</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li>
                <a href="/admin/users">Users</a>
            </li>
            <li class="active">
                <strong>Edit User</strong>
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
                <h5>Edit User</h5>
            </div>
            <div class="ibox-content">
                {{if .flash.error}}
                <div class="alert alert-danger alert-dismissable">
                    <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
                    {{.flash.error}}.
                </div>
                {{end}}
                <form method="post" action="/admin/users/edit/{{.user.Id}}" class="form-horizontal">
                    {{ .xsrfdata }}
                    <div class="form-group"><label class="col-sm-2 control-label">Name Surname</label>
                        <div class="col-sm-10"><input value="{{.user.NameSurname}}" type="text" name="name_surname" id="name_surname" class="form-control"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Username</label>
                        <div class="col-sm-10"><input value="{{.user.UserName}}" type="text" name="user_name" id="user_name" class="form-control"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Email</label>
                        <div class="col-sm-10"><input value="{{.user.Email}}" type="text" name="email" id="email" class="form-control"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Password</label>
                        <div class="col-sm-10"><input value="{{.user.Password}}" type="text" name="password" id="password" class="form-control"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Url</label>
                        <div class="col-sm-10"><input value="{{.user.Url}}" type="text" name="url" id="url" class="form-control"></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group"><label class="col-sm-2 control-label">Info</label>
                        <div class="col-sm-10"><textarea name="info" id="info" class="form-control">{{.user.Info}}</textarea></div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <div class="col-sm-6 col-sm-offset-2">
                            <button class="btn btn-primary" type="submit">Save changes</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
{{ end }}